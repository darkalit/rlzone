package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/internal/items"
)

func NewMySqlDB(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPass,
		c.DBHost,
		c.DBPort,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", c.DBName))
	db = db.Exec(fmt.Sprintf("USE %s", c.DBName))

	rawDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	rawDB.SetMaxOpenConns(c.DBMaxOpenConns)
	rawDB.SetConnMaxLifetime(time.Duration(c.DBConnMaxLifetime) * time.Second)
	rawDB.SetMaxIdleConns(c.DBMaxIdleConn)
	rawDB.SetConnMaxIdleTime(time.Duration(c.DBConnMaxIdleTime) * time.Second)

	err = rawDB.Ping()
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&items.Item{},
		&items.Stock{},
	)

	return db, nil
}
