package config

import (
	"github.com/spf13/viper"
)

// V represent config values after binding from env
var V = &M{}

func init() {
	viper.AutomaticEnv()
	viper.RegisterAlias("HEROKUSLUG", "HEROKU_APP_NAME")
	viper.RegisterAlias("DATABASEHOST", "DATABASE_URL")
	viper.SetDefault("PORT", "3000")

	if err := Load(V); err != nil {
		panic(err)
	}
}

// Load read config for provided model.
func Load(model interface{}) error {
	if err := viper.Unmarshal(model); err != nil {
		return err
	}
	return nil
}
