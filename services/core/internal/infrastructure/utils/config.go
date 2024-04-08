package utils

import (
	"github.com/spf13/viper"
)

// It uses mapstructure under the hood
type Config struct {
	DBSource  string `mapstructure:"DB_SOURCE"`
	JwtSecret string `mapstructure:"JWT_SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app") // looking for app.env
	viper.SetConfigType("env") // json, xml . . . .

	// viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		/* ----------------------- PRIORITIZE THE ENV FROM OS ----------------------- */
		viper.BindEnv("DB_SOURCE")
		viper.BindEnv("JWT_SECRET")
		viper.AutomaticEnv()
		/* -------------------------------------------------------------------------- */

		// return // no need to return, send it to be unmarshalled
	}

	err = viper.Unmarshal(&config)
	return
}
