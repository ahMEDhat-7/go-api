package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct{
	addr string
	db *sql.DB
}



func NewApiServer(addr string , db *sql.DB) *APIServer{
	return &APIServer{
		addr: addr,
		db: db,
	}
}


func (apiServer *APIServer ) Run() error{
	
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/vi").Subrouter()
	
	log.Println("Listening On",apiServer.addr)
	return http.ListenAndServe(apiServer.addr,router)
}