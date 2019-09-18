package c_type

import (
	"encoding/hex"
	"fmt"
)

type PKr [96]byte

func (self *PKr) ZPKr() (ret Uint256) {
	copy(ret[:], self[:32])
	return
}

func (self *PKr) VPKr() (ret Uint256) {
	copy(ret[:], self[32:64])
	return
}

func (self *PKr) BASEr() (ret Uint256) {
	copy(ret[:], self[64:])
	return
}

func (self *PKr) IsEndEmpty() bool {
	end := Uint256{}
	copy(end[:], self[64:])
	if end == Empty_Uint256 {
		return true
	} else {
		return false
	}
}

func NewPKrByBytes(bs []byte) (ret PKr) {
	copy(ret[:], bs[:])
	return
}

var Empty_PKr = PKr{}

func (self *PKr) ToUint512() (ret Uint512) {
	copy(ret[:], self[:])
	return
}

func (self PKr) NewRef() (ret *PKr) {
	ret = &PKr{}
	copy(ret[:], self[:])
	return ret
}

func (b PKr) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b[:])
	return result, nil
}

func (b *PKr) UnmarshalText(input []byte) error {
	if len(input) < 2 {
		return fmt.Errorf("hex string length must > 2 : current is %d", len(input))
	}
	raw := input[2:]
	if len(raw) == 0 {
		return nil
	}
	dec := PKr{}
	if len(raw)/2 != len(dec[:]) {
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(dec[:])*2, "PKr")
	}
	if _, err := hex.Decode(dec[:], raw); err != nil {
		return err
	} else {
		*b = dec
	}
	return nil
}

type NetType uint8

const (
	NET_Dev   NetType = 0
	NET_Alpha NetType = 1
	NET_Beta  NetType = 2
)
