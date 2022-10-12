package main

import "src/database"

func main() {
	var test database.Data_Hangman
	test.ChoseWord()
	test.DisplayLetters()
}
