package stdParser

type (
	IExp interface {
		Resolve() (int, error)
	}

	Exp struct {
		All  []IExp
		Oper Operator
	}

	VExp struct {
		Value int
	}

	Operator int
)

func NewVExp(value int) *VExp {
	return &VExp{Value: value}
}

func NewExp(all []IExp, oper Operator) *Exp {
	return &Exp{All: all, Oper: oper}
}

func (e Exp) Resolve() (int, error) {
	//TODO implement me
	panic("implement me")
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
