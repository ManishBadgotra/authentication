package helpers

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	ConfigFilePath = "config.json"
)

var cnf *viper.Viper

// Load config file from given path
func InitConfig() {
	v := viper.New()
	v.SetConfigType("json")
	v.SetConfigFile(ConfigFilePath)
	// v.AddConfigPath(ConfigDirPath)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found", err)
		}
		fmt.Println("config file not found", err)
		fmt.Println("config file not found", err)
	}

	cnf = v
}

func GetString(key string) string {
	return cnf.GetString(key)
}
