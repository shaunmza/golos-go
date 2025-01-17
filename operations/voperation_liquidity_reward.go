package operations

import (
	"github.com/shaunmza/golos-go/types"
)

//LiquidityRewardOperation represents liquidity_reward operation data.
type LiquidityRewardOperation struct {
	Owner  string      `json:"owner"`
	Payout types.Asset `json:"payout"`
}

//Type function that defines the type of operation LiquidityRewardOperation.
func (op *LiquidityRewardOperation) Type() OpType {
	return TypeLiquidityReward
}

//Data returns the operation data LiquidityRewardOperation.
func (op *LiquidityRewardOperation) Data() interface{} {
	return op
}
