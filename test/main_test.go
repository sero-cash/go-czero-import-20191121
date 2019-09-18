// Copyright 2015 The sero.cash Authors
// This file is part of the sero.cash library.
//
// The libzero library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The libzero library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the libzero library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"testing"

	"github.com/sero-cash/go-czero-import/c_czero"
	"github.com/sero-cash/go-czero-import/c_type"
	"github.com/sero-cash/go-czero-import/cpt"
	"github.com/sero-cash/go-czero-import/keys"
)

func TestCpt(t *testing.T) {
	rad := c_czero.Random()
	base58 := c_czero.Base58Encode(rad[:])
	if base58 == nil {
		t.FailNow()
	}
	rad_ret := c_type.Uint256{}
	e := c_czero.Base58Decode(base58, rad_ret[:])
	if e != nil {
		t.FailNow()
	}
	if rad_ret != rad {
		t.FailNow()
	}
}

func TestSk(t *testing.T) {
	seed := c_czero.Random()
	sk := keys.Seed2Sk(&seed)
	fmt.Println(sk)
}

func TestPKr(t *testing.T) {
	seed := c_czero.Random()
	pk := keys.Seed2Addr(&seed)
	pkr := c_czero.Addr2PKr(&pk, nil)
	pk_58 := c_czero.Base58Encode(pk[:])
	pkr_58 := c_czero.Base58Encode(pkr[:])

	p768 := c_type.PKr{}
	p512 := c_type.Uint512{}

	cpt.Base58Decode(pk_58, p768[:])
	cpt.Base58Decode(pkr_58, p512[:])

	p512_str := c_czero.Base58Encode(p768[:])
	p256_str := c_czero.Base58Encode(p512[:])

	fmt.Println(p512_str)
	fmt.Println(p256_str)
}

func TestKeys(t *testing.T) {
	seed := c_czero.Random()
	tk := keys.Seed2Tk(&seed)
	pk := keys.Seed2Addr(&seed)

	if tk == pk {
		t.FailNow()
	}

	r := c_czero.Random()
	pkr := c_czero.Addr2PKr(&pk, &r)
	is_my_pkr := keys.IsMyPKr(&tk, &pkr)
	if !is_my_pkr {
		t.FailNow()
	}

	seed1 := c_czero.Random()
	pk1 := keys.Seed2Addr(&seed1)
	tk1 := keys.Seed2Tk(&seed1)
	pkr1 := c_czero.Addr2PKr(&pk1, &r)
	is_my_pkr = keys.IsMyPKr(&tk1, &pkr1)
	if !is_my_pkr {
		t.FailNow()
	}
	is_my_pkr_err := keys.IsMyPKr(&tk, &pkr1)
	if is_my_pkr_err {
		t.FailNow()
	}

	h := c_czero.Random()
	sign, err := keys.SignPKr(&seed, &h, &pkr)
	if err != nil {
		t.FailNow()
	}

	v_ok := keys.VerifyPKr(&h, &sign, &pkr)
	if !v_ok {
		t.FailNow()
	}

	v_ok_err := keys.VerifyPKr(&h, &sign, &pkr1)
	if v_ok_err {
		t.FailNow()
	}
}

func TestMain(m *testing.M) {
	//cpt.ZeroInit("", c_czero.NET_Dev)
	cpt.ZeroInit_OnlyInOuts()
	m.Run()
}
