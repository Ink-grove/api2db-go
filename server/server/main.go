package main

import (
	"api2db-server/config"
	"api2db-server/initialize"
	"api2db-server/log"
	"api2db-server/middleware/db"
	"api2db-server/middleware/rpc"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	if err := config.Init(); err != nil {
		log.Fatalf("init config failed, err:%v\n", err)
	}
	log.InitLog()
	rpc.InitSvrConn()
	log.Info("log init success...")
}

func Run() error {
	router := initialize.NewRouter()
	router.RegisterRouter()
	err := router.RunByPort()
	if err != nil {
		return err
	}
	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	return nil
}

func main() {
	Init()
	defer log.Sync()
	defer db.CloseDB()
	if err := Run(); err != nil {
		log.Errorf("Svr run err:%v", err)
	}
}
