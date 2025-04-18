package parser

import (
	"compilers/langs/Mantis/lexer"
	"compilers/stdParser"
	"compilers/utils"
)

type FuncStmt struct {
	Name string
	Body stdParser.Scope
	MantisStmtBase
}

func (this FuncStmt) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | FuncStmt@WriteMemASM")
}

func (this FuncStmt) GetTitle() string {
	return this.Title
}

func NewFuncStmt(name string, body stdParser.Scope, pos utils.Pos, parser *MantisParser) *FuncStmt {
	return &FuncStmt{
		Name: name,
		Body: body,
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "FuncStmt",
			Pos:    pos,
		},
	}
}

func (parser *MantisParser) ParseFunction() (err error) {
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
	if _, err = parser.HasNextConsume(stdParser.NoSpaceMode, lexer.L_BRACE); err != nil {
		return utils.GetExpectedTokenErrOr(parser.Filename, "left brace", err.Error(), h0.Pos)
	}
	ast, err := parser.ParseScope(utils.FuncScope)
	if err != nil {
		return err
	}
	if _, err = parser.HasNextConsume(stdParser.NoSpaceMode, lexer.R_BRACE); err != nil {
		return utils.GetExpectedTokenErrOr(parser.Filename, "right brace", err.Error(), h0.Pos)
	}
	parser.Inject(NewFuncStmt(string(nameTk.Value), ast, h0.Pos, parser))
	return nil
}
