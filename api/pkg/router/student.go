package router

import (
	"github.com/shyam-unnithan/eduwiz/api/pkg/controller"

	"github.com/gorilla/mux"
)

//StudentRouter - Router for student API
type StudentRouter struct{}

//AddRoutes - Add student routes to router
func (c StudentRouter) AddRoutes(router *mux.Router) *mux.Router {
	StudentController := controller.StudentController{}
	router.Handle("/api/students", controller.ResponseHandler(StudentController.PostStudent)).Methods("POST")
	router.Handle("/api/students", controller.ResponseHandler(StudentController.GetStudents)).Methods("GET")
	return router
}
