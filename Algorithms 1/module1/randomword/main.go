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
	if len(scanner.Text()) > 0 {
		fmt.Println(strings.Split(scanner.Text(), " ")[rand.Intn(len(strings.Split(scanner.Text(), " ")))])
	}
}