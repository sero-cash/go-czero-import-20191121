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

func EncInfo(key *c_type.Uint256, asset *c_type.Asset, memo *c_type.Uint512, ar *c_type.Uint256) (einfo c_type.Einfo, e error) {
	ret := C.superzk_enc_info(
		(*C.uchar)(unsafe.Pointer(&key[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&memo[0])),
		(*C.uchar)(unsafe.Pointer(&ar[0])),
		(*C.uchar)(unsafe.Pointer(&einfo[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("enc info error: %d", int(ret))
		return
	}
	return
}

type EncInfoDesc struct {
	//---in---
	Key   c_type.Uint256
	Asset c_type.Asset
	Memo  c_type.Uint512
	Ar    c_type.Uint256
	//---out---
	Einfo c_type.Einfo
}

func EncOutput(desc *EncInfoDesc) (e error) {
	ret := C.superzk_enc_info(
		(*C.uchar)(unsafe.Pointer(&desc.Key[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Memo[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Ar[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Einfo[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("enc info error: %d", int(ret))
		return
	}
	return
}

func DecEInfo(key *c_type.Uint256, einfo *c_type.Einfo) (asset c_type.Asset, memo c_type.Uint512, ar c_type.Uint256, e error) {
	ret := C.superzk_dec_info(
		(*C.uchar)(unsafe.Pointer(&key[0])),
		(*C.uchar)(unsafe.Pointer(&einfo[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&memo[0])),
		(*C.uchar)(unsafe.Pointer(&ar[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("dec info error: %d", int(ret))
		return
	}
	return
}

type DecInfoDesc struct {
	//---in---
	Key   c_type.Uint256
	Einfo c_type.Einfo
	//---out---
	Asset_ret c_type.Asset
	Memo_ret  c_type.Uint512
	Ar_ret    c_type.Uint256
}

func DecOutput(desc *DecInfoDesc) (e error) {
	ret := C.superzk_dec_info(
		(*C.uchar)(unsafe.Pointer(&desc.Key[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Einfo[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset_ret.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset_ret.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset_ret.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset_ret.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Memo_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Ar_ret[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("dec info error: %d", int(ret))
		return
	}
	return
}

func GenKey(pkr *c_type.PKr) (key c_type.Uint256, rpk c_type.Uint256, rsk c_type.Uint256, e error) {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	rsk = c_type.RandUint256()
	ret := C.superzk_gen_key(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&rsk[0])),
		(*C.uchar)(unsafe.Pointer(&key[0])),
		(*C.uchar)(unsafe.Pointer(&rpk[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("gen key error: %d", int(ret))
		return
	}
	return
}

func FetchKey(pkr *c_type.PKr, tk *c_type.Tk, rpk *c_type.Uint256) (key c_type.Uint256, vskr c_type.Uint256, e error) {
	assertPKr(pkr)
	assertTk(tk)
	pkr = ClearPKr(pkr)
	tk = ClearTk(tk)
	ret := C.superzk_fetch_key(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&rpk[0])),
		(*C.uchar)(unsafe.Pointer(&key[0])),
		(*C.uchar)(unsafe.Pointer(&vskr[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("fetch key error: %d", int(ret))
		return
	}
	return
}
