package seed

import (
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"
	"github.com/aliromei/re-project/connection"
)

type Province struct {
	Id					int							`bson:"id"`
	Name				string					`bson:"name"`
	Latitude		float32					`bson:"latitude"`
	Longitude		float32					`bson:"longitude"`
	Cities			[]City					`bson:"cities"`
}

type City struct {
	Id					int							`bson:"id"`
	Name				string					`bson:"name"`
	Latitude		float32					`bson:"latitude"`
	Longitude		float32					`bson:"longitude"`
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
			fmt.Println("Provinces Seed Error: Inserting to Collection")
		}
	}

	connection.Disconnect()

	fmt.Println("Provinces Seed Completed")
}