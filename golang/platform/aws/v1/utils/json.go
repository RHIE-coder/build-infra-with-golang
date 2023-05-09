package utils

import "encoding/json"

func StringifyJSON(jsonData interface{}) string {
	b, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
