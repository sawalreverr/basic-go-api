package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBConn     string `mapstructure:"DB_CONN"`
}

func Load() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	var cfg Config
	err := viper.Unmarshal(&cfg)
	return &cfg, err
}
