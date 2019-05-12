package keys

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

type Uint256 [32]byte

func Uint64_To_Uint256(v uint64) (ret Uint256) {
	bys := big.NewInt(0).SetUint64(v).Bytes()
	l := len(ret)
	copy(ret[l-len(bys):], bys)
	return
}
func Uint256_To_Uint64(v *Uint256) (ret uint64) {
	ret = big.NewInt(0).SetBytes(v[:]).Uint64()
	return
}

func Seeds2Tks(seeds []Uint256) (tks []Uint512) {
	for _, seed := range seeds {
		tks = append(tks, Seed2Tk(&seed))
	}
	return
}

func (b *Uint128) ToUint256() (ret Uint256) {
	copy(ret[:], b[:])
	return
}

func (self Uint256) NewRef() (ret *Uint256) {
	ret = &Uint256{}
	copy(ret[:], self[:])
	return ret
}

func (self Uint256) LogOut() {
	logBytes(self[:])
}

func (b Uint256) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b[:])
	return result, nil
}

func (b *Uint256) UnmarshalText(input []byte) error {
	raw := input[2:]
	if len(raw) == 0 {
		return nil
	}
	dec := Uint256{}
	if len(raw)/2 != len(dec[:]) {
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(dec[:])*2, "Uint128")
	}
	if _, err := hex.Decode(dec[:], raw); err != nil {
		return err
	} else {
		*b = dec
	}
	return nil
}

var Empty_Uint256 = Uint256{}

type Uint512 [64]byte

var Empty_Uint512 = Uint512{}

func (self Uint512) NewRef() (ret *Uint512) {
	ret = &Uint512{}
	copy(ret[:], self[:])
	return ret
}

func (self Uint512) LogOut() {
	logBytes(self[:])
}

func (b Uint512) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b[:])
	return result, nil
}

func (b *Uint512) UnmarshalText(input []byte) error {
	raw := input[2:]
	if len(raw) == 0 {
		return nil
	}
	dec := Uint512{}
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

type Uint128 [16]byte

func (b Uint128) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b[:])
	return result, nil
}

func (b *Uint128) UnmarshalText(input []byte) error {
	raw := input[2:]
	if len(raw) == 0 {
		return nil
	}
	dec := Uint128{}
	if len(raw)/2 != len(dec[:]) {
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(dec[:])*2, "Uint128")
	}
	if _, err := hex.Decode(dec[:], raw); err != nil {
		return err
	} else {
		*b = dec
	}
	return nil
}

type PKr [96]byte

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

func (self PKr) LogOut() {
	logBytes(self[:])
}

func (b PKr) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b[:])
	return result, nil
}

func (b *PKr) UnmarshalText(input []byte) error {
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
