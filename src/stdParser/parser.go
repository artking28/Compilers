package stdParser

import (
	"compilers/stdLexer"
	"compilers/utils"
	"errors"
)

type Parser[T comparable] struct {
	Filename   string
	OutputFile string
	Tokens     []stdLexer.Token[T]
	Scopes     map[uint64]*Scope
	Variables  map[string]*Variable
	Cursor     int
}

func NewParser[T comparable](filename, output string, tokens []stdLexer.Token[T]) (*Parser[T], error) {

	return &Parser[T]{
		Filename:   filename,
		OutputFile: output,
		Tokens:     tokens,
		Scopes:     map[uint64]*Scope{},
		Variables:  map[string]*Variable{},
		Cursor:     0,
	}, nil
}

//func (this *Parser[T]) Inject(stmts ...Stmt) {
//	this.Output.Statements = append(this.Output.Statements, stmts...)
//}

func (this *Parser[T]) At() utils.Pos {
	return this.Tokens[this.Cursor].Pos
}

func (this *Parser[T]) Get(n int) *stdLexer.Token[T] {
	if this.Cursor+n >= len(this.Tokens) {
		return nil
	}
	return &this.Tokens[this.Cursor+n]
}

func (this *Parser[T]) Consume(n int) {
	if this.Cursor+n >= len(this.Tokens) {
		return
	}
	this.Cursor += n
}

func (this *Parser[T]) GetFirstAfter(afterOf ...T) (*stdLexer.Token[T], error) {
	all := map[T]bool{}
	for _, t := range afterOf {
		all[t] = true
	}

	token := this.Get(0)
	for i := 1; token != nil; i++ {
		if all[token.Kind] {
			token = this.Get(i)
			continue
		}
		return token, nil
	}
	return nil, errors.New("no token has been found")
}

const (
	NoSpaceMode = iota
	OptionalSpaceMode
	MandatorySpaceMode
)

func (this *Parser[T]) HasNextConsume(spaceMode int, fillOf T, kinds ...T) (*stdLexer.Token[T], error) {
	if spaceMode < NoSpaceMode || spaceMode > MandatorySpaceMode {
		panic("invalid argument in function 'HasNextConsume', unknown space mode")
	}
	if len(kinds) <= 0 {
		panic("invalid argument in function 'HasNextConsume', kinds is null or empty")
	}
	for hasSpace := false; ; {
		token := this.Get(0)
		if token == nil {
			// Fim dos tokens sem encontrar um tipo esperado
			return nil, errors.New("no token has been found")
		}

		for _, kind := range kinds {
			if token.Kind == kind {
				// Se espaços eram obrigatórios mas não foram encontrados, falha
				if spaceMode == MandatorySpaceMode && !hasSpace {
					return nil, errors.New("rule expects spaces but none has been found")
				}
				this.Consume(1)
				return token, nil
			}
		}

		if token.Kind == fillOf {
			// Se espaços não eram permitidos
			if spaceMode == NoSpaceMode {
				return nil, errors.New("space(s) has been found when it actually isn't allowed here")
			}
			hasSpace = true
			this.Consume(1)
			continue
		}

		// Se espaços eram obrigatórios e encontrou outro token, falha
		if spaceMode == MandatorySpaceMode {
			return nil, errors.New("rule expects spaces but none has been found")
		}

		return nil, errors.New("unknown error") // Qualquer outro caso não esperado falha
	}
}
