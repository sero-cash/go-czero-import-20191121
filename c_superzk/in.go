package c_superzk

/*

#include "zero.h"

*/
import "C"
import (
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

func GenNil(tk *c_type.Uint512, root_cm *c_type.Uint256) (til c_type.Uint256) {
	C.zero_til(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&til[0])),
	)
	return
}

func SignNil(tk *c_type.Uint512, hash *c_type.Uint256, root_cm *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	return
}

func VerifyNil(hash *c_type.Uint256, sign *c_type.Uint512, nil *c_type.Uint256, root_cm *c_type.Uint256, pkr *c_type.PKr) (ret bool) {
	return
}

func SignZPKa(sk *c_type.Uint512, hash *c_type.Uint256, a *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	return
}

func VerifyZPKa(hash *c_type.Uint256, sign *c_type.Uint512, zpka *c_type.Uint256) (ret bool) {
	return
}
