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

func (this *MantisVExp) Resolve() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (this *MantisExp) Resolve() (int, error) {
	//TODO implement me
	panic("implement me | MantisExp@WriteMemASM")
}

func (this MantisVExp) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | MantisVExp@WriteMemASM")
}

func (this MantisVExp) GetTitle() string {
	return this.Title
}

func (this MantisExp) WriteMemASM() ([]uint16, error) {
	//TODO implement me
	panic("implement me | MantisExp@WriteMemASM")
}

func (this MantisExp) GetTitle() string {
	return this.Title
}

func NewMantisVExp(value int, pos utils.Pos, parser *MantisParser) *MantisVExp {
	return &MantisVExp{
		VExp: *stdParser.NewVExp(value),
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

func (parser *MantisParser) ParseExpression(endAt ...lexer.MantisTokenKind) (stdParser.IExp, error) {
	h0 := parser.Get(0)
	if h0 == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}

	var expList MantisExp
	var expKind = ""
	var lastSig = stdParser.Operator(0)
	lastWasSig := false

	for i := 0; h0 != nil; i++ {
		if h0.Kind == lexer.ADD || h0.Kind == lexer.SUB {
			for {
				i++
				hs, e := parser.GetFirstAfter(lexer.SPACE)
				if e != nil {
					return nil, e
				}
				if hs.Kind == lexer.ADD || hs.Kind == lexer.SUB {
					if hs.Kind == h0.Kind {
						h0.Kind = lexer.ADD
						continue
					}
					h0.Kind = lexer.SUB
				} else {
					lastSig = Operators[h0.Kind]
					break
				}
			}

			if h0.Kind == lexer.NUMBER {
				if !lastWasSig {
					return nil, utils.GetConsecutiveValuesErr(parser.Filename, h0.Pos)
				}
				lastWasSig = false
				if expKind == "" {
					expKind = "number"
				}
				if expKind != "number" {
					return nil, utils.GetMismatchedTypesErr(parser.Filename, "bool", "number", h0.Pos)
				}
				expList.All = append(expList.All, NewMantisVExp(int(h0.Value[0]), h0.Pos, parser))
			} else if h0.Kind == lexer.TRUE || h0.Kind == lexer.FALSE {
				if !lastWasSig {
					return nil, utils.GetConsecutiveValuesErr(parser.Filename, h0.Pos)
				}
				lastWasSig = false
				if expKind == "" {
					expKind = "bool"
				}
				if expKind != "bool" {
					return nil, utils.GetMismatchedTypesErr(parser.Filename, "number", "bool", h0.Pos)
				}
				v := 1
				if h0.Kind == lexer.FALSE {
					v = 0
				}
				expList.All = append(expList.All, NewMantisVExp(v, h0.Pos, parser))
			} else if h0.Kind == lexer.MUL || h0.Kind == lexer.SHIFT_LEFT || h0.Kind == lexer.SHIFT_RIGHT {
				if lastWasSig {
					return nil, utils.GetConsecutiveOperatorsErr(parser.Filename, h0.Pos)
				}
				levelComp := lastSig.Compare(Operators[h0.Kind])
				lastWasSig = true
				if levelComp >= 0 {
					expList = *NewMantisExpChain([]stdParser.IExp{&expList}, Operators[h0.Kind], h0.Pos, parser)
				} else {
					expList.All = append(expList.All, NewMantisVExp(int(h0.Value[0]), h0.Pos, parser))
				}
			} else {
				return nil, utils.GetUnexpectedTokenErr(parser.Filename, string(h0.Value), h0.Pos)
			}
		}
		h0 = parser.Get(i)
	}
	return &expList, nil
}
