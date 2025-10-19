// Thamizh Unicode consonant representation correctness verification

package script_test // Black-box test package

import (
	"testing"

	script "github.com/ThamizhLearner/Thamizh"
)

func TestUConsonant(t *testing.T) {
	// வரிசை படி மெய் எழுத்துக்கள்; Ordered consonant list (w/o dot decoration)
	var ucodes [18]rune = [...]rune{
		'க', 'ங', 'ச', 'ஞ', 'ட', 'ண',
		'த', 'ந', 'ப', 'ம', 'ய', 'ர',
		'ல', 'வ', 'ழ', 'ள', 'ற', 'ன',
	}
	for i, uconsonant := range script.IndexedUConsonants() {
		got, want := uconsonant.BaseFormRep(), ucodes[i]
		if got != want {
			t.Errorf("UConsonant order mismatch; got '%c', want '%c'", got, want)
		}
	}
}

func TestUConsonant2(t *testing.T) {
	// வரிசை படி மெய் எழுத்துக்கள்; Ordered consonant list
	var ucodes [18]string = [...]string{
		"க்", "ங்", "ச்", "ஞ்", "ட்", "ண்",
		"த்", "ந்", "ப்", "ம்", "ய்", "ர்",
		"ல்", "வ்", "ழ்", "ள்", "ற்", "ன்",
	}
	for i, uconsonant := range script.IndexedUConsonants() {
		got, want := uconsonant.Rep(), ucodes[i]
		if got != want {
			t.Errorf("UConsonant order mismatch; got '%s', want '%s'", got, want)
		}
	}
}

func TestStrongUConsonant(t *testing.T) {
	// வரிசை படி வல்லின மெய் எழுத்துக்கள்; Ordered strong vocalization-strength consonant list
	var ucodes [6]rune = [...]rune{'க', 'ச', 'ட', 'த', 'ப', 'ற'}

	var itemIdx uint = 0
	for c := range script.UConsonants() {
		if c.VocalStrength() != script.Strong {
			continue
		}
		got, want := c.BaseFormRep(), ucodes[itemIdx]
		if got != want {
			t.Errorf("Strong UConsonant order mismatch; got '%c', want '%c'", got, want)
		}
		itemIdx++
	}
}

func TestMildUConsonant(t *testing.T) {
	// வரிசை படி மெல்லின மெய் எழுத்துக்கள்; Ordered mild vocalization-strength consonant list
	var ucodes [6]rune = [...]rune{'ங', 'ஞ', 'ண', 'ந', 'ம', 'ன'}

	var itemIdx uint = 0
	for c := range script.UConsonants() {
		if c.VocalStrength() != script.Mild {
			continue
		}
		got, want := c.BaseFormRep(), ucodes[itemIdx]
		if got != want {
			t.Errorf("Mild UConsonant order mismatch; got '%c', want '%c'", got, want)
		}
		itemIdx++
	}
}

func TestMediumUConsonant(t *testing.T) {
	// வரிசை படி இடையின மெய் எழுத்துக்கள்; Ordered medium vocalization-strength consonant list
	var ucodes [6]rune = [...]rune{'ய', 'ர', 'ல', 'வ', 'ழ', 'ள'}

	var itemIdx uint = 0
	for c := range script.UConsonants() {
		if c.VocalStrength() != script.Medium {
			continue
		}
		got, want := c.BaseFormRep(), ucodes[itemIdx]
		if got != want {
			t.Errorf("Medium UConsonant order mismatch; got '%c', want '%c'", got, want)
		}
		itemIdx++
	}
}
