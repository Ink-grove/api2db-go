package handle

// Run 执行函数
func Run(handler HandlerIntf) error {
	err := handler.HandleInput()
	if err != nil {
		return err
	}
	err = handler.HandleProcess()
	return err
}

// HandlerIntf handler接口
type HandlerIntf interface {
	HandleInput() error
	HandleProcess() error
}
