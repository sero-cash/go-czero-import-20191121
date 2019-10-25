package seroparam

func SIP1() uint64 {
	if is_dev {
		return 0
	} else {
		return uint64(130000) //for miner rewards
	}
}

func SIP2() uint64 {
	if is_dev {
		return 0
	} else {
		return uint64(606006)
	}
}

func SIP3() uint64 {
	if is_dev {
		return uint64(940410)
	} else {
		return uint64(940410)
	}
}

func VP1() uint64 {
	if is_dev {
		return 0
	} else {
		return uint64(829000)
	}
}

func VP0() uint64 {
	if is_dev {
		return 0
	} else {
		return uint64(788888)
	}
}

func SIP4() uint64 {
	if is_dev {
		return 50
	} else {
		return uint64(1300000)
	}
}

func SIP5() uint64 {
	if is_dev {
		return 100
	} else {
		return uint64(1958696)
	}
}

const MAX_O_INS_LENGTH = int(2500)

const MAX_O_OUT_LENGTH = int(10)

const MAX_Z_OUT_LENGTH_OLD = int(6)

const MAX_Z_OUT_LENGTH_SIP2 = int(500)

const MAX_CONTRACT_OUT_COUNT_LENGTH = int(256)

const LOWEST_STAKING_NODE_FEE_RATE = uint32(2500)

const HIGHEST_STAKING_NODE_FEE_RATE = uint32(7500)
