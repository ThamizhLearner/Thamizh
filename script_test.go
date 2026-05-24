package script_test

import (
	"testing"

	script "github.com/ThamizhLearner/Thamizh"
)

func TestSuffixTrimmed(t *testing.T) {
	seq := script.MustLetterSeqFrom("வழிகள்")
	sfx := script.MustLetterSeqFrom("கள்")

	res, ok := script.SuffixTrimmed(seq, sfx)
	if !ok || res.String() != "வழி" {
		t.FailNow()
	}

	// Full trim
	res, ok = script.SuffixTrimmed(seq, seq)
	if !ok || res.Len() != 0 {
		t.FailNow()
	}

	// No trim
	sfx2 := script.MustLetterSeqFrom("கல்")
	_, ok = script.SuffixTrimmed(seq, sfx2)
	if ok {
		t.FailNow()
	}
}

func TestSuffixAppended(t *testing.T) {
	seq := script.MustLetterSeqFrom("வழி")
	sfx := script.MustLetterSeqFrom("கள்")

	res := script.SuffixAppended(seq, sfx)
	if res.String() != "வழிகள்" {
		t.FailNow()
	}
}

func TestVSuffixTrimmed(t *testing.T) {
	seq := script.MustLetterSeqFrom("தமிழிலக்கணம்")
	sfx := script.MustLetterSeqFrom("இலக்கணம்")

	res, ok := script.VSuffixTrimmed(seq, sfx)
	if !ok || res.String() != "தமிழ்" {
		t.FailNow()
	}

	seq = script.MustLetterSeqFrom("ண")
	sfx = script.MustLetterSeqFrom("அ")
	res, ok = script.VSuffixTrimmed(seq, sfx)
	if !ok || res.String() != "ண்" {
		t.FailNow()
	}
}

func TestVSuffixAppended(t *testing.T) {
	seq := script.MustLetterSeqFrom("தமிழ்")
	sfx := script.MustLetterSeqFrom("இலக்கணம்")

	res, ok := script.VSuffixAppended(seq, sfx)
	if !ok || res.String() != "தமிழிலக்கணம்" {
		t.FailNow()
	}

	seq = script.MustLetterSeqFrom("ண்")
	sfx = script.MustLetterSeqFrom("அ")
	res, ok = script.VSuffixAppended(seq, sfx)
	if !ok || res.String() != "ண" {
		t.FailNow()
	}
}

func TestSyllabification(t *testing.T) {
	seq := script.MustLetterSeqFrom("தமிழ்")
	res, c := script.SyllabifiedUStr(seq, "-")
	if c != 2 || res != "த-மிழ்" {
		t.FailNow()
	}
}
