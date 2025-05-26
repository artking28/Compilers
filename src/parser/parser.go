package parser

import (
	"compilers/lexer"
	models "compilers/sharedModels"
	"compilers/utils"
)

type (
	MantisParser struct {
		Filename       string
		OutputFile     string
		Scopes         map[uint64]*Scope
		Cursor         int
		VariablesIDs   map[uint64]*Variable
		VariablesNames map[string]*Variable
		Tokens         []models.Token
	}

	MantisStmtBase struct {
		Parser *MantisParser `json:"-"`
		Title  string        `json:"title"`
		Pos    utils.Pos     `json:"pos"`
	}
)

func (this MantisStmtBase) GetTitle() string {
	return this.Title
}

func NewMantisParser(filename, output string) (*MantisParser, error) {

	tokens, err := lexer.Tokenize(filename)
	if err != nil {
		return nil, err
	}

	ret := MantisParser{
		Tokens:         tokens,
		VariablesIDs:   map[uint64]*Variable{},
		VariablesNames: map[string]*Variable{},
		Filename:       filename,
		OutputFile:     output,
		Scopes:         map[uint64]*Scope{},
		Cursor:         0,
	}

	var ts []models.Token
	for _, t := range tokens {
		ts = append(ts, models.Token(t))
	}
	return &ret, err
}
