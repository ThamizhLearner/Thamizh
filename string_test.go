package script_test // Black box test

import (
	"testing"

	script "github.com/ThamizhLearner/Thamizh"
)

func TestDecode(t *testing.T) {
	ustrs := []string{
		"அ", "க", "ழ்",
		"தமிழ்",
		"ஒட்டுமொத்தமாகப்பார்த்துக்கொண்டிருந்தாள்",
	}
	for _, ustr := range ustrs {
		s := script.MustDecode(ustr)
		if s.String() != ustr {
			t.Errorf("Round trip: Expected %s Got %s", ustr, s)
		}
	}
}

func TestLetterAt(t *testing.T) {
	s := script.MustDecode("தமிழ்")
	if s.FirstLetter().String() != "த" {
		t.Errorf("First letter")
	}
	if s.LastLetter().String() != "ழ்" {
		t.Errorf("Last letter")
	}
	for i, expectStr := range []string{"த", "மி", "ழ்"} {
		if s.LetterAt(i).String() != expectStr {
			t.Errorf("Letter At %d, expected %s, got %s", i, expectStr, s.LetterAt(i))
		}
	}
}

func TestLetterIter(t *testing.T) {
	s := script.MustDecode("தமிழ்")
	lStrs := []string{"த", "மி", "ழ்"}
	i := 0
	for letter := range s.Letters() {
		if letter.String() != lStrs[i] {
			t.Errorf("Iteration At %d, expected %s, got %s", i, lStrs[i], letter)
		}
		i++
	}
}

func TestTrimEnd(t *testing.T) {
	tests := []struct {
		ustr  string
		utrim string
		ures  string
	}{
		{"தமிழிலக்கனம்", "இலக்கனம்", "தமிழ்"},
	}
	for i, tc := range tests {
		s := script.MustDecode(tc.ustr)
		trimStr := script.MustDecode(tc.utrim)
		resStr, ok := s.TrimEnd(trimStr)
		if !ok || resStr.String() != tc.ures {
			t.Errorf("Trim end %d, expected %s, got %s", i, tc.ures, resStr)
		}
	}
}

func TestAppend(t *testing.T) {
	tests := []struct {
		ustr    string
		uappend string
		ures    string
	}{
		{"தமிழ்", "இலக்கனம்", "தமிழிலக்கனம்"},
	}
	for i, tc := range tests {
		s := script.MustDecode(tc.ustr)
		appendStr := script.MustDecode(tc.uappend)
		resStr := s.Append(appendStr)
		if resStr.String() != tc.ures {
			t.Errorf("Append %d, expected %s, got %s", i, tc.ures, resStr)
		}
	}
}

func TestSyllabification(t *testing.T) {
	tests := []struct {
		inp  string
		want string
	}{
		{"அ", "அ"},
		{"ழ்", "ழ்"},
		{"க", "க"},
		{"தமிழ்", "த-மிழ்"},
		{"ஒட்டுமொத்தமாகப்பார்த்துக்கொண்டிருந்தாள்", "ஒட்-டு-மொத்-த-மா-கப்-பார்த்-துக்-கொண்-டி-ருந்-தாள்"},
	}
	for _, testCase := range tests {
		got := script.MustDecode(testCase.inp).SyllabifiedUStr("-")
		if got != testCase.want {
			t.Errorf("Syllabification: Expected %s, Got %s", testCase.want, got)
		}
	}
}
