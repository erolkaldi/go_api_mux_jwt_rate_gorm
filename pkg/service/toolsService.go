package service

import (
	"encoding/json"
	"fmt"
)

func JsonToString(jsonObject interface{}) string {
	b, _ := json.Marshal(jsonObject)
	return string(b)
}

func StringToJson(jsonString string, obj interface{}) {
	err := json.Unmarshal([]byte(jsonString), &obj)
	if err != nil {
		fmt.Println("Json Convert error:", err.Error())
	}
}
