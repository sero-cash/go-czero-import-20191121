package c_type

import (
	"encoding/hex"
	"fmt"

	"github.com/sero-cash/go-sero/crypto/sha3"
)

type Proof [PROOF_WIDTH]byte

func (b Proof) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b[:])
	return result, nil
}

func (b *Proof) UnmarshalText(input []byte) error {
	raw := input[2:]
	if len(raw) == 0 {
		return nil
	}
	dec := Proof{}
	if len(raw)/2 != len(dec[:]) {
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(dec[:])*2, "Proof")
	}
	if _, err := hex.Decode(dec[:], raw); err != nil {
		return err
	} else {
		*b = dec
	}
	return nil
}

func (self *Proof) ToHash() (ret Uint256) {
	d := sha3.NewKeccak256()
	d.Write(self[:])
	copy(ret[:], d.Sum(nil))
	return
}
