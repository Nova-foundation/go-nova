package utils

import "math/big"

// ToNvt number of NVT to Wei
func ToNvt(nvt uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(nvt), big.NewInt(1e18))
}
