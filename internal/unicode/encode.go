package unicode

import (
	"strings"

	base "github.com/ThamizhLearner/Thamizh/internal/unicode/internal"
)

const (
	UDot rune = '\u0BCD' // Unicode Thamizh 'dot' (புள்ளி) code point
)

// Letter index => Letter (Unicode encoded) string
var letterEncodings [246]string = func() [246]string {
	var strs [246]string
	for i, v := range base.PrimaryVowelUCodes {
		strs[i] = string(v)
	}

	runes := []rune{base.UNull, UDot}
	for i, c := range base.BaseConsonantUCodes {
		runes[0] = c
		strs[12+i] = string(runes)
	}
	for ci, uc := range base.BaseConsonantUCodes {
		rowBaseIdx := 30 + ci*12
		strs[rowBaseIdx] = string(uc)
		runes[0] = uc
		for ai, a := range base.AttachedVowelUCodes[1:] {
			ai++ // Compensate for sub-slicing [1:]
			runes[1] = a
			strs[rowBaseIdx+ai] = string(runes)
		}
	}
	return strs
}()

func EncodeLetter(idx uint8) string { return letterEncodings[idx] }

func Encode(idxs []uint8) string {
	var sb strings.Builder
	for _, idx := range idxs {
		sb.WriteString(letterEncodings[idx])
	}
	return sb.String()
}
