package plasma

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Output represents the outputs of a transaction
type Output struct {
	Owner  common.Address `json:"Owner"`
	Amount *big.Int       `json:"Amount"`
}

func NewOutput(owner common.Address, amount *big.Int) *Output {
	return &Output{
		Owner:  owner,
		Amount: amount,
	}
}