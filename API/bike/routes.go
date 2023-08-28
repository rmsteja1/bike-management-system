package bike

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	//routes for user
	r.HandleFunc("/getAllUsers",getUsers).Methods("GET")
	r.HandleFunc("/getOneUser/{id}",getUser).Methods("GET")
	r.HandleFunc("/addOneUser",addUser).Methods("POST")
	r.HandleFunc("/deleteOneUser/{email}",deleteUser).Methods("DELETE")

	//routers for bikes
	r.HandleFunc("/getAllBikes",getBikes).Methods("GET")
	r.HandleFunc("/getOneBike/{id}",getBike).Methods("GET")
	r.HandleFunc("/updateOneBike",updateBike).Methods("PUT")

	fmt.Println("Starting server at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}