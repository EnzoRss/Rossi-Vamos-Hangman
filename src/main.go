package main

import (
	"os"
	"src/database"
)

func main() {
	var test database.Data_Hangman

	if len(os.Args) > 1 && os.Args[1] == "--startWith" {
		test.Resume()
	} else {
		test.Init()
	}

	test.Input()
}
