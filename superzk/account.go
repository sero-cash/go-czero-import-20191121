package superzk

import "C"
import (
	"errors"

	"github.com/sero-cash/go-czero-import/c_czero"
	"github.com/sero-cash/go-czero-import/c_superzk"
	"github.com/sero-cash/go-czero-import/c_type"
	"github.com/sero-cash/go-czero-import/seroparam"
)

func Seed2Sk(seed *c_type.Uint256, version int) (ret c_type.Uint512) {
	if version == 1 {
		return c_superzk.Czero_Seed2Sk(seed)
	} else if version == 2 {
		return c_superzk.Seed2Sk(seed)
	} else {
		panic(errors.New("seed2sk version must in [1,2]"))
	}
}

func Sk2Tk(sk *c_type.Uint512) (tk c_type.Tk, e error) {
	if c_superzk.IsSzkSk(sk) {
		return c_superzk.Sk2Tk(sk)
	} else {
		return c_superzk.Czero_sk2Tk(sk)
	}
}

func Tk2Pk(tk *c_type.Tk) (ret c_type.Uint512, e error) {
	if c_superzk.IsSzkTk(tk) {
		return c_superzk.Tk2Pk(tk)
	} else {
		return c_superzk.Czero_Tk2Pk(tk)
	}
}

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
	if c_superzk.IsSzkTk(tk) && c_superzk.IsSzkPKr(pkr) {
		return c_superzk.IsMyPKr(tk, pkr)
	} else {
		if (!c_superzk.IsSzkTk(tk)) && (!c_superzk.IsSzkPKr(pkr)) {
			if c_superzk.Czero_isMyPKr(tk, pkr) == nil {
				return true
			}
		}
	}
	return false
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

func FetchRootCM(tk *c_type.Tk, nl *c_type.Uint256, baser *c_type.Uint256) (root_cm c_type.Uint256, e error) {
	if c_superzk.IsSzkTk(tk) && c_superzk.IsSzkNil(nl) {
		return c_superzk.Czero_fetchRootCM(tk, baser)
	} else {
		if !c_superzk.IsSzkTk(tk) && !c_superzk.IsSzkNil(nl) {
			return c_superzk.FetchRootCM(tk, nl, baser)
		}
	}
	e = errors.New("superzk fetch rootcm error: tk nl not match")
	return
}
