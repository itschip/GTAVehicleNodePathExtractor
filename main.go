package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	rawXml, err := ioutil.ReadFile("paths.xml")
	if err != nil {
		log.Fatal(err.Error())
	}

	// var scene Scene
	scene, err := Unmarshal(rawXml)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	nodesv1 := [][]float64{}

	for _, obj := range scene.Objects.ObjectRoot {
		if obj.Class == "vehiclenode" {
			fmt.Println("Found vehicle node")

			posY, err := strconv.ParseFloat(obj.Transform.Object.Position.Y, 32)
			if err != nil {
				log.Fatal(err.Error())
			}

			posX, err := strconv.ParseFloat(obj.Transform.Object.Position.X, 32)
			if err != nil {
				log.Fatal(err.Error())
			}

			nodesv1 = append(nodesv1, []float64{posY, posX})
		}
	}

	file, err := json.MarshalIndent(nodesv1, "", " ")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = ioutil.WriteFile("paths.json", file, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}
