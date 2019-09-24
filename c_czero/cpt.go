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

var init_chan = make(chan bool)

func ZeroInit(account_dir string, netType c_type.NetType) error {
	go func() {
		C.zero_init(
			(C.CString)(account_dir),
			C.uchar(netType),
		)
		init_chan <- true
	}()
	<-init_chan
	return nil
}

func ZeroInit_NoCircuit() error {
	go func() {
		C.zero_init_no_circuit()
		init_chan <- true
	}()
	<-init_chan
	return nil
}

func ZeroInit_OnlyInOuts() error {
	go func() {
		C.zero_init_inouts()
		init_chan <- true
	}()
	<-init_chan
	return nil
}

type Pre struct {
	I uint32
	R c_type.Uint256
	Z [2]c_type.Uint256 //I256
}
type Extra struct {
	Pre
	O      [2]c_type.Uint256 //I256
	S1_ret c_type.Uint256
}

type Out struct {
	Addr           c_type.Uint512
	Value          c_type.Uint256 //U256
	Info           c_type.Uint512
	EText_ret      [c_type.ETEXT_WIDTH]byte
	Currency       c_type.Uint256
	Commitment_ret c_type.Uint256
}

type In struct {
	EText      [c_type.ETEXT_WIDTH]byte
	Commitment c_type.Uint256
	Path       [c_type.DEPTH * 32]byte
	S1         c_type.Uint256
	Index      uint32
	Currency   c_type.Uint256
	Anchor     c_type.Uint256
	Nil_ret    c_type.Uint256
	Trace_ret  c_type.Uint256
}

type Common struct {
	Seed     c_type.Uint256
	Hash_O   c_type.Uint256
	Currency c_type.Uint256
	C        [2]c_type.Uint256
}

func GenOutCM(
	asset c_type.Asset,
	memo *c_type.Uint512,
	pkr *c_type.PKr,
	rsk *c_type.Uint256,
) (cm c_type.Uint256) {
	C.zero_out_commitment(
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&memo[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&rsk[0])),
		(*C.uchar)(unsafe.Pointer(&cm[0])),
	)
	return
}

func GenRootCM(
	index uint64,
	out_cm *c_type.Uint256,
) (cm c_type.Uint256) {
	C.zero_root_commitment(
		C.ulong(index),
		(*C.uchar)(unsafe.Pointer(&out_cm[0])),
		(*C.uchar)(unsafe.Pointer(&cm[0])),
	)
	return
}

type ConfirmOutputDesc struct {
	Asset  c_type.Asset
	Memo   c_type.Uint512
	Pkr    c_type.PKr
	Rsk    c_type.Uint256
	Out_cm c_type.Uint256
}

func ConfirmOutput(desc *ConfirmOutputDesc) (e error) {
	ret := C.zero_output_confirm(
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Memo[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Pkr[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Rsk[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Out_cm[0])),
	)
	if ret == 0 {
		return
	} else {
		e = errors.New("confirm output error")
		return
	}
}

type OutputDesc struct {
	//---in---
	Tkn_currency c_type.Uint256
	Tkn_value    c_type.Uint256
	Tkt_category c_type.Uint256
	Tkt_value    c_type.Uint256
	Memo         c_type.Uint512
	Pkr          c_type.PKr
	Height       uint64
	//---out---
	Asset_cm_ret c_type.Uint256
	Ar_ret       c_type.Uint256
	Key_ret      c_type.Uint256
	Out_cm_ret   c_type.Uint256
	Einfo_ret    c_type.Einfo
	RPK_ret      c_type.Uint256
	Proof_ret    c_type.Proof
}

func GenOutputProof(desc *OutputDesc) (e error) {
	var is_v1 int
	if desc.Height >= seroparam.SIP2() {
		is_v1 = 1
	} else {
		is_v1 = 0
	}
	ret := C.zero_output(
		//---in---
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Memo[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Pkr[0])),
		(C.int)(is_v1),
		//---out---
		(*C.uchar)(unsafe.Pointer(&desc.Asset_cm_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Ar_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Key_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Out_cm_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Einfo_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.RPK_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Proof_ret[0])),
	)
	if ret == 0 {
		return
	} else {
		e = errors.New("gen output proof error")
		return
	}
}

