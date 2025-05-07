package parser

import (
	"compilers/langs/Mantis/lexer"
	"compilers/stdParser"
	"compilers/utils"
)

type IfStmt struct {
	Condition stdParser.IExp
	Body      stdParser.Scope
	MantisStmtBase
}

func (this IfStmt) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | IfStmt@WriteMemASM")
}

func NewIfStmt(condition stdParser.IExp, body stdParser.Scope, pos utils.Pos, parser *MantisParser) *IfStmt {
	return &IfStmt{
		Condition: condition,
		Body:      body,
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "IfStmt",
			Pos:    pos,
		}}
}

func (parser *MantisParser) ParseIfStatement() (ret *IfStmt, err error) {
	h0 := parser.Get(0)
	if h0 == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	parser.Consume(1)

	condition, typeOf, err := parser.ParseExpression()
	if err != nil {
		return nil, err
	}
	if typeOf != "bool" {
		//TODO implement me
		panic("implement me | MantisParser@ParseIfStatement")
	}

	if _, err = parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, lexer.L_BRACE); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "left brace", err.Error(), h0.Pos)
	}
	ast, err := parser.ParseScope(utils.IfScope)
	if err != nil {
		return nil, err
	}
	if _, err = parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, lexer.R_BRACE); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "right brace", err.Error(), h0.Pos)
	}
	return NewIfStmt(condition, ast, h0.Pos, parser), nil
}
