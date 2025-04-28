package lexer

import (
	"compilers/stdLexer"
	"compilers/utils"
	"fmt"
	"strconv"
)

type MantisToken stdLexer.Token[MantisTokenKind]

func NewToken(pos utils.Pos, kind MantisTokenKind, repeat int, value ...rune) MantisToken {
	return MantisToken(stdLexer.NewToken(pos, kind, repeat, value...))
}

func (this MantisToken) IsSignal() bool {
	return this.Kind == MUL ||
		this.Kind == MOD ||
		this.Kind == ADD ||
		this.Kind == SUB ||
		this.Kind == SHIFT_LEFT ||
		this.Kind == SHIFT_RIGHT ||
		this.Kind == AND_BIT ||
		this.Kind == OR_BIT ||
		this.Kind == XOR_BIT ||
		this.Kind == LOWER_THEN ||
		this.Kind == LOWER_EQUAL_THEN ||
		this.Kind == GREATER_THEN ||
		this.Kind == GREATER_EQUAL_THEN ||
		this.Kind == AND_BOOL ||
		this.Kind == OR_BOOL ||
		this.Kind == XOR_BOOL
}

func AppendToken(filename string, tokens *[]MantisToken, token MantisToken, subset int) error {
	if tokens == nil {
		tokens = &[]MantisToken{}
	}
	count := len(*tokens)
	if count > 0 {
		last := (*tokens)[count-1]
		if last.Kind == token.Kind && string(last.Value) == string(token.Value) {
			(*tokens)[count-1].Repeat = last.Repeat + 1
			last = (*tokens)[count-1]
		}
		if c := CombineTokens(last, token); c != UNKNOW {
			(*tokens)[count-1].Kind = c
			return nil
		}
	}
	if int(token.Kind)%100 > subset {
		return utils.GetInvalidTokenPerSubset(filename, string(token.Value), token.Pos)
	}
	*tokens = append(*tokens, token)
	return nil
}

func ResolveTokenId(filename string, token MantisToken) (MantisToken, error) {
	if token.Kind != ID {
		return token, nil
	}
	value := string(token.Value)

	if value == "_" {
		return NewToken(token.Pos, UNDERLINE, 1, token.Value...), nil
	}

	if tk := FindKeyword(value); tk != UNKNOW {
		return NewToken(token.Pos, tk, 1, token.Value...), nil
	}

	if n, err := strconv.ParseInt(value, 0, 64); err == nil {
		return NewToken(token.Pos, NUMBER, 1, []rune{rune(n)}...), nil
	}

	return token, nil
}

func FindKeyword(word string) MantisTokenKind {
	switch {
	case word == "fun":
		return KEY_FUN
	case word == "for":
		return KEY_FOR
	case word == "if":
		return KEY_IF
	case word == "else":
		return KEY_ELSE
	case word == "var":
		return KEY_VAR
	case word == "return":
		return KEY_RETURN
	case word == "repeat":
		return KEY_REPEAT
	case word == "match":
		return KEY_MATCH
	case word == "when":
		return KEY_WHEN
	case word == "in":
		return KEY_IN
	case word == "null":
		return NIL
	case word == "nil":
		return NIL
	case word == "true":
		return TRUE
	case word == "false":
		return FALSE
	default:
		return UNKNOW
	}
}

func (this *MantisToken) String() string {
	s := this.Kind.String()
	v := string(this.Value)
	if this.Kind == BREAK_LINE {
		v = "\\n"
	} else if this.Kind == TAB {
		v = "\\t"
	} else if this.Kind == EOF {
		v = "\\0"
	} else if this.Kind == NUMBER {
		v = strconv.Itoa(int(this.Value[0]))
	}
	return fmt.Sprintf("Token{%s, \"%s\", %.2d}", s, v, this.Repeat)
}
