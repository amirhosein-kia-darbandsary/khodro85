// @title           Khodro85 API
// @version         1.0
// @description     REST API for Khodro85.
// @host            localhost:9090
// @BasePath        /api/v1
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
