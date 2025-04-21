package lexer

import (
	"compilers/utils"
	"os"
	"strings"
)

func Tokenize(filename string) ([]MantisToken, error) {

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(bytes) <= 0 {
		return nil, utils.GetEmptyFileErr(filename)
	}

	var ret []MantisToken
	column, line := 0, 1
	isComment, islabel := false, false
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
			AppendToken(&ret, tk)
			if islabel {
				islabel = false
			}
			buffer.Reset()
		}
		switch run {
		case '\n':
			line += 1
			column = 0
			isComment = false
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, BREAK_LINE, 1, run)
			AppendToken(&ret, tk)
			break
		case '\t':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, TAB, 1, run)
			AppendToken(&ret, tk)
			break
		case '(':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, L_PAREN, 1, run)
			AppendToken(&ret, tk)
			break
		case ')':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, R_PAREN, 1, run)
			AppendToken(&ret, tk)
			break
		case '{':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, L_BRACE, 1, run)
			AppendToken(&ret, tk)
			break
		case '}':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, R_BRACE, 1, run)
			AppendToken(&ret, tk)
			break
		case '=':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, ASSIGN, 1, run)
			AppendToken(&ret, tk)
			break
		case '+':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, ADD, 1, run)
			AppendToken(&ret, tk)
			break
		case '-':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, SUB, 1, run)
			AppendToken(&ret, tk)
			break
		case '*':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, MUL, 1, run)
			AppendToken(&ret, tk)
			break
		case '<':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, LOWER_THEN, 1, run)
			AppendToken(&ret, tk)
			break
		case '>':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, GREATER_THEN, 1, run)
			AppendToken(&ret, tk)
			break
		case ' ':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, SPACE, 1, run)
			// s := []Token[*MantisToken](ret)
			AppendToken(&ret, tk)
			break
		case ',':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, COMMA, 1, run)
			AppendToken(&ret, tk)
			break
		case ':':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, COLON, 1, run)
			AppendToken(&ret, tk)
			break
		case ';':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, SEMICOLON, 1, run)
			AppendToken(&ret, tk)
			break
		case '/':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, SLASH, 1, run)
			if runes[i+1] == '/' {
				isComment = true
				i += 1
				column++
				AppendToken(&ret, tk)
				break
			}
			AppendToken(&ret, tk)
			break
		default:
			if isComment {
				buffer.WriteRune(run)
				continue
			}
			return nil, utils.GetUnexpectedTokenErr(filename, string(run), utils.Pos{Line: int64(line) - 1, Column: int64(column)})
		}
	}

	pos := utils.Pos{Line: int64(line), Column: int64(column)}
	tk := NewToken(pos, EOF, 1, '0')
	AppendToken(&ret, tk)
	return ret, nil
}
