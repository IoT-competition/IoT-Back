package entity

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	LED      int `gorm:"column:led;type=int;not null"`
	FanM     int `gorm:"column:fan_m;type=int;not null"`
	FanN     int `gorm:"column:fan_n;type=int;not null"`
	Engine   int `gorm:"column:engine;type=int;not null"`
	Smoke    int `gorm:"column:smoke;type=int;not null"`
	Infrared int `gorm:"column:infrared;type=int;not null"`
	Card     int `gorm:"column:card;type=int;not null"`
}

func (Option) TableName() string {
	return "option"
}

func GetOption(db *gorm.DB) *gorm.DB {
	return db.Model(new(Option))
}
