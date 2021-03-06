# JParse

### What is it?
JParse is a dynamic JSON file decoder/parser that removes the need for custom structs to parse JSON data. This idea was inspired by how simple JSON data can be parsed in python, dynamically searching through the JSON for the desired input and returning a splice of string values. Note you must first convert your JSON object to a string in order for it to work, see below example for refrence. 

## Installation
``` bash
go get github.com/nishgowda/JParse 
```
## Usage

An example use of this would be:

```go
package main
import(
    "json/encoding"
    "fmt"
    jparse "github.com/nishgowda/JParse"
)
func main(){
    jsonFile, err := os.Open("users.json")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened users.json")
    defer jsonFile.Close()
    body, err := ioutil.ReadAll(jsonFile)
    js := string(body)
    value := []string{"id", "name", "department"}
    embeddedValue := []string{"city", "state"}
    embeddedObj := []string{"address"}
    a := jparse.EmbeddedObjArrayParse(value, js, embeddedObj, embeddedValue)
    fmt.Println(a)
}
```

### Functions 
Currently, there exists 4 functions to decode JSON data, each with their own simplicities.

```go
jparse.SimpleParse(value []string, j string)
```
The simplest case for decoding a JSON file is a scenario with no array of values, no embedded objects, nothing. Just a plain and simple JSON object. Taking in an array of values to search for and a json string, the function will return a string slice of the parsed data.

```go
jparse.SimpleArrayParse(value []string, j string)
```
The second case for decoding JSON data is an aray of simple JSON objects. A bit more complicated, but the same idea applies. 

```go
jparse.EmbeddedObjParse(value []string, j string, embeddedObj []string, embeddedValue []string)
```
The third case is designed for dealing with an embedded object in a simple JSON object. Here, the user must pass in an array of embedded terms for the function to search for and the subsequent values of that embedded object on top of the previous parameters. 

```go
EmbeddedObjArrayParse(value []string, j string, embeddedObj []string, embeddedValue []string)
```
The final case is an array of JSON data with embedded objects. The same parameters must be passed in as the previous function. 

### To Do:
- [ ] Run more tests
- [ ] Add more error cases
- [ ] Deploy
  
