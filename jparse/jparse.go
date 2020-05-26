package jparse

import (
	"encoding/json"
	"fmt"
)

// parsing function that takes in json and array of values to parse
func SimpleParse(value []string, j string) []string {
	var parsed_json []string
	var result map[string]interface{}
	json.Unmarshal([]byte(j), &result)
	for i := 0; i < len(value); i++ {
		str := fmt.Sprint(result[value[i]])
		parsed_json = append(parsed_json, (str))
	}
	return parsed_json
}

func SimpleArrayParse(value []string, j string) []string {
	var parsed_json []string
	var results []map[string]interface{}
	json.Unmarshal([]byte(j), &results)
	for _, result := range results {
		for i := 0; i < len(value); i++ {
			str := fmt.Sprint(result[value[i]])
			parsed_json = append(parsed_json, str)
		}
	}
	return parsed_json
}

func EmbeddedObjParse(value []string, j string, embeddedObj []string, embeddedValue []string) []string {
	var parsed_json []string
	var result map[string]interface{}
	json.Unmarshal([]byte(j), &result)

	for _, emb := range embeddedObj {
		embedded := result[emb].(map[string]interface{})
		for i := 0; i < len(value); i++ {
			str := fmt.Sprint(result[value[i]])
			parsed_json = append(parsed_json, str)
		}
		for i := 0; i < len(embeddedValue); i++ {
			str := fmt.Sprint(embedded[embeddedValue[i]])
			parsed_json = append(parsed_json, str)
		}
	}

	return parsed_json
}

func EmbeddedObjArrayParse(value []string, j string, embeddedObj []string, embeddedValue []string) []string {
	var parsed_json []string
	var results []map[string]interface{}
	json.Unmarshal([]byte(j), &results)

	for _, result := range results {
		for _, emb := range embeddedObj {
			embedded := result[emb].(map[string]interface{})
			for i := 0; i < len(value); i++ {
				str := fmt.Sprint(result[value[i]])
				parsed_json = append(parsed_json, str)
			}
			for i := 0; i < len(embeddedValue); i++ {
				str := fmt.Sprint(embedded[embeddedValue[i]])
				parsed_json = append(parsed_json, str)
			}
		}
	}
	return parsed_json
}

/*
func main(){
	//jsonString()
	json_file := `[
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
	value := []string{"id", "name","department"}
	emmbeddedValue := []string{"city", "state"}
	embeddedObj := []string{"address"}
	a := embeddedObjArrayParse(value, json_file, embeddedObj, emmbeddedValue)
	//a := arrayParse(value, json_file)
	fmt.Println(a)
}
*/
