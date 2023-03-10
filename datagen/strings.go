package datagen

import (
	"strings"
)

const (
	AlphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// Alphanumeric returns a random string of alphanumeric characters.
func Alphanumeric(length int) string {
	var builder strings.Builder
	builder.Grow(length)
	alphaLen := len(AlphaNumeric)
	for i := 0; i < length; i++ {
		builder.WriteByte(AlphaNumeric[Rand().Int(alphaLen)])
	}
	return builder.String()
}

// AlphanumericBetween returns a random string of alphanumeric characters between min and max.
func AlphanumericBetween(min, max int) string {
	return Alphanumeric(Rand().IntBetween(min, max))
}

// Words returns a random string of words.
func Words(length int) string {
	var builder strings.Builder
	builder.Grow(length * 5)

	for i := 0; i < length; i++ {
		switch Rand().Int(3) {
		case 0:
			builder.WriteString(Dict().RandomAdjective())
		case 1:
			builder.WriteString(Dict().RandomVerb())
		case 2:
			builder.WriteString(Dict().RandomNoun())
		}
		if i < length-1 {
			builder.WriteByte(' ')
		}
	}
	return builder.String()
}

func Domain() string {
	return Dict().RandomDomain()
}
