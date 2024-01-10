package lt

import (
	"strings"
	"testing"

	"github.com/gammban/numtow/curtow/cur"
	testdata "github.com/gammban/numtow/internal/testdata/lang_lt"
)

func TestCurrencyFloat64_EUR(t *testing.T) {
	for k, v := range testdata.TestCaseLangLTCurrencyEURFloat {
		got, err := CurrencyFloat64(k, WithCur(cur.EUR))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, v) {
			t.Errorf("%.2f: \nexp: '%s' \ngot: '%s'", k, v, got)
			return
		}
	}
}
