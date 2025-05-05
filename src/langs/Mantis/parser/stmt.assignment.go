package parser

import (
	"compilers/langs/Mantis/lexer"
	"compilers/stdParser"
	"compilers/utils"
)

type AssignStmt struct {
	VariableId uint64
	ScopeId    uint64
	Expression stdParser.IExp
	MantisStmtBase
}

func NewAssignStmt(id uint64, scopeId uint64, exp stdParser.IExp, pos utils.Pos, parser *MantisParser) *AssignStmt {
	return &AssignStmt{
		VariableId: id,
		ScopeId:    scopeId,
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

func (parser *MantisParser) ParseArgAssign(scopeId uint64, kind lexer.MantisTokenKind) (ret *AssignStmt, err error) {
	h0 := parser.Get(0)
	if h0 == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}

	parser.Consume(1)
	if _, err = parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, kind); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "assignment", err.Error(), h0.Pos)
	}

	parser.Consume(1)
	assignValue, _, err := parser.ParseExpression()
	if err != nil {
		return nil, err
	}

	variable := parser.Variables[string(h0.Value)]
	//if variable.Type != "number" || exp.GetType() != "number" {
	//
	//}

	var exp stdParser.IExp
	switch kind {
	case lexer.ASSIGN_ADD:
		exp = NewMantisExpChain([]stdParser.IExp{variable, assignValue}, Operators[lexer.ADD], h0.Pos, parser)
		break
	case lexer.ASSIGN_SUB:
		exp = NewMantisExpChain([]stdParser.IExp{variable, assignValue}, Operators[lexer.SUB], h0.Pos, parser)
		break
	case lexer.ASSIGN_MUL:
		exp = NewMantisExpChain([]stdParser.IExp{variable, assignValue}, Operators[lexer.MUL], h0.Pos, parser)
		break
	case lexer.ASSIGN_MOD:
		exp = NewMantisExpChain([]stdParser.IExp{variable, assignValue}, Operators[lexer.MOD], h0.Pos, parser)
		break
	case lexer.ASSIGN_AND_BIT:
		exp = NewMantisExpChain([]stdParser.IExp{variable, assignValue}, Operators[lexer.AND_BIT], h0.Pos, parser)
		break
	case lexer.ASSIGN_XOR_BIT:
		exp = NewMantisExpChain([]stdParser.IExp{variable, assignValue}, Operators[lexer.XOR_BIT], h0.Pos, parser)
		break
	case lexer.ASSIGN_OR_BIT:
		exp = NewMantisExpChain([]stdParser.IExp{variable, assignValue}, Operators[lexer.OR_BIT], h0.Pos, parser)
		break
	case lexer.ASSIGN_SHIFT_RIGHT:
		exp = NewMantisExpChain([]stdParser.IExp{variable, assignValue}, Operators[lexer.SHIFT_RIGHT], h0.Pos, parser)
		break
	case lexer.ASSIGN_SHIFT_LEFT:
		exp = NewMantisExpChain([]stdParser.IExp{variable, assignValue}, Operators[lexer.SHIFT_LEFT], h0.Pos, parser)
		break
	case lexer.ASSIGN:
		exp = assignValue
		break
	default:
		panic("implement me | ParseArgAssign switch case")
	}

	if variable == nil {
		return nil, utils.GetUnkownVariableErr(parser.Filename, string(h0.Value), h0.Pos)
	}
	return NewAssignStmt(variable.Id, scopeId, exp, h0.Pos, parser), nil
}
