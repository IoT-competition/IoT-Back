package daos

import (
	"github.com/IoTCompetition/IoBack/app/entity"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var IoTDaoSet = wire.NewSet(wire.Struct(new(IoTDao), "*"))

type IoTDao struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

type UpdateInfoParams struct {
	Illumination string
	Temperature  string
	Humidity     string
}

func (i *IoTDao) UpdateInfo(params UpdateInfoParams) error {
	db := entity.GetInfo(i.DB)

	tmp := entity.Info{
		Illumination: params.Illumination,
		Temperature:  params.Temperature,
		Humidity:     params.Humidity,
	}

	if err := db.Create(&tmp).Error; err != nil {
		return err
	}
	return nil
}

func (i *IoTDao) GetInfo() (*[]entity.Info, error) {
	db := entity.GetInfo(i.DB)

	res := new([]entity.Info)
	err := db.Order("created_at desc").Limit(10).Find(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
