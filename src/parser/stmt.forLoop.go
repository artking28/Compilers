package parser

import (
	models "compilers/sharedModels"
	"compilers/utils"
)

type ForStmt struct {
	Condition IExp
	Body      Scope
	MantisStmtBase
}

func (this ForStmt) WriteMemASM() (string, error) {
	//TODO implement me
	panic("implement me | ForStmt@WriteMemASM")
}

func NewForStmt(condition IExp, body Scope, pos utils.Pos, parser *MantisParser) *ForStmt {
	return &ForStmt{
		Condition: condition,
		Body:      body,
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "ForStmt",
			Pos:    pos,
		}}
}

func (parser *MantisParser) ParseForLoopStatement() (ret *ForStmt, err error) {
	h0 := parser.Get(0)
	if h0 == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	parser.Consume(1)

	condition, err := parser.ParseExpression(false)
	if err != nil {
		return nil, err
	}
	if TypeOf(condition, parser) != "bool" {
		//TODO implement me
		panic("implement me | MantisParser@ParseForStatement")
	}

	if _, err = parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.L_BRACE); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "left brace", err.Error(), h0.Pos)
	}
	ast, err := parser.ParseScope(ForScope)
	if err != nil {
		return nil, err
	}
	if _, err = parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.R_BRACE); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "right brace", err.Error(), h0.Pos)
	}
	return NewForStmt(condition, ast, h0.Pos, parser), nil
}
