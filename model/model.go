package model

import "time"

// 离线任务报表，小时表
type ChartJobHour struct {
	Id            int64     `json:"id"  xorm:"not null pk autoincr BIGINT"`
	StartTime     int64     `json:"start_time" xorm:"not null default 0 index(start_time) comment('采样时间，豪秒级unix timestamp') BIGINT"`
	JobId         string    `json:"job_id" xorm:"not null default '' VARCHAR(64)"`
	ProjectId     string    `json:"project_id" xorm:"not null default '' unique(job_name) index VARCHAR(64)"`
	JobType       int       `json:"job_type" xorm:"not null default 1 comment('1-离线任务；2-开发环境；3-可视化任务') TINYINT"`
	TrainType     int       `json:"train_type" xorm:"not null default 1 comment('1单实例2分布式3分布式horovid/mpi') TINYINT"`
	SpaceId       string    `json:"space_id" xorm:"not null default '' unique(job_name) index(space_id) VARCHAR(16)"`
	JobenvId      int64     `json:"jobenv_id" xorm:"not null default 0 BIGINT"`
	VisualenvId   int64     `json:"visualenv_id" xorm:"not null default 0 BIGINT"`
	CreateUserId  int       `json:"create_user_id" xorm:"not null default 0 index(create_user_id) INT"`
	IsFirstCommit int       `json:"is_first_commit" xorm:"not null default 0 index(is_first_commit) comment('是否首次提交任务; 0-否; 1-是') TINYINT"`
	IsFirstFail   int       `json:"is_first_fail" xorm:"not null default 0 index(is_first_fail) comment('是否首次失败; 0-否; 1-是') TINYINT"`
	GpuDuration   int64     `json:"gpu_duration" xorm:"not null default 0 comment('GPU卡时,单位h/100,ex:139表示1.39h') BIGINT"`
	JobDuration   int64     `json:"job_duration" xorm:"not null default 0 comment('任务训练时长,单位h/100,ex:139表示1.39h') BIGINT"`
	CreateTime    time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
	UpdateTime    time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
}

func (h *ChartJobHour) TableName() string {
	return "chart_job_hour"
}

// 离线任务报表，日表
type ChartJobDay struct {
	Id            int64     `json:"id"  xorm:"not null pk autoincr BIGINT"`
	StartTime     int64     `json:"start_time" xorm:"not null default 0 index(start_time) comment('采样时间，豪秒级unix timestamp') BIGINT"`
	JobId         string    `json:"job_id" xorm:"not null default '' VARCHAR(64)"`
	ProjectId     string    `json:"project_id" xorm:"not null default '' unique(job_name) index VARCHAR(64)"`
	JobType       int       `json:"job_type" xorm:"not null default 1 comment('1-离线任务；2-开发环境；3-可视化任务') TINYINT"`
	TrainType     int       `json:"train_type" xorm:"not null default 1 comment('1单实例2分布式3分布式horovid/mpi') TINYINT"`
	SpaceId       string    `json:"space_id" xorm:"not null default '' unique(job_name) index(space_id) VARCHAR(16)"`
	JobenvId      int64     `json:"jobenv_id" xorm:"not null default 0 BIGINT"`
	VisualenvId   int64     `json:"visualenv_id" xorm:"not null default 0 BIGINT"`
	CreateUserId  int       `json:"create_user_id" xorm:"not null default 0 index(create_user_id) INT"`
	IsFirstCommit int       `json:"is_first_commit" xorm:"not null default 0 index(is_first_commit) comment('是否首次提交任务; 0-否; 1-是') TINYINT"`
	IsFirstFail   int       `json:"is_first_fail" xorm:"not null default 0 index(is_first_fail) comment('是否首次失败; 0-否; 1-是') TINYINT"`
	GpuDuration   int64     `json:"gpu_duration" xorm:"not null default 0 comment('GPU卡时,单位h/100,ex:139表示1.39h') BIGINT"`
	JobDuration   int64     `json:"job_duration" xorm:"not null default 0 comment('任务训练时长,单位h/100,ex:139表示1.39h') BIGINT"`
	CreateTime    time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
	UpdateTime    time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
}

func (d *ChartJobDay) TableName() string {
	return "chart_job_day"
}

