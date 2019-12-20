package iban

//IBAN is the international bank number representation
type IBAN struct {
	Number string
	Length int
}

//NewIBAN creates a new IBAN from the given IBAN string
func NewIBAN(number string) *IBAN {
	iban := IBAN{Number: number, Length: len(number)}
	return &iban
}
