package parser

import (
	"compilers/langs/Mantis/lexer"
	"compilers/stdLexer"
	"compilers/stdParser"
	"compilers/utils"
	"errors"
)

type (
	MantisParser struct {
		Subset int
		stdParser.Parser[lexer.MantisTokenKind]
	}

	MantisStmtBase struct {
		Parser *MantisParser `json:"-"`
		Title  string        `json:"title"`
		Pos    utils.Pos     `json:"pos"`
	}
)

func NewMantisParser(filename, output string, subset int) (*MantisParser, error) {

	tokens, err := lexer.Tokenize(filename)
	if err != nil {
		return nil, err
	}

	p, err := stdParser.NewParser[lexer.MantisTokenKind](filename, output, any(tokens).([]stdLexer.Token[lexer.MantisTokenKind]))
	ret := any(p).(*MantisParser)
	ret.Subset = subset
	return ret, err
}

func (parser *MantisParser) ParseScope(scopeType utils.ScopeType) (ret stdParser.Scope, err error) {

	tk := parser.Get(0)
	for tk != nil && tk.Kind != lexer.EOF {

		// Parses some statement on root context of the file
		switch tk.Kind {

		// Parses a comment section
		case lexer.SLASH:
			err = errors.Join(err, parser.ParseComment())
			break

		// Parses a function
		case lexer.KEY_FUN:
			err = errors.Join(err, parser.ParseFunction())
			break

		// Parses a global variable
		case lexer.KEY_VAR:
			err = errors.Join(err, parser.ParseVariable())
			break

		// Parses a variable definition, assigment or function call
		case lexer.ID:
			if scopeType == utils.RootScope {
				return ret, utils.GetUnexpectedTokenNoPosErr(parser.Filename, string(tk.Value))
			}

			break

		// Parses an if statement
		case lexer.KEY_IF:
			if scopeType == utils.RootScope {
				return ret, utils.GetUnexpectedIfStatementInRoot(parser.Filename, tk.Pos)
			}
			err = errors.Join(err, parser.ParseIfStatement())
			break

		// Parses a for loop statement
		case lexer.KEY_FOR:
			if scopeType == utils.RootScope {
				return ret, utils.GetUnexpectedForLoopStatementInRoot(parser.Filename, tk.Pos)
			}
			err = errors.Join(err, parser.ParseForLoopStatement())
			break

		// Ends any kind of AST structure calling the scope parse
		case lexer.R_BRACE:
			return ret, err

		// Default handler
		default:
			break
		}

		if err != nil {
			return ret, err
		}

		// Advances the parser cursor and update latest token
		parser.Consume(1)
		tk = parser.Get(0)
	}
	return ret, err
}
