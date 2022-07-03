package models

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pincode" bson:"pincode"`
}

type User struct {
	Name    string  `json:"name" bson:"name"`
	Age     int     `json:"age" bson:"age"`
	Address Address `json:"address" bson:"address"`
}

// type User struct {
// 	Id     bson.ObjectId `json:"id" bson:"_id"`
// 	Name   string        `json:"name" bson:"name"`
// 	Gender string        `json:"gender" bson:"gender"`
// 	Age    int           `json:"age" bson:"age"`
// }
