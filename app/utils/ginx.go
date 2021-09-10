package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	"net/http"
)

// GinxSet Ginx 的 DI 注入
var GinxSet = wire.NewSet(wire.Struct(new(Ginx), "*"))

// Ginx 是 gin 框架相关的辅助方法
type Ginx struct {
	Logger *zap.Logger
}

// ParseJSON 格式化 json
func (g *Ginx) ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		g.Logger.Warn("json 解析失败", zap.Error(err))
		return err
	}

	return nil
}

// SuccessWarp 处理成功包装
func (g *Ginx) SuccessWarp(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"err_code": "0000",
		"err_msg":  "处理成功",
		"data":     data,
	})

	c.Abort()
}

// FailWarp 处理失败包装
func (g *Ginx) FailWarp(c *gin.Context, err error) {
	defaultErr := returnError{
		code: "0001",
		msg:  "未知错误",
	}
	if v, ok := errorMap[err]; ok {
		defaultErr = v
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": defaultErr.code,
		"err_msg":  defaultErr.msg,
		"data":     nil,
	})

	c.Abort()
}
