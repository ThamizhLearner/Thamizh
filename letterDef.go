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

// Indicates a Thamizh vowel letter (உயிர் எழுத்து)
func (chr Letter) IsV() bool { return chr.idx < 12 }

// Indicates a Thamizh consonant letter (மெய் எழுத்து)
func (chr Letter) IsC() bool { return chr.idx >= 12 && chr.idx < 30 }

// Indicates a Thamizh vowelized-consonant letter (உயிர் மெய் எழுத்து)
func (chr Letter) IsCV() bool { return chr.idx >= 30 } // Note: No need for upper limit check!

// Indicates a Thamizh primary letter (முதல் எழுத்து)
func (chr Letter) IsPrimary() bool { return chr.idx < 30 }

// Indicates if this letter has strong-strength vocalization characteric (வல்லின எழுத்து)
func (chr Letter) IsStrongVocal() bool { return qBitFields[chr.idx]&bit_Strong != 0 }

// Indicates if this letter has medium-strength vocalization characteric (இடையின எழுத்து)
func (chr Letter) IsMediumVocal() bool { return qBitFields[chr.idx]&bit_Medium != 0 }

// Indicates if this letter has mild-strength vocalization characteric (மெல்லின எழுத்து)
func (chr Letter) IsMildVocal() bool { return qBitFields[chr.idx]&bit_Mild != 0 }

// Indicates if this letter has short-duration vocalization characteric (குறில் எழுத்து)
func (chr Letter) IsShortVocal() bool { return qBitFields[chr.idx]&bit_Short != 0 }

// Indicates if this letter has long-duration vocalization characteric (நெடில் எழுத்து)
func (chr Letter) IsLongVocal() bool { return qBitFields[chr.idx]&bit_Long != 0 }

// Indicates the letter is same as the specified letter
func (chr Letter) Is(l Letter) bool { return chr.idx == l.idx }

// Indicates the letter is the same as specified in the given Thamizh letter Unicode string
func (chr Letter) IsLetter(ustr string) bool {
	idxs := unicode.Decode(ustr)
	if idxs == nil {
		panic(fmt.Sprintf("Invalid letter literal %s", ustr))
	}
	if len(idxs) > 1 {
		panic("Expected single Thamizh letter Unicode string")
	}
	return idxs[0] == chr.idx
}

// Unpair Thamizh vowelized-consonant letter
func (cv Letter) SplitCV() (c Letter, v Letter) {
	if !cv.IsCV() {
		panic("expected CV letter")
	}
	cvIdx := cv.idx - 30
	return Letter{idx: 12 + cvIdx/12}, Letter{idx: cvIdx % 12}
}

// Forms vowelized-consonant letter
func (c Letter) JoinCV(v Letter) Letter {
	if !c.IsC() || !v.IsV() {
		panic("expected C and V letters")
	}
	return Letter{idx: 30 + (c.idx-12)*12 + v.idx}
}

// Stringer interface implementation
func (chr Letter) String() string { return unicode.EncodeLetter(chr.idx) }
