package c_type

import (
	"encoding/hex"
	"fmt"
)

type SignN [96]byte

var Empty_SignN = SignN{}

func (self *SignN) ToUint512() (ret Uint512) {
	copy(ret[:], self[:])
	return
}

func (self SignN) NewRef() (ret *PKr) {
	ret = &PKr{}
	copy(ret[:], self[:])
	return ret
}

func (b SignN) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b[:])
	return result, nil
}

func (b *SignN) UnmarshalText(input []byte) error {
	if len(input) < 2 {
		return fmt.Errorf("hex string length must > 2 : current is %d", len(input))
	}
	raw := input[2:]
	if len(raw) == 0 {
		return nil
	}
	dec := SignN{}
	if len(raw)/2 != len(dec[:]) {
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(dec[:])*2, "SignN")
	}
	if _, err := hex.Decode(dec[:], raw); err != nil {
		return err
	} else {
		*b = dec
	}
	return nil
}
