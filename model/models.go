package model

import "time"

const (
	BUS_STATUS_NOT_READY = 0
	BUS_STATUS_READY = 1
	BUS_STATUS_GONE = 2
)

type Passenger struct {
	PassengerId			string				`json:"passengerId,string" bson:"passengerId"`
	Seat						int						`json:"seat,number" bson:"seat"`
}

type Bus struct {
	Id							string				`json:"_id,string" bson:"_id"`
	Model						string				`json:"model,number" bson:"model"`
	Seats						int						`json:"seats,number" bson:"seats"`
	OpenSeatCount		int						`json:"openSeats,number" bson:"openSeats"`
	AvailableSeats	[]int					`json:"availableSeats" bson:"availableSeats"`
	Passengers			[]Passenger		`json:"passengers" bson:"passenger"`
	Owner						string				`json:"owner" bson:"owner"`
	Status					int						`json:"status,number" bson:"status"`
	CreatedAt				time.Time			`json:"createdAt,string" bson:"createdAt"`
	UpdatedAt				time.Time			`json:"updatedAt,string" bson:"updatedAt"`
}

type User struct {
	Id							string					`json:"_id,string" bson:"_id"`
	Name						string					`json:"name,string" bson:"name"`
	Email						string					`json:"email,string" bson:"email"`
	Password				string					`json:"password,string" bson:"password"`
	Reservations	 	[]Reservation		`json:"reservations" bson:"reservations"`
	Token						string					`json:"token,string" bson:"token"`
	IsAdmin					bool						`json:"isAdmin,boolean" bson:"isAdmin"`
	CreatedAt				time.Time				`json:"createdAt,string" bson:"createdAt"`
	UpdatedAt				time.Time				`json:"updatedAt,string" bson:"updatedAt"`
}

type Reservation struct {
	BusId 					string					`json:"busId,string" bson:"busId"`
	Origin					int							`json:"origin,number" `
	Destination			int							`json:"destination,number" bson:"destination"`
	Seat						int							`json:"seat,number" bson:"seat"`
	Status					int							`json:"status,number" bson:"status"`
}