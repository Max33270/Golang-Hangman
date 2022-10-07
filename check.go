package piscine

import "fmt"

func Check(l string, a int, u []rune, d bool) int {
	if d {
		Font(u)
	} else {
		a--
		if len(l) > 1 && a > 0 {
			a--
		}
		fmt.Println("Not present in the word,", a, "attempts remaining")
		if Read("killjose.txt")[9] == byte(13) {
			fmt.Println(string(Read("killjose.txt")[(9-a)*79 : (9-a)*79+77]))
		} else {
			fmt.Println(string(Read("killjose.txt")[(9-a)*71 : (9-a)*71+70]))
		}
	}
	return a
}
