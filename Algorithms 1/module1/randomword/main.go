package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var words []string
	var champion string
	if len(scanner.Text()) > 0 {
		if len(scanner.Text()) > 0 {
			for i := 0; i < len(strings.Fields(scanner.Text())); i++ {
				words = append(words, strings.Fields(scanner.Text())[i])
			}
			randomIdx := rand.Intn(len(strings.Split(scanner.Text(), " ")))
			champion = words[randomIdx]
		}
	}
	fmt.Println(champion)
}