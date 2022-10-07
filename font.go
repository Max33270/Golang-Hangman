package piscine

import (
	"encoding/json"
	"fmt"
	"os"
)

func Font(s []rune) {
	t, _ := os.ReadFile("save.txt")
	var save Sauvegarde
	err2 := json.Unmarshal(t, &save)
	if err2 != nil {
		fmt.Println(err2)
	}
	if len(os.Args) == 4 || (len(os.Args) == 3 && save.Police != "" ) {
		b := [9]string{}
		for v := range s {
			b = Ascii(int(s[v]), b)
		}
		for v := range b {
			fmt.Println(b[v])
		}
	} else {
		for v := range s {
			fmt.Print(string(s[v]))
			fmt.Print(" ")
		}
	}

	fmt.Println()
	fmt.Println()
}
