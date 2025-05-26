package parser

import (
	models "compilers/sharedModels"
	"compilers/utils"
)

type IfStmt struct {
	Condition IExp
	Body      Scope
	MantisStmtBase
}

func (this IfStmt) WriteMemASM() (string, error) {
	//TODO implement me
	panic("implement me | IfStmt@WriteMemASM")
}

func NewIfStmt(condition IExp, body Scope, pos utils.Pos, parser *MantisParser) *IfStmt {
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
	ast, err := parser.ParseScope(IfScope)
	if err != nil {
		return nil, err
	}
	if _, err = parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.R_BRACE); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "right brace", err.Error(), h0.Pos)
	}
	return NewIfStmt(condition, ast, h0.Pos, parser), nil
}