// 离线任务报表，周表
type ChartJobWeek struct {
	Id            int64     `json:"id"  xorm:"not null pk autoincr BIGINT"`
	StartTime     int64     `json:"start_time" xorm:"not null default 0 index(start_time) comment('采样时间，豪秒级unix timestamp') BIGINT"`
	JobId         string    `json:"job_id" xorm:"not null default '' VARCHAR(64)"`
	ProjectId     string    `json:"project_id" xorm:"not null default '' unique(job_name) index VARCHAR(64)"`
	JobType       int       `json:"job_type" xorm:"not null default 1 comment('1-离线任务；2-开发环境；3-可视化任务') TINYINT"`
	TrainType     int       `json:"train_type" xorm:"not null default 1 comment('1单实例2分布式3分布式horovid/mpi') TINYINT"`
	SpaceId       string    `json:"space_id" xorm:"not null default '' unique(job_name) index(space_id) VARCHAR(16)"`
	JobenvId      int64     `json:"jobenv_id" xorm:"not null default 0 BIGINT"`
	VisualenvId   int64     `json:"visualenv_id" xorm:"not null default 0 BIGINT"`
	CreateUserId  int       `json:"create_user_id" xorm:"not null default 0 index(create_user_id) INT"`
	IsFirstCommit int       `json:"is_first_commit" xorm:"not null default 0 index(is_first_commit) comment('是否首次提交任务; 0-否; 1-是') TINYINT"`
	IsFirstFail   int       `json:"is_first_fail" xorm:"not null default 0 index(is_first_fail) comment('是否首次失败; 0-否; 1-是') TINYINT"`
	GpuDuration   int64     `json:"gpu_duration" xorm:"not null default 0 comment('GPU卡时,单位h/100,ex:139表示1.39h') BIGINT"`
	JobDuration   int64     `json:"job_duration" xorm:"not null default 0 comment('任务训练时长,单位h/100,ex:139表示1.39h') BIGINT"`
	CreateTime    time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
	UpdateTime    time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
}

func (w *ChartJobWeek) TableName() string {
	return "chart_job_week"
}

// 离线任务报表，月表
type ChartJobMonth struct {
	Id            int64     `json:"id"  xorm:"not null pk autoincr BIGINT"`
	StartTime     int64     `json:"start_time" xorm:"not null default 0 index(start_time) comment('采样时间，豪秒级unix timestamp') BIGINT"`
	JobId         string    `json:"job_id" xorm:"not null default '' VARCHAR(64)"`
	ProjectId     string    `json:"project_id" xorm:"not null default '' unique(job_name) index VARCHAR(64)"`
	JobType       int       `json:"job_type" xorm:"not null default 1 comment('1-离线任务；2-开发环境；3-可视化任务') TINYINT"`
	TrainType     int       `json:"train_type" xorm:"not null default 1 comment('1单实例2分布式3分布式horovid/mpi') TINYINT"`
	SpaceId       string    `json:"space_id" xorm:"not null default '' unique(job_name) index(space_id) VARCHAR(16)"`
	JobenvId      int64     `json:"jobenv_id" xorm:"not null default 0 BIGINT"`
	VisualenvId   int64     `json:"visualenv_id" xorm:"not null default 0 BIGINT"`
	CreateUserId  int       `json:"create_user_id" xorm:"not null default 0 index(create_user_id) INT"`
	IsFirstCommit int       `json:"is_first_commit" xorm:"not null default 0 index(is_first_commit) comment('是否首次提交任务; 0-否; 1-是') TINYINT"`
	IsFirstFail   int       `json:"is_first_fail" xorm:"not null default 0 index(is_first_fail) comment('是否首次失败; 0-否; 1-是') TINYINT"`
	GpuDuration   int64     `json:"gpu_duration" xorm:"not null default 0 comment('GPU卡时,单位h/100,ex:139表示1.39h') BIGINT"`
	JobDuration   int64     `json:"job_duration" xorm:"not null default 0 comment('任务训练时长,单位h/100,ex:139表示1.39h') BIGINT"`
	CreateTime    time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
	UpdateTime    time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
}

func (m *ChartJobMonth) TableName() string {
	return "chart_job_month"
}
