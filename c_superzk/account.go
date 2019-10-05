package c_superzk

/*

#include "csuperzk.h"

*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

func Seed2Sk(seed *c_type.Uint256) (sk c_type.Uint512) {
	C.superzk_seed2sk(
		(*C.uchar)(unsafe.Pointer(&seed[0])),
		(*C.uchar)(unsafe.Pointer(&sk[0])),
	)
	return
}

func Sk2Tk(sk *c_type.Uint512) (tk c_type.Tk, e error) {
	ret := C.superzk_sk2tk(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&tk[0])),
	)
	if ret != C.int(0) {
		e = errors.New("sk2tk error")
		return
	}
	return
}

func Seed2Tk(seed *c_type.Uint256) (tk c_type.Tk, e error) {
	sk := Seed2Sk(seed)
	tk, e = Sk2Tk(&sk)
	return
}
