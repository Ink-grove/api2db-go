package initialize

import (
	"api2db-server/config"
	"api2db-server/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/niuniumart/gosdk/middleware/cors"
	"github.com/niuniumart/gosdk/middleware/logprint"
	recoverSdk "github.com/niuniumart/gosdk/middleware/recover"
	"github.com/niuniumart/gosdk/middleware/utils"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"net/http"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	engine := gin.Default()
	engine.Use(recoverSdk.PanicRecover())
	engine.Use(logprint.InfoLog())
	engine.Use(cors.Cors())
	engine.GET(utils.UrlMetrics, gin.WrapH(promhttp.Handler()))
	engine.GET(utils.UrlHeartBeat, heartBeat)

	return &Router{
		engine,
	}
}

// RegisterRouter 注册路由
func (r *Router) RegisterRouter() {
	v1 := r.Engine.Group("/v1")
	{
		// 创建任务接口，前面是路径，后面是执行的函数，跳进去
		v1.POST("/create_task", service.CreateTask)
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

}

func (r *Router) RunByPort() error {
	cfg := config.GetGlobalConfig()
	port := fmt.Sprintf("%d", cfg.SvrConfig.Port)

	var runPort string
	if port[0] == ':' {
		runPort = port
	} else {
		runPort = fmt.Sprintf(":%s", port)
	}

	return r.Engine.Run(runPort)
}

func heartBeat(c *gin.Context) {
	c.String(http.StatusOK, "SUCCESS")
}
