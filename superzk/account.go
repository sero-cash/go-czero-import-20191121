package superzk

import "C"
import (
	"github.com/sero-cash/go-czero-import/c_czero"
	"github.com/sero-cash/go-czero-import/c_superzk"
	"github.com/sero-cash/go-czero-import/c_type"
)

func Seed2Tk(seed *c_type.Uint256) (tk c_type.Uint512) {
	sk := c_superzk.Seed2Sk(seed)
	return c_superzk.Sk2Tk(&sk)
}

func Sk2Tk(sk *c_type.Uint512) (tk c_type.Uint512) {
	return c_superzk.Sk2Tk(sk)
}

func Seed2Sk(seed *c_type.Uint256) (sk c_type.Uint512) {
	return c_superzk.Seed2Sk(seed)
}

func Pk2PKr(addr *c_type.Uint512, r *c_type.Uint256) (pkr c_type.PKr) {
	if c_superzk.IsSzkPK(addr) {
		return c_superzk.Pk2PKr(addr, r)
	} else {
		return c_czero.Pk2PKr(addr, r)
	}
}

func IsPKValid(pk *c_type.Uint512) bool {
	if c_superzk.IsSzkPK(pk) {
		return c_superzk.IsPKValid(pk)
	} else {
		return c_czero.IsPKValid(pk)
	}
}

func IsPKrValid(pkr *c_type.PKr) bool {
	if c_superzk.IsSzkPKr(pkr) {
		return c_superzk.IsPKrValid(pkr)
	} else {
		return c_czero.IsPKrValid(pkr)
	}
}

func HashPKr(pkr *c_type.PKr) (ret [20]byte) {
	return c_superzk.HashPKr(pkr)
}

func IsMyPKr(tk *c_type.Uint512, pkr *c_type.PKr) (succ bool) {
	if c_superzk.IsSzkPKr(pkr) {
		return c_superzk.IsMyPKr(tk, pkr)
	} else {
		return c_czero.IsMyPKr(tk, pkr)
	}
}

func SignPKrBySk(sk *c_type.Uint512, data *c_type.Uint256, pkr *c_type.PKr) (sign c_type.Uint512, e error) {
	if c_superzk.IsSzkPKr(pkr) {
		return c_superzk.SignPKrBySk(sk, data, pkr)
	} else {
		return c_czero.SignPKrBySk(sk, data, pkr)
	}
}

func VerifyPKr(data *c_type.Uint256, sign *c_type.Uint512, pkr *c_type.PKr) bool {
	if c_superzk.IsSzkPKr(pkr) {
		return c_superzk.VerifyPKr(data, sign, pkr)
	} else {
		return c_czero.VerifyPKr(data, sign, pkr)
	}
}
