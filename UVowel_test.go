package script_test // Black-box test package

import (
	"testing"

	script "github.com/ThamizhLearner/Thamizh"
)

func TestUVowel(t *testing.T) {
	// வரிசை படி உயிர் எழுத்துக்கள்; Ordered vowel list (w/o dot decoration)
	var ucodes [12]rune = [...]rune{
		'அ', 'ஆ', 'இ', 'ஈ', 'உ', 'ஊ',
		'எ', 'ஏ', 'ஐ', 'ஒ', 'ஓ', 'ஔ',
	}
	for i, uvowel := range script.IndexedUVowels() {
		got, want := uvowel.Rep(), string(ucodes[i])
		if got != want {
			t.Errorf("UVowel order mismatch; idx %v, got '%s', want '%s'", i, got, want)
		}
	}
}

func TestAttachedUVowel(t *testing.T) {
	var ucodes [12]rune = [...]rune{
		script.UNull, 'ா', 'ி', 'ீ', 'ு', 'ூ',
		'ெ', 'ே', 'ை', 'ொ', 'ோ', 'ௌ',
	}
	for i, uvowel := range script.IndexedUVowels() {
		if i == 0 { // First attached-vowel form does not exist
			got, want := uvowel.AttachedFormRep(), ""
			if got != want {
				t.Errorf("Attached UVowel order mismatch; idx %v, got '%s', want '%s'", i, got, want)
			}
			continue
		}
		got, want := uvowel.AttachedFormRep(), string(ucodes[i])
		if got != want {
			t.Errorf("Attached UVowel order mismatch; idx %v, got '%s', want '%s'", i, got, want)
		}
	}
}

func TestShortUVowel(t *testing.T) {
	// வரிசை படி குறில் உயிர் எழுத்துக்கள்; Ordered short vocalization-duration vowel list
	var ucodes [5]rune = [...]rune{'அ', 'இ', 'உ', 'எ', 'ஒ'}

	var itemIdx uint = 0
	for v := range script.UVowels() {
		if v.VocalDuration() != script.Short {
			continue
		}
		got, want := v.Rep(), string(ucodes[itemIdx])
		if got != want {
			t.Errorf("Short UVowel order mismatch; idx %v, got '%s', want '%s'", itemIdx, got, want)
		}
		itemIdx++
	}
}

func TestLongUVowel(t *testing.T) {
	// வரிசை படி நெடில் உயிர் எழுத்துக்கள்; Ordered long vocalization-duration vowel list
	var ucodes [7]rune = [...]rune{'ஆ', 'ஈ', 'ஊ', 'ஏ', 'ஐ', 'ஓ', 'ஔ'}

	var itemIdx uint = 0
	for v := range script.UVowels() {
		if v.VocalDuration() != script.Long {
			continue
		}
		got, want := v.Rep(), string(ucodes[itemIdx])
		if got != want {
			t.Errorf("Long UVowel order mismatch idx %v, got '%s', want '%s'", itemIdx, got, want)
		}
		itemIdx++
	}
}
