package c_superzk

/*

#include "zero.h"

*/
import "C"
import (
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

type AssetDesc struct {
	Asset        c_type.Asset
	Ar           c_type.Uint256
	Asset_cc_ret c_type.Uint256
	Asset_cm_ret c_type.Uint256
}

func GenAssetCC(desc *AssetDesc) {
	copy(desc.Asset_cc_ret[:], desc.Asset.Tkn_value[:])
}

func GenAssetCM(desc *AssetDesc) {
	copy(desc.Asset_cc_ret[:], desc.Asset.Tkn_value[:])
}

func PtrOfSlice(s []byte) *C.uchar {
	if len(s) > 0 {
		return (*C.uchar)(unsafe.Pointer(&s[0]))
	} else {
		return (*C.uchar)(unsafe.Pointer(uintptr(0)))
	}
}

func SignBalance(desc *c_type.BalanceDesc) (e error) {
	copy(desc.Bcr[:], desc.Hash[:])
	copy(desc.Bsign[:], desc.Hash[:])
	return
}

func VerifyBalance(desc *c_type.BalanceDesc) (ret bool) {
	ret = true
	return
}
