package Mantis

func GetZeroValue(typeOf string) any {
	switch typeOf {
	case "number":
		return 0
	case "bool":
		return 0
	default:
		return nil
	}
}
