package mysql

import (
	"fmt"

	"github.com/mrrizal/todo-api-clean-arch/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	db *gorm.DB
}

func NewMysqlDB(config config.Config) (*MysqlDB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.MYSQL_USERNAME,
		config.MYSQL_PASSWORD, config.MYSQL_HOST, config.MYSQL_PORT, config.MYSQL_DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&ActivityGroups{}, &Todo{})
	return &MysqlDB{db: db}, nil
}
