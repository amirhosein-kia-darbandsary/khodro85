package main

import (
	"github.com/amirhosein-kia-darbandsary/khodro85/api"
	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/amirhosein-kia-darbandsary/khodro85/data/cache"
	"github.com/amirhosein-kia-darbandsary/khodro85/data/database"
)

func main() {
	config := config.GetConfig()
	cache.InitRedis(&config)
	database.InitPostgres(&config)
	api.InitServer(&config)
	// config := config.GetConfig()
	// fmt.Println(config)
}
