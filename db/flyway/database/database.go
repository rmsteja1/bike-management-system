package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() *sql.DB{
	db,dbError:=sql.Open("mysql","root:@Rmsteja1.@tcp(localhost:3306)/bike_management_system")
	if dbError!=nil{
		log.Fatal(dbError)
	}
	pingError:=db.Ping()

	if pingError!=nil{
		log.Fatal(pingError)
	}
	//fmt.Println("Successfully connected to Database")
	return db

}