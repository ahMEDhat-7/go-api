package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySqlStorage(cfg mysql.Config) (*sql.DB ,error)  {
	db,err := sql.Open("mysql",cfg.FormatDSN())
	err = db.Ping()
	if err != nil{
		log.Fatal(err)
		return nil,err
	}
	log.Println("[+] DB Connection is Established")
	return db,nil
}
