package stdParser

type (
	
	IExpValue interface {
		bool | int
	}
	
	ExpValue int

	Exp interface {
		GetValue() any
	}
)