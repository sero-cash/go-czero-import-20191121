package c_superzk

import (
	"testing"

	"github.com/sero-cash/go-czero-import/c_type"
)

func TestCzeroAccount(t *testing.T) {
	seed := c_type.RandUint256()
	sk := Czero_Seed2Sk(&seed)
	tk, e := Czero_sk2Tk(&sk)
	if e != nil {
		t.Fatal(e)
	}
	pk, e := Czero_Tk2Pk(&tk)
	if e != nil {
		t.Fatal(e)
	}
	if !Czero_isPKValid(&pk) {
		t.Fatal("is pk valid error")
	}
	r := c_type.RandUint256()
	pkr, e := Czero_PK2PKr(&pk, &r)
	if e != nil {
		t.Fatal(e)
	}
	if !Czero_isPKrValid(&pkr) {
		t.Fatal("is pkr valid error")
	}
	ret := Czero_isMyPKr(&tk, &pkr)
	if ret != nil {
		t.Fatal("is my pkr error")
	}

	h := c_type.RandUint256()
	sign, e := SignPKr_P0(&h, &sk, &pkr)
	if e != nil {
		t.Fatal(e)
	}
	if VerifyPKr_P0(&h, &sign, &pkr) != nil {
		t.Fatal("verify pkr error")
	}

	root_cm := RandomPt()
	nl, e := Czero_genNil(&sk, &root_cm)
	if e != nil {
		t.Fatal(e)
	}
	nl_sign, e := SignNil_P0(&h, &sk, &pkr, &root_cm)
	if e != nil {
		t.Fatal(e)
	}
	if e := VerifyNil_P0(&h, &nl_sign, &pkr, &root_cm, &nl); e != nil {
		t.Fatal(e)
	}
}
