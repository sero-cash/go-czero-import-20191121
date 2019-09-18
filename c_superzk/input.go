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

func ProveInput(
	asset_cm_new *c_type.Uint256,
	zpka *c_type.Uint256,
	nil *c_type.Uint256,
	anchor *c_type.Uint256,
	asset_cc *c_type.Uint256,
	ar_old *c_type.Uint256,
	ar_new *c_type.Uint256,
	index uint64,
	zpkr *c_type.Uint256,
	vskr *c_type.Uint256,
	baser *c_type.Uint256,
	a *c_type.Uint256,
	paths *[c_type.DEPTH * 32]byte,
	pos uint64,
) (proof c_type.Proof, e error) {
	nil = ClearNil(nil)
	baser = ClearNil(baser)
	ret := C.superzk_prove_input(
		(*C.uchar)(unsafe.Pointer(&asset_cm_new[0])),
		(*C.uchar)(unsafe.Pointer(&zpka[0])),
		(*C.uchar)(unsafe.Pointer(&nil[0])),
		(*C.uchar)(unsafe.Pointer(&anchor[0])),
		(*C.uchar)(unsafe.Pointer(&asset_cc[0])),
		(*C.uchar)(unsafe.Pointer(&ar_old[0])),
		(*C.uchar)(unsafe.Pointer(&ar_new[0])),
		C.ulong(index),
		(*C.uchar)(unsafe.Pointer(&zpkr[0])),
		(*C.uchar)(unsafe.Pointer(&vskr[0])),
		(*C.uchar)(unsafe.Pointer(&baser[0])),
		(*C.uchar)(unsafe.Pointer(&a[0])),
		(*C.uchar)(unsafe.Pointer(&paths[0])),
		C.ulong(pos),
		(*C.uchar)(unsafe.Pointer(&proof[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("prove input error: %d", int(ret))
	}
	return
}

func VerifyInput(
	proof *c_type.Proof,
	asset_cm_new *c_type.Uint256,
	zpka *c_type.Uint256,
	nil *c_type.Uint256,
	anchor *c_type.Uint256,
) (e error) {
	nil = ClearNil(nil)
	ret := C.superzk_verify_input(
		(*C.uchar)(unsafe.Pointer(&proof[0])),
		(*C.uchar)(unsafe.Pointer(&asset_cm_new[0])),
		(*C.uchar)(unsafe.Pointer(&zpka[0])),
		(*C.uchar)(unsafe.Pointer(&nil[0])),
		(*C.uchar)(unsafe.Pointer(&anchor[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("verify input error: %d", int(ret))
	}
	return
}
