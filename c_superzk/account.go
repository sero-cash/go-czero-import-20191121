package c_superzk

/*

#include "zero.h"

*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

func Seed2Sk(seed *c_type.Uint256) (sk c_type.Uint512) {
	C.zero_seed2sk(
		(*C.uchar)(unsafe.Pointer(&seed[0])),
		(*C.uchar)(unsafe.Pointer(&sk[0])),
	)
	return
}

func Sk2Tk(sk *c_type.Uint512) (tk c_type.Uint512) {
	C.zero_sk2tk(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&tk[0])),
	)
	return
}

func Tk2Pk(tk *c_type.Uint512) (pk c_type.Uint512) {
	C.zero_tk2pk(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&pk[0])),
	)
	return
}

func Pk2PKr(addr *c_type.Uint512, r *c_type.Uint256) (pkr c_type.PKr) {
	if r == nil {
		t := c_type.RandUint256()
		r = &t
	} else {
		if (*r) == c_type.Empty_Uint256 {
			panic("gen pkr, but r is empty")
		}
	}

	C.zero_pk2pkr(
		(*C.uchar)(unsafe.Pointer(&addr[0])),
		(*C.uchar)(unsafe.Pointer(&r[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	return
}

func IsMyPKr(tk *c_type.Uint512, pkr *c_type.PKr) (succ bool) {
	ret := C.zero_ismy_pkr(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&tk[0])),
	)
	if ret == C.char(0) {
		succ = true
		return
	} else {
		succ = false
		return
	}
}

func FetchKey(pkr *c_type.PKr, tk *c_type.Uint512, rpk *c_type.Uint256) (ret c_type.Uint256) {
	C.zero_fetch_key(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&rpk[0])),
		(*C.uchar)(unsafe.Pointer(&ret[0])),
	)
	return
}

func GenKey(pkr *c_type.PKr) (key c_type.Uint256, rpk c_type.Uint256, rsk c_type.Uint256) {
	return
}

func GenZPKa(pkr *c_type.PKr) (zpka c_type.Uint256, a c_type.Uint256) {
	return
}

func SignPKrBySk(sk *c_type.Uint512, data *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	C.zero_sign_pkr_by_sk(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
	)
	if sign == c_type.Empty_Uint512 {
		e = errors.New("SignOAddr: sign is empty")
		return
	} else {
		return
	}
}

func VerifyPKr(data *c_type.Uint256, sign *c_type.Uint512, pkr *c_type.PKr) bool {
	ret := C.zero_verify_pkr(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret == C.char(0) {
		return true
	} else {
		return false
	}
}
