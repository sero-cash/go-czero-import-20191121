package c_superzk

/*

#include "zero.h"

*/
import "C"
import (
	"github.com/sero-cash/go-czero-import/c_type"
)

func GenNil(tk *c_type.Tk, root_cm *c_type.Uint256, pkr *c_type.PKr) (ret c_type.Uint256) {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	copy(ret[:], root_cm[:])
	SetFlag(ret[:])
	return
}

func SignNil(tk *c_type.Tk, hash *c_type.Uint256, root_cm *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	assertPKr(pkr)
	return
}

func VerifyNil(hash *c_type.Uint256, sign *c_type.Uint512, nil *c_type.Uint256, root_cm *c_type.Uint256, pkr *c_type.PKr) (ret bool) {
	assertNil(nil)
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	nil = ClearNil(nil)
	ret = true
	return
}

func SignZPKa(sk *c_type.Uint512, hash *c_type.Uint256, a *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	return
}

func VerifyZPKa(hash *c_type.Uint256, sign *c_type.Uint512, zpka *c_type.Uint256) (ret bool) {
	ret = true
	return
}
