// Copyright 2015 The sero.cash Authors
// This file is part of the sero.cash library.
//
// The libzero library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The libzero library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the libzero library. If not, see <http://www.gnu.org/licenses/>.

package c_czero

/*

#include "zero.h"

*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

func Tk2Pk(tk *c_type.Tk) (pk c_type.Uint512) {
	C.zero_tk2pk(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&pk[0])),
	)
	return
}

func IsPKValid(pk *c_type.Uint512) bool {
	ret := C.zero_pk_valid(
		(*C.uchar)(unsafe.Pointer(&pk[0])),
	)

	if ret == C.char(0) {
		return true
	} else {
		return false
	}
}

func Pk2PKr(addr *c_type.Uint512, r *c_type.Uint256) (pkr c_type.PKr) {
	if r == nil {
		t := c_type.RandUint256()
		r = &t
	} else {
		if (*r) == c_type.Empty_Uint256 {
			panic("gen pkr, but r is empty")
		}
	}

	C.zero_pk2pkr(
		(*C.uchar)(unsafe.Pointer(&addr[0])),
		(*C.uchar)(unsafe.Pointer(&r[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	return
}

func IsMyPKr(tk *c_type.Tk, pkr *c_type.PKr) (succ bool) {
	ret := C.zero_ismy_pkr(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&tk[0])),
	)
	if ret == C.char(0) {
		succ = true
		return
	} else {
		succ = false
		return
	}
}

func FetchKey(tk *c_type.Tk, rpk *c_type.Uint256) (ret c_type.Uint256, flag bool) {
	f := C.zero_fetch_key(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&rpk[0])),
		(*C.uchar)(unsafe.Pointer(&ret[0])),
	)
	if f == C.char(0) {
		flag = false
	} else {
		flag = true
	}
	return
}

func SignPKrBySk(sk *c_type.Uint512, data *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	C.zero_sign_pkr_by_sk(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
	)
	if sign == c_type.Empty_Uint512 {
		e = errors.New("SignOAddr: sign is empty")
		return
	} else {
		return
	}
}

func VerifyPKr(data *c_type.Uint256, sign *c_type.Uint512, pkr *c_type.PKr) bool {
	ret := C.zero_verify_pkr(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret == C.char(0) {
		return true
	} else {
		return false
	}
}

func IsPKrValid(pkr *c_type.PKr) bool {
	ret := C.zero_pkr_valid(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret == C.char(0) {
		return true
	} else {
		return false
	}
}
