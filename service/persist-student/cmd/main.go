package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/shyam-unnithan/go-micro/pb"
	_ "github.com/shyam-unnithan/go-micro/service/persist-student/pkg/bootstrapper"
	"github.com/shyam-unnithan/go-micro/service/persist-student/pkg/datastore"
	"github.com/shyam-unnithan/go-micro/util"
)

const (
	subj = "register-student"
)

func printMsg(m *nats.Msg, i int) {
	student := pb.Student{}
	err := proto.Unmarshal(m.Data, &student)
	if err != nil {
		log.Println("Failed unmarshalling message data")
	}
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, student.String())
}

func main() {

	var showTime = flag.Bool("t", false, "Display timestamps")
	res := pb.Response{Success: true, Message: ""}

	stream := getStream()
	i := 0

	stream.Conn.QueueSubscribe(subj, util.NatsConfig.Queue, func(msg *nats.Msg) {
		i++
		student := pb.Student{}
		err := proto.Unmarshal(msg.Data, &student)
		if err != nil {
			res.Success = false
			res.Message = "Failed Unmarshalling"
		}
		data, err := proto.Marshal(&res)
		printMsg(msg, i)
		studentStore := datastore.StudentStore{}
		studentStore.Create(student)

		msg.Respond(data)
	})
	stream.Conn.Flush()

	if err := stream.Conn.LastError(); err != nil {
		log.Println(err)
	}

	log.Printf("Listening on [%s]", subj)
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
		User:          util.NatsConfig.User,
		Password:      util.NatsConfig.Password,
		URI:           util.NatsConfig.URI,
		Queue:         util.NatsConfig.Queue,
		Name:          util.NatsConfig.Name,
		WaitInMinutes: util.NatsConfig.WaitTimeInMinutes,
	}

	stream, err := util.NewStream(config)
	if err != nil {
		util.Logger.Fatal(err)
	}
	return stream
}
