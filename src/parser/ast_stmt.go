package parser

import (
	"compilers/utils"
)

type (
	Stmt interface {
		WriteMemASM() (string, error)
		GetTitle() string
	}

	StmtBase struct {
		Parser *MantisParser `json:"-"`
		Title  string        `json:"title"`
		Pos    utils.Pos     `json:"pos"`
	}
)

func (this StmtBase) GetTitle() string {
	return this.Title
}
