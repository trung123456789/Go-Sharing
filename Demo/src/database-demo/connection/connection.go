package connection

import (
	"config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"structdemo"
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

	db, err := sql.Open("postgres", psqlInfo)
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
