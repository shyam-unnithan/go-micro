package util

//Configuration - configuration for the database
type Configuration struct {
	Server   string // WebServer Host
	LogLevel int    // Log Level: 0 - 4
	// Config for DataBase
	DBHost, DBPort, DBUser, DBPassword, Database string
}

//NatsConfiguration - Configuration for NATS server
type NatsConfiguration struct {
	User, Password, URI, Queue, Name string
	WaitTimeInMinutes                int
}

// AppConfig holds the configurations used for web app
var AppConfig Configuration

// NatsConfig holds configuration information for connecting to Nats Server
var NatsConfig NatsConfiguration

func init() {
	AppConfig = Configuration{}
	NatsConfig = NatsConfiguration{}
}
