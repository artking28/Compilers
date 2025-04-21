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

func AppendToken(tokens *[]MantisToken, token MantisToken) {
	if tokens == nil {
		tokens = &[]MantisToken{}
	}
	count := len(*tokens)
	if count > 0 && (*tokens)[count-1].Kind == token.Kind && string((*tokens)[count-1].Value) == string(token.Value) {
		(*tokens)[count-1].Repeat = (*tokens)[count-1].Repeat + 1
		return
	}
	*tokens = append(*tokens, token)
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
