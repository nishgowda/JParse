package main

import (
	"fmt"
	"io/ioutil"
	"os"

	jparse "github.com/nishgowda/JParse"
)

func main() {
	jsonFile, err := os.Open("test-3.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Json file")
	defer jsonFile.Close()

	body, err := ioutil.ReadAll(jsonFile)
	js := string(body)
	value := []string{"id", "name", "department"}
	embeddedValue := []string{"city", "state"}
	embeddedObj := []string{"address"}
	a := jparse.EmbeddedObjArrayParse(value, js, embeddedObj, embeddedValue)
	fmt.Println(a)
}
