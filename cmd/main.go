package main

import (
	"log"

	"github.com/ahMEDhat-7/go-api/cmd/api"
	"github.com/ahMEDhat-7/go-api/config"
	"github.com/ahMEDhat-7/go-api/db"
	"github.com/go-sql-driver/mysql"
)



func main(){
	
	db,err := db.NewMySqlStorage(mysql.Config{
		User:					config.Envs.DBUser,
		Passwd:					config.Envs.DBPassword,
		Addr:					config.Envs.DBAddress,
		DBName:					config.Envs.DBName,
		Net:					"tcp",
		AllowNativePasswords: 	true,
		ParseTime: 				true,
	})


	if err != nil {
		log.Fatal(err)
	}

	server := api.NewApiServer(":8000",db)
	if err := server.Run() ; err != nil{
		log.Fatal(err)
	}
}