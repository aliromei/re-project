package seed

import (
	"fmt"
	"os"
	"encoding/json"
)

type province struct {
	id					int							`json:"id,number" bson:"id"`
	name				string					`json:"name,string" bson:"name"`
	latitude		float32					`json:"latitude,number" bson:"latitude"`
	longitude		float32					`json:"longitude,number" bson:"longitude"`
	cities			[]city					`json:"cities" bson:"cities"`
}

type city struct {
	id					int							`json:"id,number" bson:"id"`
	name				string					`json:"name,string" bson:"name"`
	latitude		float32					`json:"latitude,number" bson:"latitude"`
	longitude		float32					`json:"longitude,number" bson:"longitude"`
}

func Run() {
	defer fmt.Println("Provinces Seed Completed")

	file, err := os.Open("seed/provinces.json")
	if err != nil {
		fmt.Printf("Provinces Seed Error: %v\n", err)
		os.Exit(1)
	}

	decoder := json.NewDecoder(file)

	var p []province

	decoder.Decode(&p)

	//provinceC := connection.GetConnection().DB(connection.Database).C("provinces")

	fmt.Println(p)

	//for _, province := range p {
	//	fmt.Println(province.id)
	//	err = provinceC.Insert(&province)
	//	if err != nil {
	//		fmt.Println("Provinces Seed Error: Inserting to Collection")
	//	}
	//}
}