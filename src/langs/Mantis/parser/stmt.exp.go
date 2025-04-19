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
	MantisExp struct {
		stdParser.Exp
		MantisStmtBase
	}

	MantisVExp struct {
		stdParser.VExp
		MantisStmtBase
	}
)

func (this *MantisVExp) Resolve() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (this *MantisExp) Resolve() (int, error) {
	//TODO implement me
	panic("implement me | MantisExp@WriteMemASM")
}

func (this MantisVExp) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | MantisVExp@WriteMemASM")
}

func (this MantisVExp) GetTitle() string {
	return this.Title
}

func (this MantisExp) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | MantisExp@WriteMemASM")
}

func (this MantisExp) GetTitle() string {
	return this.Title
}

func NewMantisVExp(value int, pos utils.Pos, parser *MantisParser) *MantisVExp {
	return &MantisVExp{
		VExp: *stdParser.NewVExp(value),
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "MantisVExp",
			Pos:    pos,
		},
	}
}

func NewMantisExp(exp stdParser.Exp, pos utils.Pos, parser *MantisParser) *MantisExp {
	return &MantisExp{
		Exp: exp,
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "MantisExp",
			Pos:    pos,
		},
	}
}

func (parser *MantisParser) ParseExpression(endAt lexer.MantisTokenKind) (stdParser.IExp, error) {
	h0 := parser.Get(0)
	h1 := parser.Get(1)
	if h0 == nil || h1 == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	if h1.Kind == endAt {
		return NewMantisVExp(int(h1.Value[0]), h1.Pos, parser), nil
	}

	var e stdParser.Exp
	if h0.Kind == lexer.ID || h0.Kind == lexer.NUMBER || h0.Kind == lexer.NIL {
		exp, err := ParseExpressionReturn(endAt)
		if err == nil {
			return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
		}
		e = any(exp).(stdParser.Exp)
	} else if h0.Kind == lexer.TRUE || h0.Kind == lexer.FALSE {
		exp, err := ParseExpressionReturn(endAt)
		if err == nil {
			return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
		}
		e = any(exp).(stdParser.Exp)
	} else {
		return nil, utils.GetExpectedTokenErr(parser.Filename, "valid expression like id, numbers, booleans or nil", h0.Pos)
	}
	return NewMantisExp(e, h0.Pos, parser), nil
}

func ParseExpressionReturn(endAt ...lexer.MantisTokenKind) (stdParser.IExp, error) {
	if len(endAt) <= 0 {
		panic("invalid argument in function 'ParseExpressionReturn', endAt is null or empty")
	}
	return &MantisVExp{
		VExp:           stdParser.VExp{Value: 0},
		MantisStmtBase: MantisStmtBase{},
	}, nil
}
