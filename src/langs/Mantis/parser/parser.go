package parser

import (
	"compilers/langs/Mantis/lexer"
	"compilers/stdParser"
)

type (
	MantisParser stdParser.Parser[lexer.MantisTokenKind]
)

func NewMantisParser(filename, output string, subset int) (*MantisParser, error) {
	p, err := stdParser.NewParser[lexer.MantisTokenKind](filename, output, subset)
	return any(p).(*MantisParser), err
}

func (this *MantisParser) ParseAll() (stdParser.Ast, error) {
	return stdParser.Ast{}, nil
}
