package bike

import (
	"bikeManagementSystem/model"
	"database/sql"
	"fmt"
	"log"
)

func GetBikesDB(db *sql.DB) []model.Bike{
	var bikeArray []model.Bike
	query:=(`select * from bike`)
	rows,fetchErr:=db.Query(query)
	defer rows.Close()
	if iterationErr:=rows.Err();fetchErr!=nil{
		log.Fatal(iterationErr)
	}
	for rows.Next(){
		var bikeObj model.Bike
		scanErr:=rows.Scan(&bikeObj.BikeId,&bikeObj.BikeName,&bikeObj.Mileage,&bikeObj.Cc,&bikeObj.BikeDesc)
		if scanErr!=nil{
			log.Fatal(scanErr)
		}
		bikeArray=append(bikeArray, bikeObj)
	}
	return bikeArray
}

func GetBikeDB(db *sql.DB, id int) model.Bike{
	var bikeObj model.Bike
	query:=(`select * from bike where bike_id=?`)
	rows,fetchErr:=db.Query(query,id)
	defer rows.Close()
	if iterationErr:=rows.Err();fetchErr!=nil{
		log.Fatal(iterationErr)
	}
	if rows.Next(){
		scanErr:=rows.Scan(&bikeObj.BikeId,&bikeObj.BikeName,&bikeObj.Mileage,&bikeObj.Cc,&bikeObj.BikeDesc)
		if scanErr!=nil{
			log.Fatal(scanErr)
		}
	}
	return bikeObj
}

func UpdateBikeDB(db *sql.DB,bikeObj model.Bike) bool{
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	
	query:=(`UPDATE bike SET bike_id=?,bike_name=?,mileage=?,cc=?,bike_desc=? where bike_id=?`)
	_,InsertErr:=tx.Exec(query,bikeObj.BikeId,bikeObj.BikeName,bikeObj.Mileage,bikeObj.Cc,bikeObj.BikeDesc,bikeObj.BikeId)
	if InsertErr!=nil{
		tx.Rollback()
		log.Fatal(InsertErr)
		fmt.Println("---DB ERROR---")
		return false
	}
	if err:= tx.Commit();err!=nil{
		log.Fatal(err)
		return false
	}
	return true
}