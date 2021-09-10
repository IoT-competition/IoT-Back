package routers

import (
	"github.com/IoTCompetition/IoBack/app/controllers"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// RouterSet 注入router
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine)
	Prefixes() []string
}

// Router 路由管理器
type Router struct {
	InfoAPI   *controllers.IoTController
	OptionAPI *controllers.OptionController
}

// Register 注册路由
func (a *Router) Register(app *gin.Engine) {
	a.RegisterAPI(app)
}

// Prefixes 路由前缀列表
func (a *Router) Prefixes() []string {
	return []string{
		"/api",
	}
}
