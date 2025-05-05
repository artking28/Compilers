package stdParser

type (
	Ast struct {
		Statements []Stmt `json:"statements"`
	}

	Variable struct {
		Id    uint64
		Name  string
		Owner uint64
		IExp
	}

	Scope struct {
		Id   uint64
		Body Ast
	}
)

func NewVariable(id uint64, name string, value IExp, owner uint64) *Variable {
	if value == nil {
		return nil
	}
	return &Variable{Id: id, Name: name, Owner: owner, IExp: value}
}
