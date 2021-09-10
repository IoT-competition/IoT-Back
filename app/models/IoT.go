package models

import (
	"github.com/IoTCompetition/IoBack/app/daos"
	"github.com/IoTCompetition/IoBack/app/schema"
	"github.com/IoTCompetition/IoBack/app/utils"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

var IoTModelSet = wire.NewSet(wire.Struct(new(IoTModel), "*"))

type IoTModel struct {
	Logger *zap.Logger
	IoT    *daos.IoTDao
}

func (i *IoTModel) UpdateInfo(params schema.InfoRequest) error {
	data := daos.UpdateInfoParams{
		Illumination: params.Illumination,
		Temperature:  params.Temperature,
		Humidity:     params.Humidity,
	}
	err := i.IoT.UpdateInfo(data)
	if err != nil {
		return utils.ErrInfoUpdateError
	}
	return nil
}

func (i *IoTModel) GetInfo() (*[]schema.InfoResponse, error) {
	res, err := i.IoT.GetInfo()
	if err != nil {
		return nil, utils.ErrInfoGetNull
	}
	result := new([]schema.InfoResponse)
	err = copier.Copy(result, res)
	if err != nil {
		return nil, utils.ErrInfoGetNull
	}
	return result, nil
}
