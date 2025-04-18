package stdParser

import "compilers/utils"

type (
	Stmt interface {
		WriteMemASM() ([]uint16, error)
		GetTitle() string
	}

	StmtBase struct {
		Parser *Parser[any] `json:"-"`
		Title  string       `json:"title"`
		Pos    utils.Pos    `json:"pos"`
	}
)
