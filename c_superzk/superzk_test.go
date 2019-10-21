package c_superzk

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/sero-cash/go-czero-import/c_type"
)

func TestMain(m *testing.M) {
	InitParams()
	m.Run()
}

func TestAccount(t *testing.T) {
	seed := c_type.RandUint256()
	fmt.Println("seed: " + hex.EncodeToString(seed[:]))
	sk := Seed2Sk(&seed)
	fmt.Println("sk: " + hex.EncodeToString(sk[:]))
	tk, e := Sk2Tk(&sk)
	if e != nil {
		t.Fatal(e)
	}
	fmt.Println("tk: " + hex.EncodeToString(tk[:]))
	pk, e := Tk2Pk(&tk)
	if e != nil {
		t.Fatal(e)
	}
	fmt.Println("pk: " + hex.EncodeToString(pk[:]))
	if !IsPKValid(&pk) {
		t.Fatal("is pk valid error")
	}
	r := c_type.RandUint256()
	fmt.Println("r: " + hex.EncodeToString(r[:]))
	pkr, e := Pk2PKr(&pk, &r)
	if e != nil {
		t.Fatal(e)
	}
	fmt.Println("pkr: " + hex.EncodeToString(pkr[:]))
	if !IsPKrValid(&pkr) {
		t.Fatal("is pkr valid error")
	}
	ret := IsMyPKr(&tk, &pkr)
	if !ret {
		t.Fatal("is my pkr error")
	}

	key, rpk, _, e := GenKey(&pkr)
	if e != nil {
		t.Fatal(e)
	}
	fkey, _, e := FetchKey(&pkr, &tk, &rpk)
	if e != nil {
		t.Fatal(e)
	}
	fmt.Println(key)
	fmt.Println(fkey)
	if key != fkey {
		t.Fatal("the key is not the same")
	}

	h := c_type.RandUint256()
	sign, e := SignPKr_P(&sk, &h, &pkr)
	if e != nil {
		t.Fatal(e)
	}
	if !VerifyPKr_P(&h, &sign, &pkr) {
		t.Fatal("verify pkr error")
	}

	a := RandomFr()
	zpka, e := GenZPKa(&pkr, &a)
	if e != nil {
		t.Fatal(e)
	}
	zpka_sign, e := SignZPKa(&sk, &h, &a, &pkr)
	if e != nil {
		t.Fatal(e)
	}
	if !VerifyZPKa(&h, &zpka_sign, &zpka) {
		t.Fatal("verify zpka error")
	}

	root_cm := RandomPt()
	nl, e := GenNil(&tk, &root_cm, &pkr)
	if e != nil {
		t.Fatal(e)
	}
	nl_sign, e := SignNil(&tk, &h, &root_cm, &pkr)
	if e != nil {
		t.Fatal(e)
	}
	if e := VerifyNil(&h, &nl_sign, &nl, &root_cm, &pkr); e != nil {
		t.Fatal(e)
	}
}

func TestInfo(t *testing.T) {
	asset := c_type.Asset{}
	asset.Tkn_currency.UnmarshalText([]byte("0xffff"))
	asset.Tkn_value.UnmarshalText([]byte("0x100"))
	asset.Tkt_category.UnmarshalText([]byte("0xffff"))
	asset.Tkt_value.UnmarshalText([]byte("0x1234"))

	desc := EncInfoDesc{}
	desc.Asset = asset
	desc.Memo = c_type.RandUint512()
	desc.Ar = RandomFr()
	desc.Key = c_type.RandUint256()

	if e := EncOutput(&desc); e != nil {
		t.Fatal(e)
	}

	d_desc := DecInfoDesc{}
	d_desc.Einfo = desc.Einfo
	d_desc.Key = desc.Key
	if e := DecOutput(&d_desc); e != nil {
		t.Fatal(e)
	}

	if desc.Memo != d_desc.Memo_ret {
		t.Fatal("memo error")
	}

	if desc.Ar != d_desc.Ar_ret {
		t.Fatal()
	}

	if desc.Asset.Tkn_currency != d_desc.Asset_ret.Tkn_currency {
		t.Fatal()
	}

	if desc.Asset.Tkn_value != d_desc.Asset_ret.Tkn_value {
		t.Fatal()
	}

	if desc.Asset.Tkt_category != d_desc.Asset_ret.Tkt_category {
		t.Fatal()
	}

	if desc.Asset.Tkt_value != d_desc.Asset_ret.Tkt_value {
		t.Fatal()
	}

}

