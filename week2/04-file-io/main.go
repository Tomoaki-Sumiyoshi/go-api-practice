package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func ReadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data := make([]byte, 1024)
	count, err := file.Read(data)
	if err != nil {
		return "", err
	}

	return string(data[:count]), nil
}

func WriteFile(filePath string, text string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := []byte(text)
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	jsonFilePath := "./week2/04-file-io/users.json"

	prevJson, err := ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var users []User
	err = DecodeJson(prevJson, &users)
	if err != nil {
		fmt.Println(err)
		return
	}

	length := len(users)
	users = append(users, User{
		ID:    length,
		Name:  "User" + strconv.Itoa(length),
		Email: "user" + strconv.Itoa(length) + "@example.com",
	})

	currJson, err := EncodeJson(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = WriteFile(jsonFilePath, currJson)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Complete")
}
