package bike

import (
	"bikeManagementSystem/model"
	"database/sql"
	"fmt"
	"log"
)

func GetUsersDb(db *sql.DB, users []model.User) []model.User {
	fmt.Println("Test")
	rows,err:=db.Query(`SELECT * from users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	for rows.Next(){
		var userObj model.User
		err:=rows.Scan(&userObj.UserID,&userObj.UserName,&userObj.UserAddress,&userObj.UserMobile,&userObj.UserEmail)
		if err!=nil{
			log.Fatal(err)
		}else{
			users=append(users, userObj)
		}
	}

	if iterattionError:= rows.Err();err!=nil{
		log.Fatal(iterattionError)
	}

	return users
}

func AddUserDb(db *sql.DB, user model.User){
	query:=`INSERT INTO users (user_id,user_name,user_address,user_mobile,user_email) values (?,?,?,?,?)`
	row,err:=db.Exec(query,user.UserID,user.UserName,user.UserAddress,user.UserMobile,user.UserEmail)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Print(row)
}

func GetUserDB(db *sql.DB,id string) model.User{
	query:=(`select * from users where user_id=?`)
	row,readErr:=db.Query(query,id)
	var resultedUser model.User
	defer row.Close()
	if readErr!=nil{
		log.Fatal(readErr)
	}
	
	if row.Next(){
		scanErr:=row.Scan(&resultedUser.UserID,&resultedUser.UserName,&resultedUser.UserAddress,&resultedUser.UserMobile,&resultedUser.UserEmail)
		if scanErr!=nil{
			log.Fatal(scanErr)
		}
	}
	return resultedUser
}

func DelUserDB(db *sql.DB, email string){
	query:=(`delete from users where user_email=?`)
	_,deleteErr:=db.Exec(query,email)
	if deleteErr!=nil{
		log.Fatal(deleteErr)
	}
}

func IsExsistingUserDB(db *sql.DB,email string)  bool{
	query:=(`select user_email from users where user_email=?`)
	row,err:=db.Query(query,email)
	if err!=nil{
		log.Fatal(err)
	}
	if row.Next(){
		return true
	}
	return false
}