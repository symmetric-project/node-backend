package env

import (
	"github.com/spf13/viper"
	"github.com/symmetric-project/node-backend/utils"
)

type Config struct {
	MODE       string `mapstructure:"MODE"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`

	DOMAIN      string
	DOMAIN_DEV  string `mapstructure:"DOMAIN_DEV"`
	DOMAIN_PROD string `mapstructure:"DOMAIN_PROD"`
}

var CONFIG Config

func init() {
	viper.AutomaticEnv()
	viper.BindEnv("PORT")
	viper.SetDefault("PORT", 4444)

	if err := viper.ReadInConfig(); err != nil {
		utils.Stacktrace(err)
	}
	if err := viper.Unmarshal(&CONFIG); err != nil {
		utils.Stacktrace(err)
	}

	if CONFIG.MODE == "dev" {
		CONFIG.DOMAIN = CONFIG.DOMAIN_DEV
	} else {
		CONFIG.DOMAIN = CONFIG.DOMAIN_PROD
	}
}
