package main

import (
	"fmt"
	"os"
)

func main () {
	if len(os.Args) == 3{
		firstName := os.Args[1];
		secondName := os.Args[2];
		fmt.Println("Hello", firstName, "and", secondName)
		fmt.Println("Goodbye", secondName, "and", firstName)
	}
}