package superzk

import "C"
import (
	"github.com/sero-cash/go-czero-import/c_czero"
	"github.com/sero-cash/go-czero-import/c_superzk"
	"github.com/sero-cash/go-czero-import/c_type"
)

func ZeroInit(account_dir string, netType c_type.NetType) error {
	c_superzk.InitParams()
	return c_czero.ZeroInit(account_dir, netType)
}

func ZeroInit_NoCircuit() error {
	c_superzk.InitParams_NoCircuit()
	return c_czero.ZeroInit_NoCircuit()
}

func ZeroInit_OnlyInOuts() error {
	c_superzk.InitParams()
	return c_czero.ZeroInit_OnlyInOuts()
}
