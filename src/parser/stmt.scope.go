package parser

import (
	models "compilers/sharedModels"
	"compilers/utils"
	"errors"
)

func (parser *MantisParser) ParseScope(scopeType ScopeKind) (ret Scope, err error) {

	ret.Kind = scopeType
	scopeId := uint64(len(parser.Scopes) + 1)
	tk := parser.Get(0)
	for tk != nil && tk.Kind != models.EOF {

		// Parses some statement on root context of the file
		switch tk.Kind {

		// Parses a comment section
		case models.COMMENT_LINE:
			c, e := parser.ParseComment()
			err = errors.Join(e)
			ret.Body.Statements = append(ret.Body.Statements, c)
			break

		// Parses a function
		case models.KEY_FUN:
			f, e := parser.ParseFunction()
			err = errors.Join(e)
			ret.Body.Statements = append(ret.Body.Statements, f)
			break

		// Parses a global variable
		case models.KEY_VAR:
			t, e0 := parser.GetFirstAfter(models.SPACE, models.KEY_VAR)
			if e0 != nil {
				err = errors.Join(err, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF"))
				break
			}

			if t.Kind == models.ID {
				svd, e := parser.ParseSingleVarDef(scopeId)
				err = errors.Join(e)
				ret.Body.Statements = append(ret.Body.Statements, svd)
				break

				// Parses multi var definition
			} else if t.Kind == models.COMMA {
				mvd, e := parser.ParseMultiVarDef(scopeId)
				err = errors.Join(e)
				if mvd != nil {
					for _, svd := range *mvd {
						ret.Body.Statements = append(ret.Body.Statements, svd)
					}
				}
				break
			}

			err = errors.Join(err, utils.GetExpectedTokenErr(parser.Filename, "some token to create a variable definition, an assigment or function call", tk.Pos))
			break

		// Parses a variable definition, assigment or function call
		case models.ID:
			if scopeType == RootScope {
				return ret, utils.GetUnexpectedTokenNoPosErr(parser.Filename, string(tk.Value))
			}

			t, e0 := parser.GetFirstAfter(models.SPACE, models.ID)
			if e0 != nil {
				err = errors.Join(err, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF"))
				break
			}

			// Parses single var definition
			if t.Kind == models.INIT {
				svd, e := parser.ParseSingleVarDef(scopeId)
				err = errors.Join(e)
				ret.Body.Statements = append(ret.Body.Statements, svd)
				break

				// Parses multi var definition
			} else if t.Kind == models.COMMA {
				mvd, e := parser.ParseMultiVarDef(scopeId)
				err = errors.Join(e)
				if mvd != nil {
					for _, svd := range *mvd {
						ret.Body.Statements = append(ret.Body.Statements, svd)
					}
				}
				break

				// Parses assignments
			} else if t.Kind == models.ASSIGN ||
				t.Kind == models.ASSIGN_ADD ||
				t.Kind == models.ASSIGN_SUB ||
				t.Kind == models.ASSIGN_MUL ||
				t.Kind == models.ASSIGN_MOD ||
				t.Kind == models.ASSIGN_AND_BIT ||
				t.Kind == models.ASSIGN_XOR_BIT ||
				t.Kind == models.ASSIGN_OR_BIT ||
				t.Kind == models.ASSIGN_SHIFT_RIGHT ||
				t.Kind == models.ASSIGN_SHIFT_LEFT {

				assignStmt, e := parser.ParseArgAssign(scopeId, t.Kind)
				err = errors.Join(err, e)
				ret.Body.Statements = append(ret.Body.Statements, assignStmt)
				break

				// Parses function call
			} else if t.Kind == models.L_PAREN {
				fc, e := parser.ParseFuncCall(scopeId)
				err = errors.Join(e)
				ret.Body.Statements = append(ret.Body.Statements, fc)
				break

				// Error
			}
			err = errors.Join(err, utils.GetExpectedTokenErr(parser.Filename, "some token to create a variable definition, an assigment or function call", tk.Pos))
			break

		// Parses an if statement
		case models.KEY_IF:
			if scopeType == RootScope {
				return ret, utils.GetUnexpectedIfStatementInRoot(parser.Filename, tk.Pos)
			}
			fc, e := parser.ParseIfStatement()
			err = errors.Join(e)
			ret.Body.Statements = append(ret.Body.Statements, fc)
			break

		// Parses a for loop statement
		case models.KEY_FOR:
			if scopeType == RootScope {
				return ret, utils.GetUnexpectedForLoopStatementInRoot(parser.Filename, tk.Pos)
			}
			fc, e := parser.ParseForLoopStatement()
			err = errors.Join(e)
			ret.Body.Statements = append(ret.Body.Statements, fc)
			break

		// Ends any kind of AST structure calling the scope parse
		case models.R_BRACE:
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

func (this *MantisParser) At() utils.Pos {
	return this.Tokens[this.Cursor].Pos
}

func (this *MantisParser) Get(n int) *models.Token {
	if this.Cursor+n >= len(this.Tokens) {
		return nil
	}
	return &this.Tokens[this.Cursor+n]
}

func (this *MantisParser) Consume(n int) {
	if this.Cursor+n >= len(this.Tokens) {
		return
	}
	this.Cursor += n
}

func (this *MantisParser) GetFirstAfter(afterOf ...models.TokenKind) (*models.Token, error) {
	all := map[models.TokenKind]bool{}
	for _, t := range afterOf {
		all[t] = true
	}

	token := this.Get(1)
	for i := 1; token != nil; i++ {
		if all[token.Kind] {
			token = this.Get(i)
			continue
		}
		return token, nil
	}
	return nil, errors.New("no token has been found")
}

const (
	NoSpaceMode = iota
	OptionalSpaceMode
	MandatorySpaceMode
)

func (this *MantisParser) HasNextConsume(spaceMode int, fillOf models.TokenKind, kinds ...models.TokenKind) (*models.Token, error) {
	if spaceMode < NoSpaceMode || spaceMode > MandatorySpaceMode {
		panic("invalid argument in function 'HasNextConsume', unknown space mode")
	}
	if len(kinds) <= 0 {
		panic("invalid argument in function 'HasNextConsume', kinds is null or empty")
	}
	for hasSpace := false; ; {
		token := this.Get(0)
		if token == nil {
			// Fim dos tokens sem encontrar um tipo esperado
			return nil, errors.New("no token has been found")
		}

		for _, kind := range kinds {
			if token.Kind == kind {
				// Se espaços eram obrigatórios mas não foram encontrados, falha
				if spaceMode == MandatorySpaceMode && !hasSpace {
					return nil, errors.New("rule expects spaces but none has been found")
				}
				this.Consume(1)
				return token, nil
			}
		}

		if token.Kind == fillOf {
			// Se espaços não eram permitidos
			if spaceMode == NoSpaceMode {
				return nil, errors.New("space(s) has been found when it actually isn't allowed here")
			}
			hasSpace = true
			this.Consume(1)
			continue
		}

		// Se espaços eram obrigatórios e encontrou outro token, falha
		if spaceMode == MandatorySpaceMode {
			return nil, errors.New("rule expects spaces but none has been found")
		}

		return nil, errors.New("unknown error") // Qualquer outro caso não esperado falha
	}
}
