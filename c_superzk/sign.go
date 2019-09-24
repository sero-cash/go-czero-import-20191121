package c_superzk

/*

#include "csuperzk.h"

*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_czero"

	"github.com/sero-cash/go-czero-import/c_type"
)

func GenZPKa(pkr *c_type.PKr) (zpka c_type.Uint256, a c_type.Uint256, e error) {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	a = c_type.RandUint256()
	a = c_czero.Force_Fr(&a)
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
	assertPKr(pkr)
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

func SignPKrBySk(sk *c_type.Uint512, data *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	assertPKr(pkr)
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

func VerifyPKr(data *c_type.Uint256, sign *c_type.Uint512, pkr *c_type.PKr) bool {
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
