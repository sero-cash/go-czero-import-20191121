package c_superzk

/*

#include "zero.h"

*/
import "C"
import (
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

type EncInfoDesc struct {
	//---in---
	Key   c_type.Uint256
	Asset c_type.Asset
	Memo  c_type.Uint512
	//---out---
	Einfo c_type.Einfo
}

func EncOutput(desc *EncInfoDesc) {
	rsk := c_type.Uint256{}
	C.zero_enc_info(
		//--in--
		(*C.uchar)(unsafe.Pointer(&desc.Key[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&rsk[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Memo[0])),
		//--out--
		(*C.uchar)(unsafe.Pointer(&desc.Einfo[0])),
	)
}

type DecInfoDesc struct {
	//---in---
	Key   c_type.Uint256
	Einfo c_type.Einfo
	//---out---
	Asset_ret c_type.Asset
	Memo      c_type.Uint512
}

func DecOutput(desc *DecInfoDesc) {
	flag := C.char(1)
	rsk := c_type.Uint256{}
	C.zero_dec_einfo(
		//--in--
		(*C.uchar)(unsafe.Pointer(&desc.Key[0])),
		flag,
		(*C.uchar)(unsafe.Pointer(&desc.Einfo[0])),
		//--out--
		(*C.uchar)(unsafe.Pointer(&desc.Asset_ret.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset_ret.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset_ret.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset_ret.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&rsk[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Memo[0])),
	)
}
