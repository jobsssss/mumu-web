package model

import (
	"fmt"
	"gorm.io/gorm/schema"
	"mumu/pkg/app"
	"mumu/pkg/config"
	"mumu/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {

	var err error

	conf := mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
			config.GetString("database.mysql.username"),
			config.GetString("database.mysql.password"),
			config.GetString("database.mysql.host"),
			config.GetString("database.mysql.port"),
			config.GetString("database.mysql.database"),
			config.GetString("database.mysql.charset")),
	})
	//var logMode gormlogger.LogLevel
	logMode := gormlogger.Error
	if app.IsLocal() {
		logMode = gormlogger.Info
	} else if app.IsProduction() {
		logMode = gormlogger.Warn
	}

	// 准备数据库连接池
	DB, err = gorm.Open(conf, &gorm.Config{
		Logger: gormlogger.Default.LogMode(logMode),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: config.GetString("database.mysql.prefix"),
			SingularTable: true,
		},
	})

	logger.LogIf(err)

	return DB
}
