package main

import (
	"fmt"
	"os"
	"os/user"
	"persistio/evaluator"
	"persistio/program"
	"persistio/repl"
)

func main() {
	if len(os.Args) > 1 {
		eval := evaluator.Eval
		filePathArgument := os.Args[1]
		program.CreateProgramFromFile(filePathArgument, eval)
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is Ivan's programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
