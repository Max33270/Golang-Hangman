package piscine

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Random() string {
	rand.Seed(time.Now().UnixNano())
	file, err := os.Open(os.Args[1])
	s := ""
	var a int
	if os.Args[1] == "words.txt" {
		a = rand.Intn(37) + 1
	} else if os.Args[1] == "words2.txt" {
		a = rand.Intn(23) + 1
	} else if os.Args[1] == "words3.txt"{
		a = rand.Intn(24) + 1
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
		if i == a {
			s = scanner.Text()
		}
	}
	return s
}
