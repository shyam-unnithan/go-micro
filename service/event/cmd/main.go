package main

import (
	"flag"
	"github.com/shyam-unnithan/eduwiz/util"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/shyam-unnithan/eduwiz/service/event/pkg/datastore"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/shyam-unnithan/eduwiz/pb"
	_ "github.com/shyam-unnithan/eduwiz/service/event/pkg/bootstrapper"
)

const (
	event   = "rest-event"
	subj    = "register-student"
	timeout = 10 * time.Second
)

func printMsg(m *nats.Msg, i int) {
	event := pb.Event{}
	err := proto.Unmarshal(m.Data, &event)
	if err != nil {
		log.Println("Failed unmarshalling message data")
	}
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(event.Data))
}

func main() {

	var showTime = flag.Bool("t", false, "Display timestamps")
	res := pb.Response{Success: true, Message: ""}

	stream := getStream()
	i := 0
	stream.Conn.QueueSubscribe(event, util.NatsConfig.Queue, func(msg *nats.Msg) {
		event := pb.Event{}
		err := proto.Unmarshal(msg.Data, &event)
		if err != nil {
			res.Success = false
			res.Message = "Failed unmarshal Event"
		}

		printMsg(msg, i)
		eventStore := datastore.EventStore{}
		_, err = eventStore.Create(event)
		if err != nil {
			res.Success = false
			res.Message = "Failed persisting event"
		}

		responseMsg, err := stream.Conn.Request(event.EventType, []byte(event.Data), timeout)

		if err != nil {
			res.Success = false
			res.Message = "Failed NATS request for student creation"
		}

		data, err := proto.Marshal(&res)
		if err != nil {
			log.Println("If we reached here, we most probably are in our graves, and hence cannot respond to this request")
		}

		if !res.Success {
			msg.Respond(data)
		} else {
			msg.Respond(responseMsg.Data)
		}
	})
	stream.Conn.Flush()

	if err := stream.Conn.LastError(); err != nil {
		log.Println(err)
	}

	log.Printf("Listening on [%s]", event)
	if *showTime {
		log.SetFlags(log.LstdFlags)
	}

	// Setup the interrupt handler to drain so we don't miss
	// requests when scaling down.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Printf("Draining...")
	stream.Conn.Drain()
	log.Printf("Exiting")
}

//Function to get EventStream
func getStream() util.Stream {
	config := util.StreamConfig{
		User:     util.NatsConfig.User,
		Password: util.NatsConfig.Password,
		URL:      util.NatsConfig.URL,
		Queue:    util.NatsConfig.Queue,
		Name:	util.NatsConfig.Name,
		WaitInMinutes: util.NatsConfig.WaitTimeInMinutes,
	}

	stream, err := util.NewStream(config)
	if err != nil {
		util.Logger.Fatal(err)
	}
	return stream
}
