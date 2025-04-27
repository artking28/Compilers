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

func (this StmtBase) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | StmtBase@WriteMemASM")
}

func (this StmtBase) GetTitle() string {
	return this.Title
}
