package models

import (
	"github.com/IoTCompetition/IoBack/app/daos"
	"github.com/IoTCompetition/IoBack/app/schema"
	"github.com/IoTCompetition/IoBack/app/utils"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var OptionsModelSet = wire.NewSet(wire.Struct(new(OptionModel), "*"))

type OptionModel struct {
	Logger *zap.Logger
	Option *daos.OptionDao
}

func (o *OptionModel) SetOption(params schema.OptionsRequest) error {
	data := daos.OptionUpdate{
		LED:      params.LED,
		FanN:     params.FanN,
		FanM:     params.FanM,
		Engine:   params.Engine,
		Smoke:    params.Smoke,
		Infrared: params.Infrared,
		Card:     params.Card,
	}

	err := o.Option.UpdateOptions(data)
	if err != nil {
		return utils.ErrSetOptionFailed
	}
	return nil
}

func (o *OptionModel) GetOption() (*schema.OptionsResponse, error) {
	res, err := o.Option.GetOption()
	if err != nil {
		return nil, utils.ErrGetOptionFailed
	}
	return &schema.OptionsResponse{
		LED:      (*res)[0].LED,
		FanN:     (*res)[0].FanN,
		FanM:     (*res)[0].FanM,
		Engine:   (*res)[0].Engine,
		Smoke:    (*res)[0].Smoke,
		Infrared: (*res)[0].Infrared,
		Card:     (*res)[0].Card,
	}, nil
}
