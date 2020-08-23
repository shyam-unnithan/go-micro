package datastore

import (
	"github.com/pkg/errors"
	"github.com/shyam-unnithan/go-micro/pb"
	"github.com/shyam-unnithan/go-micro/util"
	"log"
)

//studentStore - Datastore to persist student information
type EventStore struct{}

//Create - Create a repository record in cockroach database
func (eventStore EventStore) Create(event pb.Event) (pb.Event, error) {
	sqlStatement := `INSERT INTO events (id, type, aggregate, aggregate_type, data) 
						VALUES ($2, $2, $3, $4, $5)
						RETURNING id;`
	var id string
	err := getStore().Db.QueryRow(sqlStatement, event.EventId, event.EventType, event.AggregateId, event.AggregateType, event.Data).Scan(&id)
	if err != nil {
		log.Printf("Failed inserting data: %s", err)
		return event, errors.Wrap(err, "Error while inserting on students")
	}
	return event, nil
}

func getStore() util.DataStore {
	config := util.Config{
		Host:     util.AppConfig.DBHost,
		Port:     util.AppConfig.DBPort,
		User:     util.AppConfig.DBUser,
		Password: util.AppConfig.DBPassword,
		Database: util.AppConfig.Database,
	}

	store, err := util.New(config)
	if err != nil {
		util.Logger.Fatal(err)
	}
	return store
}
