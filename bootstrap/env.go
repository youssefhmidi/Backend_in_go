package bootstrap

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	AccessTokenSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpiry  int    `mapstructure:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry int    `mapstructure:"REFRESH_TOKEN_EXPIRY"`
	ContextTimeout     int    `mapstructure:"CONTEXT_TIMEOUT"`
	ReleaseMode        bool   `mapstructure:"RELEASEMODE"`
}

func NewEnv(dst string) Env {
	var env Env

	viper.SetConfigFile(dst)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		fmt.Println(err)
	}

	return env
}
