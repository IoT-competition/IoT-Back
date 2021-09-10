package controllers

import (
	"github.com/IoTCompetition/IoBack/app/models"
	"github.com/IoTCompetition/IoBack/app/schema"
	"github.com/IoTCompetition/IoBack/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var OptionControllerSet = wire.NewSet(wire.Struct(new(OptionController), "*"))

type OptionController struct {
	Logger *zap.Logger
	Util   *utils.Ginx
	Option *models.OptionModel
}

func (co *OptionController) SetOption(c *gin.Context) {
	var data schema.OptionsRequest
	if err := co.Util.ParseJSON(c, &data); err != nil {
		co.Util.FailWarp(c, utils.ErrBodyCanNotParser)
		return
	}
	err := co.Option.SetOption(data)
	if err != nil {
		co.Util.FailWarp(c, err)
		return
	}
	co.Util.SuccessWarp(c, nil)
}

func (co *OptionController) GetOption(c *gin.Context) {
	res, err := co.Option.GetOption()
	if err != nil {
		co.Util.FailWarp(c, err)
		return
	}
	co.Util.SuccessWarp(c, res)
}
