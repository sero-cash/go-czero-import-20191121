// copyright 2018 The sero.cash Authors
// This file is part of the go-sero library.
//
// The go-sero library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-sero library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-sero library. If not, see <http://www.gnu.org/licenses/>.

package seroparam

var is_dev = false

func Init_Dev(dev bool) {
	is_dev = dev
}
func Is_Dev() bool {
	return is_dev
}

var confirmedBlock = uint64(12)

func InitComfirmedBlock(delay uint64) {
	confirmedBlock = delay
}

func DefaultConfirmedBlock() uint64 {
	if is_dev {
		return confirmedBlock
	} else {
		return confirmedBlock
	}
}

var is_exchange = false

func InitExchange(b bool) {
	is_exchange = b
}

func IsExchange() bool {
	return is_exchange
}

var is_exchangeValueStr = false

func InitExchangeValueStr(b bool) {
	is_exchangeValueStr = b
}

func IsExchangeValueStr() bool {
	return is_exchangeValueStr
}

var currentBlockNumber = uint64(0)

func InitCurrentBlockNumber(number uint64) {
	currentBlockNumber = number
}
func DefaultCurrentBlock() uint64 {
	return currentBlockNumber
}

var is_offline = false

func Init_Offline(offline bool) {
	is_offline = offline
}
func Is_Offline() bool {
	return is_offline
}
