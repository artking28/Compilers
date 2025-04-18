package utils

type (
	Pos struct {
		Line   int64 `json:"line"`
		Column int64 `json:"column"`
	}

	ScopeType int
)

const (
	RootScope ScopeType = iota
	FuncScope
	ForScope
	IfScope
	ElseScope
)
