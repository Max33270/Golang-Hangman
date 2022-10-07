package piscine

import (
	"encoding/json"
	"fmt"
	"os"
)

type Sauvegarde struct {
	Attempts int
	Pick     string
	UdScore  []rune
	Police   string
}

func Compare(u []rune, a int, w []rune, p []rune) {
	var letter string
	proposition := []string{}
	for string(u) != ToUpper(string(p)) && a > 0 {
		fmt.Print("Choose : ")
		fmt.Scanln(&letter)
		d := false
		letter = Rmv_accent(letter)
		proposition = append(proposition, ToUpper(letter))
		if len(letter) == 1 { // Letter
			for v := range w {
				if (w[v] == rune(letter[0]) || w[v] == rune(letter[0]+32)) && u[v] == '_' {
					u[v] = rune(p[v] - 32)
					d = true
				}
			}
		} else if letter == string(w) || letter == ToUpper(string(w)) { // word
			for v := range w {
				u[v] = rune(p[v] - 32)
			}
			d = true
		} else if letter == "STOP" || letter == "Stop" || letter == "stop" {
			var m Sauvegarde
			if len(os.Args) == 4 {
				m = Sauvegarde{Attempts: a, Pick: string(p), UdScore: u, Police: os.Args[3]}
			} else {
				m = Sauvegarde{Attempts: a, Pick: string(p), UdScore: u}
			}
			b, err := json.Marshal(m)
			os.WriteFile("save.txt", b, 0666)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			os.Exit(0)
		}
		repeat := 0
		for v := range proposition {
			if proposition[v] == ToUpper(letter) {
				repeat++
			}
		}

		if repeat > 1 {
			fmt.Println("Oh sorry man you've already written this!!!!")
			d = true
		}
		a = Check(letter, a, u, d)
	}
	if a < 1 {
		fmt.Println("Oh Snap !! José is dead! RIP")
		fmt.Println("The word was", string(p))
	} else {
		fmt.Println("Congrats !")
	}
}
