// Thamizh letter string model

package script

import (
	"iter"
	"slices"

	"github.com/ThamizhLearner/Thamizh/internal/unicode"
)

// The Thamizh letter string
type String struct {
	idxs []uint8
}

func TryDecode(s string) *String {
	idxs := unicode.Decode(s)
	if idxs == nil {
		return nil
	}
	return &String{idxs: idxs}
}

func MustDecode(s string) String { return String{idxs: unicode.Decode(s)} }

func (s String) String() string { return unicode.Encode(s.idxs) }

func Map[V, U any](s iter.Seq[V], mapFn func(V) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for item := range s {
			if !yield(mapFn(item)) {
				return
			}
		}
	}
}

func (s String) Values() iter.Seq[Letter] {
	mapp := func(idx uint8) Letter { return Letter{idx: idx} }
	return Map(slices.Values(s.idxs), mapp)
}

// Standard string-like methods

func (s String) Clone() String {
	tmp := make([]uint8, len(s.idxs))
	copy(tmp, s.idxs)
	return String{idxs: tmp}
}
func (s String) Compare(s2 String) int               { panic("") }
func (s String) HasPrefix(prefix String) bool        { panic("") }
func (s String) HasSuffix(suffix String) bool        { panic("") }
func (s String) TrimPrefix(prefix String) String     { panic("") }
func (s String) TrimSuffix(suffix String) String     { panic("") }
func (s String) PrefixTrimmed(prefix String) *String { panic("") }
func (s String) SuffixTrimmed(suffix String) *String { panic("") }

// Additional methods

func (s String) Syllables() []String { panic("") }
