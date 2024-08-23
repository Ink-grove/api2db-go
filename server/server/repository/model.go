package repository

import (
	"time"
)

// RespComm 通用的响应消息
type RespComm struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// CreateTaskReq 请求消息
type CreateTaskReq struct {
	TaskData TaskData `json:"task_data"`
}

// CreateTaskResp 响应消息
type CreateTaskResp struct {
	RespComm
	TaskId string `json:"task_id"`
}

type TaskData struct {
	ID          int64     `json:"id" gorm:"id" 	`
	TaskType    string    `json:"task_type" gorm:"task_type"`     // 项目
	InitParam   InitParam `json:"init_param" gorm:"init_param"`   // 初始化参数，主要为数据源、目标数据库以及对应字段映射情况参数
	Param       Context   `json:"param" gorm:"param"`             // 任务参数，
	Status      int32     `json:"status" gorm:"status"`           // 状态
	Description string    `json:"description" gorm:"description"` // 任务描述
	LastTime    time.Time `json:"last_time" gorm:"last_time"`     // 上次执行时间
	CreateTime  time.Time `json:"create_time" gorm:"create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"update_time"`
}

type Context struct {
	HTTPMethod        string      `json:"http_method"`
	CallMethod        string      `json:"call_method"`
	Path              string      `json:"path"`
	Body              interface{} `json:"body"`
	DataFiled         string      `json:"data_filed"`
	PageCountFiled    string      `json:"page_count_filed"`
	PageFiledPosition struct {
		Type          int    `json:"type"`
		NumPosition   string `json:"num_position"`
		CountPosition string `json:"count_position"`
	} `json:"page_filed_position"`
	RequestInfo struct {
		BaseServer string `json:"base_server"`
	} `json:"request_info"`
}

type InitParam struct {
	FirstPageNum    int64             `json:"first_page_num"`
	PageSize        int64             `json:"page_size"`
	PageCount       int64             `json:"page_count"`
	UniqueFieldName string            `json:"unique_field_name"`
	Md5FiledList    string            `json:"md5_filed_list"`
	SaveFiledList   string            `json:"save_filed_list"`
	FiledAlias      map[string]string `json:"filed_alias"`
}

// TableName 表名称
func (*TaskData) TableName() string {
	return "t_schedule_cfg"
}
