package parser

import (
	"compilers/utils"
)

//type MantisVariable[T any] struct {
//	stdParser.Variable[T]
//	MantisStmtBase
//}
//
//func (this MantisVariable[T]) WriteMemASM() ([]uint16, error) {
//	//TODO implement me
//	panic("implement me | VariableStmt@WriteMemASM")
//}
//
//func (this MantisVariable[T]) GetTitle() string {
//	return this.Title
//}
//
//func NewVariableStmt[T comparable](name string, pos utils.Pos, value stdParser.IExp[T], owner uint64, parser *MantisParser) *MantisVariable[T] {
//	ret := &MantisVariable[T]{
//		Variable: stdParser.Variable[T]{
//			Id:    uint64(len(parser.Variables) + 1),
//			Name:  name,
//			Value: value,
//			Owner: owner,
//		},
//		MantisStmtBase: MantisStmtBase{
//			Parser: parser,
//			Title:  "VariableStmt",
//			Pos:    pos,
//		},
//	}
//	parser.Variables[name] = &ret.Variable
//	return ret
//}

func (parser *MantisParser) ParseVariable(ownerId uint64) error {
	vars := map[string]MantisExp[any]{}
	pos := map[string]utils.Pos{}

	for k, v := range vars {
		parser.Inject(NewVariableStmt[any](k, pos[k], v, ownerId, parser))
	}
	return nil
}
