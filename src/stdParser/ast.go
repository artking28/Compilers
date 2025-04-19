package stdParser

type (
	Ast struct {
		Statements []Stmt `json:"statements"`
	}

	Variable struct {
		Id    uint64
		Name  string
		Value IExp
		Owner uint64
	}

	Scope struct {
		Id   uint64
		Body Ast
	}
)

func NewVariable(id uint64, name string, value IExp, owner uint64) *Variable {
	return &Variable{Id: id, Name: name, Value: value, Owner: owner}
}
