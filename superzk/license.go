package superzk

import "C"
import (
	"github.com/sero-cash/go-czero-import/c_czero"
	"github.com/sero-cash/go-czero-import/c_superzk"

	"github.com/sero-cash/go-czero-import/c_type"
	"github.com/sero-cash/go-czero-import/seroparam"
)

func Pk2PKrAndLICr(addr *c_type.Uint512, height uint64) (pkr c_type.PKr, licr c_type.LICr, ret bool) {
	if height >= seroparam.SIP5() {
		if IsPKValid(addr) {
			pkr = Pk2PKr(addr, c_superzk.RandomFr().NewRef())
			ret = true
		}
		return
	} else {
		return c_czero.Pk2PKrAndLICr(addr, height)
	}
	return
}

func CheckLICr(pkr *c_type.PKr, licr *c_type.LICr, height uint64) bool {
	if height >= seroparam.SIP5() {
		return IsPKrValid(pkr)
	}
	return c_czero.CheckLICr(pkr, licr, height)
}
