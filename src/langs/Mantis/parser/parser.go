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
		Variables map[string]*MantisVariable
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

	scopeId := uint64(len(parser.Scopes) + 1)
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
			err = errors.Join(err, parser.ParseSingleVarDef(scopeId))
			break

		// Parses a variable definition, assigment or function call
		case lexer.ID:
			if scopeType == utils.RootScope {
				return ret, utils.GetUnexpectedTokenNoPosErr(parser.Filename, string(tk.Value))
			}

			kind, err := parser.GetFirstAfter(lexer.SPACE)
			if err != nil {
				err = errors.Join(err, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF"))
				break
			}

			// Parses single var definition
			if kind == lexer.COLON {
				err = errors.Join(err, parser.ParseSingleVarDef(scopeId))
				break

				// Parses multi var definition
			} else if kind == lexer.COMMA {
				err = errors.Join(err, parser.ParseMultiVarDef(scopeId))
				break

				// Parses assignment
			} else if kind == lexer.EQUAL {
				err = errors.Join(err, parser.ParseAssign(scopeId))
				break

				// Parses augmented assignment
			} else if kind == lexer.ADD || kind == lexer.SUB || kind == lexer.MUL {
				err = errors.Join(err, parser.ParseArgAssign(scopeId))
				break

				// Parses function call
			} else if kind == lexer.L_PAREN {
				err = errors.Join(err, parser.ParseFuncCall(scopeId))
				break

				// Error
			} else {
				err = errors.Join(err, utils.GetExpectedTokenErr(parser.Filename, "some token to create a variable definition, assigment or function call", tk.Pos))
				break
			}

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
