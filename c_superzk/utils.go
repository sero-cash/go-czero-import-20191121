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

func IsSzkSk(sk *c_type.Uint512) bool {
	return IsFlagSet(sk[:])
}
func ClearSk(sk *c_type.Uint512) (ret *c_type.Uint512) {
	ret = &c_type.Uint512{}
	ClearFlag(ret[:], sk[:])
	return
}
func assertSk(sk *c_type.Uint512) {
	if !IsSzkSk(sk) {
		panic(errors.New("sk is not the superzk sk"))
	}
}

func IsSzkTk(pk *c_type.Tk) bool {
	return IsFlagSet(pk[:])
}
func ClearTk(tk *c_type.Tk) (ret *c_type.Tk) {
	ret = &c_type.Tk{}
	ClearFlag(ret[:], tk[:])
	return
}
func assertTk(tk *c_type.Tk) {
	if !IsSzkTk(tk) {
		panic(errors.New("tk is not the superzk tk"))
	}
}

func IsSzkPK(pk *c_type.Uint512) bool {
	return IsFlagSet(pk[:])
}
func ClearPK(pkr *c_type.Uint512) (ret *c_type.Uint512) {
	ret = &c_type.Uint512{}
	ClearFlag(ret[:], pkr[:])
	return
}
func assertPK(pk *c_type.Uint512) {
	if !IsSzkPK(pk) {
		panic(errors.New("pk is not the superzk pk"))
	}
}

func IsSzkPKr(pkr *c_type.PKr) bool {
	return IsFlagSet(pkr[:])
}
func ClearPKr(pkr *c_type.PKr) (ret *c_type.PKr) {
	ret = &c_type.PKr{}
	ClearFlag(ret[:], pkr[:])
	return
}
func assertPKr(pkr *c_type.PKr) {
	if !IsSzkPKr(pkr) {
		panic(errors.New("pkr is not the superzk pkr"))
	}
}

func IsSzkNil(nil *c_type.Uint256) bool {
	return IsFlagSet(nil[:])
}

func ClearNil(pkr *c_type.Uint256) (ret *c_type.Uint256) {
	ret = &c_type.Uint256{}
	ClearFlag(ret[:], pkr[:])
	return
}
func assertNil(nil *c_type.Uint256) {
	if !IsSzkNil(nil) {
		panic(errors.New("nil is not the superzk nil"))
	}
}
