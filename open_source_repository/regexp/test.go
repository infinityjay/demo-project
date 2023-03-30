package main

import "fmt"
import "strings"

func main() {
	str := "http://test.bj.bcebos.com/path/to/file"
	indexStr := "test.bj.bcebos.com/"

	// 查找 "test.bj.bcebos.com" 在原始字符串中的位置
	index := strings.Index(str, indexStr)

	if index != -1 {
		// 使用 str[index:] 获取从该位置开始到字符串末尾的子字符串
		subStr := str[index+len(indexStr):]

		fmt.Println(subStr) // 输出 "test.bj.bcebos.com/path/to/file"
	}
}
