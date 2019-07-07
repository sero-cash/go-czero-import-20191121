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
		return 0
	} else {
		return uint64(1293528)
	}
}

const MAX_O_INS_LENGTH = int(2500)

const MAX_TX_OUT_COUNT_LENGTH = int(256)
