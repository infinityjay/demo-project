package main

import "fmt"

func main() {
	list := []int{12, 23, 42, 54, 65, 43, 34, 98, 903, 432, 456, 234, 254, 333}
	for left := 0; left < len(list); {
		right := left + 3
		if right > len(list) {
			right = len(list)
		}
		newlist := list[left:right]
		fmt.Println(newlist)
		fmt.Printf("right = %v\n", right)
		left = right
	}
}
