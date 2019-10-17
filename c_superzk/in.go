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

func GenNil(tk *c_type.Tk, root_cm *c_type.Uint256, pkr *c_type.PKr) (nil c_type.Uint256, e error) {
	assertTk(tk)
	assertPKr(pkr)
	tk = ClearTk(tk)
	pkr = ClearPKr(pkr)
	ret := C.superzk_gen_nil(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&nil[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("gen nil error: %d", int(ret))
		return
	}
	SetFlag(nil[:])
	return
}

func SignNil(tk *c_type.Tk, hash *c_type.Uint256, root_cm *c_type.Uint256, pkr *c_type.PKr) (sign c_type.SignN, e error) {
	assertTk(tk)
	assertPKr(pkr)
	tk = ClearTk(tk)
	pkr = ClearPKr(pkr)
	ret := C.superzk_sign_nil(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&hash[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("sign nil error: %d", int(ret))
		return
	}
	return
}

func VerifyNil(hash *c_type.Uint256, sign *c_type.SignN, nl *c_type.Uint256, root_cm *c_type.Uint256, pkr *c_type.PKr) error {
	assertNil(nl)
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	nl = ClearNil(nl)
	ret := C.superzk_verify_nil(
		(*C.uchar)(unsafe.Pointer(&hash[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
		(*C.uchar)(unsafe.Pointer(&nl[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret != C.int(0) {
		return fmt.Errorf("verify nil error: %d", int(ret))
	}
	return nil
}
