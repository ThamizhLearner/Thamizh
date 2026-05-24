// Thamizh script API

package script

import (
	"strings"

	"github.com/ThamizhLearner/Thamizh/internal/unicode"
)

// -------------------------------------------------------------
// Thamizh Unicode block to Thamizh Letter, LetterSeq converters
// -------------------------------------------------------------

/*
Instantiates a Thamizh letter object
initialized using the given Unicode string representing a Thamizh letter.

Usage:

	import script "github.com/ThamizhLearner/Thamizh"
	var letter script.Letter := script.UnsafeGetLetter ("அ")
*/
func UnsafeGetLetter(ustr string) Letter {
	letterIdxs, ok := unicode.Decode(ustr)
	if !ok {
		panic("Invalid Thamizh Unicode string")
	}
	if len(letterIdxs) != 1 {
		panic("Expected single Thamizh letter Unicode string")
	}
	return Letter{idx: letterIdxs[0]}
}

func UnsafeGetLetterSeq(ustr string) LetterSeq { return LetterSeq{letters: unsafeNewLetterSlice(ustr)} }

func GetLetterSeq(ustr string) (LetterSeq, bool) {
	letters, ok := newLetterSlice(ustr)
	if !ok {
		return LetterSeq{}, false
	}
	return LetterSeq{letters: letters}, true
}

// -------------------------------------------------------------
// Internal Thamizh Unicode to Thamizh Letter array inter-converters
// -------------------------------------------------------------

func unsafeNewLetterSlice(ustr string) []Letter {
	letterIdxs, _ := unicode.Decode(ustr)
	letters := make([]Letter, len(letterIdxs))
	for i, letterIdx := range letterIdxs {
		letters[i] = Letter{idx: letterIdx}
	}
	return letters
}

func newLetterSlice(ustr string) ([]Letter, bool) {
	if len(ustr) == 0 {
		return nil, false
	}
	letterIdxs, _ := unicode.Decode(ustr)
	if len(letterIdxs) == 0 {
		return nil, false
	}
	letters := make([]Letter, len(letterIdxs))
	for i, letterIdx := range letterIdxs {
		letters[i] = Letter{idx: letterIdx}
	}
	return letters, true
}

func letterSliceToUnicode(letters []Letter) string {
	idxs := make([]uint8, len(letters))
	for i, letter := range letters {
		idxs[i] = letter.idx
	}
	return unicode.Encode(idxs)
}

// -------------------------------------------------------------
// Specialized methods dealing with LetterSeq [Extrinsic methods]
// (Basic) Linguistic domain operation methods
//	 - CV join/split
//  - (basic CV awareness level) suffix/prefix match ('n trim)
//  - (basic CV awareness level) suffix/prefix attachment
// -------------------------------------------------------------

// Match and trim specified suffix sequence as is.
func SuffixTrimmed(seq, sfx LetterSeq) (LetterSeq, bool) {
	if !seq.EndsWith(sfx) {
		return LetterSeq{}, false
	}
	return seq.TakeFirstN(seq.Len() - sfx.Len()), true
}

// Suffix concatenated as is.
func SuffixAppended(seq, sfx LetterSeq) LetterSeq { return seq.Appended(sfx) }

// Match and trim specified suffix sequence starting with a Vowel.
func VSuffixTrimmed(seq, sfx LetterSeq) (LetterSeq, bool) {
	seqLen, sfxLen := seq.Len(), sfx.Len()
	if seqLen < sfxLen || seqLen == 0 || sfxLen == 0 {
		return LetterSeq{}, false
	}
	// Second sequence must start with a Vowel
	v := sfx.First()
	if !v.IsV() {
		return LetterSeq{}, false
	}
	// First sequence must have a CV at the appropriate loacation!
	cvIdx := seqLen - sfxLen
	cv := seq.Nth(cvIdx)
	if !cv.IsCV() {
		return LetterSeq{}, false
	}
	c2, v2 := cv.SplitCV()
	if !v.Equals(v2) {
		return LetterSeq{}, false
	}
	// Match the rest of the letter sequence
	matchLen := sfxLen - 1
	if !seq.TakeLastN(matchLen).Equals(sfx.TakeLastN(matchLen)) {
		return LetterSeq{}, false
	}
	sb := NewLetterSeqBuilder(cvIdx + 1)
	sb.Append(seq.TakeFirstN(cvIdx))
	sb.AppendLetter(c2)
	return sb.Build(), true
}

// Attempts VSuffix concatenation, with CV join.
// [ ] There is another case: Join CV if possible, or simply concat/append
func VSuffixAppended(seq, sfx LetterSeq) (LetterSeq, bool) {
	if seq.Len() == 0 || sfx.Len() == 0 {
		return LetterSeq{}, false
	}
	c, v := seq.Last(), sfx.First()
	if !c.IsC() || !v.IsV() {
		return LetterSeq{}, false
	}

	sb := NewLetterSeqBuilder(seq.Len() + sfx.Len() - 1)
	sb.AppendFirstN(seq, seq.Len()-1)
	sb.AppendLetter(c.UnsafeJoinCV(v))
	sb.AppendLastN(sfx, sfx.Len()-1)
	return sb.Build(), true
}

// -------------------------------------------------------------
// Misc methods
// -------------------------------------------------------------

func Syllables(seq LetterSeq) []LetterSeq {
	// Note: Only CV and V both may be followed by 0 or more Cs...
	var subs []LetterSeq
	sIdx, cIdx := 0, -1
	for letter := range seq.Letters() {
		cIdx++
		if letter.IsC() || cIdx == 0 { // Note: cIdx == 0 means the collection hasn't even begun...
			continue
		}
		// For V and CV
		subs = append(subs, seq.TakeRangeN(sIdx, cIdx)) // Capture the syllable we got
		sIdx = cIdx
	}
	return append(subs, seq.TakeLastN(seq.Len()-sIdx)) // Capture the syllable we got
}

func SyllabifiedUStr(seq LetterSeq, sep string) (string, int) {
	var sb strings.Builder // Note: Bit more efficient than strings.Join
	syls := Syllables(seq)
	lastIdx := len(syls) - 1
	for i, syl := range syls {
		sb.WriteString(syl.String())
		if i == lastIdx {
			break
		}
		sb.WriteString(sep)
	}
	return sb.String(), len(syls)
}
