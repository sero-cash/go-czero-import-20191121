package c_superzk

/*

#include "csuperzk.h"

*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

func GenZPKa(pkr *c_type.PKr, a *c_type.Uint256) (zpka c_type.Uint256, e error) {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	ret := C.superzk_gen_zpka(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&a[0])),
		(*C.uchar)(unsafe.Pointer(&zpka[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("gen zpka error: %d", int(ret))
		return
	}
	return
}

func SignZPKa(sk *c_type.Uint512, data *c_type.Uint256, a *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	assertSk(sk)
	assertPKr(pkr)
	sk = ClearSk(sk)
	pkr = ClearPKr(pkr)
	ret := C.superzk_sign_zpka(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&a[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("sign zpka error: %d", int(ret))
		return
	}
	return
}

func VerifyZPKa(data *c_type.Uint256, sign *c_type.Uint512, zpka *c_type.Uint256) bool {
	ret := C.superzk_verify_zpka(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
		(*C.uchar)(unsafe.Pointer(&zpka[0])),
	)
	if ret != C.int(0) {
		return false
	}
	return true
}

func SignPKr_P(sk *c_type.Uint512, data *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	assertSk(sk)
	assertPKr(pkr)
	sk = ClearSk(sk)
	pkr = ClearPKr(pkr)
	ret := C.superzk_sign_pkr(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("sign pkr error: %d", int(ret))
		return
	}
	return
}

func VerifyPKr_P(data *c_type.Uint256, sign *c_type.Uint512, pkr *c_type.PKr) bool {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	ret := C.superzk_verify_pkr(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret != C.int(0) {
		return false
	}
	return true
}

func SignNil_P0(h *c_type.Uint256, sk *c_type.Uint512, pkr *c_type.PKr, root_cm *c_type.Uint256) (sign c_type.SignN, e error) {
	ret := C.czero_sign_nil_ex(
		(*C.uchar)(unsafe.Pointer(&h[0])),
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("czero sign nil error: %d", int(ret))
		return
	}
	return
}

func VerifyNil_P0(h *c_type.Uint256, sign *c_type.SignN, pkr *c_type.PKr, root_cm *c_type.Uint256, nl *c_type.Uint256) (e error) {
	ret := C.czero_verify_nil_ex(
		(*C.uchar)(unsafe.Pointer(&h[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&nl[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("czero verify nil error: %d", int(ret))
		return
	}
	return
}

func SignPKr_P0(h *c_type.Uint256, sk *c_type.Uint512, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	ret := C.czero_sign_pkr(
		(*C.uchar)(unsafe.Pointer(&h[0])),
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("czero sign pkr error: %d", int(ret))
		return
	}
	return
}

func VerifyPKr_P0(h *c_type.Uint256, sign *c_type.Uint512, pkr *c_type.PKr) (e error) {
	ret := C.czero_verify_pkr(
		(*C.uchar)(unsafe.Pointer(&h[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("czero verify pkr error: %d", int(ret))
		return
	}
	return
}

func SignPKr_X(sk *c_type.Uint512, data *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	if IsSzkPKr(pkr) {
		return SignPKr_P(sk, data, pkr)
	} else {
		return SignPKr_P0(data, sk, pkr)
	}
}

func VerifyPKr_X(data *c_type.Uint256, sign *c_type.Uint512, pkr *c_type.PKr) bool {
	if IsSzkPKr(pkr) {
		return VerifyPKr_P(data, sign, pkr)
	} else {
		if e := VerifyPKr_P0(data, sign, pkr); e != nil {
			return false
		} else {
			return true
		}
	}
}
