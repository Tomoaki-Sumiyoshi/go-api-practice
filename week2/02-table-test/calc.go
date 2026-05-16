package calc

import "errors"

func Add(a, b int) int {
	return a + b
}

func IsEven(n int) bool {
	return n%2 == 0
}

func ValidateAge(age int) error {
	if age < 0 {
		return errors.New("Age is less than 0")
	}
	return nil
}
