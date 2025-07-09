package godrinth

import (
	"encoding/json"
)

func IndentedJson(data interface{}) string {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(val)
}
