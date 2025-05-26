package parser

import (
	models "compilers/sharedModels"
	"compilers/utils"
	"fmt"
)

type (
	IExp interface {
		Stmt
		Resolve() (uint64, error)
		Count() int
		String() string
	}

	ExpChain struct {
		StmtBase
		All    []IExp           `json:"all"`
		Father *ExpChain        `json:"father"`
		Signal models.TokenKind `json:"oper"`
	}

	VExp uint64

	BExp bool

	IdExp uint64
)

func (this BExp) WriteMemASM() (string, error) {
	//TODO implement me
	panic("implement me | BExp@WriteMemASM")
}

func (this BExp) Count() int {
	return 1
}

func (this BExp) String() string {
	ret := "false"
	if this {
		ret = "true"
	}
	return ret
}

func TypeOf(exp IExp, parser *MantisParser) string {
	switch exp.GetTitle() {
	case "IdExp":
		id := uint64(exp.(IdExp))
		return TypeOf(parser.VariablesIDs[id].IExp, parser)
	case "VExp":
		return "number"
	case "BExp":
		return "bool"
	case "Exp", "ExpChain":
		switch exp.(*ExpChain).Signal {
		case models.LOWER_THEN, models.LOWER_EQUAL_THEN, models.GREATER_THEN, models.GREATER_EQUAL_THEN, models.EQUAL, models.AND_BOOL, models.OR_BOOL, models.XOR_BOOL:
			return "bool"
		case models.ADD, models.SUB, models.MUL, models.DIV, models.MOD, models.AND_BIT, models.OR_BIT, models.XOR_BIT, models.SHIFT_LEFT, models.SHIFT_RIGHT:
			return "number"
		default:
			return "unknown"
		}
	default:
		return "unknown"
	}
}

func (this IdExp) WriteMemASM() (string, error) {
	return fmt.Sprintf("GET %s\n", this.String()), nil
}

func (this IdExp) Count() int {
	return 1
}

func (this IdExp) String() string {
	return fmt.Sprintf("m%d", this)
}

func (this VExp) String() string {
	return fmt.Sprintf("%d", this)
}

func (this VExp) WriteMemASM() (string, error) {
	return fmt.Sprintf("GET %d\n", this), nil
}

func (this ExpChain) String() (ret string) {
	//if this.Count() == 1 {
	//	return this.All[0].String()
	//}
	//ret = fmt.Sprintf("( %s", this.All[0].String())
	//for _, v := range this.All[1:] {
	//	ret += fmt.Sprintf(" %s %s", this.Signal.GetSymbol(), v.String())
	//}
	//return fmt.Sprintf("%s )", ret)
	return ""
}

func (this ExpChain) WriteMemASM() (string, error) {
	//TODO implement me
	panic("implement me | ExpChain@WriteMemASM")
}

func NewExp(values []IExp, father *ExpChain, operator models.TokenKind, base StmtBase) *ExpChain {
	return &ExpChain{
		All:      values,
		Signal:   operator,
		Father:   father,
		StmtBase: base,
	}
}

func NewExpP(exps []IExp, father *ExpChain, oper models.TokenKind, pos utils.Pos, parser *MantisParser) *ExpChain {
	return NewExp(exps, father, oper, StmtBase{
		Parser: parser,
		Title:  "Exp",
		Pos:    pos,
	})
}

func (this *ExpChain) DeriveInclusiveExp(signal models.TokenKind) (*ExpChain, error) {
	last := this.All[len(this.All)-1]
	e := NewExp([]IExp{last}, this, signal, this.StmtBase)
	this.All[len(this.All)-1] = e
	return e, nil
}

func (this *ExpChain) AddTerm(term IExp) {
	this.All = append(this.All, term)
}

func (this *ExpChain) RootFather() *ExpChain {
	if this.Father != nil {
		return this.Father.RootFather()
	}
	return this
}

func NewIdExp(value uint64) *IdExp {
	n := IdExp(value)
	return &n
}

func NewVExp(value uint64) *VExp {
	n := VExp(value)
	return &n
}

func NewBExp(value bool) *BExp {
	n := BExp(value)
	return &n
}

func (this VExp) GetTitle() string {
	return "VExp"
}

func (this IdExp) GetTitle() string {
	return "IdExp"
}

func (this BExp) GetTitle() string {
	return "BExp"
}

func (this IdExp) Resolve() (uint64, error) {
	return 0, nil
}

func (this BExp) Resolve() (uint64, error) {
	return 0, nil
}

func (this VExp) Resolve() (uint64, error) {
	return uint64(this), nil
}

func (this VExp) Count() int {
	return 1
}

func (this ExpChain) Resolve() (uint64, error) {
	//TODO implement me
	panic("implement me | Exp@Resolve")
}

func (this ExpChain) Count() int {
	return len(this.All)
}
