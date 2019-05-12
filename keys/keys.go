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

package keys

/*
#cgo CFLAGS: -I ../czero/include

#cgo LDFLAGS: -L ../czero/lib

#cgo LDFLAGS: -lczero

#include "zero.h"
*/
import "C"

import (
	"crypto/rand"
	"errors"
	"unsafe"
)

func logBytes(bytes []byte) {
	C.zero_log_bytes(
		(*C.uchar)(unsafe.Pointer(&bytes[0])),
		(C.int)(len(bytes)),
	)
	return
}

func Sk2PK(sk *Uint512) (addr Uint512) {
	C.zero_sk2pk(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&addr[0])),
	)
	return
}

func Seed2Tk(seed *Uint256) (tk Uint512) {
	C.zero_seed2tk(
		(*C.uchar)(unsafe.Pointer(&seed[0])),
		(*C.uchar)(unsafe.Pointer(&tk[0])),
	)
	return
}

func Seed2Addr(seed *Uint256) (addr Uint512) {
	C.zero_seed2pk(
		(*C.uchar)(unsafe.Pointer(&seed[0])),
		(*C.uchar)(unsafe.Pointer(&addr[0])),
	)
	return
}

func Seed2PKr(seed *Uint256, rnd *Uint256) (pkr PKr) {
	addr := Seed2Addr(seed)
	var from_r Uint256
	if rnd != nil {
		copy(from_r[:], rnd[:])
	} else {
		from_r = RandUint256()
	}
	pkr = Addr2PKr(&addr, &from_r)
	return
}

func IsPKValid(pk *Uint512) bool {
	ret := C.zero_pk_valid(
		(*C.uchar)(unsafe.Pointer(&pk[0])),
	)

	if ret == C.char(0) {
		return true
	} else {
		return false
	}
}

func RandUint512() (hash Uint512) {
	rand.Read(hash[:])
	return
}

func RandUint256() (hash Uint256) {
	rand.Read(hash[:])
	return
}

func RandUint128() (hash Uint128) {
	rand.Read(hash[:])
	return
}

func Addr2PKr(addr *Uint512, r *Uint256) (pkr PKr) {
	if r == nil {
		t := RandUint256()
		r = &t
	} else {
		if (*r) == Empty_Uint256 {
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

func HashPKr(pkr *PKr) (ret [20]byte) {
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

func Addr2PKrAndLICr(addr *Uint512, height uint64) (pkr PKr, licr LICr, ret bool) {
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
	if r == C.char(0) {
		ret = true
	} else {
		ret = false
	}
	return
}

func CheckLICr(pkr *PKr, licr *LICr, height uint64) bool {
	//log.Info("CHECKLICr", "height:", height, "L", licr.L)
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

func IsMyPKr(tk *Uint512, pkr *PKr) (succ bool) {
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

func FetchKey(tk *Uint512, rpk *Uint256) (ret Uint256, flag bool) {
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

func SignPKrBySk(sk *Uint512, data *Uint256, pkr *PKr) (sign Uint512, e error) {
	C.zero_sign_pkr_by_sk(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
	)
	if sign == Empty_Uint512 {
		e = errors.New("SignOAddr: sign is empty")
		return
	} else {
		return
	}
}

func SignPKr(seed *Uint256, data *Uint256, pkr *PKr) (sign Uint512, e error) {
	C.zero_sign_pkr(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&seed[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])),
	)
	if sign == Empty_Uint512 {
		e = errors.New("SignOAddr: sign is empty")
		return
	} else {
		return
	}
}

func VerifyPKr(data *Uint256, sign *Uint512, pkr *PKr) bool {
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

func PKrValid(pkr *PKr) bool {
	ret := C.zero_pkr_valid(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret == C.char(0) {
		return true
	} else {
		return false
	}
}
