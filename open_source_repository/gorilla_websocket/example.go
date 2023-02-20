package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Accepting all requests
	},
}

type request struct {
	RealTime int `json:"realTime"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//upgrade get request to websocket protocol
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer ws.Close()
		count := 1
		stopCh := make(chan struct{})
		for {
			//Read Message from client
			mt, message, err := ws.ReadMessage()
			var req request
			json.Unmarshal(message, &req)
			if err != nil || mt == websocket.CloseMessage || req.RealTime == 2 {
				close(stopCh)
				return
			}

			//Response message to client
			count++
			fmt.Printf("count:%v\n", count)
			go write(ws, count, stopCh)
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func write(connection *websocket.Conn, count int, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			break
		default:
			connection.WriteMessage(websocket.TextMessage, []byte("res"))
			fmt.Printf("go count:%v\n", count)
			time.Sleep(2 * time.Second)
		}
	}
}
