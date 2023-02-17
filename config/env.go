package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetEnv(name, fileType, path string, stru interface{}) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(fileType)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err = viper.Unmarshal(&stru)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file struct map: %w \n", err))
	}
}
