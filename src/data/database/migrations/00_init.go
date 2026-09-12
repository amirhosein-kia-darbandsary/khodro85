package migrations

import (
	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/amirhosein-kia-darbandsary/khodro85/data/database"
	"github.com/amirhosein-kia-darbandsary/khodro85/data/models"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/logging"
)

var cfg = config.GetConfig()
var logger = logging.NewLogger(&cfg)

func Up00() {
	database := database.GetPostgresConnection()

	tables := []interface{}{}

	var country = models.Country{}
	var city = models.City{}

	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "Migration has processed", nil)
}

func Down00() {

}
