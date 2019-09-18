package c_superzk

/*

#include "zero.h"

*/
import "C"
import (
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

func GenRootCM_P(
	index uint64,
	asset *c_type.Asset,
	pkr *c_type.PKr,
	rsk *c_type.Uint256,
) (cm c_type.Uint256) {
	memo := c_type.Uint512{}
	out_cm := c_type.Uint256{}
	C.zero_out_commitment(
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&memo[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&rsk[0])),
		(*C.uchar)(unsafe.Pointer(&out_cm[0])),
	)
	C.zero_root_commitment(
		C.ulong(index),
		(*C.uchar)(unsafe.Pointer(&out_cm[0])),
		(*C.uchar)(unsafe.Pointer(&cm[0])),
	)
	return
}

func GenRootCM_C(
	index uint64,
	asset_cm *c_type.Uint256,
	pkr *c_type.PKr,
	rpk *c_type.Uint256,
) (cm c_type.Uint256) {
	C.zero_root_commitment(
		C.ulong(index),
		(*C.uchar)(unsafe.Pointer(&asset_cm[0])),
		(*C.uchar)(unsafe.Pointer(&cm[0])),
	)
	return
}
