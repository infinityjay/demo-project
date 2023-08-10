package main

import (
	"fmt"
	"time"
)

type SysComponent struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT"`
	ComponentName  string    `json:"component_name" xorm:"not null unique VARCHAR(128)"`
	Version        string    `json:"version" xorm:"VARCHAR(128)"`
	Description    string    `json:"description" xorm:"VARCHAR(512)"`
	BuildInfo      string    `json:"build_info" xorm:"VARCHAR(128)"`
	BuildTime      string    `json:"build_time" xorm:"VARCHAR(128)"`
	SortId         int       `json:"sort_id" xorm:"not null default 0 INT"`
	UpdateTime     time.Time `json:"update_time" xorm:"not null DATETIME"`
	IsManualDelete int       `json:"is_manual_delete" xorm:"not null INT"` // 1=提供手动销毁；2=不提供手动销毁；
}

func main() {
	//sysComponentList := make([]SysComponent, 0)
	//csvFile, err := os.Open("/Users/jay/Desktop/testcsv.csv")
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	//defer csvFile.Close()
	//csvLines, err := csv.NewReader(csvFile).ReadAll()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//sortId := 1
	//for i, line := range csvLines {
	//	if line[0] == "" || i == 0 {
	//		continue
	//	}
	//	sysComponent := SysComponent{
	//		ComponentName: line[0],
	//		Description:   line[2],
	//		SortId:        sortId,
	//	}
	//	if line[1] == "是" {
	//		sysComponent.IsManualDelete = 1
	//	} else if line[1] == "否" {
	//		sysComponent.IsManualDelete = 2
	//	} else {
	//		sysComponent.IsManualDelete = 1
	//	}
	//	fmt.Println(sysComponent)
	//	sysComponentList = append(sysComponentList, sysComponent)
	//	sortId++
	//}
	now := time.Now()
	//currentYear, currentMonth, _ := now.Date()
	currentLocation := time.Local
	fmt.Println(currentLocation)

	//time.Time格式
	//currentDay := time.Date(currentYear, currentMonth, now.Day(), 0, 0, 0, 0, currentLocation)
	//fmt.Println(currentDay.Format("2006年01月01日"))
	//
	//firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, currentLocation)
	//firstOfYear := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, currentLocation)
	//fmt.Println(currentDay)
	//irstOfLastDay := currentDay.AddDate(0, 0, -1)
	//fmt.Println(irstOfLastDay)
	//fmt.Println(irstOfLastDay.Add(23 * time.Hour))
	//fmt.Println(firstOfMonth)
	//fmt.Println(firstOfYear)
	//if currentDay == firstOfMonth {
	//	fmt.Println("本月第一天")
	//} else {
	//	fmt.Println("不是本月第一天")
	//}

	timeTime := time.Unix(1665763200, 0)
	fmt.Println(timeTime)
	offset := int(time.Monday - timeTime.Weekday())
	fmt.Println(offset)
	firstOfWeek := time.Date(timeTime.Year(), timeTime.Month(), timeTime.Day(), 0, 0, 0, 0, currentLocation).AddDate(0, 0, offset)
	fmt.Println(firstOfWeek)
	//fmt.Println(time.Monday)
	//if currentDay == firstOfWeek {
	//	fmt.Println("本周第一天")
	//} else {
	//	fmt.Println("不是本周第一天")
	//}
	//
	//millFirstOfWeek := firstOfWeek.UnixMilli()
	//fmt.Println(millFirstOfWeek)

	firstOfLastWeek := firstOfWeek.AddDate(0, 0, -7)
	fmt.Println(firstOfLastWeek)

	firstOfLastMonth := time.Date(now.Year(), now.Month()-time.Month(1), 1, 0, 0, 0, 0, currentLocation)
	fmt.Println(firstOfLastMonth)
	time1, _ := time.Parse(time.RFC3339, "2099-01-01T00:00:00Z")
	TS1 := time1.Unix()
	fmt.Println(time1)
	fmt.Println(TS1)
	//time1 := time.Now().Unix()
	//time2 := time.Now().Add(-3346 * time.Second).Unix()
	//minus := time1 - time2
	//fmt.Println(minus)
	//var hour float64

	//hour = float64(minus*100) / float64(3600)
	//fmt.Println(hour)
	//b := math.Round(hour)
	//fmt.Println(b)
	//fmt.Println(int64(b))
	//hour2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(minus*100)/float64(3600)), 64)
	//fmt.Println(hour2)
	//a := math.Round(hour2)
	//fmt.Println(a)
	//
	//var t time.Time
	//fmt.Println(t)
	//fmt.Println(t.Unix())
	//fmt.Println(t.IsZero())
	//fmt.Println(currentDay.IsZero())
	hour2 := fmt.Sprintf("%.2f", float64(2314)/float64(36)) + "%"
	fmt.Println(hour2)

	//type Tables struct {
	//	TableName string
	//	ModelType interface{}
	//}
	//
	//var tableMap = map[string]Tables{
	//	"day":   {"chart_job_day", new([]*model.ChartJobDay)},
	//	"week":  {"chart_job_week", new([]*model.ChartJobWeek)},
	//	"month": {"chart_job_month", new([]*model.ChartJobMonth)},
	//}
	//fmt.Println(tableMap["day"].TableName)
	//fmt.Println(tableMap["day"].ModelType)

	fmt.Println(">>>>>>>>>>> test month <<<<<<<<<<<<")
	now = time.Unix(1669824000, 0)
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, currentLocation)
	fmt.Printf("firstOfMonth: %v\n", firstOfMonth)
	monthStartTime := time.Date(now.Year(), now.Month()-time.Month(1), 1, 0, 0, 0, 0, currentLocation)
	fmt.Printf("monthStartTime: %v\n", monthStartTime)

	endTime := firstOfLastMonth.AddDate(0, 1, 0).Format("2006-01-02 15:04:05")
	fmt.Println("startTime: ", firstOfLastMonth)
	fmt.Println("endTime: ", endTime)
}
