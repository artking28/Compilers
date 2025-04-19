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
		Tokens    []lexer.MantisToken
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

	//ret := MantisParser{
	//	Subset:    subset,
	//	Tokens:    tokens,
	//	Variables: nil,
	//	Parser: stdParser.Parser[lexer.MantisTokenKind]{
	//		Filename:   filename,
	//		OutputFile: output,
	//		Scopes:     map[uint64]*stdParser.Scope{},
	//		Output:     stdParser.Ast{},
	//		Cursor:     0,
	//	},
	//}
	//return &ret, err

	var ts []stdLexer.Token[lexer.MantisTokenKind]
	for _, t := range tokens {
		ts = append(ts, stdLexer.Token[lexer.MantisTokenKind](t))
	}
	p, err := stdParser.NewParser[lexer.MantisTokenKind](filename, output, ts)
	if err != nil {
		return nil, err
	}
	ret := &MantisParser{}
	ret.Variables = map[string]*MantisVariable{}
	ret.Parser = *p
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
			c, e := parser.ParseComment()
			err = errors.Join(e)
			ret.Body.Statements = append(ret.Body.Statements, c)
			break

		// Parses a function
		case lexer.KEY_FUN:
			f, e := parser.ParseFunction()
			err = errors.Join(e)
			ret.Body.Statements = append(ret.Body.Statements, f)
			break

		// Parses a global variable
		case lexer.KEY_VAR:
			vd, e := parser.ParseSingleVarDef(scopeId)
			err = errors.Join(e)
			ret.Body.Statements = append(ret.Body.Statements, vd)
			break

		// Parses a variable definition, assigment or function call
		case lexer.ID:
			if scopeType == utils.RootScope {
				return ret, utils.GetUnexpectedTokenNoPosErr(parser.Filename, string(tk.Value))
			}
			parser.Consume(1)
			kind, e0 := parser.GetFirstAfter(lexer.SPACE)
			if e0 != nil {
				err = errors.Join(err, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF"))
				break
			}

			parser.Consume(-1)
			// Parses single var definition
			if kind == lexer.COLON {
				svd, e := parser.ParseSingleVarDef(scopeId)
				err = errors.Join(e)
				ret.Body.Statements = append(ret.Body.Statements, svd)
				break

				// Parses multi var definition
			} else if kind == lexer.COMMA {
				mvd, e := parser.ParseMultiVarDef(scopeId)
				err = errors.Join(e)
				for _, svd := range *mvd {
					ret.Body.Statements = append(ret.Body.Statements, svd)
				}
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
				fc, e := parser.ParseFuncCall(scopeId)
				err = errors.Join(e)
				ret.Body.Statements = append(ret.Body.Statements, fc)
				break

				// Error
			} else {
				err = errors.Join(err, utils.GetExpectedTokenErr(parser.Filename, "some token to create a variable definition, an assigment or function call", tk.Pos))
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
