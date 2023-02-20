package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var up = websocket.Upgrader{
	//允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//Websocket服务
func WebsocketServer(c echo.Context) error {
	var (
		wsConn        *websocket.Conn
		conn          *util.Connection
		pubSub        *redis.PubSub
		enterexitId   []byte
		redisChanData <-chan *redis.Message
		err           error
	)
	//注册Websocket
	wsConn, err = upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("UpgradeErr:", err)
		return err
	}

	//初始化Connection
	conn, err = util.InitConnection(wsConn)
	if err != nil {
		log.Println("InitConnection_Err:", err)
		goto ERROR
	}

	//读取客户端发来的 enterexitId
	enterexitId, err = conn.ReadMessage()
	if err != nil {
		log.Println(conn.SocketId+" >>> ReadMessage_Err:", err)
		goto ERROR
	}
	log.Println("ReadMessage:", string(enterexitId))

	//注册消息订阅监听,订阅推送
	pubSub = db.Redis().Subscribe(string(enterexitId))
	defer pubSub.Close()

	_, err = pubSub.Receive()
	if err != nil {
		log.Println("Receive_Err:", err)
		goto ERROR
	}

	//获取推送的Channel
	redisChanData = pubSub.Channel()

	//独立协程，不停的发送心跳30秒一次
	go sendHeartBeat(conn, "heart_beat:"+string(enterexitId))

	//等待消息发送
	for value := range redisChanData {

		err = conn.WriteMessage([]byte(value.Payload))
		if err != nil {
			log.Println("WriteMessage_Err:", err)
			goto ERROR
		}
	}

ERROR:
	//todo:关闭连接操作
	conn.Close()
	return nil
}

//发送心跳包，30秒频率
func sendHeartBeat(conn *util.Connection, data string) {
	var err error
	for {
		err = conn.WriteMessage([]byte(data))
		if err != nil {
			log.Println(data+":", err)
			return
		}
		time.Sleep(30 * time.Second)
	}
}
