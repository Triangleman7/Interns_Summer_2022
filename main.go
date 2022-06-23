package main

import (
	"encoding/json"
	"fmt"
)

func ProcessInput(jsondata string) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsondata), &result)

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func main() {
	var testjson string = `{"textInput": "Hello World!"}`
	data := ProcessInput(testjson)

	fmt.Println(data["textInput"])
}
