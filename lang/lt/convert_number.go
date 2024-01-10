package lt

import (
	"github.com/gammban/numtow/internal/ds"
)

// Int64 converts int64 to words or returns error.
//
//	result, err := Int64(1) // result: бiр
func Int64(number int64, options ...OptFunc) (result string, err error) {
	return convert(ds.ParseInt64(number), Male)
}

// MustInt64 converts int64 to words or returns empty string.
//
//	result, err := MustInt64(1) // result: бiр
func MustInt64(number int64, options ...OptFunc) string {
	result, err := Int64(number, options...)
	if err != nil {
		return ""
	}

	return result
}
