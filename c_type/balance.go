package c_type

type BalanceDesc struct {
	Zin_acms  []byte
	Zin_ars   []byte
	Zout_acms []byte
	Zout_ars  []byte
	Oin_accs  []byte
	Oout_accs []byte
	Hash      Uint256
	Bcr       Uint256
	Bsign     Uint512
}

type Asset struct {
	Tkn_currency Uint256
	Tkn_value    Uint256
	Tkt_category Uint256
	Tkt_value    Uint256
}

func (self Asset) NewRef() (ret *Asset) {
	return &self
}
