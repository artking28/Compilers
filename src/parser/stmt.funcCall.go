package parser

import (
	models "compilers/sharedModels"
	"compilers/utils"
)

type FuncCall struct {
	Name string
	from uint64
	MantisStmtBase
}

func (this FuncCall) WriteMemASM() (string, error) {
	//TODO implement me
	panic("implement me | FuncCall@WriteMemASM")
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

func (parser *MantisParser) ParseFuncCall(from uint64) (ret *FuncCall, err error) {
	nameTk := parser.Get(0)
	if nameTk == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	parser.Consume(1)
	if _, err = parser.HasNextConsume(NoSpaceMode, models.SPACE, models.L_PAREN); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "left parenthesis", err.Error(), nameTk.Pos)
	}
	if _, err = parser.HasNextConsume(NoSpaceMode, models.SPACE, models.R_PAREN); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "right parenthesis", err.Error(), nameTk.Pos)
	}
	return NewFuncCall(string(nameTk.Value), from, nameTk.Pos, parser), nil
}
