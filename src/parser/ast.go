package parser

type (
	Ast struct {
		Statements []Stmt `json:"statements"`
	}

	Variable struct {
		Id   uint64 `json:"id"`
		Name string `json:"name"`
		StmtBase
		IExp
	}

	Scope struct {
		Id   uint64    `json:"id"`
		Kind ScopeKind `json:"kind"`
		Body Ast       `json:"body"`
		StmtBase
	}

	ScopeKind int
)

const (
	RootScope ScopeKind = iota
	FuncScope
	IfScope
	ElseScope
	ForScope
)

func (this Scope) GetTitle() string {
	//TODO implement me
	return "Scope"
}

func (this Scope) WriteMemASM() (string, error) {
	//TODO implement me
	panic("implement me | Scope@WriteMemASM")
}

func (this Variable) GetTitle() string {
	//TODO implement me
	return "Variable"
}

func (this Variable) WriteMemASM() (ret string, err error) {
	//TODO implement me
	panic("implement me | Variable@WriteMemASM")
}

func NewVariable(name string, value IExp, parser *MantisParser) *Variable {
	if value == nil {
		return nil
	}
	id := uint64(len(parser.VariablesNames) + 1)
	ret := &Variable{Id: id, Name: name, IExp: value, StmtBase: StmtBase{
		Parser: parser,
		Title:  "Variable",
		Pos:    parser.At(),
	}}
	parser.VariablesNames[name] = ret
	parser.VariablesIDs[id] = ret
	return ret
}
