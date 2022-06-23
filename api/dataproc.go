package api

import (
	"encoding/json"
	"fmt"
)

func ProcessInput(jsondata string) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsondata), &result)

	if err != nil {
		panic(err)
	}

	return result
}
