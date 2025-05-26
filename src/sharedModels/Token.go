package models

import (
	"compilers/utils"
	"fmt"
	"strconv"
)

type Token struct {
	Pos    utils.Pos `json:"-"`
	Kind   TokenKind `json:"kind"`
	Value  []rune    `json:"value"`
	Repeat int       `json:"repeat"`
}

func NewToken(pos utils.Pos, kind TokenKind, repeat int, value ...rune) Token {
	return Token{Pos: pos, Kind: kind, Value: value, Repeat: repeat}
}

func (this Token) IsSignal() bool {
	return this.Kind == ADD ||
		this.Kind == SUB ||
		this.Kind == MUL ||
		this.Kind == DIV ||
		this.Kind == MOD ||
		this.Kind == AND_BIT ||
		this.Kind == OR_BIT ||
		this.Kind == XOR_BIT ||
		this.Kind == SHIFT_LEFT ||
		this.Kind == SHIFT_RIGHT ||
		this.Kind == LOWER_THEN ||
		this.Kind == LOWER_EQUAL_THEN ||
		this.Kind == GREATER_THEN ||
		this.Kind == GREATER_EQUAL_THEN ||
		this.Kind == EQUAL ||
		this.Kind == AND_BOOL ||
		this.Kind == OR_BOOL ||
		this.Kind == XOR_BOOL
}

func ResolveTokenId(filename string, token Token) (Token, error) {
	if token.Kind != ID {
		return token, nil
	}
	value := string(token.Value)

	if len(token.Value) == 1 && token.Value[0] == '_' {
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

func FindKeyword(word string) TokenKind {
	switch word {
	case "true":
		return TRUE
	case "false":
		return FALSE
	case "nil":
		return NIL
	case "null":
		return NIL
	case "fun":
		return KEY_FUN
	case "for":
		return KEY_FOR
	case "if":
		return KEY_IF
	case "else":
		return KEY_ELSE
	case "var":
		return KEY_VAR
	case "break":
		return KEY_BREAK
	default:
		return UNKNOW
	}
}

func (this *Token) String() string {
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
