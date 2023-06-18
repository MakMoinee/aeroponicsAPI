package config

import (
	"aeroponicsAPI/internal/aeroponicsapi/common"
	"log"

	"github.com/spf13/viper"
)

var Registry *viper.Viper

func Set() {
	log.Println("Loading Configurations ...")
	var err error
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	Registry = viper.GetViper()
	common.SERVER_PORT = Registry.GetString("SERVER_PORT")
}
