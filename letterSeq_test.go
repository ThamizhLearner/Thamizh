package script_test

import (
	"testing"

	script "github.com/ThamizhLearner/Thamizh"
)

// [c] API coverage (Ideal: API must have full test coverage)

// [c] Empty letter sequence
func TestEmptyLetterSeq(t *testing.T) {
	t.Run("Default LetterSeq construction", func(t *testing.T) {
		seq := script.LetterSeq{}
		if !seq.IsEmpty() {
			t.FailNow()
		}
	})

	t.Run("Empty LetterSeq initialization", func(t *testing.T) {
		seq := script.NewLetterSeq(nil)
		if !seq.IsEmpty() {
			t.FailNow()
		}

		seq2 := script.NewLetterSeq(make([]script.Letter, 0, 1))
		if !seq2.IsEmpty() {
			t.FailNow()
		}
	})
}

// [c] Letter sequence initialization & Unicode interop
func TestLetterSeqInit(t *testing.T) {
	t.Run("LetterSeq initialization", func(t *testing.T) {
		seq := script.MustLetterSeqFrom("தமிழ்")
		if seq.Len() != 3 || seq.String() != "தமிழ்" {
			t.FailNow()
		}
	})
}

// [c] Extract single Letter
func TestLetterIndexers(t *testing.T) {
	t.Run("Letter indexers", func(t *testing.T) {
		seq := script.MustLetterSeqFrom("தமிழ்")

		// First letter from the letter sequence
		if !seq.First().Equals(script.MustLetterFrom("த")) {
			t.FailNow()
		}

		// Last letter from the letter sequence
		if !seq.Last().Equals(script.MustLetterFrom("ழ்")) {
			t.FailNow()
		}

		// Nth letter from the letter sequence
		for i, ustr := range []string{"த", "மி", "ழ்"} {
			if !seq.Nth(i).Equals(script.MustLetterFrom(ustr)) {
				t.FailNow()
			}
		}
	})
}

// [c] Iter - Letters
func TestLetterIterators(t *testing.T) {
	t.Run("Index, Letter iterator", func(t *testing.T) {
		seq := script.MustLetterSeqFrom("தமிழ்")

		uLetterStrs := []string{"த", "மி", "ழ்"}

		for i, letter := range seq.IndexedLetters() {
			if !letter.Equals(script.MustLetterFrom(uLetterStrs[i])) {
				t.FailNow()
			}
		}
	})

	t.Run("Letter iterator", func(t *testing.T) {
		seq := script.MustLetterSeqFrom("தமிழ்")

		uLetterStrs := []string{"த", "மி", "ழ்"}

		i := 0
		for letter := range seq.Letters() {
			if !letter.Equals(script.MustLetterFrom(uLetterStrs[i])) {
				t.FailNow()
			}
			i++
		}
	})
}

// [c] Take* - Sub-range letter sequence
func TestLetterSubSequencing(t *testing.T) {
	t.Run("Take First N", func(t *testing.T) {
		seq := script.MustLetterSeqFrom("தமிழ்")
		seq2 := seq.TakeFirstN(0)
		if seq2.Len() != 0 || seq2.String() != "" {
			t.FailNow()
		}
		seq2 = seq.TakeFirstN(3)
		if seq2.Len() != 3 || seq2.String() != "தமிழ்" {
			t.FailNow()
		}
		seq2 = seq.TakeFirstN(2)
		if seq2.Len() != 2 || seq2.String() != "தமி" {
			t.FailNow()
		}
	})

	t.Run("Take Last N", func(t *testing.T) {
		seq := script.MustLetterSeqFrom("தமிழ்")
		seq2 := seq.TakeLastN(0)
		if seq2.Len() != 0 || seq2.String() != "" {
			t.FailNow()
		}
		seq2 = seq.TakeLastN(3)
		if seq2.Len() != 3 || seq2.String() != "தமிழ்" {
			t.FailNow()
		}
		seq2 = seq.TakeLastN(2)
		if seq2.Len() != 2 || seq2.String() != "மிழ்" {
			t.FailNow()
		}
	})

	t.Run("Take Range N", func(t *testing.T) {
		seq := script.MustLetterSeqFrom("தமிழ்")
		seq2 := seq.TakeRangeN(0, 0)
		if seq2.Len() != 0 || seq2.String() != "" {
			t.FailNow()
		}
		seq2 = seq.TakeRangeN(0, 3)
		if seq2.Len() != 3 || seq2.String() != "தமிழ்" {
			t.FailNow()
		}
		seq2 = seq.TakeRangeN(1, 3)
		if seq2.Len() != 2 || seq2.String() != "மிழ்" {
			t.FailNow()
		}
		seq2 = seq.TakeRangeN(2, 2)
		if seq2.Len() != 0 || seq2.String() != "" {
			t.FailNow()
		}
	})
}

// [c] Append* - Join LetterStr, Letter
func TestLetterSeqExtend(t *testing.T) {
	seq := script.MustLetterSeqFrom("கசட")
	seq2 := script.MustLetterSeqFrom("தபற")
	seqRes := seq.Appended(seq2)
	if seqRes.String() != "கசடதபற" {
		t.FailNow()
	}
	seqRes = seq.LetterAppended(script.MustLetterFrom("த"))
	if seqRes.String() != "கசடத" {
		t.FailNow()
	}
	seqRes = seq.LetterPrepended(script.MustLetterFrom("த"))
	if seqRes.String() != "தகசட" {
		t.FailNow()
	}
}

// [ ] Match - Letter sequence
func TestLetterSeqMatch(t *testing.T) {
	t.Run("Quick match", func(t *testing.T) {
		seq := script.MustLetterSeqFrom("தமிழ்")
		seq2 := script.MustLetterSeqFrom("தமிழ்")
		if !seq.Equals(seq2) {
			t.FailNow()
		}
		if seq.IndexOf(seq2) != 0 {
			t.FailNow()
		}
		if seq.LastIndexOf(seq2) != 0 {
			t.FailNow()
		}
		if !seq.Contains(seq2) {
			t.FailNow()
		}
		if !seq.StartsWith(seq2) {
			t.FailNow()
		}
		if !seq.EndsWith(seq2) {
			t.FailNow()
		}
	})

}
