package c_superzk

/*

#cgo CFLAGS: -I ../czero/szk_include

#cgo LDFLAGS: -L ../czero/lib -lsuperzkd

*/
import "C"

func Is_czero_debug() bool {
	return false
}
