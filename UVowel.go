package script

import (
	"iter"
	"slices"
)

// Unicode code point "None"
const UNone rune = '\000'

// Verifies if given rune is UNone
func IsUNone(r rune) bool { return r == UNone }

// Thamizh vowel vocalization-duration enum
type VocalDuration uint8

const (
	// குறில் எழுத்து; Short vocalization-duration vowel
	Short VocalDuration = iota

	// நெடில் எழுத்து; Long vocalization-duration vowel
	Long
)

// Abstract Thamizh vowel Unicode code point
type UVowel struct {
	// Thamizh vowel Unicode code point
	uForm rune
	// Unicode code point for attached-vowel variant of Thamizh vowel
	uAttachedForm rune

	// Thamizh vowel vocalization-duration enum
	vocalDuration VocalDuration
}

// Thamizh vowel representative string
func (v UVowel) Rep() string { return string(v.uForm) }

// Thamizh attached-vowel representative string; or empty string if none
func (v UVowel) AttachedFormRep() string {
	if v.uAttachedForm == UNone {
		return ""
	} else {
		return string(v.uAttachedForm)
	}
}

// Thamizh vowel vocalization-duration enum
func (v UVowel) VocalDuration() VocalDuration { return v.vocalDuration }

// (Internal protected) Ordered Thamizh Unicode vowel list
var internalVees [12]UVowel = [...]UVowel{
	{'அ', UNone, Short},
	{'ஆ', 'ா', Long},
	{'இ', 'ி', Short},
	{'ஈ', 'ீ', Long},
	{'உ', 'ு', Short},
	{'ஊ', 'ூ', Long},
	{'எ', 'ெ', Short},
	{'ஏ', 'ே', Long},
	{'ஐ', 'ை', Long},
	{'ஒ', 'ொ', Short},
	{'ஓ', 'ோ', Long},
	{'ஔ', 'ௌ', Long},
}

// Ordered abstract Thamizh vowel Unicode code points (indexed) iterator
func IndexedUVowels() iter.Seq2[int, UVowel] { return slices.All(internalVees[:]) }

// Ordered abstract Thamizh vowel Unicode code points iterator
func UVowels() iter.Seq[UVowel] { return slices.Values(internalVees[:]) }
