package router

import (
	"github.com/gorilla/mux"
)

//InitRoutes ...
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = StudentRouter{}.AddRoutes(router)
	return router
}
