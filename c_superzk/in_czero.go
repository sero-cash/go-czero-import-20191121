package c_superzk

/*

#include "zero.h"

*/
import "C"
import (
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_czero"

	"github.com/sero-cash/go-czero-import/c_type"
)

func GenCzeroTrace(tk *c_type.Tk, root_cm *c_type.Uint256) (til c_type.Uint256) {
	til = c_czero.GenTil(tk, root_cm)
	return
}

func GenCzeroNil(sk *c_type.Uint512, root_cm *c_type.Uint256) (til c_type.Uint256) {
	copy(til[:], root_cm[:])
	til = c_czero.GenNil(sk, root_cm)
	return
}

func SignCzeroNil(sk *c_type.Uint512, hash *c_type.Uint256, root_cm *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	return
}

func VerifyCzeroNil(hash *c_type.Uint256, sign *c_type.Uint256, nil *c_type.Uint256, root_cm *c_type.Uint256, pkr *c_type.PKr) (ret bool) {
	ret = true
	return
}

func CzeroIsMyPKr(tk *c_type.Tk, pkr *c_type.PKr) (succ bool) {
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

func FetchCzeroKey(tk *c_type.Tk, rpk *c_type.Uint256) (ret c_type.Uint256, flag bool) {
	f := C.zero_fetch_key(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&rpk[0])),
		(*C.uchar)(unsafe.Pointer(&ret[0])),
	)
	if f == C.char(0) {
		flag = false
	} else {
		flag = true
	}
	return
}
