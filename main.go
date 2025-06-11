package main

import (
	"fmt"
	"os"
	"os/user"
	"persistio/evaluator"
	"persistio/lexer"
	"persistio/object"
	"persistio/parser"
	"persistio/repl"
)

func main() {
	if len(os.Args) > 1 {
		filePathArgument := os.Args[1]
		utilsContent, err := os.ReadFile("utils/utils.prs")
		if err != nil {
			panic(err)
		}
		content, err := os.ReadFile(filePathArgument)
		if err != nil {
			panic(err)
		}
		l := lexer.New(string(utilsContent) + string(content))
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		evaluator.Eval(program, env)
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
