package c_superzk

/*

#include "zero.h"

*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

type AssetDesc struct {
	Asset        c_type.Asset
	Asset_cc_ret c_type.Uint256
	Asset_cm_ret c_type.Uint256
	Asset_ar_ret c_type.Uint256
}

func GenAssetCC(desc *AssetDesc) {
	C.zero_gen_asset_cc(
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_value[0])),
		//--out--
		(*C.uchar)(unsafe.Pointer(&desc.Asset_cc_ret[0])),
	)
}

func GenAssetCM(desc *AssetDesc) {
	C.zero_gen_asset_cc(
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_value[0])),
		//--out--
		(*C.uchar)(unsafe.Pointer(&desc.Asset_cc_ret[0])),
	)
}

func PtrOfSlice(s []byte) *C.uchar {
	if len(s) > 0 {
		return (*C.uchar)(unsafe.Pointer(&s[0]))
	} else {
		return (*C.uchar)(unsafe.Pointer(uintptr(0)))
	}
}

func SignBalance(desc *c_type.BalanceDesc) {
	C.zero_sign_balance(
		C.int(len(desc.Zin_acms)/32),
		PtrOfSlice(desc.Zin_acms),
		PtrOfSlice(desc.Zin_ars),
		C.int(len(desc.Zout_acms)/32),
		PtrOfSlice(desc.Zout_acms),
		PtrOfSlice(desc.Zout_ars),
		C.int(len(desc.Oin_accs)/32),
		PtrOfSlice(desc.Oin_accs),
		C.int(len(desc.Oout_accs)/32),
		PtrOfSlice(desc.Oout_accs),
		(*C.uchar)(unsafe.Pointer(&desc.Hash[0])),
		//---out--
		(*C.uchar)(unsafe.Pointer(&desc.Bsign[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Bcr[0])),
	)
	return
}

func VerifyBalance(desc *c_type.BalanceDesc) (e error) {
	ret := C.zero_verify_balance(
		C.int(len(desc.Zin_acms)/32),
		PtrOfSlice(desc.Zin_acms),
		C.int(len(desc.Zout_acms)/32),
		PtrOfSlice(desc.Zout_acms),
		C.int(len(desc.Oin_accs)/32),
		PtrOfSlice(desc.Oin_accs),
		C.int(len(desc.Oout_accs)/32),
		PtrOfSlice(desc.Oout_accs),
		(*C.uchar)(unsafe.Pointer(&desc.Hash[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Bcr[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Bsign[0])),
	)
	if ret == 0 {
		return
	} else {
		e = errors.New("verify balance error")
		return
	}
}
