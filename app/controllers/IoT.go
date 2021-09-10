package controllers

import (
	"github.com/IoTCompetition/IoBack/app/models"
	"github.com/IoTCompetition/IoBack/app/schema"
	"github.com/IoTCompetition/IoBack/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var IoTControllerSet = wire.NewSet(wire.Struct(new(IoTController), "*"))

type IoTController struct {
	Logger *zap.Logger
	Util   *utils.Ginx
	IoT    *models.IoTModel
}

func (i *IoTController) SetInfo(c *gin.Context) {
	var data schema.InfoRequest
	if err := i.Util.ParseJSON(c, &data); err != nil {
		i.Util.FailWarp(c, utils.ErrBodyCanNotParser)
		return
	}
	err := i.IoT.UpdateInfo(data)
	if err != nil {
		i.Util.FailWarp(c, utils.ErrTypeMismatch)
		return
	}
	i.Util.SuccessWarp(c, nil)
}

func (i *IoTController) GetInfo(c *gin.Context) {
	res, err := i.IoT.GetInfo()
	if err != nil {
		i.Util.FailWarp(c, utils.ErrTypeMismatch)
		return
	}
	i.Util.SuccessWarp(c, res)
}
