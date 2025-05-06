package initializers

import (
	"time"

	"github.com/spf13/viper"
)


type Config struct {
	DBPath       string        `mapstructure:"DB_PATH" default:"./mydb.sqlite"`
	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`
	ClientOrigin string        `mapstructure:"CLIENT_ORIGIN"`
}



func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}