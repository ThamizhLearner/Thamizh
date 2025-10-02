package script

import (
	"iter"
	"slices"
)

// Unicode code point U+0BCD
const UAttachedDot rune = '்'

// Thamizh consonant vocalization-strength enum
type VocalStrength uint8

const (
	// வல்லின எழுத்து; Strong vocalization-strength consonant
	Strong VocalStrength = iota

	// இடையின எழுத்து; Medium vocalization-strength consonant
	Medium

	// மெல்லின எழுத்து; Mild vocalization-strength consonant
	Mild
)

// Abstract Thamizh consonant Unicode code point
type UConsonant struct {
	// Thamizh consonant Unicode code point (w/o attached-dot decoration)
	uBaseForm rune

	// Thamizh consonant vocalization-strength enum
	vocalStrength VocalStrength
}

// Thamizh consonant representative string
func (c UConsonant) Rep() string { return string([]rune{c.uBaseForm, UAttachedDot}) }

// Thamizh base consonant representative string
func (c UConsonant) BaseFormRep() string { return string(c.uBaseForm) }

// Thamizh consonant vocalization-strength enum
func (c UConsonant) VocalStrength() VocalStrength { return c.vocalStrength }

// (Internal protected) Ordered Thamizh Unicode consonant list
var internalCees [18]UConsonant = [...]UConsonant{
	{'க', Strong},
	{'ங', Mild},
	{'ச', Strong},
	{'ஞ', Mild},
	{'ட', Strong},
	{'ண', Mild},
	{'த', Strong},
	{'ந', Mild},
	{'ப', Strong},
	{'ம', Mild},
	{'ய', Medium},
	{'ர', Medium},
	{'ல', Medium},
	{'வ', Medium},
	{'ழ', Medium},
	{'ள', Medium},
	{'ற', Strong},
	{'ன', Mild},
}

// Ordered abstract Thamizh consonant Unicode code points (indexed) iterator
func IndexedUConsonants() iter.Seq2[int, UConsonant] { return slices.All(internalCees[:]) }

// Ordered abstract Thamizh consonant Unicode code points iterator
func UConsonants() iter.Seq[UConsonant] { return slices.Values(internalCees[:]) }
