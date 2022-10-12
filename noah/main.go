package main

import (
	"bufio"
	"fmt"
	"os"
	
)

func main() {
	lettre := "mots caché"
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Ecrit la lettre de ton choix : ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Tu as ecrit : %q", input)
	for char := 'a'; char <= 'z'; char++ {
		if "%q" == char && "%q" == lettre {
			fmt.Println("Vous avez découvert une lettre :")
	}else{
		fmt.Println("Le caractère que vous avez rentre n'est pas valide")
		break
	}else if{
		continue
	}

}
}