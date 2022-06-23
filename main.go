package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var testjson string = `{"textInput": "Hello World!"}`

	var result map[string]interface{}
	json.Unmarshal([]byte(testjson), &result)

	fmt.Println(result["textInput"])
}
