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

func GenRootCM_P(
	index uint64,
	asset *c_type.Asset,
	ar *c_type.Uint256,
	pkr *c_type.PKr,
) (cm c_type.Uint256, e error) {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	ret := C.superzk_gen_root_cm_p(
		C.ulong(index),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&ar[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&cm[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("gen root cm p error: %d", int(ret))
	}
	return
}

func GenRootCM_C(
	index uint64,
	asset_cm *c_type.Uint256,
	pkr *c_type.PKr,
) (cm c_type.Uint256, e error) {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	ret := C.superzk_gen_root_cm_c(
		C.ulong(index),
		(*C.uchar)(unsafe.Pointer(&asset_cm[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&cm[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("gen root cm c error: %d", int(ret))
	}
	return
}
