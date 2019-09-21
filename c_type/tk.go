package c_type

import (
	"encoding/hex"
	"fmt"
)

type Tk [64]byte

var Empty_Tk = Tk{}

func (self Tk) NewRef() (ret *Tk) {
	ret = &Tk{}
	copy(ret[:], self[:])
	return ret
}

func (b Tk) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b[:])
	return result, nil
}

func (b *Tk) UnmarshalText(input []byte) error {
	raw := input[2:]
	if len(raw) == 0 {
		return nil
	}
	dec := Tk{}
	if len(raw)/2 != len(dec[:]) {
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(dec[:])*2, "Uint512")
	}
	if _, err := hex.Decode(dec[:], raw); err != nil {
		return err
	} else {
		*b = dec
	}
	return nil
}
