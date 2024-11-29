// Package configs ...
package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

//Configs ... for Viper
type Configs struct {
	LogLevel string `mapstructure:"LOG_LEVEL"`
	//WordTemplate []byte `mapstructure:"WORD_TEMPLATE"`
	//ReplacedWordTemplate []byte `mapstructure:"REPLACED_WORD_TEMPLATE"`
	TimesheetTemplate string `mapstructure:"TIMESHEET_TEMPLATE"`
	ReplacedTimesheetTemplate string `mapstructure:"REPLACED_TIMESHEET_TEMPLATE"`
	ResetPassword string `mapstructure:"RESET_PASSWORD"`
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