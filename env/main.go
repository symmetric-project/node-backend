package env

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/symmetric-project/node-backend/utils"
)

type Config struct {
	MODE         string `mapstructure:"MODE"`
	JWT_SECRET   string `mapstructure:"JWT_SECRET"`
	DATABASE_URL string `mapstructure:"DATABASE_URL"`

	DOMAIN      string
	DOMAIN_DEV  string `mapstructure:"DOMAIN_DEV"`
	DOMAIN_PROD string `mapstructure:"DOMAIN_PROD"`
}

var CONFIG Config

func init() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		utils.Stacktrace(err)
	}
	if err := viper.Unmarshal(&CONFIG); err != nil {
		utils.Stacktrace(err)
	}

	if CONFIG.MODE == "dev" {
		CONFIG.DOMAIN = CONFIG.DOMAIN_DEV
		gin.SetMode("debug")

	} else {
		CONFIG.DOMAIN = CONFIG.DOMAIN_PROD
		gin.SetMode("release")

	}
}
