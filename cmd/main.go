package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/YJinHai/MyIm/config"
	"github.com/YJinHai/MyIm/internal/app/router"
)

var (
	cfg = pflag.StringP("config", "c", "", "video config file path.")
)

func main() {
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	config.Init("im")
	r := router.Load()

	r.Run("0.0.0.0:4990")
}
