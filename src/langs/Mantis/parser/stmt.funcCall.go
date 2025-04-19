package parser

import (
	"compilers/langs/Mantis/lexer"
	"compilers/stdParser"
	"compilers/utils"
)

type FuncCall struct {
	Name string
	from uint64
	MantisStmtBase
}

func (this FuncCall) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | FuncCall@WriteMemASM")
}

func (this FuncCall) GetTitle() string {
	return this.Title
}

func NewFuncCall(name string, from uint64, pos utils.Pos, parser *MantisParser) *FuncCall {
	return &FuncCall{
		Name: name,
		from: from,
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "FuncCall",
			Pos:    pos,
		}}
}

func (parser *MantisParser) ParseFuncCall(from uint64) (err error) {
	h0 := parser.Get(0)
	if h0 == nil {
		return utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	nameTk, _ := parser.HasNextConsume(stdParser.MandatorySpaceMode, lexer.ID)
	if nameTk == nil {
		return utils.GetExpectedTokenErr(parser.Filename, "function name", h0.Pos)
	}
	if _, err = parser.HasNextConsume(stdParser.NoSpaceMode, lexer.L_PAREN); err != nil {
		return utils.GetExpectedTokenErrOr(parser.Filename, "left parenthesis", err.Error(), h0.Pos)
	}
	if _, err = parser.HasNextConsume(stdParser.NoSpaceMode, lexer.R_PAREN); err != nil {
		return utils.GetExpectedTokenErrOr(parser.Filename, "right parenthesis", err.Error(), h0.Pos)
	}
	parser.Inject(NewFuncCall(string(nameTk.Value), from, h0.Pos, parser))
	return nil
}
