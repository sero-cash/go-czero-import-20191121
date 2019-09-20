package c_superzk

/*

#include "zero.h"

*/
import "C"
import (
	"math/big"

	"github.com/sero-cash/go-czero-import/c_type"
)

func GenRootCM_P(
	index uint64,
	asset *c_type.Asset,
	pkr *c_type.PKr,
	rsk *c_type.Uint256,
) (cm c_type.Uint256) {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	copy(cm[:], big.NewInt(int64(index+1)).Bytes())
	return
}

func GenRootCM_C(
	index uint64,
	asset_cm *c_type.Uint256,
	pkr *c_type.PKr,
	rpk *c_type.Uint256,
) (cm c_type.Uint256) {
	assertPKr(pkr)
	pkr = ClearPKr(pkr)
	copy(cm[:], big.NewInt(int64(index+1)).Bytes())
	return
}
