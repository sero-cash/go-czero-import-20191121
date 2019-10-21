package c_superzk

/*

#include "csuperzk.h"

*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/sero-cash/go-czero-import/c_type"
)

func Czero_Seed2Sk(seed *c_type.Uint256) (sk c_type.Uint512) {
	C.superzk_seed2sk(
		(*C.uchar)(unsafe.Pointer(&seed[0])),
		(*C.uchar)(unsafe.Pointer(&sk[0])),
	)
	return
}

func Czero_sk2Tk(sk *c_type.Uint512) (tk c_type.Tk, e error) {
	ret := C.superzk_sk2tk(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&tk[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("sk2tk error: %d", int(ret))
		return
	}
	return
}

func Czero_Tk2Pk(tk *c_type.Tk) (pk c_type.Uint512, e error) {
	ret := C.czero_tk2pk(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&pk[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("czero tk2pk error: %d", int(ret))
		return
	}
	return
}

func Czero_PK2PKr(pk *c_type.Uint512, r *c_type.Uint256) (pkr c_type.PKr, e error) {
	if r == nil || *r == c_type.Empty_Uint256 {
		r = RandomFr().NewRef()
	}
	ret := C.czero_pk2pkr(
		(*C.uchar)(unsafe.Pointer(&pk[0])),
		(*C.uchar)(unsafe.Pointer(&r[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("czero pk2pkr error: %d", int(ret))
		return
	}
	return
}

func Czero_isMyPKr(tk *c_type.Tk, pkr *c_type.PKr) (e error) {
	if IsSzkPKr(pkr) {
		e = fmt.Errorf("czero ismypkr error")
		return
	}
	if IsSzkTk(tk) {
		e = fmt.Errorf("czero ismypkr error")
		return
	}
	ret := C.czero_ismy_pkr(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("czero ismypkr error: %d", int(ret))
		return
	}
	return
}

func Czero_isPKValid(pk *c_type.Uint512) bool {
	if IsSzkPK(pk) {
		return false
	}
	ret := C.superzk_pk_valid(
		(*C.uchar)(unsafe.Pointer(&pk[0])),
	)
	if ret != C.int(0) {
		return false
	}
	return true
}

func Czero_isPKrValid(pkr *c_type.PKr) bool {
	if IsSzkPKr(pkr) {
		return false
	}
	ret := C.superzk_pkr_valid(
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
	)
	if ret != C.int(0) {
		return false
	}
	return true
}

func Czero_fetchKey(tk *c_type.Tk, rpk *c_type.Uint256) (key c_type.Uint256, flag bool, e error) {
	if IsSzkTk(tk) {
		e = fmt.Errorf("czero fetch key but the tk is szk")
		return
	}
	ret := C.czero_fetch_key(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&rpk[0])),
		(*C.uchar)(unsafe.Pointer(&key[0])),
	)
	if int(ret) < 0 {
		e = fmt.Errorf("czero fetch key error: %d", int(ret))
		return
	}
	if int(ret) == 0 {
		flag = false
	} else {
		flag = true
	}
	return
}

func Czero_decEInfo(key *c_type.Uint256, flag bool, einfo *c_type.Einfo) (asset c_type.Asset, rsk c_type.Uint256, memo c_type.Uint512, e error) {
	var f C.char
	if flag {
		f = C.char(1)
	} else {
		f = C.char(0)
	}
	ret := C.czero_dec_einfo(
		(*C.uchar)(unsafe.Pointer(&key[0])),
		f,
		(*C.uchar)(unsafe.Pointer(&einfo[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&rsk[0])),
		(*C.uchar)(unsafe.Pointer(&memo[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("czero dec einfo error: %d", int(ret))
		return
	}
	return
}

func Czero_genNil(sk *c_type.Uint512, root_cm *c_type.Uint256) (nl c_type.Uint256, e error) {
	ret := C.czero_gen_nil(
		(*C.uchar)(unsafe.Pointer(&sk[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&nl[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("czero gen nil error: %d", int(ret))
		return
	}
	return
}

func Czero_genTrace(tk *c_type.Tk, root_cm *c_type.Uint256) (trace c_type.Uint256, e error) {
	ret := C.czero_gen_trace(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
		(*C.uchar)(unsafe.Pointer(&trace[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("czero gen trace error: %d", int(ret))
		return
	}
	return
}

func Czero_genOutCM(asset *c_type.Asset, memo *c_type.Uint512, pkr *c_type.PKr, rsk *c_type.Uint256) (out_cm c_type.Uint256, e error) {
	ret := C.czero_gen_out_cm(
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_currency[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkn_value[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_category[0])),
		(*C.uchar)(unsafe.Pointer(&asset.Tkt_value[0])),
		(*C.uchar)(unsafe.Pointer(&memo[0])),
		(*C.uchar)(unsafe.Pointer(&pkr[0])),
		(*C.uchar)(unsafe.Pointer(&rsk[0])),
		(*C.uchar)(unsafe.Pointer(&out_cm[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("czero gen out cm error: %d", int(ret))
		return
	}
	return
}

func Czero_genRootCM(index uint64, out_cm *c_type.Uint256) (cm c_type.Uint256) {
	C.czero_gen_root_cm(
		C.ulong(index),
		(*C.uchar)(unsafe.Pointer(&out_cm[0])),
		(*C.uchar)(unsafe.Pointer(&cm[0])),
	)
	return
}

func Czero_combine(l *c_type.Uint256, r *c_type.Uint256) (out c_type.Uint256) {
	C.czero_merkle_combine(
		(*C.uchar)(unsafe.Pointer(&l[0])),
		(*C.uchar)(unsafe.Pointer(&r[0])),
		(*C.uchar)(unsafe.Pointer(&out[0])),
	)
	return
}

func Czero_fetchRootCM(tk *c_type.Tk, til *c_type.Uint256) (root_cm c_type.Uint256, e error) {
	if IsSzkTk(tk) {
		e = errors.New("czero fetch rootcm error: tk is szk")
		return
	}
	if IsSzkNil(til) {
		e = errors.New("czero fetch rootcm error: til is szk")
		return
	}
	ret := C.czero_til2cm(
		(*C.uchar)(unsafe.Pointer(&tk[0])),
		(*C.uchar)(unsafe.Pointer(&til[0])),
		(*C.uchar)(unsafe.Pointer(&root_cm[0])),
	)
	if ret != C.int(0) {
		e = fmt.Errorf("czero fetch rootcm error: %d", int(ret))
		return
	}
	return
}
