package lexer

import (
	"compilers/stdLexer"
	"compilers/utils"
	"crypto/aes"
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

	var ret []stdLexer.Token[MantisTokenKind]
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
			tk = NewToken(pos, MantisToken_ID, 1, []rune(buffer.String())...)
			if !isComment {
				tk, err = models.ResolveTokenId(filename, tk)
				if err != nil {
					return nil, err
				}
			}
			stdLexer.AppendToken(&ret, tk)
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
			tk := NewToken(pos, MantisToken_BREAK_LINE, 1, run)
			stdLexer.AppendToken[MantisTokenKind](&ret, tk)
			break
		case '\t':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, MantisToken_TAB, 1, run)
			stdLexer.AppendToken(&ret, tk)
			break
		case ' ':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, MantisToken_SPACE, 1, run)
			// s := []stdLexer.Token[*MantisToken](ret)
			stdLexer.AppendToken[MantisTokenKind](&ret, tk)
			break
		case ',':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, MantisToken_COMMA, 1, run)
			stdLexer.AppendToken(&ret, tk)
			break
		case ':':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, MantisToken_COLON, 1, run)
			stdLexer.AppendToken(&ret, tk)
			break
		case '#':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, MantisToken_HASHTAG, 1, run)
			stdLexer.AppendToken(&ret, tk)
			islabel = true
			break
		case '/':
			pos := utils.Pos{Line: int64(line), Column: int64(column)}
			tk := NewToken(pos, MantisToken_SLASH, 1, run)
			if runes[i+1] == '/' {
				isComment = true
				i += 1
				column++
				stdLexer.AppendToken(&ret, tk)
				break
			}
			stdLexer.AppendToken(&ret, tk)
			break
		default:
			if isComment {
				buffer.WriteRune(run)
				continue
			}
			return nil, utils.GetUnexpectedTokenErr(filename, string(run), utils.Pos{Line: int64(line), Column: int64(column)})
		}
	}

	pos := utils.Pos{Line: int64(line), Column: int64(column)}
	tk := NewToken(pos, MantisToken_EOF, 1, '0')
	stdLexer.AppendToken(&ret, tk)
	return ret, nil
}
