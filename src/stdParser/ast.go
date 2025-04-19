package stdParser

type (
	Ast struct {
		Statements []Stmt `json:"statements"`
	}

	Variable[T any] struct {
		Id    uint64
		Name  string
		Value IExp[T]
		Owner uint64
	}

	Scope struct {
		Id   uint64
		Body []Ast
	}
)

func NewVariable[T any](id uint64, name string, value IExp[T], owner uint64) *Variable[T] {
	return &Variable[T]{Id: id, Name: name, Value: value, Owner: owner}
}
