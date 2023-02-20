package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 1
	fmt.Printf("原始map的内存地址是：%p\n", m)
	modify(m)
	fmt.Println("map值被修改了，新值为:", m)
}
func modify(m interface{}) {
	fmt.Printf("函数里接收到map的内存地址是：%p\n", m)
	p := m.(map[int]int)
	p[1] = 2
}
