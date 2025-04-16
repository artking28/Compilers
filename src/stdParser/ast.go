package stdParser

type (
	Ast struct {
		Statements []Stmt `json:"statements"`
	}
	
	Variable struct {
		Id    uint64
		Name  string 
		Value Exp
		Owner uint64
	}
	
	Scope struct {
		Id        uint64
		Body      []Ast
	}
)