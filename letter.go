// The Thamizh letter unit model

package script

import (
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

// Indicates if this letter has strong-strength vocalization characteristic (வல்லின எழுத்து)
func (chr Letter) IsStrongVocal() bool { return qBitFields[chr.idx]&bit_Strong != 0 }

// Indicates if this letter has medium-strength vocalization characteristic (இடையின எழுத்து)
func (chr Letter) IsMediumVocal() bool { return qBitFields[chr.idx]&bit_Medium != 0 }

// Indicates if this letter has mild-strength vocalization characteristic (மெல்லின எழுத்து)
func (chr Letter) IsMildVocal() bool { return qBitFields[chr.idx]&bit_Mild != 0 }

// Indicates if this letter has short-duration vocalization characteristic (குறில் எழுத்து)
func (chr Letter) IsShortVocal() bool { return qBitFields[chr.idx]&bit_Short != 0 }

// Indicates if this letter has long-duration vocalization characteristic (நெடில் எழுத்து)
func (chr Letter) IsLongVocal() bool { return qBitFields[chr.idx]&bit_Long != 0 }

// Indicates the letter is same as the specified letter
func (chr Letter) Is(l Letter) bool     { return chr.idx == l.idx }
func (chr Letter) Equals(l Letter) bool { return chr.idx == l.idx }

// Indicates the letter is the same as specified in the given Thamizh letter Unicode string
func (chr Letter) IsLetter(ustr string) bool {
	idxs, ok := unicode.Decode(ustr)
	if !ok || len(idxs) != 1 {
		return false
	}
	return idxs[0] == chr.idx
}

// Unpair Thamizh vowelized-consonant letter
func (cv Letter) SplitCV() (c Letter, v Letter, ok bool) {
	if !cv.IsCV() {
		return Letter{}, Letter{}, false
	}
	c, v = cv.UnsafeSplitCV()
	return c, v, true
}

// Unpair Thamizh vowelized-consonant letter
func (cv Letter) MustSplitCV() (c Letter, v Letter) {
	c, v, ok := cv.SplitCV()
	if !ok {
		panic("Expected CV letter")
	}
	return c, v
}

// Unpair Thamizh vowelized-consonant letter
func (cv Letter) UnsafeSplitCV() (c Letter, v Letter) {
	cvIdx := cv.idx - 30
	return Letter{idx: 12 + cvIdx/12}, Letter{idx: cvIdx % 12}
}

// Forms vowelized-consonant letter
func (c Letter) JoinCV(v Letter) (Letter, bool) {
	if !c.IsC() || !v.IsV() {
		return Letter{}, false
	}
	return c.UnsafeJoinCV(v), true
}

// Forms vowelized-consonant letter
func (c Letter) MustJoinCV(v Letter) Letter {
	cv, ok := c.JoinCV(v)
	if !ok {
		panic("Expected C and V letters")
	}
	return cv
}

// Forms vowelized-consonant letter
func (c Letter) UnsafeJoinCV(v Letter) Letter { return Letter{idx: 30 + (c.idx-12)*12 + v.idx} }

// Stringer interface implementation
func (chr Letter) String() string { return unicode.EncodeLetter(chr.idx) }
