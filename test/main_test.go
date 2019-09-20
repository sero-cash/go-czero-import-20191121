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
	"github.com/sero-cash/go-czero-import/superzk"
)

func TestSk(t *testing.T) {
	seed := c_type.RandUint256()
	sk := superzk.Seed2Sk(&seed)
	fmt.Println(sk)
}

func TestKeys(t *testing.T) {
	seed := c_type.RandUint256()
	sk := superzk.Seed2Sk(&seed)
	tk := superzk.Seed2Tk(&seed)
	pk := c_czero.Tk2Pk(&tk)

	if tk == pk {
		t.FailNow()
	}

	r := c_type.RandUint256()
	pkr := superzk.Pk2PKr(&pk, &r)
	is_my_pkr := superzk.IsMyPKr(&tk, &pkr)
	if !is_my_pkr {
		t.FailNow()
	}

	seed1 := c_type.RandUint256()
	tk1 := superzk.Seed2Tk(&seed1)
	pk1 := c_czero.Tk2Pk(&tk1)
	pkr1 := superzk.Pk2PKr(&pk1, &r)
	is_my_pkr = superzk.IsMyPKr(&tk1, &pkr1)
	if !is_my_pkr {
		t.FailNow()
	}
	is_my_pkr_err := superzk.IsMyPKr(&tk, &pkr1)
	if is_my_pkr_err {
		t.FailNow()
	}

	h := c_type.RandUint256()
	sign, err := superzk.SignPKrBySk(&sk, &h, &pkr)
	if err != nil {
		t.FailNow()
	}

	v_ok := superzk.VerifyPKr(&h, &sign, &pkr)
	if !v_ok {
		t.FailNow()
	}

	v_ok_err := superzk.VerifyPKr(&h, &sign, &pkr1)
	if v_ok_err {
		t.FailNow()
	}
}

func TestMain(m *testing.M) {
	superzk.ZeroInit_OnlyInOuts()
	m.Run()
}
