package main

import (
	"fmt"
)

const testjson string = `{
	"testString": "Hello World!",
	"testInt": 42,
	"testFloat": 3.14,
	"testBoolean": true,
	"testArray": [1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144]
}`

func main() {
	data := api.ProcessInput(testjson)

	fmt.Println(data)
}
