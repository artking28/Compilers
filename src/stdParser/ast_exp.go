package stdParser

type (
	IExp interface {
		Stmt
		Resolve() (int, error)
		GetType() string
		Values() int
	}

	Exp struct {
		StmtBase
		All  []IExp
		Oper Operator
	}

	VExp struct {
		StmtBase
		Value int
		Type  string
	}

	Operator int
)

func NewVExp(value int, typeOf string) *VExp {
	return &VExp{
		Value: value,
		Type:  typeOf,
	}
}

func NewExp(all []IExp, oper Operator) *Exp {
	return &Exp{All: all, Oper: oper}
}

func (this VExp) Resolve() (int, error) {
	return this.Value, nil
}

func (this VExp) GetType() string {
	return this.Type
}

func (this VExp) Values() int {
	return 1
}

func (this Exp) Resolve() (int, error) {
	//TODO implement me
	panic("implement me | Exp@Resolve")
}

func (this Exp) GetType() string {
	switch this.Oper & (NumberOutput | BoolOutput) {
	case NumberOutput:
		return "number"
	case BoolOutput:
		return "bool"
	default:
		return "unknown"
	}
}

func (this Exp) Values() int {
	return len(this.All)
}

const (
	NumberInput = 1 << iota
	NumberOutput
	BoolInput
	BoolOutput
)

const (
	None Operator = iota
	Mul           = iota*10 + NumberInput + NumberOutput
	Mod
	Sum
	Sub
	ShiftL
	ShiftR
	LowerThen = iota*10 + NumberInput + BoolOutput
	LowerEqualThen
	GreaterThen
	GreaterEqualThen
	AndBit
	OrBit
	XorBit
	BoolAnd = iota*10 + BoolInput + BoolOutput
	BoolOr
	BoolXor
)

var OperLevel = map[Operator]int{
	Mul:              0,
	Mod:              0,
	Sum:              1,
	Sub:              1,
	ShiftL:           2,
	ShiftR:           2,
	LowerThen:        3,
	LowerEqualThen:   3,
	GreaterThen:      3,
	GreaterEqualThen: 3,
	AndBit:           4,
	XorBit:           5,
	OrBit:            6,
	BoolAnd:          7,
	BoolXor:          8,
	BoolOr:           9,
}

func (this Operator) Compare(input Operator) int {
	if OperLevel[this] < OperLevel[input] {
		return 1
	} else if OperLevel[this] == OperLevel[input] {
		return 0
	}
	return -1
}
