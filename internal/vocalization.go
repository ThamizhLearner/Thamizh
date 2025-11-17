// Thamizh letter vocal characteristics

package base

// Thamizh vowel vocalization-duration enum
type VocalDuration uint8

// Thamizh vowel vocalization-duration enums
const (
	// குறில் எழுத்து; Short vocalization-duration vowel
	Short VocalDuration = iota

	// நெடில் எழுத்து; Long vocalization-duration vowel
	Long
)

// Thamizh consonant vocalization-strength enum
type VocalStrength uint8

// Thamizh consonant vocalization-strength enums
const (
	// வல்லின எழுத்து; Strong vocalization-strength consonant
	Strong VocalStrength = iota

	// இடையின எழுத்து; Medium vocalization-strength consonant
	Medium

	// மெல்லின எழுத்து; Mild vocalization-strength consonant
	Mild
)

var VowelDurations = [12]VocalDuration{
	Short, Long, Short, Long, Short, Long,
	Short, Long, Long, Short, Long, Long,
}

var ConsonantStengths = [18]VocalStrength{
	Strong, Mild, Strong, Mild, Strong, Mild,
	Strong, Mild, Strong, Mild, Medium, Medium,
	Medium, Medium, Medium, Medium, Strong, Mild,
}
