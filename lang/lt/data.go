package lt

const (
	minus       = "minus"
	negative    = "neigiamas"
	zero        = "nulis"
	sep         = " "
	hundred     = "šimtas"
	hundreds    = "šimtai"
	integerPart = "taškas"
	and         = "ir"
)

//nolint:gochecknoglobals
var (
	megsSingular   = [16]string{"", "tūkstantis", "milijonas", "milijardas", "trilijonas", "kvadrilijonas", "kvintilijonas", "sikstilijonas", "septilijonas", "oktilijonas", "nonilijonas"}
	megsPlural     = [16]string{"", "tūkstančiai", "milijonai", "milijardai", "trilijonai", "kvadrilijonai", "kvintilijonai", "sikstilijonai", "septilijonai", "oktilijonai", "nonilijonai"}
	megsPlural2    = [16]string{"", "tūkstančių", "milijonų", "milijardų", "trilijonų", "kvadrilijonų", "kvintilijonų", "sikstilijonų", "septilijonų", "oktilijonų", "nonilijonų"}
	unitsMasculine = [10]string{"", "vienas", "du", "trys", "keturi", "penki", "šeši", "septyni", "aštuoni", "devyni"}
	unitsFeminine  = [10]string{"", "viena", "dvi", "trys", "keturios", "penkios", "šešios", "septynios", "aštuonios", "devynios"}
	tens           = [10]string{"", "dešimt", "dvidešimt", "trisdešimt", "keturiasdešimt", "penkiasdešimt", "šešiasdešimt", "septyniasdešimt", "aštuoniasdešimt", "devyniasdešimt"}
	teens          = [10]string{"dešimt", "vienuolika", "dvylika", "trylika", "keturiolika", "penkiolika", "šešiolika", "septyniolika", "aštuoniolika", "devyniolika"}
	units          = map[Gender][10]string{
		Male:   unitsMasculine,
		Female: unitsFeminine,
	}
)
