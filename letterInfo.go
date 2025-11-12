// Thamizh letter vocalization query

package script

import (
	base "github.com/ThamizhLearner/Thamizh/internal"
)

// Letter vocalization query bits
type qBitField uint8

const (
	bit_Strong qBitField = 1 << iota
	bit_Medium
	bit_Mild
	bit_Short
	bit_Long
)

func vocalStrengthQBit(vs base.VocalStrength) qBitField {
	switch vs {
	case base.Strong:
		return bit_Strong
	case base.Medium:
		return bit_Medium
	case base.Mild:
		return bit_Mild
	default:
		panic("Unexpected vocal strength value")
	}
}

func vocalDurationQBit(vd base.VocalDuration) qBitField {
	switch vd {
	case base.Short:
		return bit_Short
	case base.Long:
		return bit_Long
	default:
		panic("Unexpected vocal duration value")
	}
}

var qBitFields [246]qBitField = func() [246]qBitField {
	var fields [246]qBitField
	for i, v := range base.VowelDurations {
		fields[i] = vocalDurationQBit(v)
	}
	for i, c := range base.ConsonantStengths {
		fields[12+i] = vocalStrengthQBit(c)
	}
	for ci, c := range base.ConsonantStengths {
		offset := 30 + ci*12
		vstrength := vocalStrengthQBit(c)
		for vi, v := range base.VowelDurations {
			fields[offset+vi] = vstrength | vocalDurationQBit(v)
		}
	}
	return fields
}()
