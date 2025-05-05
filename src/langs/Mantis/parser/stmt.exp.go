package parser

import (
	"compilers/langs/Mantis/lexer"
	"compilers/stdParser"
	"compilers/utils"
)

var ExpTokens = []lexer.MantisTokenKind{
	lexer.ID,
	lexer.NUMBER,
	lexer.NIL,
	lexer.TRUE,
	lexer.FALSE,
}

var Operators = map[lexer.MantisTokenKind]stdParser.Operator{
	lexer.MUL:                stdParser.Mul,
	lexer.MOD:                stdParser.Mod,
	lexer.ADD:                stdParser.Sum,
	lexer.SUB:                stdParser.Sub,
	lexer.SHIFT_LEFT:         stdParser.ShiftL,
	lexer.SHIFT_RIGHT:        stdParser.ShiftR,
	lexer.AND_BIT:            stdParser.AndBit,
	lexer.OR_BIT:             stdParser.OrBit,
	lexer.XOR_BIT:            stdParser.XorBit,
	lexer.LOWER_THEN:         stdParser.LowerThen,
	lexer.LOWER_EQUAL_THEN:   stdParser.LowerEqualThen,
	lexer.GREATER_THEN:       stdParser.GreaterThen,
	lexer.GREATER_EQUAL_THEN: stdParser.GreaterEqualThen,
	lexer.AND_BOOL:           stdParser.BoolAnd,
	lexer.OR_BOOL:            stdParser.BoolOr,
	lexer.XOR_BOOL:           stdParser.BoolXor,
}

type (
	MantisExp struct {
		stdParser.Exp
		MantisStmtBase
	}

	MantisVExp struct {
		stdParser.VExp
		MantisStmtBase
	}
)

func NewMantisVExp(value int, pos utils.Pos, typeOf string, parser *MantisParser) *MantisVExp {
	return &MantisVExp{
		VExp: *stdParser.NewVExp(value, typeOf),
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "MantisVExp",
			Pos:    pos,
		},
	}
}

func NewMantisExpChain(exps []stdParser.IExp, oper stdParser.Operator, pos utils.Pos, parser *MantisParser) *MantisExp {
	return &MantisExp{
		Exp: stdParser.Exp{
			All:  exps,
			Oper: oper,
		},
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "MantisExp",
			Pos:    pos,
		},
	}
}

func (parser *MantisParser) ParseExpression() (stdParser.IExp, string, error) {
	h0 := parser.Get(0)
	if h0 == nil {
		return nil, "", utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}

	var expList stdParser.IExp
	var expKind = ""
	var lastSig = Operators[lexer.ADD]
	lastWasSig := false

	for i := 0; h0 != nil; i++ {
		if h0.Kind == lexer.ADD || h0.Kind == lexer.SUB {
			for {
				i++
				hs, e := parser.GetFirstAfter(lexer.SPACE)
				if e != nil {
					return nil, "", e
				}
				if hs.Kind == lexer.ADD || hs.Kind == lexer.SUB {
					if hs.Kind == h0.Kind {
						h0.Kind = lexer.ADD
						parser.Consume(1)
						continue
					}
					h0.Kind = lexer.SUB
					lastSig = Operators[h0.Kind]
				} else {
					//lastSig = Operators[h0.Kind]
					//lastWasSig = true
					break
				}
			}
		}

		if h0.Kind == lexer.UNDERLINE {
			parser.Consume(1)
			return nil, "", nil
		} else if h0.Kind == lexer.NUMBER {
			if expList != nil && !lastWasSig {
				return nil, "", utils.GetConsecutiveValuesErr(parser.Filename, h0.Pos)
			}
			lastWasSig = false
			if expKind == "" {
				expKind = "number"
			}
			if expKind != "number" {
				return nil, "", utils.GetMismatchedTypesErr(parser.Filename, "bool", "number", h0.Pos)
			}
			n := int(h0.Value[0])
			if expList == nil {
				expList = NewMantisVExp(n, h0.Pos, expKind, parser)
			} else {
				e, ok := (expList).(*MantisExp)
				if ok {
					e.All = append(e.All, NewMantisVExp(n, h0.Pos, expKind, parser))
				} else {
					expList = NewMantisExpChain([]stdParser.IExp{expList, NewMantisVExp(n, h0.Pos, expKind, parser)}, lastSig, h0.Pos, parser)
				}
			}
		} else if h0.Kind == lexer.TRUE || h0.Kind == lexer.FALSE {
			if expList != nil && !lastWasSig {
				return nil, "", utils.GetConsecutiveValuesErr(parser.Filename, h0.Pos)
			}
			lastWasSig = false
			if expKind == "" {
				expKind = "bool"
			}
			if expKind != "bool" {
				return nil, "", utils.GetMismatchedTypesErr(parser.Filename, "number", "bool", h0.Pos)
			}
			v := 1
			if h0.Kind == lexer.FALSE {
				v = 0
			}
			if expList == nil {
				expList = NewMantisVExp(v, h0.Pos, expKind, parser)
			} else {
				e := (expList).(*MantisExp)
				e.All = append((expList).(MantisExp).All, NewMantisVExp(v, h0.Pos, expKind, parser))
			}
		} else if h0.Kind == lexer.SPACE {
			parser.Consume(1)
			h0 = parser.Get(0)
			continue
		} else if (*lexer.MantisToken)(h0).IsSignal() {
			if expList != nil && lastWasSig {
				return nil, "", utils.GetConsecutiveOperatorsErr(parser.Filename, h0.Pos)
			}
			levelComp := lastSig.Compare(Operators[h0.Kind])
			lastWasSig = true
			if lastSig != Operators[h0.Kind] {
				if levelComp >= 0 {
					if expList != nil && expList.Values() > 1 {
						expList = NewMantisExpChain([]stdParser.IExp{expList}, Operators[h0.Kind], h0.Pos, parser)
					} else if expList != nil && expList.Values() == 1 {
						n, err := expList.Resolve()
						if err != nil {
							return nil, "", err
						}
						expList = NewMantisVExp(n, h0.Pos, expKind, parser)
					} else {
						expList = NewMantisVExp(0, h0.Pos, expKind, parser)
					}
				} else {
					if expList != nil && expList.Values() > 1 {
						e := (expList).(*MantisExp)
						e.All = append((expList).(MantisExp).All, NewMantisVExp(int(h0.Value[0]), h0.Pos, expKind, parser))
					} else if expList != nil && expList.Values() == 1 {
						n, err := expList.Resolve()
						if err != nil {
							return nil, "", err
						}
						expList = NewMantisVExp(n, h0.Pos, expKind, parser)
					} else {
						expList = NewMantisVExp(0, h0.Pos, expKind, parser)
					}
				}
			}
			lastSig = Operators[h0.Kind]
		} else {
			return expList, expKind, nil
		}

		parser.Consume(1)
		h0 = parser.Get(0)
	}
	return expList, expKind, nil
}
