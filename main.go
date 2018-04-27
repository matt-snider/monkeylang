package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/matt-snider/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome %s\n", user.Name)
	fmt.Println("This is the Monkey Language")
	repl.Run(os.Stdin, os.Stdout)
	fmt.Println("")
}
