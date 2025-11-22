package script_test // Black box test

import (
	"slices"
	"testing"

	script "github.com/ThamizhLearner/Thamizh"
)

func TestLetterType(t *testing.T) {
	s := script.MustDecode("தமிழ்")
	letters := slices.Collect(s.Letters())
	if !letters[0].IsCV() {
		t.Errorf("Expected vowelized-consonant CV")
	}
	if !letters[2].IsC() {
		t.Errorf("Expected consonant C")
	}
	vowelLetter := script.MustNewLetter("அ")
	if !vowelLetter.IsV() {
		t.Errorf("Expected vowel V")
	}
}

func TestLetterIs(t *testing.T) {
	letters := slices.Collect(script.MustDecode("இல்").Letters())
	if !letters[0].IsLetter("இ") {
		t.Errorf("Letter IsLetter")
	}

	testLetter := script.MustNewLetter("இ")
	if !letters[0].Is(testLetter) {
		t.Errorf("Letter Is")
	}
}

func TestVocalization(t *testing.T) {
	letter := script.MustNewLetter("அ")
	if !letter.IsShortVocal() {
		t.Errorf("Short vocal letter")
	}

	letter = script.MustNewLetter("க்")
	if !letter.IsStrongVocal() {
		t.Errorf("Strong vocal letter")
	}

	letter = script.MustNewLetter("க")
	if !letter.IsStrongVocal() || !letter.IsShortVocal() {
		t.Errorf("Short and strong vocal lettter")
	}
}

func TestSplitJoinCV(t *testing.T) {
	cv := script.MustNewLetter("க")
	c, v := cv.SplitCV()
	if !c.IsLetter("க்") && !v.IsLetter("அ") {
		t.Errorf("Split CV")
	}

	cv = c.JoinCV(v)
	if !cv.IsLetter("க") {
		t.Errorf("Join CV")
	}
}
