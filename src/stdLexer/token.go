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