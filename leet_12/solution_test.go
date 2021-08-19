package leet12

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestRoma(t *testing.T) {
	res := intToRoman(1994)
	assert.Equal(t, res, "MCMXCIV")
}
