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

package cpt

/*

#cgo CFLAGS: -I ../czero/include

#cgo LDFLAGS: -L ../czero/lib -lczero

#include "zero.h"

*/
import "C"
import (
	"encoding/hex"
	"errors"
	"fmt"
	"unsafe"

	"github.com/sero-cash/go-sero/crypto/sha3"

	"github.com/sero-cash/go-czero-import/keys"
)

func Is_czero_debug() bool {
	return false
}

var init_chan = make(chan bool)

type NetType uint8

const (
	NET_Dev   NetType = 0
	NET_Alpha NetType = 1
	NET_Beta  NetType = 2
)

func ZeroInit(account_dir string, netType NetType) error {
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

func Random() (out keys.Uint256) {
	C.zero_random32(
		(*C.uchar)(unsafe.Pointer(&out[0])),
	)
	return
}

func Force_Fr(data *keys.Uint256) (fr keys.Uint256) {
	C.zero_force_fr(
		(*C.uchar)(unsafe.Pointer(&data[0])),
		(*C.uchar)(unsafe.Pointer(&fr[0])),
	)
	return
}

func Combine(l *keys.Uint256, r *keys.Uint256) (out keys.Uint256) {
	C.zero_merkle_combine(
		(*C.uchar)(unsafe.Pointer(&l[0])),
		(*C.uchar)(unsafe.Pointer(&r[0])),
		(*C.uchar)(unsafe.Pointer(&out[0])),
	)
	return
}

func Base58Encode(bytes []byte) (ret *string) {
	str := C.zero_base58_enc(
		(*C.uchar)(unsafe.Pointer(&bytes[0])),
		C.int(len(bytes)),
	)
	if str != nil {
		defer C.zero_fee_str(str)
		s := C.GoString(str)
		ret = &s
		return
	} else {
		return
	}
}

func Base58Decode(str *string, bytes []byte) (e error) {
	ret := C.zero_base58_dec(
		C.CString(*str),
		(*C.uchar)(unsafe.Pointer(&bytes[0])),
		C.int(len(bytes)),
	)
	if ret == C.char(0) {
		return
	} else {
		e = errors.New("base58 can not decode string")
		return
	}
}

type Pre struct {
	I uint32
	R keys.Uint256
	Z [2]keys.Uint256 //I256
}
type Extra struct {
	Pre
	O      [2]keys.Uint256 //I256
	S1_ret keys.Uint256
}

type Out struct {
	Addr           keys.Uint512
	Value          keys.Uint256 //U256
	Info           keys.Uint512
	EText_ret      [ETEXT_WIDTH]byte
	Currency       keys.Uint256
	Commitment_ret keys.Uint256
}

type In struct {
	EText      [ETEXT_WIDTH]byte
	Commitment keys.Uint256
	Path       [DEPTH * 32]byte
	S1         keys.Uint256
	Index      uint32
	Currency   keys.Uint256
	Anchor     keys.Uint256
	Nil_ret    keys.Uint256
	Trace_ret  keys.Uint256
}

type Proof [PROOF_WIDTH]byte

func (b Proof) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b[:])
	return result, nil
}

func (b *Proof) UnmarshalText(input []byte) error {
	raw := input[2:]
	if len(raw) == 0 {
		return nil
	}
	dec := Proof{}
	if len(raw)/2 != len(dec[:]) {
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(dec[:])*2, "Proof")
	}
	if _, err := hex.Decode(dec[:], raw); err != nil {
		return err
	} else {
		*b = dec
	}
	return nil
}

func (self *Proof) ToHash() (ret keys.Uint256) {
	d := sha3.NewKeccak256()
	d.Write(self[:])
	copy(ret[:], d.Sum(nil))
	return
}

type Common struct {
	Seed     keys.Uint256
	Hash_O   keys.Uint256
	Currency keys.Uint256
	C        [2]keys.Uint256
}

