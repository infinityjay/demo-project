package main

import (
	"fmt"
	"time"
)

func main() {
	submitTime := time.Unix(1684425600, 0)
	submitString := submitTime.Format("2006-01-02")
	fmt.Println("The submit day is : ", submitString)
	securityTime := time.Unix(1688136367, 0)
	future := securityTime.AddDate(0, 0, 220).Format("2006-01-02")
	fmt.Println("The estimated day of finish check: ", future)

}
