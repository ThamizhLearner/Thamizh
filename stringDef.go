// The Thamizh letter string model

package script

import (
	"iter"
	"slices"
	"strings"

	"github.com/ThamizhLearner/Thamizh/internal/unicode"
)

/*
	Basic rules:
	0. Never allow invalid Letter [Program sanity concern]
	1. Never allow zero-length String [Program sanity concern]
	2. Reuse the slice when (simple) trimming [Memory efficiency concern]
	3.	Must use new slice when mutating values [Avoidable bug concern]
*/

// The Thamizh letter string
type String struct {
	idxs []uint8
}

// Attempts decoding the given Thamizh Unicode string
func Decode(ustr string) (String, bool) {
	idxs := unicode.Decode(ustr)
	if idxs == nil {
		return String{}, false // Simply return nil string, which cannot be used meaningfully!
	}
	return String{idxs: idxs}, true
}

// Decodes the given (structurally valid) Thamizh Unicode string
//
// Use it for Thamizh Unicode literals in code, which are expected to be valid Thamizh Unicode strings.
func MustDecode(ustr string) String {
	idxs := unicode.Decode(ustr)
	if idxs == nil {
		panic("invalid Thamizh Unicode string")
	}
	return String{idxs: idxs}
}

// Indicates if given Unicode string is a (structurally valid) Thamizh Unicode string
func IsValidThamizhUnicode(ustr string) bool { // TODO: Needs performace optimization!
	idxs := unicode.Decode(ustr)
	return idxs != nil
}

func (s String) FirstLetter() Letter     { return Letter{idx: s.idxs[0]} }
func (s String) LastLetter() Letter      { return Letter{idx: s.idxs[len(s.idxs)-1]} }
func (s String) LetterAt(idx int) Letter { return Letter{idx: s.idxs[idx]} }

// Stringer interface implementation
func (s String) String() string { return unicode.Encode(s.idxs) }

// Count of letters in the string
func (s String) Len() int { return len(s.idxs) }

func mapper[V, U any](s iter.Seq[V], mapFn func(V) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for item := range s {
			if !yield(mapFn(item)) {
				return
			}
		}
	}
}

// Iterator over letters making up Thamizh String object
func (sa String) Letters() iter.Seq[Letter] {
	mapp := func(idx uint8) Letter { return Letter{idx: idx} }
	return mapper(slices.Values(sa.idxs), mapp)
}

func (s String) TrimEnd(trim String) (String, bool) {
	strLen, trimLen := len(s.idxs), len(trim.idxs)
	if strLen < trimLen {
		return s, false
	}
	matchIdx := strLen - trimLen // sub-string match start index, along the given string
	// Simple idx match should do! Except for the first idx.
	if trimLen > 1 {
		idxs := s.idxs[matchIdx+1:]
		idxs2 := trim.idxs[1:]
		for i, v := range idxs {
			if idxs2[i] != v {
				return s, false
			}
		}
	}
	if s.idxs[matchIdx] == trim.idxs[0] {
		return String{idxs: s.idxs[:matchIdx]}, true // Reusing the original slice!
	}

	// Check if we have CV and V match...
	suffix0 := trim.FirstLetter()
	if !suffix0.IsV() {
		return s, false
	}
	str0 := s.LetterAt(matchIdx)
	if !str0.IsCV() {
		return s, false
	}
	c, v := str0.SplitCV()
	if v.idx != suffix0.idx {
		return s, false
	}
	// Carefully compose the trimmed string
	// Make a copy of the slice, since we are going to modify it
	matchIdx++ // Account for the overlapping CV/V letters
	idxs2 := make([]uint8, matchIdx)
	copy(idxs2, s.idxs[:matchIdx])
	idxs2[matchIdx-1] = c.idx
	return String{idxs: idxs2}, true
}

// Attempts matching the given trim, and on match replaces that with given replacement string
func (s String) ReplaceEnd(trim String, append String) (String, bool) {
	st, ok := s.TrimEnd(trim)
	if !ok {
		return s, false
	}
	return st.Append(append), true
}

// Simple concatenation!
func (s String) appendRaw(s2 String) String {
	idxs := make([]uint8, len(s.idxs)+len(s2.idxs))
	copy(idxs, s.idxs)
	copy(idxs[len(s.idxs):], s2.idxs)
	return String{idxs: idxs}
}

// Deep concatenation.
//
// Merges any trailing C (on the first) with leading V (on the second), forming CV letter
func (s String) Append(a String) String {
	lenA := len(s.idxs)
	letterA := s.LastLetter()
	letterB := a.FirstLetter()
	if !letterA.IsC() || !letterB.IsV() {
		return s.appendRaw(a)
	}
	idxs := make([]uint8, lenA+len(a.idxs)-1)
	copy(idxs, s.idxs)
	copy(idxs[lenA-1:], a.idxs)
	idxs[lenA-1] = letterA.JoinCV(letterB).idx
	return String{idxs: idxs}
}

func (s String) Syllables() []String {
	// CV and V both may be followed by 0 or more Cs...
	var subs []String
	sIdx, cIdx := 0, -1
	for letter := range s.Letters() {
		cIdx++
		if letter.IsC() || cIdx == 0 { // Note: cIdx == 0 means the collection hasn't even begun...
			continue
		}
		// For V and CV
		subs = append(subs, String{idxs: s.idxs[sIdx:cIdx]}) // Capture the syllable we got
		sIdx = cIdx
	}
	return append(subs, String{idxs: s.idxs[sIdx:]}) // Capture the syllable we got
}

func (s String) SyllabifiedUStr(usep string) string {
	var sb strings.Builder
	subs := s.Syllables()
	lastIdx := len(subs) - 1
	for i, syl := range subs {
		sb.WriteString(syl.String())
		if i == lastIdx {
			break
		}
		sb.WriteString(usep)
	}
	return sb.String()
}
