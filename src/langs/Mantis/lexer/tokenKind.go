package lexer

import (
	"fmt"
)

type (
	MantisTokenKind int
)

const (
	SUBSET_0   = 0
	SUBSET_1   = 1
	SUBSET_MAX = SUBSET_1
)

var Subsets = map[int]bool{
	SUBSET_0: true,
	SUBSET_1: true,
}

const (
	EOF MantisTokenKind = (iota+1)*100 + SUBSET_0
	UNKNOW
	BREAK_LINE
	TAB
	SPACE
	ID
	NIL
	TRUE
	FALSE
	NUMBER
	UNDERLINE
	COMMA
	COLON
	SEMICOLON
	SLASH
	COMMENT_LINE
	L_PAREN
	R_PAREN
	L_BRACE
	R_BRACE
	INIT
	ASSIGN
	EQUAL
	LOWER_THEN
	LOWER_EQUAL_THEN
	SHIFT_LEFT
	ASSIGN_SHIFT_LEFT
	GREATER_THEN
	GREATER_EQUAL_THEN
	SHIFT_RIGHT
	ASSIGN_SHIFT_RIGHT
	MOD
	ASSIGN_MOD
	ADD
	ASSIGN_ADD
	SUB
	ASSIGN_SUB
	AND_BIT
	ASSIGN_AND_BIT
	AND_BOOL
	XOR_BIT
	ASSIGN_XOR_BIT
	XOR_BOOL
	OR_BIT
	ASSIGN_OR_BIT
	OR_BOOL
	KEY_FUN
	KEY_FOR
	KEY_IF
	KEY_ELSE
	KEY_VAR
	KEY_RETURN = (iota+1)*100 + SUBSET_1
	KEY_MATCH
	KEY_WHEN
	KEY_REPEAT
	KEY_IN
	MUL
	ASSIGN_MUL
)

func (this *MantisTokenKind) String() (s string) {
	switch *this {
	case EOF:
		return "EOF"
	case BREAK_LINE:
		return "BREAK_LINE"
	case TAB:
		return "TAB"
	case SPACE:
		return "SPACE"
	case ID:
		return "ID"
	case NIL:
		return "NIL"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case NUMBER:
		return "NUMBER"
	case UNDERLINE:
		return "UNDERLINE"
	case COMMA:
		return "COMMA"
	case COLON:
		return "COLON"
	case SEMICOLON:
		return "SEMICOLON"
	case SLASH:
		return "SLASH"
	case COMMENT_LINE:
		return "COMMENT_LINE"
	case L_PAREN:
		return "L_PAREN"
	case R_PAREN:
		return "R_PAREN"
	case L_BRACE:
		return "L_BRACE"
	case R_BRACE:
		return "R_BRACE"
	case INIT:
		return "INIT"
	case EQUAL:
		return "EQUAL"
	case LOWER_THEN:
		return "LOWER_THEN"
	case LOWER_EQUAL_THEN:
		return "LOWER_EQUAL_THEN"
	case SHIFT_LEFT:
		return "SHIFT_LEFT"
	case GREATER_THEN:
		return "GREATER_THEN"
	case GREATER_EQUAL_THEN:
		return "GREATER_EQUAL_THEN"
	case SHIFT_RIGHT:
		return "SHIFT_RIGHT"
	case ASSIGN:
		return "ASSIGN"
	case MOD:
		return "MOD"
	case ASSIGN_MOD:
		return "ASSIGN_MOD"
	case ADD:
		return "ADD"
	case ASSIGN_ADD:
		return "ASSIGN_ADD"
	case SUB:
		return "SUB"
	case ASSIGN_SUB:
		return "ASSIGN_SUB"
	case AND_BIT:
		return "AND_BIT"
	case AND_BOOL:
		return "AND_BOOL"
	case XOR_BIT:
		return "XOR_BIT"
	case XOR_BOOL:
		return "XOR_BOOL"
	case OR_BIT:
		return "OR_BIT"
	case OR_BOOL:
		return "OR_BOOL"
	case KEY_FUN:
		return "KEY_FUN"
	case KEY_FOR:
		return "KEY_FOR"
	case KEY_IF:
		return "KEY_IF"
	case KEY_ELSE:
		return "KEY_ELSE"
	case KEY_VAR:
		return "KEY_VAR"
	case KEY_RETURN:
		return "KEY_RETURN"
	case KEY_MATCH:
		return "KEY_MATCH"
	case KEY_WHEN:
		return "KEY_WHEN"
	case KEY_REPEAT:
		return "KEY_REPEAT"
	case KEY_IN:
		return "KEY_IN"
	case MUL:
		return "MUL"
	case ASSIGN_MUL:
		return "ASSIGN_MUL"
	default:
		s = fmt.Sprintf("UNKNOWN(%d)", *this)
	}
	return s
}

func CombineTokens(tk0, tk1 MantisToken) (MantisTokenKind, []rune) {

	if tk0.Kind == COLON && tk1.Kind == ASSIGN {
		return INIT, []rune(":=")
	} else if tk0.Kind == ADD && tk1.Kind == ASSIGN {
		return ASSIGN_ADD, []rune("+=")
	} else if tk0.Kind == SUB && tk1.Kind == ASSIGN {
		return ASSIGN_SUB, []rune("-=")
	} else if tk0.Kind == MOD && tk1.Kind == ASSIGN {
		return ASSIGN_MOD, []rune("%=")
	} else if tk0.Kind == MUL && tk1.Kind == ASSIGN {
		return ASSIGN_MUL, []rune("*=")
	} else if tk0.Kind == ASSIGN && tk1.Kind == ASSIGN {
		return EQUAL, []rune("==")
	} else if tk0.Kind == GREATER_THEN && tk1.Kind == ASSIGN {
		return GREATER_EQUAL_THEN, []rune(">=")
	} else if tk0.Kind == LOWER_THEN && tk1.Kind == ASSIGN {
		return LOWER_EQUAL_THEN, []rune("<=")
	} else if tk0.Kind == GREATER_THEN && tk1.Kind == GREATER_THEN {
		return SHIFT_RIGHT, []rune(">>")
	} else if tk0.Kind == AND_BIT && tk1.Kind == AND_BIT {
		return AND_BOOL, []rune("&&")
	} else if tk0.Kind == OR_BIT && tk1.Kind == OR_BIT {
		return OR_BOOL, []rune("||")
	} else if tk0.Kind == XOR_BIT && tk1.Kind == XOR_BIT {
		return XOR_BOOL, []rune("~~")
	} else if tk0.Kind == SLASH && tk1.Kind == SLASH {
		return COMMENT_LINE, []rune("//")
	} else {
		return UNKNOW, []rune("")
	}
}
