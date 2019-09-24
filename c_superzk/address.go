package c_superzk

/*

#include "csuperzk.h"

*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

func Tk2Pk(tk *c_type.Tk) (pk c_type.Uint512, e error) {
	ret := C.superzk_tk2pk(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&pk[0])),
	)
	if ret != C.int(0) {
		e = errors.New("tk2pk error")
		return
	}
	SetFlag(pk[:])
	return
}

func Pk2PKr(addr *c_type.Uint512, r *c_type.Uint256) (pkr c_type.PKr, e error) {
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
	ret := C.superzk_pk2pkr(
		(*C.uchar)(unsafe.Pointer(&addr[0])),
		(*C.uchar)(unsafe.Pointer(&r[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret != C.int(0) {
		e = errors.New("pk2pkr error")
		return
	}
	SetFlag(pkr[:])
	return
}

func IsPKValid(pk *c_type.Uint512) bool {
	if !IsSzkPK(pk) {
		return false
	}
	pk = ClearPK(pk)
	ret := C.superzk_pk_valid(
		(*C.uchar)(unsafe.Pointer(&pk[0])),
	)
	if ret != C.int(0) {
		return false
	}
	return true
}

func IsPKrValid(pkr *c_type.PKr) bool {
	if !IsSzkPKr(pkr) {
		return false
	}
	pkr = ClearPKr(pkr)
	ret := C.superzk_pkr_valid(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret != C.int(0) {
		return false
	}
	return true
}

func IsMyPKr(tk *c_type.Tk, pkr *c_type.PKr) bool {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	ret := C.superzk_my_pkr(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret != C.int(0) {
		return false
	}
	return true
}
