package types

type AdminChartJobMultispaceTableReply struct {
	Code int64                            `json:"code"`
	Data AdminChartJobMultispaceTableResp `json:"data"`
	Msg  string                           `json:"msg"`
}

type AdminChartJobMultispaceTableResp struct {
	Job                   Job                   `json:"job"`                   // Job
	JobCountDistribution  JobCountDistribution  `json:"jobCountDistribution"`  // 任务提交次数分布
	Jobenv                Jobenv                `json:"jobenv"`                // Jobenv
	SpaceTotalCount       int64                 `json:"spaceTotalCount"`       // 总空间数
	UserCountDistribution UserCountDistribution `json:"userCountDistribution"` // 任务提交人数分布
	UserTotalCount        int64                 `json:"userTotalCount"`        // 总用户数
	Visualenv             Visualenv             `json:"visualenv"`             // 可视化概况
}

// 离线任务概况
type Job struct {
	AverageDuation             int64  `json:"averageDuation"`             // 任务平均时长，单位h/100，前端需除以100
	AverageDuationRatio        string `json:"averageDuationRatio"`        // 任务平均时长环比，持平返回“持平”，无比较返回“-”
	AverageDuationTendency     int64  `json:"averageDuationTendency"`     // 任务平均时长环比趋势，1=上升；2=下降；3=持平；4=空值；
	GPUAverageDuration         int64  `json:"gpuAverageDuration"`         // 平均卡时，单位h/100，前端需除以100
	GPUAverageDurationRatio    string `json:"gpuAverageDurationRatio"`    // 平均卡时环比，持平返回“持平”，无比较返回“-”
	GPUAverageDurationTendency int64  `json:"gpuAverageDurationTendency"` // 平均卡时环比趋势，1=上升；2=下降；3=持平；4=空值；
	GPUTotalDuration           int64  `json:"gpuTotalDuration"`           // 总卡时，单位h/100，前端需除以100
	GPUTotalDurationRatio      string `json:"gpuTotalDurationRatio"`      // 总卡时环比，持平返回“持平”，无比较返回“-”
	GPUTotalDurationTendency   int64  `json:"gpuTotalDurationTendency"`   // 总卡时环比趋势，1=上升；2=下降；3=持平；4=空值；
	TotalCount                 int64  `json:"totalCount"`                 // 总个数
	TotalCountRatio            string `json:"totalCountRatio"`            // 总个数环比，持平返回“持平”，无比较返回“-”
	TotalCountTendency         int64  `json:"totalCountTendency"`         // 总个数环比趋势，1=上升；2=下降；3=持平；4=空值；
	TotalDuration              int64  `json:"totalDuration"`              // 任务总时长，单位h/100，前端需除以100
	TotalDurationRatio         string `json:"totalDurationRatio"`         // 任务总时长环比，持平返回“持平”，无比较返回“-”
	TotalDurationTendency      int64  `json:"totalDurationTendency"`      // 任务总时长环比趋势，1=上升；2=下降；3=持平；4=空值；
}

// 任务提交次数分布
type JobCountDistribution struct {
	JobCommitCount    int64 `json:"jobCommitCount"`    // 离线任务提交次数
	JobenvCommitCount int64 `json:"jobenvCommitCount"` // 开发环境提交次数
	JobenvFailCount   int64 `json:"jobenvFailCount"`   // 开发环境失败次数
	JobFailCount      int64 `json:"jobFailCount"`      // 离线任务失败次数
}

// 开发环境概况
type Jobenv struct {
	AverageDuation             int64  `json:"averageDuation"`             // 任务平均时长，单位h/100，前端需除以100
	AverageDuationRatio        string `json:"averageDuationRatio"`        // 任务平均时长环比，持平返回“持平”，无比较返回“-”
	AverageDuationTendency     int64  `json:"averageDuationTendency"`     // 任务平均时长环比趋势，1=上升；2=下降；3=持平；4=空值；
	GPUAverageDuration         int64  `json:"gpuAverageDuration"`         // 平均卡时，单位h/100，前端需除以100
	GPUAverageDurationRatio    string `json:"gpuAverageDurationRatio"`    // 平均卡时环比，持平返回“持平”，无比较返回“-”
	GPUAverageDurationTendency int64  `json:"gpuAverageDurationTendency"` // 平均卡时环比趋势，1=上升；2=下降；3=持平；4=空值；
	GPUTotalDuration           int64  `json:"gpuTotalDuration"`           // 总卡时，单位h/100，前端需除以100
	GPUTotalDurationRatio      string `json:"gpuTotalDurationRatio"`      // 总卡时环比，持平返回“持平”，无比较返回“-”
	GPUTotalDurationTendency   int64  `json:"gpuTotalDurationTendency"`   // 总卡时环比趋势，1=上升；2=下降；3=持平；4=空值；
	TotalCount                 int64  `json:"totalCount"`                 // 总个数
	TotalCountRatio            string `json:"totalCountRatio"`            // 总个数环比，持平返回“持平”，无比较返回“-”
	TotalCountTendency         int64  `json:"totalCountTendency"`         // 总个数环比趋势，1=上升；2=下降；3=持平；4=空值；
	TotalDuration              int64  `json:"totalDuration"`              // 任务总时长，单位h/100，前端需除以100
	TotalDurationRatio         string `json:"totalDurationRatio"`         // 任务总时长环比，持平返回“持平”，无比较返回“-”
	TotalDurationTendency      int64  `json:"totalDurationTendency"`      // 任务总时长环比趋势，1=上升；2=下降；3=持平；4=空值；
}

// 任务提交人数分布
type UserCountDistribution struct {
	JobDistributeUserCount int64 `json:"jobDistributeUserCount"` // 训练任务-分布式人数
	JobenvUserCount        int64 `json:"jobenvUserCount"`        // 开发环境人数
	JobSignleUserCount     int64 `json:"jobSignleUserCount"`     // 训练任务-单机人数
	TotalUserCount         int64 `json:"totalUserCount"`         // 任务提交总人数
	VisualenvUserCount     int64 `json:"visualenvUserCount"`     // 可视化人数
}

// 可视化概况
type Visualenv struct {
	AverageDuation         int64  `json:"averageDuation"`         // 任务平均时长，单位h/100，前端需除以100
	AverageDuationRatio    string `json:"averageDuationRatio"`    // 任务平均时长环比，持平返回“持平”，无比较返回“-”
	AverageDuationTendency int64  `json:"averageDuationTendency"` // 任务平均时长环比趋势，1=上升；2=下降；3=持平；4=空值；
	TotalCount             int64  `json:"totalCount"`             // 总个数
	TotalCountRatio        string `json:"totalCountRatio"`        // 总个数环比，持平返回“持平”，无比较返回“-”
	TotalCountTendency     int64  `json:"totalCountTendency"`     // 总个数环比趋势，1=上升；2=下降；3=持平；4=空值；
	TotalDuration          int64  `json:"totalDuration"`          // 任务总时长，单位h/100，前端需除以100
	TotalDurationRatio     string `json:"totalDurationRatio"`     // 任务总时长环比，持平返回“持平”，无比较返回“-”
	TotalDurationTendency  int64  `json:"totalDurationTendency"`  // 任务总时长环比趋势，1=上升；2=下降；3=持平；4=空值；
}
