package app

import (
	"github.com/IoTCompetition/IoBack/app/utils"
	"net/http"
	"time"

	"github.com/IoTCompetition/IoBack/app/config"
	"github.com/IoTCompetition/IoBack/app/routers"
	"github.com/gin-gonic/gin"
)

/*
NewApp 获取Gin的应用实例
*/
func NewApp(c *config.Type, router routers.IRouter) (server *http.Server) {
	// 设置运行模式
	if c.GetAppConfig().GetMode() != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()

	// 处理 panic
	app.Use(gin.Recovery())
	app.Use(utils.Cors())

	router.Register(app)

	server = &http.Server{
		Addr:         c.GetAppConfig().GetPort(),
		Handler:      app,
		WriteTimeout: time.Duration(c.GetAppConfig().GetTimeout()) * time.Second,
	}

	return
}
