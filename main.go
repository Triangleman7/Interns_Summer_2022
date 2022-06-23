package main

import (
	"encoding/json"
	"fmt"
)

const testjson string = `{
	"testString": "Hello World!",
	"testInt": 42,
	"testFloat": 3.14,
	"testBoolean": true,
	"testArray": [1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144]
}`

func ProcessInput(jsondata string) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsondata), &result)

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func main() {
	data := ProcessInput(testjson)

	fmt.Println(data)
}
