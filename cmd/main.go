package main

import (
	"bikeManagementSystem/API/bike"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := mux.NewRouter()
	// Serve Swagger documentation at /swagger/{anything}
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Your API endpoints...
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello, Swagger!"}`))
	})

	
	// Connect to the database
	bike.Routes(r)
	// Run the server
	http.ListenAndServe(":8080", r)
}