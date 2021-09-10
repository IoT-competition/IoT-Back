package daos

import (
	"github.com/IoTCompetition/IoBack/app/entity"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var OptionDaoSet = wire.NewSet(wire.Struct(new(OptionDao), "*"))

type OptionDao struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

type OptionUpdate struct {
	LED      int
	FanN     int
	FanM     int
	Engine   int
	Smoke    int
	Infrared int
	Card     int
}

func (o *OptionDao) UpdateOptions(params OptionUpdate) error {
	db := entity.GetOption(o.DB)

	tmp := entity.Option{
		Model: gorm.Model{
			ID: 1,
		},
		LED:      params.LED,
		FanM:     params.FanM,
		FanN:     params.FanN,
		Engine:   params.Engine,
		Smoke:    params.Smoke,
		Infrared: params.Infrared,
		Card:     params.Card,
	}
	db = db.Where(map[string]interface{}{"id": 1})

	err := db.Updates(tmp).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *OptionDao) GetOption() (*[]entity.Option, error) {
	db := entity.GetOption(o.DB)
	result := new([]entity.Option)

	err := db.Find(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
