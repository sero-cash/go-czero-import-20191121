package c_czero

/*

#include "zero.h"

*/
import "C"
import (
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

func Force_Fr(data *c_type.Uint256) (fr c_type.Uint256) {
	C.zero_force_fr(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&fr[0])),
	)
	return
}

func Combine(l *c_type.Uint256, r *c_type.Uint256) (out c_type.Uint256) {
	C.zero_merkle_combine(
		(*C.uchar)(unsafe.Pointer(&l[0])),
		(*C.uchar)(unsafe.Pointer(&r[0])),
		(*C.uchar)(unsafe.Pointer(&out[0])),
	)
	return
}
