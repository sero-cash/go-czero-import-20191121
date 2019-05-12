package cpt

import "github.com/sero-cash/go-czero-import/keys"

func createEmpty() (ret [DEPTH + 1]keys.Uint256) {
	ret[0] = keys.Empty_Uint256
	for i := 1; i <= DEPTH; i++ {
		ret[i] = Combine(&ret[i-1], &ret[i-1])
	}
	return
}

var emptyRoots [DEPTH + 1]keys.Uint256
var is_load bool

func EmptyRoots() []keys.Uint256 {
	if !is_load {
		is_load = true
		emptyRoots = createEmpty()
	}
	return emptyRoots[:]
}