func GenOutCM(
	tkn_currency *keys.Uint256,
	tkn_value *keys.Uint256,
	tkt_category *keys.Uint256,
	tkt_value *keys.Uint256,
	memo *keys.Uint512,
	pkr *keys.PKr,
	rsk *keys.Uint256,
) (cm keys.Uint256) {
	C.zero_out_commitment(
		(*C.uchar)(unsafe.Pointer(&tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&memo[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&rsk[0])),
		(*C.uchar)(unsafe.Pointer(&cm[0])),
	)
	return
}

func GenRootCM(
	index uint64,
	out_cm *keys.Uint256,
) (cm keys.Uint256) {
	C.zero_root_commitment(
		C.ulong(index),
		(*C.uchar)(unsafe.Pointer(&out_cm[0])),
		(*C.uchar)(unsafe.Pointer(&cm[0])),
	)
	return
}

type ConfirmOutputDesc struct {
	Tkn_currency keys.Uint256
	Tkn_value    keys.Uint256
	Tkt_category keys.Uint256
	Tkt_value    keys.Uint256
	Memo         keys.Uint512
	Pkr          keys.PKr
	Rsk          keys.Uint256
	Out_cm       keys.Uint256
}

func ConfirmOutput(desc *ConfirmOutputDesc) (e error) {
	ret := C.zero_output_confirm(
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_value[0])),
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
	Tkn_currency keys.Uint256
	Tkn_value    keys.Uint256
	Tkt_category keys.Uint256
	Tkt_value    keys.Uint256
	Memo         keys.Uint512
	Pkr          keys.PKr
	Height       uint64
	//---out---
	Asset_cm_ret keys.Uint256
	Ar_ret       keys.Uint256
	Out_cm_ret   keys.Uint256
	Einfo_ret    Einfo
	RPK_ret      keys.Uint256
	Proof_ret    Proof
}

func GenOutputProof(desc *OutputDesc) (e error) {
	var is_v1 int
	if desc.Height >= SIP2 {
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

type Einfo [INFO_WIDTH]byte

func (b Einfo) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b[:])
	return result, nil
}

func (b *Einfo) UnmarshalText(input []byte) error {
	raw := input[2:]
	if len(raw) == 0 {
		return nil
	}
	dec := Einfo{}
	if len(raw)/2 != len(dec[:]) {
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(dec[:])*2, "Einfo")
	}
	if _, err := hex.Decode(dec[:], raw); err != nil {
		return err
	} else {
		*b = dec
	}
	return nil
}

type EncOutputInfo struct {
	//---in---
	Key          keys.Uint256
	Tkn_currency keys.Uint256
	Tkn_value    keys.Uint256
	Tkt_category keys.Uint256
	Tkt_value    keys.Uint256
	Rsk          keys.Uint256
	Memo         keys.Uint512
	//---out---
	Einfo Einfo
}

func EncOutput(desc *EncOutputInfo) {
	C.zero_enc_info(
		//--in--
		(*C.uchar)(unsafe.Pointer(&desc.Key[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Rsk[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Memo[0])),
		//--out--
		(*C.uchar)(unsafe.Pointer(&desc.Einfo[0])),
	)
}

type InfoDesc struct {
	//---in---
	Key   keys.Uint256
	Flag  bool
	Einfo Einfo
	//---out---
	Tkn_currency keys.Uint256
	Tkn_value    keys.Uint256
	Tkt_category keys.Uint256
	Tkt_value    keys.Uint256
	Rsk          keys.Uint256
	Memo         keys.Uint512
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
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Rsk[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Memo[0])),
	)
}

func GenTil(tk *keys.Uint512, root_cm *keys.Uint256) (til keys.Uint256) {
	C.zero_til(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&til[0])),
	)
	return
}

func GenNil(sk *keys.Uint512, root_cm *keys.Uint256) (nil keys.Uint256) {
	C.zero_nil(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&nil[0])),
	)
	return
}

