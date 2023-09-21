package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 1 {
		fp := os.Args[1]
		if _, err := os.Stat(fp); err == nil {
			repl.Run(fp, os.Stdout)

		} else {
			fmt.Printf("File %s not found! exiting...", fp)

		}

	} else {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
