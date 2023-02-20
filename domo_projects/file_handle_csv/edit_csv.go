package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// 文件列表
	fileList := []string{
		"/Users/jay/Desktop/chart_job_202212012005",
		"/Users/jay/Desktop/chart_job_status_202212012005",
		"/Users/jay/Desktop/chart_space_user_202212012005",
	}
	changeFile(fileList)

}

func changeFile(fileList []string) {
	for _, file := range fileList {
		csvFile, err := os.Open(file + ".csv")
		log.Println("csvFile:", file)
		if err != nil {
			log.Println("open csv file error")
		}
		defer csvFile.Close()
		csvLines, err := csv.NewReader(csvFile).ReadAll()
		if err != nil {
			log.Println("open csv file error")
		}
		var newlines [][]string
		var id, startTime, period, creatTime, updateTime int
		for i, line := range csvLines {
			if i == 0 {
				newlines = append(newlines, line)
				for j, head := range line {
					switch head {
					case "id":
						id = j
						log.Println("id:", id)
					case "start_time":
						startTime = j
						log.Println("startTime:", startTime)
					case "period":
						period = j
						log.Println("period:", period)
					case "create_time":
						creatTime = j
						log.Println("creatTime:", creatTime)
					case "update_time":
						updateTime = j
						log.Println("updateTime:", updateTime)
					}
				}
			}
			// 将 2022-11-22 00:00:00 日数据转换为 2022-11-07 00:00:00 周数据
			if line[startTime] == "1669737600000" && line[period] == "day" {
				newli := line
				newli[id] = ""
				newli[startTime] = "1668960000000"
				newli[period] = "week"
				//now := time.Now().Format("2006-01-02 15:04:05")
				newli[creatTime] = ""
				if updateTime != 0 {
					newli[updateTime] = ""
				}
				newlines = append(newlines, newli)
			}
		}
		File, err := os.OpenFile(file+"_week_2.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Println("文件打开失败！")
		}
		defer File.Close()
		WriterCsv := csv.NewWriter(File)
		for _, line := range newlines {
			err1 := WriterCsv.Write(line)
			if err1 != nil {
				log.Println("WriterCsv写入文件失败")
			}
			WriterCsv.Flush()
		}
	}
}
