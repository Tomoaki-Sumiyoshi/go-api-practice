package main

import (
	"errors"
	"fmt"
)

func ValidateAge(age int) error {
	if age < 0 {
		return errors.New("年齢が0未満です。")
	}
	return nil
}

func main() {
	ageSlice := []int{-1, 2, 0, 5}

	for _, age := range ageSlice {
		err := ValidateAge(age)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(age)
	}
}
