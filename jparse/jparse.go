package jparse

import (
	"encoding/json"
	"fmt"
)

// SimpleParse is a parsing function that takes in a simple json and array of values to parse
func SimpleParse(value []string, j string) []string {
	var parsedJSON []string
	var result map[string]interface{}
	json.Unmarshal([]byte(j), &result)
	for i := 0; i < len(value); i++ {
		str := fmt.Sprint(result[value[i]])
		parsedJSON = append(parsedJSON, (str))
	}
	return parsedJSON
}

// SimpleArrayParse is a parsing function that takes in a array of json and array of values to parse
func SimpleArrayParse(value []string, j string) []string {
	var parsedJSON []string
	var results []map[string]interface{}
	json.Unmarshal([]byte(j), &results)
	for _, result := range results {
		for i := 0; i < len(value); i++ {
			str := fmt.Sprint(result[value[i]])
			parsedJSON = append(parsedJSON, str)
		}
	}
	return parsedJSON
}

// EmbeddedObjParse is a parsing function that takes in a simple json with embedded values and array of values to parse
func EmbeddedObjParse(value []string, j string, embeddedObj []string, embeddedValue []string) []string {
	var parsedJSON []string
	var result map[string]interface{}
	json.Unmarshal([]byte(j), &result)

	for _, emb := range embeddedObj {
		embedded := result[emb].(map[string]interface{})
		for i := 0; i < len(value); i++ {
			str := fmt.Sprint(result[value[i]])
			parsedJSON = append(parsedJSON, str)
		}
		for i := 0; i < len(embeddedValue); i++ {
			str := fmt.Sprint(embedded[embeddedValue[i]])
			parsedJSON = append(parsedJSON, str)
		}
	}

	return parsedJSON
}

// EmbeddedObjArrayParse is a parsing function that takes in a json array with embedded values and array of values to parse
func EmbeddedObjArrayParse(value []string, j string, embeddedObj []string, embeddedValue []string) []string {
	var parsedJSON []string
	var results []map[string]interface{}
	json.Unmarshal([]byte(j), &results)

	for _, result := range results {
		for _, emb := range embeddedObj {
			embedded := result[emb].(map[string]interface{})
			for i := 0; i < len(value); i++ {
				str := fmt.Sprint(result[value[i]])
				parsedJSON = append(parsedJSON, str)
			}
			for i := 0; i < len(embeddedValue); i++ {
				str := fmt.Sprint(embedded[embeddedValue[i]])
				parsedJSON = append(parsedJSON, str)
			}
		}
	}
	return parsedJSON
}
