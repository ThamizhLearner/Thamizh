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
	1. Reuse the slice wherever possible [Memory efficiency concern]
	2.	Must never mutate or append to existing slices (slices are shared wherever possible) [Avoidable bug concern]
*/

// The Thamizh letter string
type String struct {
	idxs []uint8
}

func newEmptyString() String { return String{idxs: make([]uint8, 0)} }

// Attempts decoding the given Thamizh Unicode string
func Decode(ustr string) (String, bool) {
	idxs := unicode.Decode(ustr)
	if idxs == nil {
		return newEmptyString(), false
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
func IsValidThamizhUStr(ustr string) bool { // TODO: Needs performace optimization!
	idxs := unicode.Decode(ustr)
	return idxs != nil
}

func (s String) FirstLetter() Letter     { return Letter{idx: s.idxs[0]} }
func (s String) LastLetter() Letter      { return Letter{idx: s.idxs[len(s.idxs)-1]} }
func (s String) LetterAt(idx int) Letter { return Letter{idx: s.idxs[idx]} }

// Stringer interface implementation
func (s String) String() string { return unicode.Encode(s.idxs) }

// Count of Thamizh letters in the string
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

// Iterator over letters making up Thamizh String
func (sa String) Letters() iter.Seq[Letter] {
	mapp := func(idx uint8) Letter { return Letter{idx: idx} }
	return mapper(slices.Values(sa.idxs), mapp)
}

func (s String) TailTrimmedRaw(tailStr String) (String, bool) {
	strLen, tailLen := len(s.idxs), len(tailStr.idxs)
	if strLen < tailLen {
		return s, false
	}
	startIdx := strLen - tailLen // sub-string match start index, along the given string
	if !matchIdxSlices(s.idxs[startIdx:], tailStr.idxs) {
		return s, false
	}
	return String{idxs: s.idxs[:startIdx]}, true
}

// Get (tail string) trimmed string on successful tail string match
func (s String) TailTrimmed(tailStr String) (String, bool) {
	strLen, tailLen := len(s.idxs), len(tailStr.idxs)
	if strLen < tailLen {
		return s, false
	}
	startIdx := strLen - tailLen // sub-string match start index, along the given string
	// Simple idx match should do! Except for the first idx.
	if tailLen > 1 {
		if !matchIdxSlices(s.idxs[startIdx+1:], tailStr.idxs[1:]) {
			return s, false
		}
	}
	if s.idxs[startIdx] == tailStr.idxs[0] {
		return String{idxs: s.idxs[:startIdx]}, true // Reusing the original slice!
	}

	// Check if we have CV and V match...
	letterV := tailStr.FirstLetter()
	if !letterV.IsV() {
		return s, false
	}
	letterCV := s.LetterAt(startIdx)
	if !letterCV.IsCV() {
		return s, false
	}
	c, v := letterCV.SplitCV()
	if v.idx != letterV.idx {
		return s, false
	}
	// Carefully compose the trimmed string
	// Make a copy of the slice, since we are going to modify it
	startIdx++ // Account for the overlapping CV/V letters
	idxs2 := make([]uint8, startIdx)
	copy(idxs2, s.idxs[:startIdx])
	idxs2[startIdx-1] = c.idx
	return String{idxs: idxs2}, true
}

func (s String) HeadTrimmedRaw(headStr String) (String, bool) {
	strLen, headLen := len(s.idxs), len(headStr.idxs)
	if strLen < headLen {
		return s, false
	}
	if !matchIdxSlices(s.idxs[:headLen], headStr.idxs) {
		return s, false
	}
	return String{idxs: s.idxs[headLen:]}, true
}

// Get (head string) trimmed string on successful head string match
func (s String) HeadTrimmed(headStr String) (String, bool) {
	panic("TODO")
}

// Get (replacement tail string) replaced string on successful tail string match
func (s String) TailReplaced(tailStr String, tailFix String) (String, bool) {
	st, ok := s.TailTrimmed(tailStr)
	if !ok {
		return s, false
	}
	return st.Appended(tailFix), true
}

// Simple concatenation! (No CV join is attempted)
func (s String) AppendedRaw(s2 String) String {
	idxs := make([]uint8, len(s.idxs)+len(s2.idxs))
	copy(idxs, s.idxs)
	copy(idxs[len(s.idxs):], s2.idxs)
	return String{idxs: idxs}
}

// Appended given string to this string
func (s String) Appended(s2 String) String {
	lenA := len(s.idxs)
	letterA := s.LastLetter()
	letterB := s2.FirstLetter()
	if !letterA.IsC() || !letterB.IsV() {
		return s.AppendedRaw(s2)
	}
	idxs := make([]uint8, lenA+len(s2.idxs)-1)
	copy(idxs, s.idxs)
	copy(idxs[lenA-1:], s2.idxs)
	idxs[lenA-1] = letterA.JoinCV(letterB).idx
	return String{idxs: idxs}
}

func (s String) StartsWith(headMatch String) bool {
	if headMatch.Len() < s.Len() {
		return false
	}
	panic("TODO")
}

func (s String) StartsExactWith(headMatch String) bool {
	if headMatch.Len() < s.Len() {
		return false
	}
	return matchIdxSlices(s.idxs[:headMatch.Len()], headMatch.idxs)
}

func (s String) EndsWith(tailMatch String) bool {
	if tailMatch.Len() < s.Len() {
		return false
	}
	panic("TODO")
}

func (s String) EndsExactWith(tailMatch String) bool {
	if tailMatch.Len() < s.Len() {
		return false
	}
	return matchIdxSlices(s.idxs[s.Len()-tailMatch.Len():], tailMatch.idxs)
}

func (s String) Syllables() []String {
	// Note: Only CV and V both may be followed by 0 or more Cs...
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

// Indicates if the given letter index slices have the same elements
func matchIdxSlices(idxs, idxs2 []uint8) bool { // TODO: Create generics variant
	if len(idxs) != len(idxs2) {
		panic("expected same length for both the slices")
	}
	for i, v := range idxs {
		if idxs2[i] != v {
			return false
		}
	}
	return true
}
