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

	"github.com/sero-cash/go-czero-import/seroparam"
)

func Sk2PK(sk *c_type.Uint512) (addr c_type.Uint512) {
	C.zero_sk2pk(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&addr[0])),
	)
	return
}

func Seed2Sk(seed *c_type.Uint256) (sk c_type.Uint512) {
	C.zero_seed2sk(
		(*C.uchar)(unsafe.Pointer(&seed[0])),
		(*C.uchar)(unsafe.Pointer(&sk[0])),
	)
	return
}

func Sk2Tk(sk *c_type.Uint512) (tk c_type.Uint512) {
	C.zero_sk2tk(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&tk[0])),
	)
	return
}

func Tk2Pk(tk *c_type.Uint512) (pk c_type.Uint512) {
	C.zero_tk2pk(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&pk[0])),
	)
	return
}

func Seed2Tk(seed *c_type.Uint256) (tk c_type.Uint512) {
	C.zero_seed2tk(
		(*C.uchar)(unsafe.Pointer(&seed[0])),
		(*C.uchar)(unsafe.Pointer(&tk[0])),
	)
	return
}

func Seed2Addr(seed *c_type.Uint256) (addr c_type.Uint512) {
	C.zero_seed2pk(
		(*C.uchar)(unsafe.Pointer(&seed[0])),
		(*C.uchar)(unsafe.Pointer(&addr[0])),
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

func Seeds2Tks(seeds []c_type.Uint256) (tks []c_type.Uint512) {
	for _, seed := range seeds {
		tks = append(tks, Seed2Tk(&seed))
	}
	return
}

func Addr2PKr(addr *c_type.Uint512, r *c_type.Uint256) (pkr c_type.PKr) {
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

func HashPKr(pkr *c_type.PKr) (ret [20]byte) {
	C.zero_hpkr(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&ret[0])),
	)
	return
}

const PROOF_WIDTH = 131

type LICr struct {
	Proof [PROOF_WIDTH]byte
	C     uint64
	L     uint64
	H     uint64
}

func Addr2PKrAndLICr(addr *c_type.Uint512, height uint64) (pkr c_type.PKr, licr LICr, ret bool) {
	r := C.zero_pk2pkr_and_licr(
		//---in---
		(*C.uchar)(unsafe.Pointer(&addr[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		C.ulong(height),
		//---out--
		(*C.ulong)(unsafe.Pointer(&licr.C)),
		(*C.ulong)(unsafe.Pointer(&licr.L)),
		(*C.ulong)(unsafe.Pointer(&licr.H)),
		(*C.uchar)(unsafe.Pointer(&licr.Proof[0])),
	)
	licr.C = 1000000000000
	if r == C.char(0) {
		ret = true
	} else {
		ret = false
	}
	return
}

func CheckLICr(pkr *c_type.PKr, licr *LICr, height uint64) bool {
	if seroparam.Is_Dev() {
		return true
	}
	if !PKrValid(pkr) {
		return false
	}
	r := C.zero_check_licr(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&licr.Proof[0])),
		(C.ulong)(licr.C),
		(C.ulong)(licr.L),
		(C.ulong)(licr.H),
		(C.ulong)(height),
	)
	if r == C.char(0) {
		return true
	} else {
		return false
	}
}

func (self *LICr) GetProp() (counteract uint64, limit uint64) {
	return 0, 0
}

func IsMyPKr(tk *c_type.Uint512, pkr *c_type.PKr) (succ bool) {
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

func FetchKey(tk *c_type.Uint512, rpk *c_type.Uint256) (ret c_type.Uint256, flag bool) {
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

func SignPKr(seed *c_type.Uint256, data *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	C.zero_sign_pkr(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&seed[0])),
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

func PKrValid(pkr *c_type.PKr) bool {
	ret := C.zero_pkr_valid(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret == C.char(0) {
		return true
	} else {
		return false
	}
}
