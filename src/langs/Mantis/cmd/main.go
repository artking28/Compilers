package main

import (
	"compilers/langs/Mantis/lexer"
	"compilers/langs/Mantis/parser"
	"compilers/utils"
	"fmt"
	mgu "github.com/artking28/myGoUtils"
)

func main() {

	//cmd := ParseInput()
	cmd := CliCommand{
		Program:    "mantis",
		Method:     Build,
		TargetFile: "../examples/example1.mnts",
		Subset:     lexer.SUBSET_0,
	}

	file := cmd.TargetFile
	p, err := parser.NewMantisParser(file, "", 0)
	if err != nil {
		panic(err.Error())
	}

	out, err := p.ParseScope(utils.RootScope)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(mgu.String(out))
}
