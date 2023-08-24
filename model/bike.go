package model

type Store struct {
	store_id       int    `json:"storeid",bun:"storeid"`
	store_name     string `json:"storename",bun:"storename"`
	store_location string `json:"storelocation",bun:"storelocation"`
	store_mobile   int    `json:"storemobile", bun:"storemobile"`
	store_email    string `json:"storeemail",bun:"storeemail"`
}

type Bike struct {
	bike_id     int    `json:"bikeid",bun:"bikeid"`
	bike_name   string `json:"bikename",bun:"bikename"`
	mileage     int    `json:"mileage",bun:"mileage"`
	cc          int    `json:"cc",bun:"cc"`
	description string `json:"description",bun:"desscription"`
}

type User struct {
	user_id      int    `json:"userid",bun:"userid"`
	user_name    string `json:"username",bun:"username"`
	user_address string `json:"useraddress",bun:"useraddress"`
	user_mobile  int    `json:"usermobile", bun:"usermobile"`
	user_email   string `json:"useremail",bun:"useremail"`
}

type Payment struct {
	payment_id   int    `json:"paymentid",bun:"paymentid"`
	store_id     int    `json:"storeid",bun:"storeid"`
	user_id      int    `json:"userid",bun:"userid"`
	bike_id      int    `json:"bikeid",bun:"bikeid"`
	amount       int    `json:"amount",bun"amount"`
	payment_type string `json:"paymenttyps",bun:"paymenttype"`
}