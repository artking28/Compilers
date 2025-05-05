package lexer

import (
	"compilers/utils"
	"errors"
	"os"
	"strings"
)

func Tokenize(filename string, subset int) ([]MantisToken, error) {

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(bytes) <= 0 {
		return nil, utils.GetEmptyFileErr(filename)
	}

	var ret []MantisToken
	column, line := 0, 1
	isComment := false
	runes := []rune(string(bytes))
	buffer := strings.Builder{}
	for i, run := range runes {
		column++
		if 'a' <= run && run <= 'z' ||
			'A' <= run && run <= 'Z' ||
			'0' <= run && run <= '9' ||
			run == '_' {
			buffer.WriteRune(run)
			continue
		}
		if buffer.Len() > 0 {
			var tk MantisToken
			pos := utils.Pos{Line: int64(line), Column: int64(column - buffer.Len())}
			tk = NewToken(pos, ID, 1, []rune(buffer.String())...)
			if !isComment {
				tk, err = ResolveTokenId(filename, tk)
				if err != nil {
					return nil, err
				}
			}
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			buffer.Reset()
		}
		switch run {
		case '\n':
			line += 1
			column = 0
			isComment = false
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, BREAK_LINE, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '\t':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, TAB, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '(':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, L_PAREN, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case ')':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, R_PAREN, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '{':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, L_BRACE, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '}':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, R_BRACE, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '=':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, ASSIGN, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '+':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, ADD, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '-':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, SUB, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '*':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, MUL, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '<':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, LOWER_THEN, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '>':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, GREATER_THEN, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case ' ':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, SPACE, 1, run)
			// s := []Token[*MantisToken](ret)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case ',':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, COMMA, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case ':':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, COLON, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case ';':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, SEMICOLON, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '|':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, OR_BIT, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '&':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, AND_BIT, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '~':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, XOR_BIT, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '%':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, MOD, 1, run)
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		case '/':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, SLASH, 1, run)
			if runes[i+1] == '/' {
				isComment = true
				i += 1
				column++
				err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
				break
			}
			err = errors.Join(err, AppendToken(filename, &ret, tk, subset))
			break
		default:
			if isComment {
				buffer.WriteRune(run)
				continue
			}
			err = utils.GetUnexpectedTokenErr(filename, string(run), utils.Pos{Line: int64(line), Column: int64(column)})
		}
		if err != nil {
			return nil, err
		}
	}

	pos := utils.Pos{Line: int64(line), Column: int64(column)}
	tk := NewToken(pos, EOF, 1, '0')
	_ = AppendToken(filename, &ret, tk, subset)
	return ret, nil
}
