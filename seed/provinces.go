package seed

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"github.com/aliromei/re-project/connection"
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

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Provinces Seed Error: Couldn't Find Working Directory")
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(dir + "/seed/provinces.json")
	if err != nil {
		fmt.Printf("Provinces Seed Error: %v\n", err)
		os.Exit(1)
	}

	//fmt.Println(string(file))

	var p []province

	json.Unmarshal(file, &p)

	provinceC := connection.GetConnection().DB(connection.Database).C("provinces")

	fmt.Println(p)

	for _, province := range p {
		fmt.Println(province.id)
		err = provinceC.Insert(&province)
		if err != nil {
			fmt.Println("Provinces Seed Error: Inserting to Collection")
		}
	}
}