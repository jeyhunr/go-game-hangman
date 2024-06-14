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
	hangmanState := 0

	for !gameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) > 1 {
			fmt.Println("Invalid input. Please use letters only ...")
			continue
		}

		letter := rune(input[0])
		if isGuessCorrect(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}

	fmt.Println("Game over ...")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You win!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You lose!")
	} else {
		panic("Invalid state. The game is over and there is no winner.")
	}
}

func getRandomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}

func initializeGuessedWord(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
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

func isGuessCorrect(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}

func gameOver(targetWord string, guessedLetters map[rune]bool, hamgmanState int) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hamgmanState)
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, char := range targetWord {
		if !guessedLetters[char] {
			return false
		}
	}

	return true
}
