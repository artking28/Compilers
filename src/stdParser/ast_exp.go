package stdParser

type (
	IExp interface {
		Resolve() (int, error)
	}

	Exp struct {
		All  []int
		Oper Operator
	}

	VExp struct {
		Value int
	}

	Operator int
)

func (e Exp) Resolve() (int, error) {
	//TODO implement me
	panic("implement me")
}

const (
	Sum Operator = iota
	Sub
	ShiftL
	ShiftR
	AndBit
	OrBit
	XorBit
)

func NewVExp(value int) *VExp {
	return &VExp{Value: value}
}

func NewExp(all []int, oper Operator) *Exp {
	return &Exp{All: all, Oper: oper}
}
