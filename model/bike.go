package model

type Store struct {
	StoreId       int    `json:"storeid" bun:"storeid"`
	StoreName     string `json:"storename" bun:"storename"`
	StoreLocation string `json:"storelocation" bun:"storelocation"`
	StoreMobile   int    `json:"storemobile" bun:"storemobile"`
	StoreEmail    string `json:"storeemail" bun:"storeemail"`
}

type Bike struct {
	BikeId   int    `json:"bikeid" bun:"bikeid"`
	BikeName string `json:"bikename" bun:"bikename"`
	Mileage  int    `json:"mileage" bun:"mileage"`
	Cc       int    `json:"cc"`
	BikeDesc string `json:"description" bun:"desscription"`
}

type User struct {
	UserID      int    `json:"userid" bun:"userid"`
	UserName    string `json:"username" bun:"username"`
	UserAddress string `json:"useraddress" bun:"useraddress"`
	UserMobile  int    `json:"usermobile"  bun:"usermobile"`
	UserEmail   string `json:"useremail" bun:"useremail"`
}

type Payment struct {
	PaymentId   int    `json:"paymentid" bun:"paymentid"`
	StoreId     int    `json:"storeid" bun:"storeid"`
	UserId      int    `json:"userid" bun:"userid"`
	BikeId      int    `json:"bikeid" bun:"bikeid"`
	Amount      int    `json:"amount"`
	PaymentType string `json:"paymenttyps" bun:"paymenttype"`
}
