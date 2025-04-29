package parser

import (
	"compilers/langs/Mantis"
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
		if _, err = parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, lexer.INIT); err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "colon token", parser.At())
		}
	} else {
		if _, err = parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, lexer.ASSIGN); err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "assign token", parser.At())
		}
	}
	parser.Consume(1)
	value, _, err := parser.ParseExpression()
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
	t, err := parser.GetFirstAfter(lexer.SPACE)
	if t == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	if t.Kind != lexer.ID || err != nil {
		return nil, utils.GetExpectedTokenErr(parser.Filename, "variable name", parser.At())
	}
	//parser.Consume(1)
	var names []string
	var pos []utils.Pos
	var values []stdParser.IExp
	for first := false; true; {
		nameTk, err = parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, lexer.ID, lexer.INIT)
		if err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "variable name", parser.At())
		}
		if nameTk.Kind == lexer.INIT {
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
		end, err := parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, lexer.COMMA, lexer.ASSIGN, lexer.INIT)
		if err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "comma", parser.At())
		}
		if end.Kind == lexer.ASSIGN || end.Kind == lexer.INIT {
			if end.Kind == lexer.INIT && !waitColon {
				return nil, utils.GetExpectedTokenErr(parser.Filename, "assign token", parser.At())
			}
			parser.Consume(1)
			break
		}
	}

	last := ""
	for first := true; true; {
		exp, typeOf, err := parser.ParseExpression()
		if err != nil {
			return nil, err
		}
		if exp != nil {
			values = append(values, exp)
			last = typeOf
		} else {
			if first {
				return nil, utils.GetExpectedTokenErr(parser.Filename, "at least one variable value", parser.At())
			}
			exp = NewMantisVExp((mantis.GetZeroValue(last)).(int), parser.At(), last, parser)
			values = append(values, exp)
		}

		first = false
		tk, err := parser.HasNextConsume(stdParser.OptionalSpaceMode, lexer.SPACE, lexer.COMMA, lexer.BREAK_LINE)
		if err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "comma or break line", parser.At())
		}
		if tk.Kind == lexer.BREAK_LINE {
			break
		}
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
