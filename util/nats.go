package util

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

//StreamConfig - Structure for NATS config info
type StreamConfig struct {
	User, Password, URI, Queue, Name string
	WaitInMinutes                    int
}

//Stream - Structure to hold NATS config and conn
type Stream struct {
	cfg  StreamConfig
	Conn *nats.Conn
}

//NewStream is to create a NATS Stream
func NewStream(cfg StreamConfig) (Stream, error) {
	stream := Stream{}
	var nc *nats.Conn
	var err error

	// Connect Options.
	opts := []nats.Option{nats.Name(cfg.Name)}
	if cfg.WaitInMinutes != 0 {
		opts = appendWaitOpts(cfg, opts)
	}

	// Provide Authentication information
	opts = append(opts, nats.UserInfo(cfg.User, cfg.Password))

	//Connect to NATS
	if cfg.URI != "" {
		nc, err = nats.Connect(cfg.URI, opts...)
	} else {
		nc, err = nats.Connect(nats.DefaultURL, opts...)
	}
	if err != nil {
		log.Println(err)
	}

	stream.Conn = nc
	return stream, err
}

func appendWaitOpts(cfg StreamConfig, opts []nats.Option) []nats.Option {
	totalWait := time.Duration(cfg.WaitInMinutes) * time.Minute
	reconnectDelay := time.Second
	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to: %s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Printf("Exiting: %v", nc.LastError())
	}))
	return opts
}
