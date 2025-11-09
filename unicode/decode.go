package unicode

import (
	"fmt"
	"strings"

	base "github.com/ThamizhLearner/Thamizh/unicode/internal"
)

const (
	uBlockBase rune  = 0x0B80 // Start of Thamizh Unicode block
	codesCount uint8 = 80     // Count of annotated code points
	uBlockEnd  rune  = uBlockBase + rune(codesCount)
)

type tGroup uint8

const (
	tNone          tGroup = iota // Out of range code point group
	tDetachedDot                 // Thamizh detached-dot code point group
	tPrimaryVowel                // Thamizh primary vowel code point group
	tBaseConsonant               // Thamizh base-consonant code point group
	tDetachedVowel               // Thamizh detached-vowel code point group
)

// Annotated code points
type annoCode struct {
	group tGroup // Code group
	idx   uint8  // Index within the code group
}

// Implicit Stringer interface implementation
func (ac annoCode) String() string {
	var sb strings.Builder
	switch ac.group {
	case tNone:
		sb.WriteString("None")
	case tDetachedDot:
		sb.WriteString("DetachedDot")
		if ac.idx != 0 {
			panic("Internal error")
		}
		sb.WriteString(fmt.Sprintf(" %c", UDot))
	case tPrimaryVowel:
		sb.WriteString("PrimaryVowel")
		sb.WriteString(fmt.Sprintf(" %c", base.PrimaryVowelUCodes[ac.idx]))
	case tBaseConsonant:
		sb.WriteString("BaseConsonant")
		sb.WriteString(fmt.Sprintf(" %c", base.BaseConsonantUCodes[ac.idx]))
	case tDetachedVowel:
		sb.WriteString("DetachedVowel")
		sb.WriteString(fmt.Sprintf(" %c", base.AttachedVowelUCodes[ac.idx]))
	default:
		panic("")
	}
	sb.WriteString(fmt.Sprintf(" %v", ac.idx))
	return sb.String()
}

// Note: Input must be Thamizh Unicode code point
func annoCodeIdx(r rune) uint8 { return uint8(r - uBlockBase) }

// Annotated code points
var ucodes [codesCount]annoCode = func() [codesCount]annoCode {
	var annoCodes [codesCount]annoCode
	annoCodes[annoCodeIdx(UDot)] = annoCode{group: tDetachedDot}
	for i, v := range base.PrimaryVowelUCodes {
		annoCodes[annoCodeIdx(v)] = annoCode{group: tPrimaryVowel, idx: uint8(i)}
	}
	for i, a := range base.AttachedVowelUCodes[1:] {
		i++ // Compensate for sub-slicing [1:]
		annoCodes[annoCodeIdx(a)] = annoCode{group: tDetachedVowel, idx: uint8(i)}
	}
	for i, c := range base.BaseConsonantUCodes {
		annoCodes[annoCodeIdx(c)] = annoCode{group: tBaseConsonant, idx: uint8(i)}
	}
	return annoCodes
}()

// Range check validation for the give Unicode code point rune
func validURange(r rune) bool { return r >= uBlockBase && r < uBlockEnd }

// Gets annotated code point corresponding to given "valid" rune
func getAnnotation(r rune) annoCode {
	if validURange(r) {
		return ucodes[r-uBlockBase]
	}
	return annoCode{}
}

func getAnnotations(s string) []annoCode {
	var codes []annoCode
	for _, r := range s {
		codes = append(codes, getAnnotation(r))
	}
	return codes
}

func DumpAnnotations(s string) {
	for _, a := range getAnnotations(s) {
		fmt.Println(a)
	}
}

func Decode(s string) []uint8 {
	// Only the BaseConsonant may be followed by atmost one AttachedDot/AttachedVowel
	prev := annoCode{}
	var idxs []uint8
	for _, r := range s {
		curr := getAnnotation(r)
		switch curr.group {
		case tBaseConsonant:
			if prev.group == tBaseConsonant {
				idxs = append(idxs, 30+prev.idx*12)
			}
		case tPrimaryVowel:
			if prev.group == tBaseConsonant {
				idxs = append(idxs, 30+prev.idx*12)
			}
			idxs = append(idxs, curr.idx)
		case tDetachedVowel, tDetachedDot:
			if prev.group != tBaseConsonant {
				return nil // Invalid input: Missing base-consonant
			}
			if curr.group == tDetachedDot {
				idxs = append(idxs, 12+prev.idx)
			} else {
				idxs = append(idxs, 30+prev.idx*12+curr.idx)
			}
		case tNone:
			return nil // Invalid input: Unexpected annotation
		}
		prev = curr
	}
	if prev.group == tBaseConsonant {
		idxs = append(idxs, 30+prev.idx*12)
	}
	return idxs
}
