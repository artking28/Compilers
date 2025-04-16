package stdParser

import "compilers/utils"

type (
	Stmt interface {
		WriteMemASM() ([]uint16, error)
		GetTitle() string
	}
	
	StmtBase struct {
		Title  string  `json:"title"`
		Pos    utils.Pos     `json:"pos"`
	}
)