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

func IsTknValid(currency *c_type.Uint256) (e error) {
	base := c_type.Uint256{}
	ret := C.superzk_gen_tkn_base(
		(*C.uchar)(unsafe.Pointer(&currency[0])),
		(*C.uchar)(unsafe.Pointer(&base[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("gen tkn base error: %d", int(ret))
		return
	}
	return
}

func IsTktValid(category *c_type.Uint256, value *c_type.Uint256) (e error) {
	base := c_type.Uint256{}
	ret := C.superzk_gen_tkt_base(
		(*C.uchar)(unsafe.Pointer(&category[0])),
		(*C.uchar)(unsafe.Pointer(&value[0])),
		(*C.uchar)(unsafe.Pointer(&base[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("gen tkt base error: %d", int(ret))
		return
	}
	return
}

func GenAssetCC(asset *c_type.Asset) (cc c_type.Uint256, e error) {
	ret := C.superzk_gen_asset_cc(
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&cc[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("gen asset cc error: %d", int(ret))
		return
	}
	return
}

func GenAssetCM_PC(asset *c_type.Asset, ar *c_type.Uint256) (cm c_type.Uint256, cc c_type.Uint256, e error) {
	ret := C.superzk_gen_asset_cm(
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&ar[0])),
		(*C.uchar)(unsafe.Pointer(&cm[0])),
		(*C.uchar)(unsafe.Pointer(&cc[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("gen asset cc error: %d", int(ret))
		return
	}
	return
}

type AssetDesc struct {
	Asset        c_type.Asset
	Ar           c_type.Uint256
	Asset_cc_ret c_type.Uint256
	Asset_cm_ret c_type.Uint256
}

func GenAssetCM(desc *AssetDesc) (e error) {
	ret := C.superzk_gen_asset_cm(
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Ar[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset_cm_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset_cc_ret[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("gen asset cc error: %d", int(ret))
		return
	}
	return
}

func PtrOfSlice(s []byte) *C.uchar {
	if len(s) > 0 {
		return (*C.uchar)(unsafe.Pointer(&s[0]))
	} else {
		return (*C.uchar)(unsafe.Pointer(uintptr(0)))
	}
}

func SignBalance(desc *c_type.BalanceDesc) (e error) {
	ret := C.superzk_sign_balance(
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
	if ret != C.int(0) {
		e = fmt.Errorf("sign balance error: %d", int(ret))
	}
	return
}

func VerifyBalance(desc *c_type.BalanceDesc) (e error) {
	ret := C.superzk_verify_balance(
		C.int(len(desc.Zin_acms)/32),
		PtrOfSlice(desc.Zin_acms),
		C.int(len(desc.Zout_acms)/32),
		PtrOfSlice(desc.Zout_acms),
		C.int(len(desc.Oin_accs)/32),
		PtrOfSlice(desc.Oin_accs),
		C.int(len(desc.Oout_accs)/32),
		PtrOfSlice(desc.Oout_accs),
		(*C.uchar)(unsafe.Pointer(&desc.Hash[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Bsign[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Bcr[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("verify balance error: %d", int(ret))
		return
	}
	return
}
