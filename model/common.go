package model

import (
	"gopkg.in/mgo.v2"
	"github.com/aliromei/re-project/connection"
)

func connect(collection string) *mgo.Collection {
  return connection.GetConnection().DB(connection.Database).C(collection)
}
