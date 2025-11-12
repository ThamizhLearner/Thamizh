// The Thamizh letter unit model

package script

import (
	"fmt"

	"github.com/ThamizhLearner/Thamizh/internal/unicode"
)

// The Thamizh letter
type Letter struct {
	idx uint8
}

// Not-a-valid-letter letter
func InvalidLetter() Letter { return Letter{idx: 255} }

// Indicates a Thamizh vowel letter
func (chr Letter) IsV() bool { return chr.idx < 12 }

// Indicates a Thamizh consonant letter
func (chr Letter) IsC() bool { return chr.idx >= 12 && chr.idx < 30 }

// Indicates a Thamizh vowelized-consonant letter
func (chr Letter) IsCV() bool { return chr.idx >= 30 } // Note: No need for upper limit check!

// Indicates a Thamizh primary letter
func (chr Letter) IsPrimary() bool { return chr.idx < 30 }

func (chr Letter) IsStrongVocal() bool { return qBitFields[chr.idx]&bit_Strong != 0 }
func (chr Letter) IsMediumVocal() bool { return qBitFields[chr.idx]&bit_Medium != 0 }
func (chr Letter) IsMildVocal() bool   { return qBitFields[chr.idx]&bit_Mild != 0 }
func (chr Letter) IsShortVocal() bool  { return qBitFields[chr.idx]&bit_Short != 0 }
func (chr Letter) IsLongVocal() bool   { return qBitFields[chr.idx]&bit_Long != 0 }

// Checks if the letter is same as the specified letter
func (chr Letter) IsLetter(l Letter) bool { return chr.idx == l.idx }

// Checks if the letter is the same as specified in the given letter literal
func (chr Letter) Is(letterLiteral string) bool {
	idxs := unicode.Decode(letterLiteral)
	if len(idxs) != 0 {
		panic(fmt.Sprintf("Invalid letter literal %s", letterLiteral))
	}
	return idxs[0] == chr.idx
}

// Unpair Thamizh vowelized-consonant letter
func (chr Letter) DetachedCV() (c Letter, v Letter) {
	cvIdx := chr.idx - 30
	return Letter{idx: cvIdx / 12}, Letter{idx: cvIdx % 12}
}

func (chr Letter) String() string { return unicode.EncodeLetter(chr.idx) }
