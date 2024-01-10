package lt

import (
	"github.com/gammban/numtow/internal/digit"
	"github.com/gammban/numtow/internal/triplet"
)

type Declension int

const (
	DeclensionSingular Declension = iota // singular nominative declension
	DeclensionPlural                     // plural nominative declension
	DeclensionPlural2                    // plural genitive declension
)

func getTripletDeclension(t triplet.Triplet) Declension {
	if t.Tens() == digit.Digit1 {
		return DeclensionPlural2
	}
	if t.Units() == digit.Digit1 {
		return DeclensionSingular
	}
	if t.Units() == digit.Digit0 {
		return DeclensionPlural2
	}
	return DeclensionPlural
}
