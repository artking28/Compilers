package stdParser

type (
	IExp[T any] interface {
		Resolve() (T, error)
	}

	Exp[T any] struct {
		All  []T
		Oper Operator
	}

	VExp[T any] struct {
		Value T
	}

	Operator int
)

const (
	Sum Operator = iota
	Sub
	ShiftL
	ShiftR
	AndBit
	OrBit
	XorBit
)

func NewVExp[T any](value T) *VExp[T] {
	return &VExp[T]{Value: value}
}

func NewExp[T any](all []T, oper Operator) *Exp[T] {
	return &Exp[T]{All: all, Oper: oper}
}
