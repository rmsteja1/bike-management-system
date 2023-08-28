package bike

import (
	//"database/sql"
	"bikeManagementSystem/db/flyway/database"
	"bikeManagementSystem/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

var dbConnection *sql.DB=database.ConnectToDatabase()

func getUsers(w http.ResponseWriter, r *http.Request){
	var allUsers []model.User
	w.Header().Set("content-type","application/json")
	dbConnection :=database.ConnectToDatabase()
	if dbConnection !=nil{
		allUsers= GetUsersDb(dbConnection,allUsers)
		if len(allUsers)>0{
			json.NewEncoder(w).Encode(allUsers)
		}else{
			json.NewEncoder(w).Encode("At present there are no users")
		}
		fmt.Println("User count: ",len(allUsers))
	}
}

func addUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-type","application/json")
	var newUser model.User
	json.NewDecoder(r.Body).Decode(&newUser)
	if isValidEmail(newUser.UserEmail){
		if !isExsistingUser(newUser.UserEmail){
			dbConnection:=database.ConnectToDatabase()
			AddUserDb(dbConnection,newUser)
			json.NewEncoder(w).Encode("Successfully added user.")
		}else{
			json.NewEncoder(w).Encode("User with the give email already exsist")
		}

	}else{
		json.NewEncoder(w).Encode("User Email is invalid")
	}	
}

func isExsistingUser(userEmail string) bool{
	dbConnection:=database.ConnectToDatabase()
	return IsExsistingUserDB(dbConnection,userEmail)
}

func isValidEmail(email string) bool {
	// Regular expression pattern for a simple email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Use the regular expression to match the email
	return re.MatchString(email)
}

func getUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-type","application/json")
	dbConnection :=database.ConnectToDatabase()
	params := mux.Vars(r)
	if dbConnection !=nil{
		user:=GetUserDB(dbConnection,params["id"])
		if isObjEmpty(user)==true{
			json.NewEncoder(w).Encode("No user with the given user id.")
		}else{
			json.NewEncoder(w).Encode(user)
		}
	}
}

func isObjEmpty(user model.User) bool{
	if user.UserID==0{
		return true
	}
	return false
}

func deleteUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-type","application/json")
	params:=mux.Vars(r)
	userEmail:=params["email"]

	if userEmail!=""{
		if !isExsistingUser(userEmail){
			json.NewEncoder(w).Encode("User with the given email does not exsist")
		}else{
			DelUserDB(dbConnection,userEmail)
			json.NewEncoder(w).Encode("User successfully deleted")
		}
		
	}
}

func getBikes(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-type","application/json")
	resultedBikes:=GetBikesDB(dbConnection)
	fmt.Print()
	if len(resultedBikes)==0{
		json.NewEncoder(w).Encode("There a re no bikes to show")
	}else{
		json.NewEncoder(w).Encode(resultedBikes)
	}
}

func getBike(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-type","application/json")
	var resultedBike model.Bike
	params:=mux.Vars(r)
	bikeId,_:=strconv.Atoi(params["id"])
	resultedBike=GetBikeDB(dbConnection,bikeId)
	if resultedBike.BikeId==0{
		json.NewEncoder(w).Encode("There is no bike with this id")
	}else{
		json.NewEncoder(w).Encode(resultedBike)
	}
}

func updateBike(w http.ResponseWriter,r *http.Request){
	var bikeObj model.Bike
	json.NewDecoder(r.Body).Decode(&bikeObj)
	bikeId:=bikeObj.BikeId
	resultedBike:=GetBikeDB(dbConnection,bikeId)
	if resultedBike.BikeId==0{
		json.NewEncoder(w).Encode("No bike wi the given id")
	}else{
		isUpdated:=UpdateBikeDB(dbConnection,bikeObj)
		if isUpdated{
			json.NewEncoder(w).Encode("Successfully updated bike details.")
		}
	}
}
