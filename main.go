package main

import (
	"fmt"
	"math/rand"
	"unicode"
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
	guessedLetters := initializeGuessedWord(targetWord)
	printGameState(targetWord, guessedLetters)

}

func initializeGuessedWord(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func getRandomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}

func printGameState(targetWord string, guessedLetters map[rune]bool) {
	for _, char := range targetWord {
		if char == ' ' {
			fmt.Print(" ")
		} else if guessedLetters[unicode.ToLower(char)] {
			fmt.Printf("%c", char)
		} else {
			fmt.Print("_")
		}
		fmt.Print(" ")
	}

	fmt.Println()
}
