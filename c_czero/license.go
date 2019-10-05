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
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_superzk"

	"github.com/sero-cash/go-czero-import/c_type"

	"github.com/sero-cash/go-czero-import/seroparam"
)

const PROOF_WIDTH = 131

type LICr struct {
	Proof [PROOF_WIDTH]byte
	C     uint64
	L     uint64
	H     uint64
}

func Pk2PKrAndLICr(addr *c_type.Uint512, height uint64) (pkr c_type.PKr, licr LICr, ret bool) {
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
	if !c_superzk.Czero_isPKrValid(pkr) {
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
