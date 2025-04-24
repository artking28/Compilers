package parser

import (
	"compilers/langs/Mantis/lexer"
	"compilers/utils"
)

type CommentStmt struct {
	Value string `json:"value"`
	MantisStmtBase
}

func NewCommentStmt(content string, pos utils.Pos, parser *MantisParser) *CommentStmt {
	return &CommentStmt{
		Value: content,
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "CommentStmt",
			Pos:    pos,
		},
	}
}

func (this CommentStmt) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | CommentStmt@WriteMemASM")
}

func (this CommentStmt) GetTitle() string {
	return this.Title
}

func (parser *MantisParser) ParseComment() (*CommentStmt, error) {
	var comment string
	h0 := parser.Get(0)
	if h0 == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	if h0.Kind != lexer.COMMENT_LINE {
		return nil, utils.GetUnexpectedTokenErr(parser.Filename, string(h0.Value), h0.Pos)
	}
	parser.Consume(1)
	for here := parser.Get(0); here != nil && here.Kind != lexer.BREAK_LINE; here = parser.Get(0) {
		comment += string(here.Value)
		parser.Consume(1)
	}
	return NewCommentStmt(comment, h0.Pos, parser), nil
}