type InfoDesc struct {
	//---in---
	Key   c_type.Uint256
	Flag  bool
	Einfo c_type.Einfo
	//---out---
	Asset c_type.Asset
	Rsk   c_type.Uint256
	Memo  c_type.Uint512
}

func DecOutput(desc *InfoDesc) {
	flag := C.char(0)
	if desc.Flag {
		flag = C.char(1)
	}
	C.zero_dec_einfo(
		//--in--
		(*C.uchar)(unsafe.Pointer(&desc.Key[0])),
		flag,
		(*C.uchar)(unsafe.Pointer(&desc.Einfo[0])),
		//--out--
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Rsk[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Memo[0])),
	)
}

func GenTil(tk *c_type.Tk, root_cm *c_type.Uint256) (til c_type.Uint256) {
	C.zero_til(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&til[0])),
	)
	return
}

func FetchRootCM(tk *c_type.Tk, til *c_type.Uint256) (root_cm c_type.Uint256) {
	C.zero_til2cm(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&til[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
	)
	return
}

func GenNil(sk *c_type.Uint512, root_cm *c_type.Uint256) (nil c_type.Uint256) {
	C.zero_nil(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&nil[0])),
	)
	return
}

type InputDesc struct {
	//---in0--
	Sk c_type.Uint512
	//---in---
	Seed  c_type.Uint256
	Pkr   c_type.PKr
	RPK   c_type.Uint256
	Einfo c_type.Einfo
	//--
	Index    uint64
	Anchor   c_type.Uint256
	Position uint32
	Path     [c_type.DEPTH * 32]byte
	//---out---
	Asset_cm_ret c_type.Uint256
	Ar_ret       c_type.Uint256
	Nil_ret      c_type.Uint256
	Til_ret      c_type.Uint256
	Proof_ret    [c_type.PROOF_WIDTH]byte
}

