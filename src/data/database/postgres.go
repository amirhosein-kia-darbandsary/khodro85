package database

import (
	"fmt"
	"log"
	"time"

	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgressConnection *gorm.DB

var cfg = config.GetConfig()
var logger = logging.NewLogger(&cfg)

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
		logger.Fatal(logging.Postgres, logging.ExternalService, err.Error(), nil)
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
		logger.Fatal(logging.Postgres, logging.ExternalService, err.Error(), nil)

	}
	sqlDB.Close()
}
