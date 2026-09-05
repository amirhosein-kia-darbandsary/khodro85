package database

import (
	"fmt"
	"log"
	"time"

	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgressConnection *gorm.DB

func InitPostgres(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	log.Println("PostgreSQL connected successfully")
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Second * 180)

}

func GetPostgresConnection() *gorm.DB {
	return PostgressConnection
}

func ClosePostgresConnection(connetion *gorm.DB) {
	sqlDB, err := connetion.DB()
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)

	}
	sqlDB.Close()
}
