package main

import (
	"fmt"
	"strings"
)

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Lowercase(s string) string {
	return strings.ToLower(s)
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func main() {
	fmt.Println(transformString("Aditya", Uppercase))
	fmt.Println(transformString("Aditya", Lowercase))
}
