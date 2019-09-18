package c_superzk

/*

#include "zero.h"

*/
import "C"
import (
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

func GenCzeroTrace(tk *c_type.Uint512, root_cm *c_type.Uint256) (til c_type.Uint256) {
	C.zero_til(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&til[0])),
	)
	return
}

func GenCzeroNil(sk *c_type.Uint512, root_cm *c_type.Uint256) (til c_type.Uint256) {
	C.zero_nil(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&til[0])),
	)
	return
}

func SignCzeroNil(sk *c_type.Uint512, hash *c_type.Uint256, root_cm *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	return
}

func VerifyCzeroNil(hash *c_type.Uint256, sign *c_type.Uint256, nil *c_type.Uint256, root_cm *c_type.Uint256, pkr *c_type.PKr) (ret bool) {
	return
}

func CzeroIsMyPKr(tk *c_type.Uint512, pkr *c_type.PKr) (succ bool) {
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

func FetchCzeroKey(tk *c_type.Uint512, rpk *c_type.Uint256) (ret c_type.Uint256, flag bool) {
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
