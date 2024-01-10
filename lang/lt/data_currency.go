package lt

import (
	"fmt"

	"github.com/gammban/numtow/curtow/cur"
)

var (
	//nolint
	currencyOpts = map[cur.Currency]currencyInfo{
		cur.KZT: {
			Name: map[Declension]string{
				DeclensionSingular: "tengė",
				DeclensionPlural:   "tengės",
				DeclensionPlural2:  "tengių",
			},
			NameGender: Female,
			UnitName: map[Declension]string{
				// maybe like this
				DeclensionSingular: "tijyna",
				DeclensionPlural:   "tijynos",
				DeclensionPlural2:  "tijynų",
			},
			UnitGender: Female,
		},
		cur.RUB: {
			Name: map[Declension]string{
				DeclensionSingular: "rublis",
				DeclensionPlural:   "rubliai",
				DeclensionPlural2:  "rublių",
			},
			NameGender: Male,
			UnitName: map[Declension]string{
				DeclensionSingular: "kapeika",
				DeclensionPlural:   "kapeikos",
				DeclensionPlural2:  "kapeikų",
			},
			UnitGender: Female,
		},
		cur.USD: {
			Name: map[Declension]string{
				DeclensionSingular: "doleris",
				DeclensionPlural:   "doleriai",
				DeclensionPlural2:  "dolerių",
			},
			NameGender: Male,
			UnitName: map[Declension]string{
				DeclensionSingular: "centas",
				DeclensionPlural:   "centai",
				DeclensionPlural2:  "centų",
			},
			UnitGender: Male,
		},
		cur.EUR: {
			Name: map[Declension]string{
				DeclensionSingular: "euras",
				DeclensionPlural:   "eurai",
				DeclensionPlural2:  "eurų",
			},
			NameGender: Male,
			UnitName: map[Declension]string{
				DeclensionSingular: "centas",
				DeclensionPlural:   "centai",
				DeclensionPlural2:  "centų",
			},
			UnitAbbreviation: "ct",
			UnitGender:       Male,
		},
	}
)

type currencyInfo struct {
	Name             map[Declension]string
	NameGender       Gender
	NameAbbreviation string
	UnitName         map[Declension]string
	UnitGender       Gender
	UnitAbbreviation string
}

func (c currencyInfo) GetCurrencyName(d Declension) string {
	res, ok := c.Name[d]
	if ok {
		return res
	}

	return ""
}

func (c currencyInfo) GetCurrencyGender() Gender {
	return c.NameGender
}

//func (c currencyInfo) GetCurrencyUnitName(d Declension) string {
//	res, ok := c.UnitName[d]
//	if ok {
//		return res
//	}
//
//	return ""
//}

func (c currencyInfo) GetCurrencyUnitName() string {
	return c.UnitAbbreviation
}

func getCurrencyInfo(c cur.Currency) (info currencyInfo, err error) {
	info, ok := currencyOpts[c]
	if ok {
		return info, nil
	}

	return info, fmt.Errorf("%w: info not found", cur.ErrBadCurrency)
}

func (c currencyInfo) GetCurrencyUnitGender() Gender {
	return c.UnitGender
}
