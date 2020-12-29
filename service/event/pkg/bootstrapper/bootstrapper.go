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
		util.Logger.Fatal("config file not found: ", err)
	}

	// Configure Postgres configuration values
	util.AppConfig.DBHost = viper.GetString("cockroach.Host")
	util.AppConfig.DBPort = viper.GetString("cockroach.Port")
	util.AppConfig.DBUser = viper.GetString("cockroach.User")
	util.AppConfig.DBPassword = viper.GetString("cockroach.Password")
	util.AppConfig.Database = viper.GetString("cockroach.Database")

	//Configure NATS authentication information
	util.NatsConfig.User = viper.GetString("nats.User")
	util.NatsConfig.Password = viper.GetString("nats.Password")
	util.NatsConfig.Queue = viper.GetString("nats.Queue")
	util.NatsConfig.Name = viper.GetString("nats.Name")
	util.NatsConfig.URI = viper.GetString("nats.URI")
	util.NatsConfig.WaitTimeInMinutes, err = strconv.Atoi(viper.GetString("nats.WaitTimeInMintues"))
	if err != nil {
		util.NatsConfig.WaitTimeInMinutes = 0
	}

}
