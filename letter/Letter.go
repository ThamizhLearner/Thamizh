// The Thamizh letter model
package letter

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	script "github.com/ThamizhLearner/Thamizh"
)

// Letter characteristics query bits
type qBits uint8

const (
	Bit_Consonant qBits = 1 << iota
	Bit_Vowel
	Bit_Strong
	Bit_Medium
	Bit_Mild
	Bit_Short
	Bit_Long
)

// Primary Thamizh letter
type Letter struct {
	form  string // Note: This gets used solely during unicode string reconstruction
	flags qBits  // attribute query bit field
}

// Constructs new vowel (V) letter
func newV(v script.UVowel) Letter {
	var flags qBits = Bit_Vowel | vocalDurationQBit(v)
	return Letter{string(v.Rep()), flags}
}

// Constructs new consonant (C) letter
func newC(c script.UConsonant) Letter {
	var flags qBits = Bit_Consonant | vocalStrengthQBit(c)
	return Letter{c.Rep(), flags}
}

// Constructs new vowelized-consonant (CV) letter
func newCV(c script.UConsonant, v script.UVowel) Letter {
	var flags qBits = Bit_Vowel | Bit_Consonant | vocalStrengthQBit(c) | vocalDurationQBit(v)
	runes := []rune{c.BaseFormRep()}
	if v.AttachedFormRep() != script.UNull {
		runes = append(runes, v.AttachedFormRep())
	}
	return Letter{string(runes), flags}
}

func vocalStrengthQBit(c script.UConsonant) qBits {
	switch c.VocalStrength() {
	case script.Strong:
		return Bit_Strong
	case script.Medium:
		return Bit_Medium
	case script.Mild:
		return Bit_Mild
	default:
		panic("Unexpected vocal strength value")
	}
}

func vocalDurationQBit(v script.UVowel) qBits {
	switch v.VocalDuration() {
	case script.Short:
		return Bit_Short
	case script.Long:
		return Bit_Long
	default:
		panic("Unexpected vocal duration value")
	}
}

const cvMask qBits = Bit_Consonant | Bit_Vowel

// Indicates primary Thamizh vowel letter
func (l Letter) IsVowel() bool { return (l.flags & cvMask) == Bit_Vowel }

// Indicates primary Thamizh consonant letter
func (l Letter) IsConsonant() bool { return (l.flags & cvMask) == Bit_Consonant }

// Indicates dependent Thamizh vowelized-consonant letter
func (l Letter) IsCV() bool { return (l.flags & cvMask) == cvMask }

func (l Letter) String() string { return l.form }

func (l Letter) Dump() string {
	var bStrs []string

	switch l.flags & cvMask {
	case Bit_Vowel:
		bStrs = append(bStrs, "Vowel")
	case Bit_Consonant:
		bStrs = append(bStrs, "Consonant")
	default:
		bStrs = append(bStrs, "Consonant|Vowel")
	}

	if (l.flags & Bit_Vowel) != 0 { // For V and CV letters
		switch {
		case (l.flags & Bit_Short) != 0:
			bStrs = append(bStrs, "Short")
		case (l.flags & Bit_Long) != 0:
			bStrs = append(bStrs, "Long")
		}
	}

	if (l.flags & Bit_Consonant) != 0 { // For C and CV letters
		switch {
		case (l.flags & Bit_Mild) != 0:
			bStrs = append(bStrs, "Mild")
		case (l.flags & Bit_Medium) != 0:
			bStrs = append(bStrs, "Medium")
		case (l.flags & Bit_Strong) != 0:
			bStrs = append(bStrs, "Strong")
		}
	}

	return fmt.Sprintf("%s, %v", l, strings.Join(bStrs, "|"))
}

func buildThamizhLetterList() [246]Letter {
	var letters = [246]Letter{}
	// Primary vowel letters (V form)
	for i, uv := range script.IndexedUVowels() {
		letters[i] = newV(uv)
	}
	// Primary consonant letters (C form)
	for i, uc := range script.IndexedUConsonants() {
		letters[12+i] = newC(uc)
	}
	// Dependent vowelized-consonant letters (CV form)
	for ci, uc := range script.IndexedUConsonants() {
		consonantRowIdx := 30 + ci*12
		for vi, uv := range script.IndexedUVowels() {
			letters[consonantRowIdx+vi] = newCV(uc, uv)
		}
	}
	return letters
}

// (Internal protected) Ordered Thamizh letter list
var internalLetterList [246]Letter = buildThamizhLetterList()

// Ordered Thamizh letter (indexed) iterator
func IndexedLetters() iter.Seq2[int, Letter] { return slices.All(internalLetterList[:]) }

// Ordered Thamizh letter iterator
func Letters() iter.Seq[Letter] { return slices.Values(internalLetterList[:]) }
