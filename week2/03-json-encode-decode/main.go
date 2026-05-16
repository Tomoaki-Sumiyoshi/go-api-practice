package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func EncodeJson(T any) (string, error) {
	b, err := json.Marshal(T)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func DecodeJson(jsonData string, T any) error {
	err := json.Unmarshal([]byte(jsonData), T)

	return err
}

func main() {
	json, encodeError := EncodeJson(User{
		ID:    0,
		Name:  "Taro Tanaka",
		Email: "t-tanaka@example.com",
	})

	if encodeError != nil {
		fmt.Println(encodeError)
		return
	} else {
		fmt.Println(json)
	}

	var user User
	decodeError := DecodeJson(`{"id":1,"name":"Hanako Yamada","email":"h-yamada@example.com"}`, &user)

	if decodeError != nil {
		fmt.Println(decodeError)
	} else {
		fmt.Println(user)
	}
}
