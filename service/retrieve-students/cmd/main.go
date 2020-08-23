package main

import (
	"flag"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/shyam-unnithan/go-micro/pb"
	_ "github.com/shyam-unnithan/go-micro/service/retrieve-students/pkg/bootstrapper"
	"github.com/shyam-unnithan/go-micro/service/retrieve-students/pkg/datastore"
	"github.com/shyam-unnithan/go-micro/util"
	"log"
	"os"
	"os/signal"
)

const (
	subj = "get-students"
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
	res := pb.GetStudentsResponse{Success: true, Message: ""}

	stream := getStream()
	i := 0

	stream.Conn.QueueSubscribe(subj, util.NatsConfig.Queue, func(msg *nats.Msg) {
		i++
		studentStore := datastore.StudentStore{}
		result, err := studentStore.Get()
		if err != nil {
			res.Success = false
			res.Message = "Failed Unmarshalling"
		}

		if err != nil {
			log.Printf("Failed Marshalling students from database to result")
		}
		res.Students = result
		data, err := proto.Marshal(&res)
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
