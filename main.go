package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)
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
	hungmanState := 0
	for {
		printGameState(targetWord, guessedLetters, hungmanState)
		input := readInput()
		if len(input) > 1 {
			fmt.Println("Invalid input. Please use letters only ...")
			continue
		}
	}

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

func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println(getHangmanDrawing(hangmanState))
}

func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""
	for _, char := range targetWord {
		if char == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(char)] {
			result += fmt.Sprintf("%c", char)
		} else {
			result += "_"
		}
		result += " "
	}

	return result
}

func getHangmanDrawing(state int) string {
	data, err := os.ReadFile(fmt.Sprintf("states/hangman%d", state))
	if err != nil {
		panic(err)
	}

	return string(data)
}

func readInput() string {
	fmt.Print("> ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}
