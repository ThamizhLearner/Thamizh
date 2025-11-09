package unicode

import (
	"fmt"
)

// Range check validation for the give Unicode code point rune
func validURange(r rune) bool { return r >= uBlockBase && r < uBlockEnd }

// Gets annotated code point coresponding to given "valid" rune
func getAnnotation(r rune) annoCode {
	if validURange(r) {
		return annotatedCodes[r-uBlockBase]
	}
	return annoCode{}
}

func getAnnotations(s string) []annoCode {
	var codes []annoCode
	for _, r := range s {
		var anno annoCode
		if validURange(r) {
			anno = annotatedCodes[r-uBlockBase]
		}
		codes = append(codes, anno)
	}
	return codes
}

func DumpAnnotations(s string) {
	for _, a := range getAnnotations(s) {
		fmt.Println(a)
	}
}

// Decodes Unicode string; returns Thamizh letter index slice
func Decode(s string) []uint8 {
	// Only the BaseConsonant may be followed by atmost one AttachedDot/AttachedVowel
	prev := annoCode{}
	var idxs []uint8
	for _, curr := range getAnnotations(s) {
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
			return nil // Invalid input: Not-supported code-point annotation
		}
		prev = curr
	}
	if prev.group == tBaseConsonant {
		idxs = append(idxs, 30+prev.idx*12)
	}
	return idxs
}
