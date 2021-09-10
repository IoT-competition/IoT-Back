package basic

import (
	"github.com/IoTCompetition/IoBack/app/entity"
	"log"

	"github.com/IoTCompetition/IoBack/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
NewDB 获得数据库连接实例
*/
func NewDB(c *config.Type) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.GetDBConfig().GetURL()), &gorm.Config{})
	if err != nil {
		log.Panicf("[Init] 数据库连接失败: %v", err.Error())
	}
	err = AutoMigrate(db)
	if err != nil {
		log.Panicf("[AutoMigrate] 数据库同步失败: %v", err.Error())
	}

	return db
}

// AutoMigrate 自动映射数据表
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Info{}, &entity.Option{})
}
