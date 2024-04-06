package main

import (
	"log"
	"newestcdd/app/db"
	"newestcdd/app/server"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
)

func main() {
	var validate *validator.Validate
	dbs,err:=db.NewMysqlServer(mysql.Config{
		User:                 "root",
		Passwd:               "r23password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "glg_restful",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewApiServer(":3000",dbs,validate)
	if err:=server.Run();err!=nil{
		log.Fatal(err)
	}
}
