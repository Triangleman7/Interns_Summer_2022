package main

import (
	"encoding/json"
	"fmt"
)

func ProcessInput(jsondata string) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(jsondata), &result)

	return result
}

func main() {
	var testjson string = `{"textInput": "Hello World!"}`
	data := ProcessInput(testjson)

	fmt.Println(data["textInput"])
}
