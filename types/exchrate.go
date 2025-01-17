package types

import (
	"github.com/shaunmza/golos-go/encoding/transaction"
)

//ExchRate is an additional structure used by other structures.
type ExchRate struct {
	Base  *Asset `json:"base"`
	Quote *Asset `json:"quote"`
}

//MarshalTransaction is a function of converting type ExchRate to bytes.
func (exch *ExchRate) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(exch.Base)
	enc.Encode(exch.Quote)
	return enc.Err()
}
