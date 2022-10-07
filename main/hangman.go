package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"piscine"
)

func main() {
	if len(os.Args) == 2 || (len(os.Args) == 4 && os.Args[2] == "--letterFile") {
		Attempts := 10
		UdScore := []rune{}
		Pick := piscine.Random()
		Picklist := []rune(Pick)
		Word := piscine.Accent(Picklist)
		for range Pick {
			UdScore = append(UdScore, '_')
		}
		for v := 0; v < len(Word)/2-1; v++ {
			random := rand.Intn(len(Word))
			if UdScore[random] == '_' {
				UdScore[random] = rune(Picklist[random] - 32)
			} else {
				v--
			}
		}
		fmt.Println()
		fmt.Println("Good Luck, you have", Attempts, "Attempts !")
		piscine.Font(UdScore)
		piscine.Compare(UdScore, Attempts, Word, Picklist)
	} else if len(os.Args) == 3 && os.Args[1] == "--startWith" && os.Args[2] == "save.txt" {
		t, _ := os.ReadFile("save.txt")
		var save piscine.Sauvegarde
		err2 := json.Unmarshal(t, &save)
		if err2 != nil {
			fmt.Println(err2)
		}
		Attempts := save.Attempts
		fmt.Println()
		fmt.Println("Welcome back, you have", Attempts, "Attempts !")
		Pick := save.Pick
		UdScore := save.UdScore
		Picklist := []rune(Pick)
		Word := piscine.Accent(Picklist)
		piscine.Font(UdScore)
		piscine.Compare(UdScore, Attempts, Word, Picklist)
	} else {
		fmt.Println("Erreur")
	}
}
