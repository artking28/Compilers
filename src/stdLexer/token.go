package stdLexer

import "compilers/utils"

type Token[T comparable] struct {
	Pos    utils.Pos `json:"-"`
	Kind   T         `json:"kind"`
	Value  []rune    `json:"value"`
	Repeat int       `json:"repeat"`
}

func NewToken[T comparable](pos utils.Pos, kind T, repeat int, value ...rune) Token[T] {
	return Token[T]{Pos: pos, Kind: kind, Value: value, Repeat: repeat}
}

func AppendToken[T comparable](tokens *[]Token[T], token Token[T]) {
	if tokens == nil {
		tokens = &[]Token[T]{}
	}
	count := len(*tokens)
	if count > 0 && (*tokens)[count-1].Kind == token.Kind && string((*tokens)[count-1].Value) == string(token.Value) {
		(*tokens)[count-1].Repeat = (*tokens)[count-1].Repeat + 1
		return
	}
	*tokens = append(*tokens, token)
}
