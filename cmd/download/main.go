package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"os"
)

func main() {
	r := gin.New()
	r.GET("/download", DownloadFile)
	r.GET("/downloadExcel", DownloadExcel)
	fmt.Println("开始监听8888端口...")
	err := r.Run()
	if err != nil {
		return
	}
}

func DownloadFile(c *gin.Context) {
	// 创建一个缓冲区用来保存压缩文件内容
	buf := new(bytes.Buffer)
	// 创建一个压缩文档
	w := zip.NewWriter(buf)
	// 将文件加入压缩文档
	var files = []struct {
		Name, Body string
	}{
		{"Golang.xlsx", "我很开心,123,你你你,ttttt"},
		{"test.csv", "我很开心,123,你你你,ttttt"},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			fmt.Println(err)
		}
	}
	// 关闭压缩文档
	err := w.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileName := "file.zip"
	// 将压缩文档内容写入文件
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	buf.WriteTo(f)
	c.Writer.Header().Set("Content-type", "application/zip")
	c.FileAttachment(fileName, fileName)
	fmt.Println("file.zip generated")
	defer os.Remove(fileName)
}

func DownloadExcel(c *gin.Context) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1")
	if err != nil {
		panic(err)
	}
	row := sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell := row.AddCell()
	cell.Value = "haha"
	cell = row.AddCell()
	cell.Value = "我怕我骗我"
	err = file.Save("file.xlsx")
	if err != nil {
		panic(err)
	}
}
