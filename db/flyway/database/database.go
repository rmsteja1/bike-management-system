package database

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectToDatabase() *sql.DB{
	db,dbError:=sql.Open("mysql","rmsteja:@Rmsteja1.@tcp(localhost:3306)/bike_management_system")
	if dbError!=nil{
		log.Fatal(dbError)
	}
	pingError:=db.Ping()

	if pingError!=nil{
		log.Fatal(pingError)
	}
	fmt.Println("Successfully connected to Database")
	return db

}