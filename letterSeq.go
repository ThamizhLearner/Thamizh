package script

import (
	"iter"
	"slices"
)

// Immutable sequence of letters.
type LetterSeq struct {
	letters []Letter
}

// -------------------------------------------------------------
// Unicode interop
// -------------------------------------------------------------

func (this LetterSeq) String() string { return letterSliceToUnicode(this.letters) }

// -------------------------------------------------------------
// LetterSeq initialization
// -------------------------------------------------------------

// Creates LetterSeq from specified array of letters.
func NewLetterSeq(letters []Letter) LetterSeq {
	// Note: Since the given array is mutable, we must clone it first.
	return LetterSeq{letters: slices.Clone(letters)}
}

// Note: Unsafe because it simply uses the given letter array itself, and not its clone!
func UnsafeNewLetterSeq(letters []Letter) LetterSeq { return LetterSeq{letters: letters} }

// -------------------------------------------------------------
// LetterSeq property query
// -------------------------------------------------------------

// Tells if this LetterSeq is empty.
func (this LetterSeq) IsEmpty() bool { return len(this.letters) == 0 }

// Count of letters in this LetterSeq.
func (this LetterSeq) Len() int { return len(this.letters) }

// -------------------------------------------------------------
// LetterSeq letter getters
// -------------------------------------------------------------

// Gets first letter in this LetterSeq.
func (this LetterSeq) First() Letter { return this.letters[0] }

// Gets last letter in this LetterSeq.
func (this LetterSeq) Last() Letter { return this.letters[len(this.letters)-1] }

// Gets nth letter in this LetterSeq.
func (this LetterSeq) Nth(idx int) Letter { return this.letters[idx] }

// -------------------------------------------------------------
// LetterSeq letters iterators
// -------------------------------------------------------------

// Gets an iterator over index-letter pairs in the usual order.
func (this LetterSeq) IndexedLetters() iter.Seq2[int, Letter] { return slices.All(this.letters) }

// Gets an iterator that yields the letters in order.
func (this LetterSeq) Letters() iter.Seq[Letter] { return slices.Values(this.letters) }

// -------------------------------------------------------------
// LetterSeq content manipulators
// [Creates new LetterSeq with modified content]
// -------------------------------------------------------------

// Creates new LetterSeq by appending specified LetterSeq to this LetterSeq.
func (this LetterSeq) Appended(seq LetterSeq) LetterSeq {
	lenA := len(this.letters)
	letters := make([]Letter, lenA+len(seq.letters)) // Create array of appropriate length
	copy(letters, this.letters)                      // Copy the first letter array letters
	copy(letters[lenA:], seq.letters)                // Copy the second letter array letters, at the appropriate offset
	return LetterSeq{letters: letters}               // Return newly constructed concatenated letter array
}

// Creates new LetterSeq by appending specified Letter to this LetterSeq.
func (this LetterSeq) LetterAppended(lt Letter) LetterSeq {
	lenA := len(this.letters)
	letters := make([]Letter, lenA+1)  // Create array of appropriate length
	copy(letters, this.letters)        // Copy the specified letter array letters
	letters[lenA] = lt                 // Copy the specified letter, at the appropriate offset
	return LetterSeq{letters: letters} // Return newly constructed concatenated letter array
}

// Creates new LetterSeq by prepending specified Letter to this LetterSeq.
func (this LetterSeq) LetterPrepended(lt Letter) LetterSeq {
	letters := make([]Letter, len(this.letters)+1) // Create array of appropriate length
	letters[0] = lt                                // Copy the specified letter first
	copy(letters[1:], this.letters)                // Copy the specified letter array letters, at the appropriate offset
	return LetterSeq{letters: letters}             // Return newly constructed concatenated letter array
}

// -------------------------------------------------------------
// LetterSeq ranged content extractors
// -------------------------------------------------------------

// Create letter sub-array by copying first N letters
func (this LetterSeq) TakeFirstN(count int) LetterSeq {
	return LetterSeq{letters: this.letters[:count]}
}

// Create letter sub-array by copying last N letters
func (this LetterSeq) TakeLastN(count int) LetterSeq {
	return LetterSeq{letters: this.letters[len(this.letters)-int(count):]}
}

// Create letter sub-array by coping N letters, starting at the specified start offset
func (this LetterSeq) TakeRangeN(startIdx int, count int) LetterSeq {
	return LetterSeq{letters: this.letters[startIdx : startIdx+count]}
}

// -------------------------------------------------------------
// LetterSeq content probe methods
// -------------------------------------------------------------

func (this LetterSeq) Equals(seq LetterSeq) bool {
	if len(this.letters) != len(seq.letters) {
		return false
	}
	for i := range len(this.letters) {
		if !this.letters[i].Equals(seq.letters[i]) {
			return false
		}
	}
	return true
}

// Gets the offset index of specified letter array in this letter array; or -1
func (this LetterSeq) IndexOf(seq LetterSeq) int {
	lenA := len(this.letters)
	lenB := len(seq.letters)
	if lenA < lenB {
		return -1
	}
	for i := range (lenA - lenB) + 1 { // Possible match checks
		match := true
		for j := range lenB {
			if this.letters[i+j] == seq.letters[j] {
				continue
			}
			match = false
		}
		if match {
			return i
		}
	}
	return -1
}

// Gets the offset index of specified letter array in this letter array; or -1
func (this LetterSeq) LastIndexOf(seq LetterSeq) int {
	lenA := len(this.letters)
	lenB := len(seq.letters)
	if lenA < lenB {
		return -1
	}

	for i := (lenA - lenB); i >= 0; i-- { // Possible match checks
		match := true
		for j := range lenB {
			if this.letters[i+j] == seq.letters[j] {
				continue
			}
			match = false
		}
		if match {
			return i
		}
	}
	return -1
}

// Tells if specified letter array is contained in this letter array
func (this LetterSeq) Contains(seq LetterSeq) bool { return this.IndexOf(seq) != -1 }

// Tell if this letter array has specified letter array as (exact) prefix
func (this LetterSeq) StartsWith(seq LetterSeq) bool {
	lenA := len(this.letters)
	lenB := len(seq.letters)
	if lenA < lenB {
		return false
	}
	for i := range lenB {
		if this.letters[i] != seq.letters[i] {
			return false
		}
	}
	return true
}

// Tell if this letter array has specified letter array as (exact) suffix
func (this LetterSeq) EndsWith(seq LetterSeq) bool {
	lenA := len(this.letters)
	lenB := len(seq.letters)
	if lenA < lenB {
		return false
	}
	offset := lenA - lenB
	for i := range lenB {
		if this.letters[offset+i] != seq.letters[i] {
			return false
		}
	}
	return true
}
