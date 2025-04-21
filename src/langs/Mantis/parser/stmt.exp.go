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

func (parser *MantisParser) ParseExpression(endAt ...lexer.MantisTokenKind) (stdParser.IExp, error) {
	h0 := parser.Get(0)
	if h0 == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}

	var expList []stdParser.IExp
	var lastSig = lexer.UNKNOW
	lastWasSig := false

	for i:=0; h0 != nil; i++ {
		if h0.Kind == lexer.ADD || h0.Kind == lexer.SUB {
			for {
				i++
				hs, e := parser.GetFirstAfter(lexer.SPACE)
				if e != nil {
					return nil, e
				}
				if hs.Kind == lexer.ADD || hs.Kind == lexer.SUB {

				} else {
					lastSig = h0.Kind
					break
				}
			}
		}

		h0 = parser.Get(i)
	}

	return new Exp(expList, lastSig)
}















