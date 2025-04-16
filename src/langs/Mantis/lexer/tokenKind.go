package lexer

import "fmt"

type (
	MantisTokenKind int
)

const (
	SUBSET_0 = 0
	SUBSET_1 = 1
	SUBSET_MAX = SUBSET_1
)

const (
	EOF MantisTokenKind = (iota + 1) * 100 + SUBSET_0
	UNKNOW
	BREAK_LINE
	TAB
	SPACE
	ID
	NUMBER
	UNDERLINE
	COMMA
	COLON
	SEMICOLON
	SLASH
	L_PAREN
	R_PAREN
	L_BRACE
	R_BRACE
	EQUAL
	LESS_THEN
	GREATER_THEN 
	ADD
	SUB
	MUL
	KEY_FUN
	KEY_FOR
	KEY_IF
	KEY_ELSE
	KEY_VAR
	KEY_RETURN = (iota + 1) * 100 + SUBSET_1
	KEY_MATCH
	KEY_WHEN
	KEY_REPEAT
	KEY_IN
)

func (this* MantisTokenKind) String() (s string) {
	switch *this {
	case EOF:
	s = "EOF"
	case BREAK_LINE:
	s = "BREAK_LINE"
	case TAB:
	s = "TAB"
	case SPACE:
	s = "SPACE"
	case ID:
	s = "ID"
	case NUMBER:
	s = "NUMBER"
	case UNDERLINE:
	s = "UNDERLINE"
	case COMMA:
	s = "COMMA"
	case COLON:
	s = "COLON"
	case SEMICOLON:
	s = "SEMICOLON"
	case SLASH:
	s = "SLASH"
	case L_PAREN:
	s = "L_PAREN"
	case R_PAREN:
	s = "R_PAREN"
	case L_BRACE:
	s = "L_BRACE"
	case R_BRACE:
	s = "R_BRACE"
	case EQUAL:
	s = "EQUAL"
	case LESS_THEN:
	s = "LESS_THEN"
	case GREATER_THEN :
	s = "GREATER_THEN"
	case ADD:
	s = "ADD"
	case SUB:
	s = "SUB"
	case MUL:
	s = "MUL"
	case KEY_FUN:
	s = "KEY_FUN"
	case KEY_FOR:
	s = "KEY_FOR"
	case KEY_IF:
	s = "KEY_IF"
	case KEY_ELSE:
	s = "KEY_ELSE"
	case KEY_RETURN:
	s = "KEY_RETURN"
	case KEY_MATCH:
	s = "KEY_MATCH"
	case KEY_WHEN:
	s = "KEY_WHEN"
	case KEY_REPEAT:
	s = "KEY_REPEAT"
	case KEY_IN:
		s = "KEY_IN"
	default:
		s = fmt.Sprintf("UNKNOWN(%d)", *this)
	}
	return s
}