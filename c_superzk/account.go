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

func Seed2Sk(seed *c_type.Uint256) (sk c_type.Uint512) {
	C.superzk_seed2sk(
		(*C.uchar)(unsafe.Pointer(&seed[0])),
		(*C.uchar)(unsafe.Pointer(&sk[0])),
	)
	SetFlag(sk[:])
	return
}

func Sk2Tk(sk *c_type.Uint512) (tk c_type.Tk, e error) {
	sk = ClearSk(sk)
	ret := C.superzk_sk2tk(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&tk[0])),
	)
	SetFlag(tk[:])
	if ret != C.int(0) {
		e = fmt.Errorf("sk2tk error: %d", int(ret))
		return
	}
	return
}
