package c_superzk

/*

#include "zero.h"

*/
import "C"
import (
	"bytes"
	"errors"
	"unsafe"

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

func Tk2Pk(tk *c_type.Tk) (pk c_type.Uint512) {
	copy(pk[:], tk[:])
	SetFlag(pk[:])
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

func Pk2PKr(addr *c_type.Uint512, r *c_type.Uint256) (pkr c_type.PKr) {
	assertPK(addr)
	addr = ClearPK(addr)
	if r == nil {
		t := c_type.RandUint256()
		r = &t
	} else {
		if (*r) == c_type.Empty_Uint256 {
			panic("gen pkr, but r is empty")
		}
	}
	copy(pkr[:], addr[:])
	copy(pkr[64:], r[:])
	SetFlag(pkr[:])
	return
}

func Sk2Tk(sk *c_type.Uint512) (tk c_type.Tk) {
	C.zero_sk2tk(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&tk[0])),
	)
	return
}

func Seed2Sk(seed *c_type.Uint256) (sk c_type.Uint512) {
	C.zero_seed2sk(
		(*C.uchar)(unsafe.Pointer(&seed[0])),
		(*C.uchar)(unsafe.Pointer(&sk[0])),
	)
	return
}

func HashPKr(pkr *c_type.PKr) (ret [20]byte) {
	pkr = ClearPKr(pkr)
	C.zero_hpkr(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&ret[0])),
	)
	return
}

func IsPKValid(pk *c_type.Uint512) bool {
	if !IsSzkPK(pk) {
		return false
	}
	return true
}

func IsPKrValid(pkr *c_type.PKr) bool {
	if !IsSzkPKr(pkr) {
		return false
	}
	return true
}

func IsMyPKr(tk *c_type.Tk, pkr *c_type.PKr) (succ bool) {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	if bytes.Compare(tk[:], pkr[:64]) == 0 {
		return true
	} else {
		return false
	}
}

func GenKey(pkr *c_type.PKr) (key c_type.Uint256, rpk c_type.Uint256, rsk c_type.Uint256) {
	assertPKr(pkr)
	return
}

func FetchKey(pkr *c_type.PKr, tk *c_type.Tk, rpk *c_type.Uint256) (ret c_type.Uint256) {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	return
}

func GenZPKa(pkr *c_type.PKr) (zpka c_type.Uint256, a c_type.Uint256) {
	assertPKr(pkr)
	return
}

func SignPKrBySk(sk *c_type.Uint512, data *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	return
}

func VerifyPKr(data *c_type.Uint256, sign *c_type.Uint512, pkr *c_type.PKr) bool {
	assertPKr(pkr)
	return true
}
