package parser

import (
	"compilers/langs/Mantis/lexer"
	"compilers/stdParser"
	"compilers/utils"
)

var ExpTokens = []lexer.MantisTokenKind{
	lexer.ID,
	lexer.NUMBER,
	lexer.NIL,
	lexer.TRUE,
	lexer.FALSE,
}

type (
	MantisExp[T any] struct {
		stdParser.Exp[T]
		MantisStmtBase
	}

	MantisVExp[T any] struct {
		stdParser.VExp[T]
		MantisStmtBase
	}
)

func (this MantisVExp[T]) Resolve() (T, error) {
	//TODO implement me
	panic("implement me")
}

func (this MantisExp[T]) Resolve() (T, error) {
	//TODO implement me
	panic("implement me | MantisExp@WriteMemASM")
}

func (this MantisVExp[T]) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | MantisVExp@WriteMemASM")
}

func (this MantisVExp[T]) GetTitle() string {
	return this.Title
}

func (this MantisExp[T]) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | MantisExp@WriteMemASM")
}

func (this MantisExp[T]) GetTitle() string {
	return this.Title
}

func NewMantisVExp[T any](value T, pos utils.Pos, parser *MantisParser) *MantisVExp[T] {
	return &MantisVExp[T]{
		VExp: *stdParser.NewVExp(value),
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "MantisVExp",
			Pos:    pos,
		},
	}
}

func NewMantisExp[T any](exp stdParser.Exp[T], pos utils.Pos, parser *MantisParser) *MantisExp[T] {
	return &MantisExp[T]{
		Exp: exp,
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "MantisExp",
			Pos:    pos,
		},
	}
}

func (parser *MantisParser) ParseExpression(endAt lexer.MantisTokenKind) error {
	h0 := parser.Get(0)
	h1 := parser.Get(1)
	if h0 == nil || h1 == nil {
		return utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	if h1.Kind == endAt {
		parser.Inject(NewMantisVExp(h1.Value, h1.Pos, parser))
	}

	var e stdParser.Exp[any]
	if h0.Kind == lexer.ID || h0.Kind == lexer.NUMBER || h0.Kind == lexer.NIL {
		exp, err := ParseExpressionReturn[int](endAt)
		if err == nil {
			return utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
		}
		e = any(exp).(stdParser.Exp[any])
	} else if h0.Kind == lexer.TRUE || h0.Kind == lexer.FALSE {
		exp, err := ParseExpressionReturn[bool](endAt)
		if err == nil {
			return utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
		}
		e = any(exp).(stdParser.Exp[any])
	} else {
		return utils.GetExpectedTokenErr(parser.Filename, "valid expression like id, numbers, booleans or nil", h0.Pos)
	}
	parser.Inject(NewMantisExp(e, h0.Pos, parser))
	return nil
}

func ParseExpressionReturn(endAt ...lexer.MantisTokenKind) (*stdParser.Exp[any], error) {
	if len(endAt) <= 0 {
		panic("invalid argument in function 'ParseExpressionReturn', endAt is null or empty")
	}
	return nil, nil
}
