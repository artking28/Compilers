package main

import (
	"compilers/parser"
	"fmt"
)

func main() {

	cmd := ParseInput()
	//cmd := CliCommand{
	//	Program:    "mantis",
	//	Method:     Build,
	//	TargetFile: "../example.mnts",
	//}

	file := cmd.TargetFile
	p, err := parser.NewMantisParser(file, "")
	if err != nil {
		panic(err.Error())
	}

	out, err := p.ParseScope(parser.RootScope)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(out)
}
