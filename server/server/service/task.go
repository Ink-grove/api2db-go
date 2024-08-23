package service

import (
	"api2db-server/constant"
	"api2db-server/log"
	"api2db-server/middleware/handle"
	"api2db-server/middleware/rpc"
	"api2db-server/pkg/pb"
	"api2db-server/repository"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTaskHandler 创建任务
type CreateTaskHandler struct {
	Req  repository.CreateTaskReq
	Resp repository.CreateTaskResp
}

// CreateTask 接口
func CreateTask(c *gin.Context) {
	var hd CreateTaskHandler
	defer func() {
		hd.Resp.Msg = constant.GetErrMsg(hd.Resp.Code)
		c.JSON(http.StatusOK, hd.Resp)
	}()

	// 解析请求包
	err := c.ShouldBind(&hd.Req)
	if err != nil {
		log.Errorf("CreateTask shouldBind err %s", err.Error())
		hd.Resp.Code = constant.ERR_SHOULD_BIND
		return
	}

	handle.Run(&hd)
}

// HandleInput 参数检查
func (p *CreateTaskHandler) HandleInput() error {
	if p.Req.TaskData.TaskType == "" {
		log.Errorf("input invalid")
		p.Resp.Code = constant.ERR_INPUT_INVALID
		return constant.ERR_HANDLE_INPUT
	}
	//if p.Req.TaskData.Priority != nil {
	//	if *p.Req.TaskData.Priority > db.MAX_PRIORITY || *p.Req.TaskData.Priority < 0 {
	//		p.Resp.Code = constant.ERR_INPUT_INVALID
	//		martlog.Errorf("input invalid")
	//		return constant.ERR_HANDLE_INPUT
	//	}
	//}
	return nil
}

// HandleProcess 处理函数
func (p *CreateTaskHandler) HandleProcess() error {
	log.Infof("into HandleProcess")
	err := repository.DbCreatTask(p.Req.TaskData)
	if err != nil {
		log.Errorf("DbCreatTask err", err.Error())
		p.Resp.Code = constant.ERR_INPUT_INVALID
		return err
	}

	// 创建了之后拆分任务给client
	rpc.GetSvrClient().AssignTasks(context.Background(), &pb.TaskData{
		Id:       p.Req.TaskData.ID,
		TaskType: p.Req.TaskData.TaskType,
		InitParam: &pb.InitParam{
			FirstPageNum:    p.Req.TaskData.InitParam.FirstPageNum,
			PageSize:        p.Req.TaskData.InitParam.PageSize,
			PageCount:       p.Req.TaskData.InitParam.PageCount,
			UniqueFieldName: p.Req.TaskData.InitParam.UniqueFieldName,
			Md5FiledList:    p.Req.TaskData.InitParam.Md5FiledList,
			SaveFiledList:   p.Req.TaskData.InitParam.SaveFiledList,
			FiledAlias:      p.Req.TaskData.InitParam.FiledAlias,
		},
		Param:  "",
		Status: p.Req.TaskData.Status,
	})
	return nil
}
