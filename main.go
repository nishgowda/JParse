package main

import (
	"fmt"

	"github.com/nishgowda/JParse/jparse"
)

func main() {
	jsonFile := `[
		{
			"id": 1,
			"name": "Mr. Boss",
			"department": "",
			"designation": "Director",
			"address": {
				"city": "Mumbai",
				"state": "Maharashtra",
				"country": "India"
			}
		},
		{
			"id": 11,
			"name": "Irshad",
			"department": "IT",
			"designation": "Product Manager",
			"address": {
				"city": "Mumbai",
				"state": "Maharashtra",
				"country": "India"
			}
		},
		{
			"id": 12,
			"name": "Pankaj",
			"department": "IT",
			"designation": "Team Lead",
			"address": {
				"city": "Pune",
				"state": "Maharashtra",
				"country": "India"
			}
		}
	]`
	value := []string{"id", "name", "department"}
	emmbeddedValue := []string{"city", "state"}
	embeddedObj := []string{"address"}
	a := jparse.EmbeddedObjArrayParse(value, jsonFile, embeddedObj, emmbeddedValue)
	fmt.Println(a)
}
