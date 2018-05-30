package seed

import (
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"
	"github.com/aliromei/re-project/connection"
)

type Province struct {
	Id					int				`json:"id,number" bson:"id"`
	Name				string		`json:"name" bson:"name"`
	Latitude		float32		`json:"latitude,number" bson:"latitude"`
	Longitude		float32		`json:"longitude,number" bson:"longitude"`
	Cities			[]City		`json:"cities,array" bson:"cities"`
}

type City struct {
	Id					int				`json:"id,number" bson:"id"`
	Name				string		`json:"name" bson:"name"`
	Latitude		float32		`json:"latitude,number" bson:"latitude"`
	Longitude		float32		`json:"longitude,number" bson:"longitude"`
}

func Run() {
	var provinces []Province

	provinceC := connection.GetConnection().DB(connection.Database).C("provinces")

	file, err := ioutil.ReadFile("seed/provinces.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(file, &provinces)

	for _, province := range provinces {
		fmt.Println(&province)
		err = provinceC.Insert(&province)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Provinces Seed Completed")
}