package c_type

type LICr struct {
	Proof [PROOF_WIDTH]byte
	C     uint64
	L     uint64
	H     uint64
}

func (self *LICr) GetProp() (counteract uint64, limit uint64) {
	return 0, 0
}
