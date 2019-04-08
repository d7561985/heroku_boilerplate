package config

import (
	"github.com/spf13/viper"
)

// V represent config values after binding from env
var V = &M{}

func init() {
	viper.AutomaticEnv()
	viper.RegisterAlias("HEROKUSLUG", "HEROKU_APP_NAME")
	viper.SetDefault("PORT", "3000")

	if err := read(); err != nil {
		panic(err)
	}
}

func read() error {
	if err := viper.Unmarshal(V); err != nil {
		return err
	}
	return nil
}