func newUint256ByText(text string) (ret c_type.Uint256) {
	copy(ret[:], []byte(text))
	return
}

func newUint256ByInt(i int) (ret c_type.Uint256) {
	b := big.NewInt(int64(i))
	bs := b.Bytes()
	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(bs)-1-i] = bs[len(bs)-1-i], bs[i]
	}
	copy(ret[:], bs)
	return
}

func TestAsset(t *testing.T) {
	oin_asset := c_type.Asset{
		newUint256ByText("SERO"),
		newUint256ByInt(1000),
		newUint256ByText("TKT"),
		newUint256ByInt(100),
	}
	var e error
	var in_cc c_type.Uint256
	if in_cc, e = GenAssetCC(&oin_asset); e != nil {
		t.Fatal(e)
	}

	oout_asset := c_type.Asset{
		newUint256ByText("SERO"),
		newUint256ByInt(1000),
		newUint256ByText("TKT"),
		newUint256ByInt(100),
	}

	var out_cc c_type.Uint256
	if out_cc, e = GenAssetCC(&oout_asset); e != nil {
		t.Fatal(e)
	}

	zin_asset := c_type.Asset{
		newUint256ByText("SERO"),
		newUint256ByInt(1000),
		newUint256ByText("TKT"),
		newUint256ByInt(100),
	}
	zin_desc := AssetDesc{}
	zin_desc.Asset = zin_asset
	zin_desc.Ar = RandomFr()
	if e := GenAssetCM(&zin_desc); e != nil {
		t.Fatal(e)
	}

	zout_asset := c_type.Asset{
		newUint256ByText("SERO"),
		newUint256ByInt(1000),
		newUint256ByText("TKT"),
		newUint256ByInt(100),
	}
	zout_desc := AssetDesc{}
	zout_desc.Asset = zout_asset
	zout_desc.Ar = RandomFr()
	if e := GenAssetCM(&zout_desc); e != nil {
		t.Fatal(e)
	}

	balance_dec := c_type.BalanceDesc{}
	balance_dec.Hash = c_type.RandUint256()

	balance_dec.Oin_accs = append(balance_dec.Oin_accs, in_cc[:]...)
	balance_dec.Oout_accs = append(balance_dec.Oout_accs, out_cc[:]...)

	balance_dec.Zin_acms = append(balance_dec.Zin_acms, zin_desc.Asset_cm_ret[:]...)
	balance_dec.Zin_ars = append(balance_dec.Zin_ars, zin_desc.Ar[:]...)

	balance_dec.Zout_acms = append(balance_dec.Zout_acms, zout_desc.Asset_cm_ret[:]...)
	balance_dec.Zout_ars = append(balance_dec.Zout_ars, zout_desc.Ar[:]...)

	if e := SignBalance(&balance_dec); e != nil {
		t.Fatal(e)
	}

	if e := VerifyBalance(&balance_dec); e != nil {
		t.Fatal(e)
	}
}

func TestRootCM(t *testing.T) {
	asset := c_type.Asset{
		newUint256ByText("SERO"),
		newUint256ByInt(1000),
		newUint256ByText("TKT"),
		newUint256ByInt(100),
	}
	seed := c_type.RandUint256()
	sk := Seed2Sk(&seed)
	tk, _ := Sk2Tk(&sk)
	pk, _ := Tk2Pk(&tk)
	r := c_type.RandUint256()
	pkr, _ := Pk2PKr(&pk, &r)

	ar := RandomFr()
	root_cm, e := GenRootCM_P(1, &asset, &ar, &pkr)
	if e != nil {
		t.Fatal(e)
	}

	asset_desc := AssetDesc{}
	asset_desc.Asset = asset
	asset_desc.Ar = ar
	if e := GenAssetCM(&asset_desc); e != nil {
		t.Fatal(e)
	}
	root_cm_c, e := GenRootCM_C(1, &asset_desc.Asset_cm_ret, &pkr)
	if e != nil {
		t.Fatal(e)
	}
	if root_cm_c != root_cm {
		t.Fatal()
	}
}

func TestCombine(t *testing.T) {
	l := c_type.RandUint256()
	r := c_type.RandUint256()
	h := Combine(&l, &r)
	if h == c_type.Empty_Uint256 {
		t.Fatal()
	}
}
