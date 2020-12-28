package bootstrapper

import (
	"strconv"

	"github.com/shyam-unnithan/go-micro/util"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		util.Logger.Fatal("Config file not found: ", err)
	}

	//Configure NATS authentication information
	util.NatsConfig.User = viper.GetString("nats.User")
	util.NatsConfig.Password = viper.GetString("nats.Password")
	util.NatsConfig.Queue = viper.GetString("nats.Queue")
	util.NatsConfig.Name = viper.GetString("nats.Name")
	util.NatsConfig.URL = viper.GetString("nats.URL")
	util.NatsConfig.WaitTimeInMinutes, err = strconv.Atoi(viper.GetString("nats.WaitTimeInMintues"))
	if err != nil {
		util.NatsConfig.WaitTimeInMinutes = 0
	}
}
