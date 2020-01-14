package connection

import (
	"config"
	"constants"
	"database/sql"
	"fmt"
	"structdemo"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// Config type
type Config structdemo.Config

// CreateConnection function
func CreateConnection() (*sql.DB, error) {
	// var cfg Config
	cfg := config.GetEnv()
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Database.DbHost, cfg.Database.DbPort,
		cfg.Database.DbUser, cfg.Database.DbPass, cfg.Database.DbName)
	db, err := sql.Open(constants.DriverName, psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")
	return db, nil
}

// GormConnection function
func GormConnection() (*gorm.DB, error) {
	// var cfg Config
	cfg := config.GetEnv()
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Database.DbHost, cfg.Database.DbPort,
		cfg.Database.DbUser, cfg.Database.DbPass, cfg.Database.DbName)
	db, err := gorm.Open(constants.DriverName, psqlInfo)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")
	return db, nil
}
