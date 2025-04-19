package parser

import (
	"compilers/langs/Mantis/lexer"
	"compilers/stdParser"
	"compilers/utils"
)

type MantisVariable struct {
	stdParser.Variable
	MantisStmtBase
}

func (this MantisVariable) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | VariableStmt@WriteMemASM")
}

func NewVariableStmt(name string, pos utils.Pos, value stdParser.IExp, owner uint64, parser *MantisParser) *MantisVariable {
	ret := &MantisVariable{
		Variable: *stdParser.NewVariable(uint64(len(parser.Variables)+1), name, value, owner),
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "VariableStmt",
			Pos:    pos,
		},
	}
	parser.Variables[name] = ret
	return ret
}

func (this MantisVariable) GetTitle() string {
	return this.Title
}

func (parser *MantisParser) ParseSingleVarDef(scopeId uint64) (ret *MantisVariable, err error) {
	waitColon, nameTk := true, parser.Get(0)
	if nameTk == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	if nameTk.Kind == lexer.KEY_VAR {
		parser.Consume(1)
		waitColon = false
	}
	nameTk, err = parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, lexer.ID)
	if nameTk == nil {
		return nil, utils.GetExpectedTokenErr(parser.Filename, "variable name", parser.At())
	}
	if waitColon {
		if _, err = parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, lexer.COLON); err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "colon token", parser.At())
		}
	}
	if _, err = parser.HasNextConsume(stdParser.NoSpaceMode, lexer.SPACE, lexer.EQUAL); err != nil {
		return nil, utils.GetExpectedTokenErr(parser.Filename, "assign token", parser.At())
	}
	if _, err = parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, ExpTokens...); err != nil {
		return nil, utils.GetExpectedTokenErr(parser.Filename, "expression value", parser.At())
	}
	value, err := ParseExpressionReturn(lexer.BREAK_LINE, lexer.SEMICOLON)
	if err != nil {
		return nil, err
	}
	return NewVariableStmt(string(nameTk.Value), parser.At(), value, scopeId, parser), nil
}

func (parser *MantisParser) ParseMultiVarDef(scopeId uint64) (*[]MantisVariable, error) {
	waitColon, nameTk := true, parser.Get(0)
	if nameTk == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	if nameTk.Kind == lexer.KEY_VAR {
		parser.Consume(1)
		waitColon = false
	}
	kind, err := parser.GetFirstAfter(lexer.SPACE)
	if kind != lexer.ID || err != nil {
		return nil, utils.GetExpectedTokenErr(parser.Filename, "variable name", parser.At())
	}

	var names []string
	var pos []utils.Pos
	var values []stdParser.IExp
	for first := false; true; {
		nameTk, err = parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, lexer.ID, lexer.COLON)
		if err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "variable name", parser.At())
		}
		if nameTk.Kind == lexer.COLON {
			if !first {
				return nil, utils.GetExpectedTokenErr(parser.Filename, "at least one variable name", parser.At())
			}
			if !waitColon {
				return nil, utils.GetExpectedTokenErr(parser.Filename, "assign token", parser.At())
			}
			break
		}

		first = true
		names = append(names, string(nameTk.Value))
		pos = append(pos, nameTk.Pos)
		end, err := parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, lexer.COMMA, lexer.EQUAL, lexer.COLON)
		if err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "comma", parser.At())
		}
		if end.Kind == lexer.EQUAL || end.Kind == lexer.COLON {
			parser.Consume(1)
			break
		}
	}
	for first := false; true; {
		exp, err := parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, append(ExpTokens, lexer.UNDERLINE)...)
		if err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "variable name", parser.At())
		}
		if exp.Kind == lexer.UNDERLINE {
			if !first {
				return nil, utils.GetExpectedTokenErr(parser.Filename, "at least one variable value", parser.At())
			}
		}

		first = true
		nextKind, err := parser.GetFirstAfter(lexer.SPACE)
		if err != nil {
			return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
		}
		if nextKind != lexer.COMMA {
			value, err := ParseExpressionReturn(lexer.BREAK_LINE, lexer.SEMICOLON, lexer.COMMA)
			if err != nil {
				return nil, err
			}
			values = append(values, value)
			h0 := parser.Get(0)
			if h0 == nil {
				return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
			}
			if h0.Kind == lexer.SEMICOLON || h0.Kind == lexer.BREAK_LINE {
				break
			}
		}
		values = append(values, NewMantisVExp(int(exp.Value[0]), parser.At(), parser))
		parser.Consume(1)
	}
	if len(values) > len(names) {
		return nil, utils.GetTooManyValuesErr(parser.Filename, parser.At().Line)
	}
	afterEqual := len(values) - 1
	var ret []MantisVariable
	for i := 0; i < len(names); i++ {
		if i > afterEqual {
			ret = append(ret, *NewVariableStmt(names[i], pos[i], values[afterEqual], scopeId, parser))
			continue
		}
		ret = append(ret, *NewVariableStmt(names[i], pos[i], values[i], scopeId, parser))
	}
	return &ret, nil
}