type InputDesc struct {
	//---in0--
	Sk keys.Uint512
	//---in---
	Seed  keys.Uint256
	Pkr   keys.PKr
	RPK   keys.Uint256
	Einfo Einfo
	//--
	Index    uint64
	Anchor   keys.Uint256
	Position uint32
	Path     [DEPTH * 32]byte
	//---out---
	Asset_cm_ret keys.Uint256
	Ar_ret       keys.Uint256
	Nil_ret      keys.Uint256
	Til_ret      keys.Uint256
	Proof_ret    [PROOF_WIDTH]byte
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

func GenInputProof(desc *InputDesc) (e error) {
	ret := C.zero_input(
		//---in---
		(*C.uchar)(unsafe.Pointer(&desc.Seed[0])),
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

type BalanceDesc struct {
	Zin_acms  []byte
	Zin_ars   []byte
	Zout_acms []byte
	Zout_ars  []byte
	Oin_accs  []byte
	Oout_accs []byte
	Hash      keys.Uint256
	Bcr       keys.Uint256
	Bsign     keys.Uint512
}

func PtrOfSlice(s []byte) *C.uchar {
	if len(s) > 0 {
		return (*C.uchar)(unsafe.Pointer(&s[0]))
	} else {
		return (*C.uchar)(unsafe.Pointer(uintptr(0)))
	}
}

func SignBalance(desc *BalanceDesc) {
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

func VerifyBalance(desc *BalanceDesc) (e error) {
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
	Tkn_currency keys.Uint256
	Tkn_value    keys.Uint256
	Tkt_category keys.Uint256
	Tkt_value    keys.Uint256
	Asset_cc     keys.Uint256
	Asset_cm     keys.Uint256
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
	AssetCM keys.Uint256
	OutCM   keys.Uint256
	RPK     keys.Uint256
	Proof   Proof
	Height  uint64
}

func VerifyOutput(desc *OutputVerifyDesc) (e error) {
	var is_v1 int
	if desc.Height >= SIP2 {
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
	AssetCM keys.Uint256
	Anchor  keys.Uint256
	Nil     keys.Uint256
	Proof   Proof
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
	Tkn_currency keys.Uint256
	Tkn_value    keys.Uint256
	Tkt_category keys.Uint256
	Tkt_value    keys.Uint256
	Memo         keys.Uint512
	Ar_ret       keys.Uint256
	Pkg_cm       keys.Uint256
}

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
	Key          keys.Uint256
	Tkn_currency keys.Uint256
	Tkn_value    keys.Uint256
	Tkt_category keys.Uint256
	Tkt_value    keys.Uint256
	Memo         keys.Uint512
	//---out---
	Asset_cm_ret keys.Uint256
	Ar_ret       keys.Uint256
	Pkg_cm_ret   keys.Uint256
	Einfo_ret    Einfo
	Proof_ret    Proof
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
	AssetCM keys.Uint256
	PkgCM   keys.Uint256
	Proof   Proof
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
}

type InputSDesc struct {
	//---in0---
	Sk keys.Uint512
	//---in---
	Ehash  keys.Uint256
	Seed   keys.Uint256
	Pkr    keys.PKr
	RootCM keys.Uint256
	//---out---
	Nil_ret  keys.Uint256
	Til_ret  keys.Uint256
	Sign_ret keys.Uint512
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

func GenInputSProof(desc *InputSDesc) (e error) {
	ret := C.zero_input_s(
		//---in---
		(*C.uchar)(unsafe.Pointer(&desc.Ehash[0])),
		(*C.uchar)(unsafe.Pointer(&desc.Seed[0])),
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
	Ehash  keys.Uint256
	RootCM keys.Uint256
	Pkr    keys.PKr
	Nil    keys.Uint256
	Sign   keys.Uint512
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

func Miner_Hash_0(in []byte, num uint64) []byte {
	var bs [64]byte
	if num >= VP1 {
		C.zero_hash_2_enter(
			(*C.uchar)(unsafe.Pointer(&in[0])),
			(*C.uchar)(unsafe.Pointer(&bs[0])),
		)
	} else if num >= SIP1 {
		C.zero_hash_1_enter(
			(*C.uchar)(unsafe.Pointer(&in[0])),
			(*C.uchar)(unsafe.Pointer(&bs[0])),
		)
	} else {
		C.zero_hash_0_enter(
			(*C.uchar)(unsafe.Pointer(&in[0])),
			(*C.uchar)(unsafe.Pointer(&bs[0])),
		)
	}
	return bs[:]
}

func Miner_Hash_1(in []byte, num uint64) []byte {
	var bs [32]byte
	if num >= VP1 {
		C.zero_hash_2_leave(
			(*C.uchar)(unsafe.Pointer(&in[0])),
			(*C.uchar)(unsafe.Pointer(&bs[0])),
		)
	} else if num >= SIP1 {
		C.zero_hash_1_leave(
			(*C.uchar)(unsafe.Pointer(&in[0])),
			(*C.uchar)(unsafe.Pointer(&bs[0])),
		)
	} else {
		C.zero_hash_0_leave(
			(*C.uchar)(unsafe.Pointer(&in[0])),
			(*C.uchar)(unsafe.Pointer(&bs[0])),
		)
	}
	return bs[:]
}
