package repositories

import (
	"bobo/config"
	"bobo/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Orm *gorm.DB

func init() {

	config := config.Get()
	dsn := config.MySql.Username + ":" + config.MySql.Password + "@tcp(" + config.MySql.Host + ":" + config.MySql.Port + ")/" + config.MySql.Database + "?charset=utf8&loc=Asia%2FShanghai&parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(config.MySql.LogLevel)),
	})
	if err != nil {
		log.Println(err)
		panic(err)
	}
	sql, _ := db.DB()
	sql.SetMaxOpenConns(config.MySql.MaxOpenCon)
	sql.SetMaxIdleConns(config.MySql.MaxIdleCon)
	sql.SetConnMaxIdleTime(time.Second * time.Duration(config.MySql.MaxConLifeTime))
	Orm = db

	// 关联,模型
	Orm.AutoMigrate(&models.Tomaha{})

	// /// 释放
	// defer sql.Close()
}
