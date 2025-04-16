package stdParser

import "compilers/stdLexer"

type Parser[T comparable] struct {
	Filename  string
	Tokens    []stdLexer.Token[T]
	Scopes    map[uint64]*Scope
	Variables map[string]*Variable
	Output    Ast
	Cursor    int
}

func NewParser[T comparable](filename, output string, subset int) (*Parser[T], error) {
	
	// tokens, err := lexer.Tokenize(filename)
    // if err != nil {
    	// return nil, err
    // }
	
	return &Parser[T]{
		Filename: filename,
		Output: Ast{},
		Cursor: 0,
		// Tokens: tokens,
	}, nil
}

func (this *Parser[T]) Inject(stmts ...Stmt) {
	this.Output.Statements = append(this.Output.Statements, stmts...)
}

func (this *Parser[T]) Get(n int) *stdLexer.Token[T] {
	if this.Cursor+n >= len(this.Tokens) {
		return nil
	}
	return &this.Tokens[this.Cursor+n]
}

func (this *Parser[T]) Consume(n int) {
	if this.Cursor >= len(this.Tokens) {
		return
	}
	this.Cursor += n
}

const (
	NoSpaceMode = iota
	OptionalSpaceMode
	MandatorySpaceMode
)

func (this *Parser[T]) HasNextConsume(spaceMode int, spaceCode T, kinds ...T) *stdLexer.Token[T] {
	if spaceMode < NoSpaceMode || spaceMode > MandatorySpaceMode {
		panic("invalid argument in function 'HasNextConsume'")
	}
	for findSpace := false; ; {
		token := this.Get(0)
		if token == nil {
			// Fim dos tokens sem encontrar um tipo esperado
			return nil
		}

		for _, kind := range kinds {
			if token.Kind == kind {
				// Se espaços eram obrigatórios mas não foram encontrados, falha
				if spaceMode == MandatorySpaceMode && !findSpace {
					return nil
				}
				this.Consume(1)
				return token
			}
		}

		if token.Kind == spaceCode {
			findSpace = true
			this.Consume(1)
			continue
		}

		// Se espaços não eram permitidos ou eram obrigatórios e encontrou outro token, falha
		if spaceMode == NoSpaceMode || spaceMode == MandatorySpaceMode {
			return nil
		}

		return nil // Qualquer outro caso não esperado falha
	}
}
