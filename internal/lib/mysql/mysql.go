package mysql

import (
	"game/internal/logger"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMysql creates a new MySQL
func NewMysql(dsn string) *gorm.DB {
	// 异常处理
	defer func() {
		if err := recover(); err != nil {
			logger.Debugf("%+v", err)
			os.Exit(1)
		}
	}()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("failed to connect database:%+v", err)
		os.Exit(1)
	}

	return db

}
