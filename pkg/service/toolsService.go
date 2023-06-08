package service

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed)
}

func ValidatePassword(password string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false
	}
	return true
}
