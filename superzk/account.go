package superzk

import "github.com/sero-cash/go-czero-import/c_type"

func IsSzkPKr(pkr *c_type.PKr) bool {
	return false
}

func IsSzkUint512(pk *c_type.Uint512) bool {
	return false
}

func IsSzkUint256(nil *c_type.Uint256) bool {
	return false
}
