package main

import "fmt"

func SliceToMap(s []string) map[string]int {
	m := make(map[string]int)

	for _, v := range s {
		m[v] = m[v] + 1
	}

	return m
}

func main() {
	fmt.Println(SliceToMap([]string{"go", "js", "go"}))
}
