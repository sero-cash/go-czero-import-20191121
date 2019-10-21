package c_superzk

/*

#include "csuperzk.h"

*/
import "C"
import (
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

func InitParams() {
	C.superzk_init_params()
}

func InitParams_NoCircuit() {
	C.superzk_init_params_no_circuit()
}

func RandomPt() (ret c_type.Uint256) {
	C.superzk_random_pt(
		(*C.uchar)(unsafe.Pointer(&ret[0])),
	)
	return
}

func RandomFr() (ret c_type.Uint256) {
	C.superzk_random_fr(
		(*C.uchar)(unsafe.Pointer(&ret[0])),
	)
	return
}

func ForceFr(data *c_type.Uint256) (ret c_type.Uint256) {
	C.superzk_force_fr(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&ret[0])),
	)
	return
}

func Combine(l *c_type.Uint256, r *c_type.Uint256) (out c_type.Uint256) {
	C.superzk_merkle_combine(
		(*C.uchar)(unsafe.Pointer(&l[0])),
		(*C.uchar)(unsafe.Pointer(&r[0])),
		(*C.uchar)(unsafe.Pointer(&out[0])),
	)
	return
}

func HashPKr(pkr *c_type.PKr) (ret [20]byte) {
	C.superzk_hpkr(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&ret[0])),
	)
	return
}
