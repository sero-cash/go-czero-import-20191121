package c_czero

/*

#include "zero.h"

*/
import "C"

/*
func ConfirmPkg(desc *ConfirmPkgDesc) (e error) {
	ret := C.zero_pkg_confirm(
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Memo[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Ar_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Pkg_cm[0])),
	)
	if ret == 0 {
		return
	} else {
		e = errors.New("confirm pkg error")
		return
	}
}

type PkgDesc struct {
	//---in---
	Key          c_type.Uint256
	Tkn_currency c_type.Uint256
	Tkn_value    c_type.Uint256
	Tkt_category c_type.Uint256
	Tkt_value    c_type.Uint256
	Memo         c_type.Uint512
	//---out---
	Asset_cm_ret c_type.Uint256
	Ar_ret       c_type.Uint256
	Pkg_cm_ret   c_type.Uint256
	Einfo_ret    c_type.Einfo
	Proof_ret    c_type.Proof
}

func GenPkgProof(desc *PkgDesc) (e error) {
	ret := C.zero_pkg(
		//---in---
		(*C.uchar)(unsafe.Pointer(&desc.Key[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Memo[0])),
		//---out---
		(*C.uchar)(unsafe.Pointer(&desc.Asset_cm_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Ar_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Pkg_cm_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Einfo_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Proof_ret[0])),
	)
	if ret == 0 {
		return
	} else {
		e = errors.New("gen pkg proof error")
		return
	}
}

type PkgVerifyDesc struct {
	AssetCM c_type.Uint256
	PkgCM   c_type.Uint256
	Proof   c_type.Proof
}

func VerifyPkg(desc *PkgVerifyDesc) (e error) {
	ret := C.zero_pkg_verify(
		(*C.uchar)(unsafe.Pointer(&desc.AssetCM[0])),
		(*C.uchar)(unsafe.Pointer(&desc.PkgCM[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Proof[0])),
	)
	if ret == 0 {
		return
	} else {
		e = errors.New("verify pkg error")
		return
	}
}*/
