// Unicode decoder/encoder mechanism support base
// Thamizh (Unicode-encoded) string <=> Thamizh letter-index sequence

package base

const (
	UNull rune = '\u0000' // Unicode 'Null' code point
)

// Unicode code points corresponding to Thamizh vowel (letter) characters
var PrimaryVowelUCodes = [12]rune{
	'அ', 'ஆ', 'இ', 'ஈ', 'உ', 'ஊ',
	'எ', 'ஏ', 'ஐ', 'ஒ', 'ஓ', 'ஔ',
}

// Unicode code points corresponding to Thamizh attached-vowel (non-letter) characters
var AttachedVowelUCodes = [12]rune{
	UNull, 'ா', 'ி', 'ீ', 'ு', 'ூ',
	'ெ', 'ே', 'ை', 'ொ', 'ோ', 'ௌ',
}

// Unicode code points corresponding to Thamizh base-consonant (letter) characters
var BaseConsonantUCodes = [18]rune{
	'க', 'ங', 'ச', 'ஞ', 'ட', 'ண',
	'த', 'ந', 'ப', 'ம', 'ய', 'ர',
	'ல', 'வ', 'ழ', 'ள', 'ற', 'ன',
}
