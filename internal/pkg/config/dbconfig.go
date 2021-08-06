package config

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(conf *AppConfig) (*gorm.DB, error) {
	dbUser := conf.GetString("db.user")
	dbPassword := conf.GetString("db.Password")
	dbPort := conf.GetString("db.port")
	dbHost := conf.GetString("db.host")
	dbName := conf.GetString("db.name")

	format := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local"

	connectionString := fmt.Sprintf(format, dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	return db, nil
}
