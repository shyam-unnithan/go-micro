package datastore

import (
	"log"
	"os"

	"github.com/shyam-unnithan/go-micro/pb"
	"github.com/shyam-unnithan/go-micro/util"
)

//StudentStore - Datastore to persist student information
type StudentStore struct{}

//Create - Create a repository record in cockroach database
func (studentStore StudentStore) Get() ([]*pb.Student, error) {
	var students []*pb.Student
	sqlStatement := `SELECT id, name, email FROM students`

	rows, err := getStore().Db.Query(sqlStatement)
	if err != nil {
		log.Printf("Failed querying students: %s", err)
	}
	defer rows.Close()

	for rows.Next(){
		var id string
		var name string
		var email string
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			log.Printf("Failed gettings students: %s", err)
		}
		student := &pb.Student{ Id: id, Name: name, Email: email}
		students = append(students, student)
	}

	return students, nil
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
