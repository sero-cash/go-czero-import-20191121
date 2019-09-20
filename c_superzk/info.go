package c_superzk

/*

#include "zero.h"

*/
import "C"
import (
	"github.com/sero-cash/go-czero-import/c_type"
)

type EncInfoDesc struct {
	//---in---
	Key   c_type.Uint256
	Asset c_type.Asset
	Memo  c_type.Uint512
	Ar    c_type.Uint256
	//---out---
	Einfo c_type.Einfo
}

func EncOutput(desc *EncInfoDesc) {
	bytes := []byte{}
	bytes = append(bytes, desc.Asset.Tkn_currency[:]...)
	bytes = append(bytes, desc.Asset.Tkn_value[:]...)
	bytes = append(bytes, desc.Asset.Tkt_category[:]...)
	bytes = append(bytes, desc.Asset.Tkt_value[:]...)
	bytes = append(bytes, desc.Memo[:]...)
	bytes = append(bytes, desc.Ar[:]...)
	copy(desc.Einfo[:], bytes)
}

type DecInfoDesc struct {
	//---in---
	Key   c_type.Uint256
	Einfo c_type.Einfo
	//---out---
	Asset_ret c_type.Asset
	Memo      c_type.Uint512
	Ar_ret    c_type.Uint256
}

func DecOutput(desc *DecInfoDesc) {
	len := 0
	copy(desc.Asset_ret.Tkn_currency[:], desc.Einfo[len:])
	len += 32
	copy(desc.Asset_ret.Tkn_value[:], desc.Einfo[len:])
	len += 32
	copy(desc.Asset_ret.Tkt_category[:], desc.Einfo[len:])
	len += 32
	copy(desc.Asset_ret.Tkt_value[:], desc.Einfo[len:])
	len += 32
	copy(desc.Memo[:], desc.Einfo[len:])
	len += 64
	copy(desc.Ar_ret[:], desc.Einfo[len:])
}
