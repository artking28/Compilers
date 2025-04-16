package lexer

import (
	"compilers/stdLexer"
	"compilers/utils"
)

type MantisToken stdLexer.Token[MantisTokenKind]

func NewToken(pos utils.Pos, kind MantisTokenKind, repeat int, value ...rune) MantisToken {
	return MantisToken(stdLexer.NewToken(pos, kind, repeat, value...))
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
		default:
		 	return UNKNOW
	}
}

func (this* MantisToken) String() string {
	s := this.Kind.String()
	v := string(this.Literal)
	if this.Kind == BREAK_LINE {
		v = "\\n"
	} else if this.Kind == TAB {
		v = "\\t"
	} else if this.Kind == EOF {
		v = "\\0"
	} else if this.Kind == NUMBER {
		v = strconv.Itoa(int(this.Literal[0]))
	}
	return fmt.Sprintf("Token{%s, \"%s\", %.2d}", s, v, this.Repeat)
}