func GenInputProofBySk(desc *InputDesc) (e error) {
	ret := C.zero_input_by_sk(
		//---in---
		(*C.uchar)(unsafe.Pointer(&desc.Sk[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Pkr[0])),
		(*C.uchar)(unsafe.Pointer(&desc.RPK[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Einfo[0])),
		//--
		C.ulong(desc.Index),
		(*C.uchar)(unsafe.Pointer(&desc.Anchor[0])),
		C.ulong(desc.Position),
		(*C.uchar)(unsafe.Pointer(&desc.Path[0])),
		//---out---
		(*C.uchar)(unsafe.Pointer(&desc.Asset_cm_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Ar_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Nil_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Til_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Proof_ret[0])),
	)
	if ret == 0 {
		return
	} else {
		e = errors.New("gen input desc error")
		return
	}
}

func PtrOfSlice(s []byte) *C.uchar {
	if len(s) > 0 {
		return (*C.uchar)(unsafe.Pointer(&s[0]))
	} else {
		return (*C.uchar)(unsafe.Pointer(uintptr(0)))
	}
}

func SignBalance(desc *c_type.BalanceDesc) {
	C.zero_sign_balance(
		C.int(len(desc.Zin_acms)/32),
		PtrOfSlice(desc.Zin_acms),
		PtrOfSlice(desc.Zin_ars),
		C.int(len(desc.Zout_acms)/32),
		PtrOfSlice(desc.Zout_acms),
		PtrOfSlice(desc.Zout_ars),
		C.int(len(desc.Oin_accs)/32),
		PtrOfSlice(desc.Oin_accs),
		C.int(len(desc.Oout_accs)/32),
		PtrOfSlice(desc.Oout_accs),
		(*C.uchar)(unsafe.Pointer(&desc.Hash[0])),
		//---out--
		(*C.uchar)(unsafe.Pointer(&desc.Bsign[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Bcr[0])),
	)
	return
}

func VerifyBalance(desc *c_type.BalanceDesc) (e error) {
	ret := C.zero_verify_balance(
		C.int(len(desc.Zin_acms)/32),
		PtrOfSlice(desc.Zin_acms),
		C.int(len(desc.Zout_acms)/32),
		PtrOfSlice(desc.Zout_acms),
		C.int(len(desc.Oin_accs)/32),
		PtrOfSlice(desc.Oin_accs),
		C.int(len(desc.Oout_accs)/32),
		PtrOfSlice(desc.Oout_accs),
		(*C.uchar)(unsafe.Pointer(&desc.Hash[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Bcr[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Bsign[0])),
	)
	if ret == 0 {
		return
	} else {
		e = errors.New("verify balance error")
		return
	}
}

type AssetDesc struct {
	Tkn_currency c_type.Uint256
	Tkn_value    c_type.Uint256
	Tkt_category c_type.Uint256
	Tkt_value    c_type.Uint256
	Asset_cc     c_type.Uint256
	Asset_cm     c_type.Uint256
}

func GenAssetCC(desc *AssetDesc) {
	C.zero_gen_asset_cc(
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_value[0])),
		//--out--
		(*C.uchar)(unsafe.Pointer(&desc.Asset_cc[0])),
	)
}

type OutputVerifyDesc struct {
	AssetCM c_type.Uint256
	OutCM   c_type.Uint256
	RPK     c_type.Uint256
	Proof   c_type.Proof
	Height  uint64
}

func VerifyOutput(desc *OutputVerifyDesc) (e error) {
	var is_v1 int
	if desc.Height >= seroparam.SIP2() {
		is_v1 = 1
	} else {
		is_v1 = 0
	}
	ret := C.zero_output_verify(
		(*C.uchar)(unsafe.Pointer(&desc.AssetCM[0])),
		(*C.uchar)(unsafe.Pointer(&desc.OutCM[0])),
		(*C.uchar)(unsafe.Pointer(&desc.RPK[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Proof[0])),
		(C.int)(is_v1),
	)
	if ret == 0 {
		return
	} else {
		e = errors.New("verify output error")
		return
	}
}

type InputVerifyDesc struct {
	AssetCM c_type.Uint256
	Anchor  c_type.Uint256
	Nil     c_type.Uint256
	Proof   c_type.Proof
}

func VerifyInput(desc *InputVerifyDesc) (e error) {
	ret := C.zero_input_verify(
		(*C.uchar)(unsafe.Pointer(&desc.AssetCM[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Anchor[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Nil[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Proof[0])),
	)
	if ret == 0 {
		return
	} else {
		e = errors.New("verify output error")
		return
	}
}

type ConfirmPkgDesc struct {
	Tkn_currency c_type.Uint256
	Tkn_value    c_type.Uint256
	Tkt_category c_type.Uint256
	Tkt_value    c_type.Uint256
	Memo         c_type.Uint512
	Ar_ret       c_type.Uint256
	Pkg_cm       c_type.Uint256
}

type InputSDesc struct {
	//---in0---
	Sk c_type.Uint512
	//---in---
	Ehash  c_type.Uint256
	Seed   c_type.Uint256
	Pkr    c_type.PKr
	RootCM c_type.Uint256
	//---out---
	Nil_ret  c_type.Uint256
	Til_ret  c_type.Uint256
	Sign_ret c_type.Uint512
}

func GenInputSProofBySk(desc *InputSDesc) (e error) {
	ret := C.zero_input_s_by_sk(
		//---in---
		(*C.uchar)(unsafe.Pointer(&desc.Ehash[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Sk[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Pkr[0])),
		(*C.uchar)(unsafe.Pointer(&desc.RootCM[0])),
		//---out---
		(*C.uchar)(unsafe.Pointer(&desc.Nil_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Til_ret[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Sign_ret[0])),
	)
	if ret == 0 {
		return
	} else {
		e = errors.New("gen input s desc error")
		return
	}
}

type VerifyInputSDesc struct {
	//---in---
	Ehash  c_type.Uint256
	RootCM c_type.Uint256
	Pkr    c_type.PKr
	Nil    c_type.Uint256
	Sign   c_type.Uint512
}

func VerifyInputS(desc *VerifyInputSDesc) (e error) {
	ret := C.zero_verify_input_s(
		//---in---
		(*C.uchar)(unsafe.Pointer(&desc.Ehash[0])),
		(*C.uchar)(unsafe.Pointer(&desc.RootCM[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Pkr[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Nil[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Sign[0])),
	)
	if ret == 0 {
		return
	} else {
		e = errors.New("gen input desc error")
		return
	}
}
