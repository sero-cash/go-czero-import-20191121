package superzk

import "C"
import (
	"errors"

	"github.com/sero-cash/go-czero-import/c_czero"
	"github.com/sero-cash/go-czero-import/c_superzk"
	"github.com/sero-cash/go-czero-import/c_type"
	"github.com/sero-cash/go-czero-import/seroparam"
)

func Pk2PKr(addr *c_type.Uint512, r *c_type.Uint256) (pkr c_type.PKr) {
	if c_superzk.IsSzkPK(addr) {
		pkr, _ = c_superzk.Pk2PKr(addr, r)
		return
	} else {
		pkr, _ = c_superzk.Czero_PK2PKr(addr, r)
		return
	}
}

func IsPKValid(pk *c_type.Uint512) bool {
	if c_superzk.IsSzkPK(pk) {
		return c_superzk.IsPKValid(pk)
	} else {
		return c_superzk.Czero_isPKValid(pk)
	}
}

func IsPKrValid(pkr *c_type.PKr) bool {
	if c_superzk.IsSzkPKr(pkr) {
		return c_superzk.IsPKrValid(pkr)
	} else {
		return c_superzk.Czero_isPKrValid(pkr)
	}
}

func IsMyPKr(tk *c_type.Tk, pkr *c_type.PKr) (succ bool) {
	if c_superzk.IsSzkPKr(pkr) {
		return c_superzk.IsMyPKr(tk, pkr)
	} else {
		if e := c_superzk.Czero_isMyPKr(tk, pkr); e != nil {
			return false
		} else {
			return true
		}
	}
}

func SignPKr_ByHeight(num uint64, sk *c_type.Uint512, data *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	if num >= seroparam.SIP5() {
		return c_superzk.SignPKr_X(sk, data, pkr)
	} else {
		if c_superzk.IsSzkPKr(pkr) {
			e = errors.New("sign pkr error: szk pkr not support < sip5")
			return
		} else {
			return c_czero.SignPKrBySk(sk, data, pkr)
		}
	}
}

func VerifyPKr_ByHeight(num uint64, data *c_type.Uint256, sign *c_type.Uint512, pkr *c_type.PKr) bool {
	if num >= seroparam.SIP5() {
		return c_superzk.VerifyPKr_X(data, sign, pkr)
	} else {
		if c_superzk.IsSzkPKr(pkr) {
			return false
		} else {
			return c_czero.VerifyPKr(data, sign, pkr)
		}
	}
}
