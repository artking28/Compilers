package stdParser

type (
	Ast struct {
		Statements []Stmt `json:"statements"`
	}

	Variable struct {
		Id    uint64
		Name  string
		Value IExp[any]
		Owner uint64
	}

	Scope struct {
		Id   uint64
		Body []Ast
	}
)
