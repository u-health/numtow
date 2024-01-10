package lt

import (
	"strings"
	"testing"

	testdata "github.com/gammban/numtow/internal/testdata/lang_lt"
)

func BenchmarkLT_ConvertInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, v := range testdata.TestCaseLangLTNumbersInt64 {
			got, err := Int64(k)
			if err != nil {
				b.Error(err)
				return
			}

			if !strings.EqualFold(got, v) {
				b.Errorf("%d: \nexp: '%s'\ngot: '%s'", k, v, got)
				return
			}
		}
	}
}

func TestMustInt64(t *testing.T) {
	if res := MustInt64(5); res != "бес" {
		t.Fatal("result mismatch, expected бір")
	}
}
