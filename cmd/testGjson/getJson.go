package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	res := &cat{
		Age:  10,
		Name: "test",
	}
	by, _ := json.Marshal(res)
	code := gjson.Get(string(by), "code").Int()
	msg := gjson.Get(string(by), "msg").String()
	fmt.Printf("code: %v, msg: %v", code, msg)
}

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type cat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
