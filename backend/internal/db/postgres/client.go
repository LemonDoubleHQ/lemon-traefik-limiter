package postgres

import (
	"fmt"
	"lemon-traefik-limiter-backend/internal/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres(){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", config.GlobalConfig.Postgres.Host, config.GlobalConfig.Postgres.UserName, config.GlobalConfig.Postgres.Password, config.GlobalConfig.Postgres.DBName)

	db, err := gorm.Open(postgres.Open(dsn))
	
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	DB = db
}

func RunDbMigration(){
	err := DB.AutoMigrate(&ApiKey{}, &APIKeyPath{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}