package main

import (
	"fmt"

	"github.com/t-sumiyoshi/go-api-practice/week2/01-package-split/calculator"
)

type CalculatorPair struct {
	a int
	b int
}

func main() {
	pairList := []CalculatorPair{{1, 2}, {3, 0}, {0, 5}, {4, 7}}

	for _, pair := range pairList {
		fmt.Println(calculator.Add(pair.a, pair.b))
		fmt.Println(calculator.Sub(pair.a, pair.b))
		fmt.Println(calculator.Mul(pair.a, pair.b))

		result, err := calculator.Div(pair.a, pair.b)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}
