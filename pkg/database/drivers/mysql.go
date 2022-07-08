package drivers

import (
	"Walker/pkg/contract"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Mysql struct {
}

func (m *Mysql) MysqlConnector(config contract.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetString("db.username"),
		config.GetString("db.password"),
		config.GetString("db.host"),
		config.GetInt("db.port"),
		config.GetString("db.database"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	if err != nil {
		panic(err.Error())
	}
	return db
}
