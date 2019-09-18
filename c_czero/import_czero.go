package c_czero

/*

#cgo CFLAGS: -I ../czero/include

#cgo LDFLAGS: -L ../czero/lib -lczero

*/
import "C"

func Is_czero_debug() bool {
	return false
}
