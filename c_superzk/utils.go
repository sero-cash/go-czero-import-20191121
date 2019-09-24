package c_superzk

import (
	"errors"

	"github.com/sero-cash/go-czero-import/c_type"
)

func SetFlag(bytes []byte) {
	bytes[len(bytes)-1] |= uint8(0x1 << 6)
}

func ClearFlag(ret []byte, bytes []byte) {
	copy(ret, bytes)
	ret[len(ret)-1] &= ^uint8(0x1 << 6)
	return
}

func IsFlagSet(bytes []byte) bool {
	flag := bytes[len(bytes)-1] & (0x1 << 6)
	if flag != 0 {
		return true
	} else {
		return false
	}
}

func IsSzkPKr(pkr *c_type.PKr) bool {
	return IsFlagSet(pkr[:])
}

func IsSzkPK(pk *c_type.Uint512) bool {
	return IsFlagSet(pk[:])
}

func IsSzkNil(nil *c_type.Uint256) bool {
	return IsFlagSet(nil[:])
}

func ClearPKr(pkr *c_type.PKr) (ret *c_type.PKr) {
	ret = &c_type.PKr{}
	ClearFlag(ret[:], pkr[:])
	return
}
func ClearPK(pkr *c_type.Uint512) (ret *c_type.Uint512) {
	ret = &c_type.Uint512{}
	ClearFlag(ret[:], pkr[:])
	return
}
func ClearNil(pkr *c_type.Uint256) (ret *c_type.Uint256) {
	ret = &c_type.Uint256{}
	ClearFlag(ret[:], pkr[:])
	return
}

func assertPK(pk *c_type.Uint512) {
	if !IsSzkPK(pk) {
		panic(errors.New("pk is not the superzk pk"))
	}
}

func assertPKr(pkr *c_type.PKr) {
	if !IsSzkPKr(pkr) {
		panic(errors.New("pkr is not the superzk pkr"))
	}
}

func assertNil(nil *c_type.Uint256) {
	if !IsSzkNil(nil) {
		panic(errors.New("nil is not the superzk nil"))
	}
}
