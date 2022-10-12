package database

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Data_Hangman struct {
	Word             string   // Word composed of '_', ex: H_ll_
	ToFind           string   // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int      // Number of attempts left
	HangmanPositions []string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func (w *Data_Hangman) ChoseWord() {
	arr, _ := os.ReadFile("database/Word.txt")
	arrWord := strings.Split(string(arr), string(13))
	w.ToFind = arrWord[rand.Intn(len(arrWord))]
	fmt.Println(w.ToFind)
}

func (w *Data_Hangman) DisplayLetters() {
	n := len(w.ToFind)/2 - 1
	var temp []string
	for i := 0; i < len(w.ToFind)-1; i++ {
		temp = append(temp, "_")
	}
	for i := 0; i < n; i++ {
		nb := rand.Intn(len(w.ToFind))
		fmt.Println(nb)
		temp[nb-1] = string(w.ToFind[nb])
	}
	for i := range temp {
		w.Word += temp[i]
	}
	fmt.Println(w.Word)
}
