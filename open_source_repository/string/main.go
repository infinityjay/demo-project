package main

import (
	"fmt"
)

var chartExportMap = map[string]string{
	"001001": "资源数量统计-节点",
	"001002": "资源使用情况统计-节点",
	"001003": "节点清单-资源使用情况",
	"001004": "节点清单-资源占用情况",
	"002001": "资源数量统计-GPU卡",
	"002002": "资源使用情况统计-GPU卡",
	"002003": "GPU卡型号分布",
	"002004": "使用环比(未虚拟化GPU卡)",
	"002005": "利用率统计(未虚拟化GPU卡)",
	"002006": "使用环比(已虚拟化GPU卡)",
	"002007": "利用率统计(已虚拟化GPU卡)",
	"002008": "GPU卡清单",
	"003001": "任务提交人数分布",
	"003002": "任务失败次数分布",
	"003003": "空间数量统计",
	"003004": "平台用户数量统计",
	"004001": "环比统计",
	"004002": "训练时长TOP5(空间)",
	"004003": "训练时长TOP5(用户)",
	"004004": "GPU卡时TOP5(空间)",
	"004005": "GPU卡时TOP5(用户)",
	"004006": "任务提交情况",
	"005001": "环比统计",
	"005002": "训练时长TOP5(空间)",
	"005003": "训练时长TOP5(用户)",
	"005004": "GPU卡时TOP5(空间)",
	"005005": "GPU卡时TOP5(用户)",
	"005006": "任务提交情况",
	"006001": "任务情况统计-开发环境",
	"006002": "任务情况统计-离线任务",
}

func main() {
	getSuffix("ima.ge.png")
	//ids := []string{"001002", "003001", "001001", "002003", "000000"}
	//for _, id := range ids {
	//	fmt.Println(id[0:3])
	//}
	//sort.Slice(ids, func(i, j int) bool {
	//	if ids[i] < ids[j] {
	//		return true
	//	}
	//	return false
	//})
	//fmt.Println(ids)
	//
	//var array = make([]string, 1)
	//array[0] = "第一行"
	//array = append(array, "dierhang")
	//fmt.Println(array)
	//
	//fmt.Println(ids[:3])
}

func getSuffix(fileName string) {
	imageType := "jpg"
	if len(fileName) > 0 {
		i := len(fileName) - 1
		for ; i >= 0; i-- {
			if fileName[i] == '.' {
				break
			}
		}
		imageType = fileName[i+1:]
	}
	fmt.Println(imageType)
}
