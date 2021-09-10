package entity

import "gorm.io/gorm"

type Info struct {
	gorm.Model
	Illumination string `gorm:"column:illumination;type=varchar(255);not null"`
	Temperature  string `gorm:"column:temperature;type=varchar(255);not null"`
	Humidity     string `gorm:"column:humidity;type=varchar(255);not null"`
}

func (Info) TableName() string {
	return "info"
}

func GetInfo(db *gorm.DB) *gorm.DB {
	return db.Model(new(Info))
}
