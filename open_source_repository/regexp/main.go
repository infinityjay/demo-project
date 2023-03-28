package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `This is a string with an image 
![example](https://example.com/image.png) ![11](https://example.com/iiiii) and a ![2222asdfs..//()]asf(https://asfdsfadfs) link (https://example.com/).`
	re := regexp.MustCompile(`!\[[^\]]*\]\((https?://[^\)]+)\)`)
	matches := re.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		fmt.Println(match[1])
	}
}
