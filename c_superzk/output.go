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

func ProveOutput(
	asset *c_type.Asset,
	ar *c_type.Uint256,
	asset_cm *c_type.Uint256,
) (proof c_type.Proof, e error) {
	ret := C.superzk_prove_output(
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&ar[0])),
		(*C.uchar)(unsafe.Pointer(&asset_cm[0])),
		(*C.uchar)(unsafe.Pointer(&proof[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("prove output error: %d", int(ret))
	}
	return
}

func VerifyOutput(asset_cm *c_type.Uint256, proof *c_type.Proof) (e error) {
	ret := C.superzk_verify_output(
		(*C.uchar)(unsafe.Pointer(&asset_cm[0])),
		(*C.uchar)(unsafe.Pointer(&proof[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("verify output error: %d", int(ret))
	}
	return
}
