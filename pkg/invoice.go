package giga

import "github.com/shopspring/decimal"

// Invoice is a request for payment created by Gigawallet.
type Invoice struct {
	// ID is the single-use address that the invoice needs to be paid to.
	ID      Address `json:"id"`      // pay-to Address (Invoice ID)
	Account Address `json:"account"` // an Account.Address (Account ID)
	TXID    string  `json:"txid"`
	Vendor  string  `json:"vendor"`
	Items   []Item  `json:"items"`
	// These are used internally to track invoice status.
	KeyIndex      uint32 `json:"-"` // which HD Wallet child-key was generated
	BlockID       string `json:"-"` // transaction seen in this mined block
	Confirmations int32  `json:"-"` // number of confirmed blocks (since block_id)
}

// CalcTotal sums up the Items listed on the Invoice.
func (i *Invoice) CalcTotal() CoinAmount {
	total := ZeroCoins
	for _, item := range i.Items {
		total = total.Add(decimal.NewFromInt(int64(item.Quantity)).Mul(item.Price))
	}
	return total
}

type Item struct {
	Name      string          `json:"name"`
	Price     decimal.Decimal `json:"price"`
	Quantity  int             `json:"quantity"`
	ImageLink string          `json:"image_link"`
}
