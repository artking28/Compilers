package parser

import (
	models "compilers/sharedModels"
	"compilers/utils"
)

type AssignStmt struct {
	VariableId uint64
	ScopeId    uint64
	Expression IExp
	MantisStmtBase
}

func NewAssignStmt(id uint64, scopeId uint64, exp IExp, pos utils.Pos, parser *MantisParser) *AssignStmt {
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

func (this AssignStmt) WriteMemASM() (string, error) {
	//TODO implement me
	panic("implement me | AssignStmt@WriteMemASM")
}

func (parser *MantisParser) ParseArgAssign(scopeId uint64, kind models.TokenKind) (ret *AssignStmt, err error) {
	h0 := parser.Get(0)
	if h0 == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}

	parser.Consume(1)
	if _, err = parser.HasNextConsume(OptionalSpaceMode, models.SPACE, kind); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "assignment", err.Error(), h0.Pos)
	}

	parser.Consume(1)
	assignValue, err := parser.ParseExpression(false)
	if err != nil {
		return nil, err
	}

	variable := parser.VariablesNames[string(h0.Value)]
	//if variable.Type != "number" || exp.GetType() != "number" {
	//
	//}

	var exp IExp
	switch kind {
	case models.ASSIGN_ADD:
		exp = NewExpP([]IExp{variable, assignValue}, nil, models.ADD, h0.Pos, parser)
		break
	case models.ASSIGN_SUB:
		exp = NewExpP([]IExp{variable, assignValue}, nil, models.SUB, h0.Pos, parser)
		break
	case models.ASSIGN_MUL:
		exp = NewExpP([]IExp{variable, assignValue}, nil, models.MUL, h0.Pos, parser)
		break
	case models.ASSIGN_MOD:
		exp = NewExpP([]IExp{variable, assignValue}, nil, models.MOD, h0.Pos, parser)
		break
	case models.ASSIGN_AND_BIT:
		exp = NewExpP([]IExp{variable, assignValue}, nil, models.AND_BIT, h0.Pos, parser)
		break
	case models.ASSIGN_XOR_BIT:
		exp = NewExpP([]IExp{variable, assignValue}, nil, models.XOR_BIT, h0.Pos, parser)
		break
	case models.ASSIGN_OR_BIT:
		exp = NewExpP([]IExp{variable, assignValue}, nil, models.OR_BIT, h0.Pos, parser)
		break
	case models.ASSIGN_SHIFT_RIGHT:
		exp = NewExpP([]IExp{variable, assignValue}, nil, models.SHIFT_RIGHT, h0.Pos, parser)
		break
	case models.ASSIGN_SHIFT_LEFT:
		exp = NewExpP([]IExp{variable, assignValue}, nil, models.SHIFT_LEFT, h0.Pos, parser)
		break
	case models.ASSIGN:
		exp = assignValue
		break
	default:
		panic("implement me | ParseArgAssign switch case")
	}

	if variable == nil {
		return nil, utils.GetunknownVariableErr(parser.Filename, string(h0.Value), h0.Pos)
	}
	return NewAssignStmt(variable.Id, scopeId, exp, h0.Pos, parser), nil
}
