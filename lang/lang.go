package lang

import "errors"

type Lang uint32

var (
	ErrBadLanguage = errors.New("bad language")
)

const (
	Unknown Lang = iota
	KZ
	RU
	EN
	LT
)

const (
	CodeKZ = "KZ"
	CodeRU = "RU"
	CodeEN = "EN"
	CodeLT = "LT"
)

// String returns language code or empty string
func (l Lang) String() string {
	switch l {
	case KZ:
		return CodeKZ
	case RU:
		return CodeRU
	case EN:
		return CodeEN
	case LT:
		return CodeLT
	case Unknown:
		return ""
	default:
		return ""
	}
}
