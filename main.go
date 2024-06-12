package main

import (
	"fmt"
	"math/rand"
)

var dictionary = []string{
	"Jeyhun",
	"Rahimli",
	"Gopher",
	"Golang",
	"Docker",
	"Kubernates",
	"AWS",
	"Cloud Computing",
}

func main() {
	targetWord := getRandomWord()

	fmt.Println(targetWord)
}

func getRandomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}
