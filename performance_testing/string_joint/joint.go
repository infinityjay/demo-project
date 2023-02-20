package string_joint

import (
	"bytes"
	"fmt"
	"strings"
)

func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}
func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func bufferConcat(n int, str string) string {
	buffer := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buffer.WriteString(str)
	}
	return buffer.String()
}

func byteConcat(n int, str string) string {
	buf := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}
