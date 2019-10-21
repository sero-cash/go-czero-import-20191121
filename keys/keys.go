package keys

import (
	"github.com/sero-cash/go-czero-import/c_superzk"
	"github.com/sero-cash/go-czero-import/c_type"
	"github.com/sero-cash/go-czero-import/superzk"
)

type Uint256 c_type.Uint256
type Uint512 c_type.Uint512
type PKr c_type.PKr

func RandUint256() (ret Uint256) {
	ret = Uint256(c_superzk.RandomFr())
	return
}

func Seed2Sk(seed *Uint256) (ret Uint512) {
	s := c_type.Uint256(*seed)
	sk := superzk.Seed2Sk(&s, 1)
	ret = Uint512(sk)
	return
}

func Seed2SkByVersion(seed *Uint256, version int) (ret Uint512) {
	s := c_type.Uint256(*seed)
	sk := superzk.Seed2Sk(&s, version)
	ret = Uint512(sk)
	return
}

func Sk2Tk(sk *Uint512) (ret Uint512) {
	s := c_type.Uint512(*sk)
	t, _ := superzk.Sk2Tk(&s)
	ret = Uint512(t)
	return
}

func Tk2Pk(tk *Uint512) (ret Uint512) {
	t := c_type.Tk([64]byte(c_type.Uint512(*tk)))
	p, _ := superzk.Tk2Pk(&t)
	ret = Uint512(p)
	return
}

func Addr2PKr(pk *Uint512, rnd *Uint256) (ret PKr) {
	p := c_type.Uint512(*pk)
	r := c_type.Uint256(*rnd)
	pr := superzk.Pk2PKr(&p, &r)
	ret = PKr(pr)
	return
}

func IsPKValid(pk *Uint512) bool {
	if pk == nil {
		return false
	}
	p := c_type.Uint512(*pk)
	return superzk.IsPKValid(&p)
}

func PKrValid(pkr *PKr) bool {
	if pkr == nil {
		return false
	}
	p := c_type.PKr(*pkr)
	return superzk.IsPKrValid(&p)
}
