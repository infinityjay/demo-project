package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	baseUrl := "https://platform-dev4.virtaicloud.com:31443/gemini/v1/gemini_api/gemini_api/admin/ops/component/logList?realTimeQuery=2&sortOrder=asc&sortField=time"
	keyword := ""
	// startTime endTime 为毫秒时间戳，时间间隔不超过31天
	startTime := "1657870417000"
	endTime := "1658475218000"
	// limit 为请求日志条数，一般不要设置过大，loki本身有限制
	limit := "10000"
	reqUrl := baseUrl + "&keywords=" + keyword + "&startTime=" + startTime + "&endTime=" + endTime + "&limit=" + limit
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		fmt.Println("http request error")
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("request body error")
	}
	fmt.Println(string(body))
}
