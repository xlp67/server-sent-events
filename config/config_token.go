package config

import (
	"github.com/spf13/viper"
)

type Token struct {
	TokenOpenAI string `mapstructure:"TOKEN_OPENAI"`
}

func LoadConfig() string {
	var token *Token
	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {panic(err)}
	viper.Unmarshal(&token)
	return token.TokenOpenAI
}