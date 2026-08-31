package main

import (
	"fmt"

	"github.com/amirhosein-kia-darbandsary/khodro85/config"
)

func main() {
	// api.InitServer()
	config := config.GetConfig()
	fmt.Println(config)
}
