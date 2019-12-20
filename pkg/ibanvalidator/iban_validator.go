package ibanvalidator

import (
	"math/big"
	"strconv"
	"strings"

	iban "github.com/alionurgeven/IBAN/pkg/iban"
)

//Validate is implemented as defined in ISO 7064
func Validate(number string) bool {
	IBAN := iban.NewIBAN(strings.ReplaceAll(number, " ", ""))
	reArrange(IBAN)
	convertToInteger(IBAN)
	return isValid(IBAN)

}

func reArrange(iban *iban.IBAN) {
	runes := []rune(iban.Number)
	firstFourChars := string(runes[0:4])
	remainingChars := string(runes[4:iban.Length])
	iban.Number = remainingChars + firstFourChars
}

func convertToInteger(iban *iban.IBAN) {
	newNumber := ""
	for _, char := range iban.Number {
		val := int(char)
		if val < 91 && val > 64 {
			val -= 55
			newNumber += strconv.Itoa(val)
		} else {
			newNumber += string(char)
		}
	}
	iban.Number = newNumber
}

func isValid(iban *iban.IBAN) bool {

	bigVal, success := new(big.Int).SetString(iban.Number, 10)
	if !success {
		return false
	}

	modVal := new(big.Int).SetInt64(97)
	resVal := new(big.Int).Mod(bigVal, modVal)

	// Check if module is equal to 1
	if resVal.Int64() != 1 {
		return false
	}

	return true
}
