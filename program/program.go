package program

import (
	"fmt"
	"os"
	"path/filepath"
	"persistio/ast"
	"persistio/lexer"
	"persistio/object"
	"persistio/parser"
	"persistio/utils"
)

func CreateProgramFromFile(filePath string, eval func(ast.Node, *object.Environment) object.Object) (*ast.Program, *object.Environment) {
	basePath := filepath.Dir(filePath)
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	utilsContent := utils.GetUtilsContent()
	l := lexer.New(utilsContent + string(content))
	p := parser.New(l)
	program := p.ParseProgramFromFile(basePath)
	if len(p.Errors()) != 0 {
		fmt.Println("Parser errors:")
		for _, msg := range p.Errors() {
			fmt.Printf("\t%s\n", msg)
		}
		os.Exit(1)
	}
	env := object.NewEnvironment()
	eval(program, env)
	return program, env
}

func CreateBaseProgram(eval func(ast.Node, *object.Environment) object.Object) (*ast.Program, *object.Environment) {
	utilsContent := utils.GetUtilsContent()
	l := lexer.New(utilsContent)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		fmt.Println("Parser errors:")
		for _, msg := range p.Errors() {
			fmt.Printf("\t%s\n", msg)
		}
		os.Exit(1)
	}
	env := object.NewEnvironment()
	eval(program, env)
	return program, env
}
