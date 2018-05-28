package connection

import (
	"gopkg.in/mgo.v2"
)

var Database string = "busReservation"
var connection *mgo.Session

func Dial() {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	connection = session
}

func Disconnect() {
	connection.Close()
}

func GetConnection() *mgo.Session {
	return connection
}
