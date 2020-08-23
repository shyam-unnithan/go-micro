package datastore

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/shyam-unnithan/go-micro/pb"
	"github.com/shyam-unnithan/go-micro/util"
)

//StudentStore - Datastore to persist student information
type StudentStore struct{}

//Create - Create a repository record in cockroach database
func (studentStore StudentStore) Create(student pb.Student) (pb.Student, error) {
	sqlStatement := `INSERT INTO students (id, name, email) 
						VALUES ($2, $2, $3)
						RETURNING id;`
	var id string
	err := getStore().Db.QueryRow(sqlStatement, student.Id, student.Name, student.Email).Scan(&id)
	if err != nil {
		log.Printf("Failed inserting data: %s", err)
		return student, errors.Wrap(err, "Error while inserting on students")
	}
	return student, nil
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
		os.Exit(1)
	}
	return store
}
