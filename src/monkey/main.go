package main

import (
	"fmt"

	"monkey/lexer"
)

func main() {
	lexer := lexer.New("a + b")
	fmt.Println(lexer)
}
