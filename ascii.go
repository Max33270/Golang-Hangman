package piscine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func Ascii(a int, alpha [9]string) [9]string {
	a = a - 32
	t, _ := os.ReadFile("save.txt")
	var save Sauvegarde
	err2 := json.Unmarshal(t, &save)
	if err2 != nil {
		fmt.Println(err2)
	}
	file, err := os.Open("standart.txt")
	if len(os.Args) == 3 {
		file, err = os.Open(save.Police)
	} else if os.Args[3] == "standart.txt" || os.Args[3] == "shadow.txt" || os.Args[3] == "thinkertoy.txt" {
		file, err = os.Open(os.Args[3])
	} else {
		fmt.Println("Erreur")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		if i < 9*a+10 && i > 9*a {
			alpha[i-9*a-1] += scanner.Text()
		}
	}
	return alpha
}
