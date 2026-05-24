package script_test // Black box test

import (
	"slices"
	"testing"

	script "github.com/ThamizhLearner/Thamizh"
)

func TestLetterType(t *testing.T) {
	seq := script.MustLetterSeqFrom("தமிழ்")
	letters := slices.Collect(seq.Letters())
	if !letters[0].IsCV() {
		t.Errorf("Expected vowelized-consonant CV property")
	}
	if !letters[2].IsC() {
		t.Errorf("Expected consonant C property")
	}
	vowelLetter := script.MustLetterFrom("அ")
	if !vowelLetter.IsV() {
		t.Errorf("Expected vowel V property")
	}
}

func TestLetterIs(t *testing.T) {
	letters := slices.Collect(script.MustLetterSeqFrom("இல்").Letters())
	if !letters[0].IsLetter("இ") {
		t.Errorf("Letter IsLetter")
	}

	testLetter := script.MustLetterFrom("இ")
	if !letters[0].Is(testLetter) {
		t.Errorf("Letter Is")
	}
}

func TestVocalization(t *testing.T) {
	letter := script.MustLetterFrom("அ")
	if !letter.IsShortVocal() {
		t.Errorf("Short vocal letter")
	}

	letter = script.MustLetterFrom("க்")
	if !letter.IsStrongVocal() {
		t.Errorf("Strong vocal letter")
	}

	letter = script.MustLetterFrom("க")
	if !letter.IsStrongVocal() || !letter.IsShortVocal() {
		t.Errorf("Short and strong vocal lettter")
	}
}

func TestSplitJoinCV(t *testing.T) {
	cv := script.MustLetterFrom("க")
	c, v := cv.MustSplitCV()
	if !c.IsLetter("க்") && !v.IsLetter("அ") {
		t.Errorf("Split CV")
	}

	cv = c.MustJoinCV(v)
	if !cv.IsLetter("க") {
		t.Errorf("Join CV")
	}
}
