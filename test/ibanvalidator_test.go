package test

import (
	"testing"

	"github.com/alionurgeven/IBAN/pkg/ibanvalidator"
	"github.com/stretchr/testify/assert"
)

func TestValidIBANs(t *testing.T) {
	IBAN := "AE070331234567890123456"

	assert.True(t, ibanvalidator.Validate(IBAN))

	IBAN = "GB82 WEST 1234 5698 7654 32"

	assert.True(t, ibanvalidator.Validate(IBAN))
}

func TestInvalidIBANs(t *testing.T) {
	IBAN := "AE070aas331231234567890123456"

	assert.False(t, ibanvalidator.Validate(IBAN))

	IBAN = "GB82 123WEST 1234 12351698 765412 32"

	assert.False(t, ibanvalidator.Validate(IBAN))
}
