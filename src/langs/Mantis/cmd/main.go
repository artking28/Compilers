package main

import (
	"compilers/langs/Mantis/parser"
	"compilers/utils"
	"fmt"
	mgu "github.com/artking28/myGoUtils"
)

func main() {

	file := "../examples/example.mnts"
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
