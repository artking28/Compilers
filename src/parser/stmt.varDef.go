package parser

import (
	models "compilers/sharedModels"
	"compilers/utils"
)

func (parser *MantisParser) ParseSingleVarDef(scopeId uint64) (ret *Variable, err error) {
	waitColon, nameTk := true, parser.Get(0)
	if nameTk == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	if nameTk.Kind == models.KEY_VAR {
		parser.Consume(1)
		waitColon = false
	}
	nameTk, err = parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.ID)
	if nameTk == nil {
		return nil, utils.GetExpectedTokenErr(parser.Filename, "variable name", parser.At())
	}
	if waitColon {
		if _, err = parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.INIT); err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "colon token", parser.At())
		}
	} else {
		if _, err = parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.ASSIGN); err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "assign token", parser.At())
		}
	}
	parser.Consume(1)
	value, err := parser.ParseExpression(false)
	if err != nil {
		return nil, err
	}

	return NewVariable(string(nameTk.Value), value, parser), nil
}

func (parser *MantisParser) ParseMultiVarDef(scopeId uint64) (*[]Variable, error) {
	waitColon, nameTk := true, parser.Get(0)
	if nameTk == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	if nameTk.Kind == models.KEY_VAR {
		parser.Consume(1)
		waitColon = false
	}
	//t, err := parser.GetFirstAfter(models.SPACE)
	//if t == nil {
	//	return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	//}
	//if t.Kind != models.ID || err != nil {
	//	return nil, utils.GetExpectedTokenErr(parser.Filename, "variable name", parser.At())
	//}

	var names []string
	var pos []utils.Pos
	var values []IExp
	for first := false; true; {
		nameTk, err := parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.ID, models.INIT)
		if err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "variable name", parser.At())
		}
		if nameTk.Kind == models.INIT {
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
		end, err := parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.COMMA, models.ASSIGN, models.INIT)
		if err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "comma", parser.At())
		}
		if end.Kind == models.ASSIGN || end.Kind == models.INIT {
			if end.Kind == models.INIT && !waitColon {
				return nil, utils.GetExpectedTokenErr(parser.Filename, "assign token", parser.At())
			}
			parser.Consume(1)
			break
		}
	}

	for first := true; true; {
		exp, err := parser.ParseExpression(false)
		if err != nil {
			return nil, err
		}
		if exp != nil {
			values = append(values, exp)
		} else {
			if first {
				return nil, utils.GetExpectedTokenErr(parser.Filename, "at least one variable value", parser.At())
			}
			values = append(values, NewVExp(0))
		}

		first = false
		tk, err := parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.COMMA, models.BREAK_LINE)
		if err != nil {
			return nil, utils.GetExpectedTokenErr(parser.Filename, "comma or break line", parser.At())
		}
		if tk.Kind == models.BREAK_LINE {
			break
		}
		parser.Consume(1)
	}
	if len(values) > len(names) {
		return nil, utils.GetTooManyValuesErr(parser.Filename, parser.At().Line)
	}
	afterEqual := len(values) - 1
	var ret []Variable
	for i := 0; i < len(names); i++ {
		if i > afterEqual {
			ret = append(ret, *NewVariable(names[i], values[afterEqual], parser))
			continue
		}
		ret = append(ret, *NewVariable(names[i], values[i], parser))
	}
	return &ret, nil
}
