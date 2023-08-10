package main

import (
	"fmt"
	"time"
)

func main() {
	firstTime := 1688140800000
	secondTime := 1690819200000
	start := time.Unix(0, int64(firstTime*1e6))
	end := time.Unix(0, int64(secondTime*1e6))
	fmt.Println("minus : %v", end.Sub(start).Hours()/24)

	a := int64(100000099)
	b := int64(32)
	ave := a / b
	var i int64
	for i = 0; i < b; i++ {
		amount := ave
		if i == b-1 {
			amount = a - ave*(b-1)
		}
		fmt.Println("results : %v", amount)
		now := start.AddDate(0, 0, int(i))
		fmt.Println("日期", now)
	}

}
