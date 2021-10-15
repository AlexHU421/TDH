package util

import (
	"encoding/json"
	"fmt"
)

func JsonUnmarshalByString(str string )map[string]interface{} {
	mMap := make (map[string]interface{})
	err := json.Unmarshal([]byte(str),&mMap)
	if err !=nil {
		fmt.Println(err)
	}
	return mMap
}

