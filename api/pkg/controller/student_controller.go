package controller

import (
	"encoding/json"
	"github.com/shyam-unnithan/go-micro/util"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"github.com/shyam-unnithan/go-micro/pb"
)

//StudentController manage student
type StudentController struct {
}

const (
	aggregate = "student"
	subj      = "rest-event"
	timeout   = 10 * time.Second
)

//PostStudent - Handler for POST request on students api
func (handler StudentController) PostStudent(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	var student pb.Student
	eventType := "register-student"

	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &student)
	if err != nil {
		log.Printf("Failed decoding")
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Failed creating student")
	}
	log.Printf("student Object: { id: %s, name: %s, email: %s }", student.Id, student.Name, student.Email)
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Printf("Failed UUID generation for student %s", body)
	}
	student.Id = uuid.String()
	studentData, err := proto.Marshal(&student)

	event, err := createEvent(eventType, studentData)
	if err != nil {
		log.Println("Failed generating event")
	}

	data, err := proto.Marshal(&event)
	if err != nil {
		log.Println("Failed marshalling event")
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Failed creating student")
	}
	defer getStream().Conn.Close()
	msg, err := getStream().Conn.Request(subj, data, timeout)
	if err != nil {
		log.Println("Failed NATS request for event creation")
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Failed creating student")
	}

	res := pb.Response{}
	err = proto.Unmarshal(msg.Data, &res)
	if err != nil {
		log.Println("Failed parsing response from NATS")
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Failed creating student")
	}
	return res, http.StatusCreated, nil
}

//GetStudents - Handler for GET request on students api
func (handler StudentController) GetStudents(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	eventType := "get-students"
	var noData []byte

	event, err := createEvent(eventType, noData)
	if err != nil {
		log.Println("Failed generating event")
	}

	data, err := proto.Marshal(&event)
	if err != nil {
		log.Println("Failed marshalling event")
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Failed getting student")
	}
	defer getStream().Conn.Close()
	msg, err := getStream().Conn.Request(subj, data, timeout)
	if err != nil {
		log.Println("Failed NATS request for event creation")
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Failed getting student")
	}

	res := pb.GetStudentsResponse{}
	err = proto.Unmarshal(msg.Data, &res)

	if err != nil {
		log.Println("Failed parsing response from NATS")
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Failed getting students")
	}

	return res, http.StatusOK, nil
}

//Helper to get EventStream
func getStream() util.Stream {
	config := util.StreamConfig{
		User:     util.NatsConfig.User,
		Password: util.NatsConfig.Password,
		URL:      util.NatsConfig.URL,
		Queue:    util.NatsConfig.Queue,
		Name: "Nats Requester",
	}

	stream, err := util.NewStream(config)
	if err != nil {
		util.Logger.Fatal(err)
	}
	return stream
}

// Helper to create an Event Object
func createEvent(eventType string, data []byte) (pb.Event, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Printf("Failed UUID generation for event %s", data)
		return pb.Event{}, err
	}
	event := pb.Event{
		EventId:       uuid.String(),
		EventType:     eventType,
		AggregateId:   uuid.String(),
		AggregateType: aggregate,
		Data:          string(data),
	}
	return event, nil
}