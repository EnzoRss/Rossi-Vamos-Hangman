package src

import (
	"io/ioutil"
	"math/rand"
	"strings"
)

type Data_Hangman struct {
	Word             string   // Word composed of '_', ex: H_ll_
	ToFind           string   // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int      // Number of attempts left
	HangmanPositions []string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func (w *Data_Hangman) ChoseWord() {
	arr, _ := ioutil.ReadFile("word.txt")
	arrWord := strings.Split(string(arr), "\n")
	w.Word = arrWord[rand.Intn(len(arrWord))]
}

func 
