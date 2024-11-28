// Package configs ...
package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

//Configs ... for Viper
type Configs struct {
	// Port string `mapstructure:"PORT"`
	// Db string `mapstructure:"DB"`
	// TestDb string `mapstructure:"TEST_DB"`
	LogLevel string `mapstructure:"LOG_LEVEL"`
	WordTemplate string `mapstructure:"WORD_TEMPLATE"`
	ReplacedWordTemplate string `mapstructure:"REPLACED_WORD_TEMPLATE"`
	// JwtSecret string `mapstructure:"childokeford"`
	// User string `mapstructure:"drbaz"`
	// Password string `mapstructure:"Hydrogen1"`
}

//LoadEnvViper ...
func LoadEnvViper() (config Configs)  {
	viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".") // optionally look for config in the working directory
	viper.SetConfigFile("app.env")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
	panic(fmt.Errorf("fatal error config file: %s", err))
}
	err = viper.Unmarshal(&config)
	if err != nil {
    panic(err)
}
	return config
}