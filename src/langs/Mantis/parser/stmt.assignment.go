package parser

import (
	"compilers/langs/Mantis/lexer"
	"compilers/stdParser"
	"compilers/utils"
)

type AssignStmt struct {
	VariableId uint64
	Expression stdParser.IExp
	MantisStmtBase
}

func NewAssignStmt(id uint64, exp stdParser.IExp, pos utils.Pos, parser *MantisParser) *AssignStmt {
	return &AssignStmt{
		VariableId: id,
		Expression: exp,
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "AssignStmt",
			Pos:    pos,
		}}
}

func (this AssignStmt) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | AssignStmt@WriteMemASM")
}

func (parser *MantisParser) ParseAssign(scopeId uint64) error {
	return nil
}

func (parser *MantisParser) ParseArgAssign(scopeId uint64, kind lexer.MantisTokenKind) (ret *AssignStmt, err error) {
	h0 := parser.Get(0)
	if h0 == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}

	nameVar, err := parser.HasNextConsume(stdParser.MandatorySpaceMode, lexer.SPACE, lexer.ID)
	if nameVar == nil {
		return nil, utils.GetExpectedTokenErr(parser.Filename, "variable name", h0.Pos)
	}

	if _, err = parser.HasNextConsume(stdParser.NoSpaceMode, lexer.SPACE, kind); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "assignment", err.Error(), h0.Pos)
	}

	parser.Consume(1)
	exp, typeof, err := parser.ParseExpression()
	if err != nil {
		return nil, err
	}

	variable := parser.Variables[typeof]
	if variable.Type != "number" || exp.GetType() != "number" {

	}

	switch kind {
	case lexer.ASSIGN_ADD:

	case lexer.ASSIGN_SUB:

	case lexer.ASSIGN_MUL:

	case lexer.ASSIGN_AND_BIT:

	case lexer.ASSIGN_XOR_BIT:

	case lexer.ASSIGN_OR_BIT:

	case lexer.ASSIGN_SHIFT_RIGHT:

	case lexer.ASSIGN_SHIFT_LEFT:

	}

	return NewAssignStmt(variable.Id, exp, h0.Pos, parser), nil
}
