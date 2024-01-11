package lt

import (
	"strconv"
	"strings"

	"github.com/gammban/numtow/internal/digit"
	"github.com/gammban/numtow/internal/ds"
)

func convert(d ds.DigitString, gender Gender) (result string, err error) {
	unit := units[gender]

	if d.IsEmpty() {
		return "", nil
	}

	if d.IsZero() {
		return zero, nil
	}

	s := strings.Builder{}

	if d.IsSignMinus {
		s.WriteString(minus)
		s.WriteString(sep)
	}

	triplets, err := d.Split()
	if err != nil {
		return result, err
	}

	if triplets.Len() > len(megsSingular) {
		return result, strconv.ErrRange
	}

	for i := len(triplets) - 1; i >= 0; i-- {
		t := triplets[i]

		if t.IsZero() {
			continue
		}

		// _xx
		switch t.Hundreds() {
		case digit.Digit0, digit.DigitUnknown:
		case digit.Digit1:
			s.WriteString(unit[t.Hundreds()])
			s.WriteString(sep)
			s.WriteString(hundred)
			s.WriteString(sep)
		case digit.Digit2, digit.Digit3, digit.Digit4, digit.Digit5, digit.Digit6, digit.Digit7, digit.Digit8, digit.Digit9:
			s.WriteString(unit[t.Hundreds()])
			s.WriteString(sep)
			s.WriteString(hundreds)
			s.WriteString(sep)
		}
		// x__
		switch t.Tens() {
		case digit.DigitUnknown:
		case digit.Digit0: // x0_
			if t.Units() == digit.Digit0 {
				break
			}
			s.WriteString(unit[t.Units()])
			s.WriteString(sep)
		case digit.Digit1: // x1_
			s.WriteString(teens[t.Units()])
			s.WriteString(sep)
		case digit.Digit2, digit.Digit3, digit.Digit4, digit.Digit5, digit.Digit6, digit.Digit7, digit.Digit8, digit.Digit9:
			s.WriteString(tens[t.Tens()])
			s.WriteString(sep)
			s.WriteString(unit[t.Units()])
			s.WriteString(sep)
		}

		// megas
		if t.Tens() == digit.Digit1 {
			s.WriteString(megsPlural2[i])
			s.WriteString(sep)
			continue
		}
		switch t.Units() {
		case digit.DigitUnknown:
		case digit.Digit0:
			s.WriteString(megsPlural2[i])
			s.WriteString(sep)
		case digit.Digit1:
			s.WriteString(megsSingular[i])
			s.WriteString(sep)
		case digit.Digit2, digit.Digit3, digit.Digit4, digit.Digit5, digit.Digit6, digit.Digit7, digit.Digit8, digit.Digit9:
			s.WriteString(megsPlural[i])
			s.WriteString(sep)
		}

	}

	return strings.TrimSpace(s.String()), nil
}
