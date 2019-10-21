package c_superzk

/*

#cgo CFLAGS: -I ../czero/szk_include

#cgo LDFLAGS: -L ../czero/lib -lsuperzk

*/
import "C"

func Is_czero_debug() bool {
	return false
}
