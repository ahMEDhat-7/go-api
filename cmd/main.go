package main

import (
	"log"

	"github.com/ahMEDhat-7/go-api/cmd/api"
)



func main(){

	server := api.NewApiServer(":8000",nil)
	if err := server.Run() ; err != nil{
		log.Fatal(err)
	}
}