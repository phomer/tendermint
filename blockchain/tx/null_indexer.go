package tx

import (
	"github.com/tendermint/tendermint/types"
)

// NullIndexer acts as a /dev/null.
type NullIndexer struct{}

// Tx gets transaction from the KV storage and returns it or nil if the
// transaction is not found.
func (indexer *NullIndexer) Tx(hash string) (*types.TxResult, error) {
	panic("You are trying to get the transaction from a null indexer")
}

// Index synchronously writes transaction to the KV storage.
func (indexer *NullIndexer) Index(hash string, txResult types.TxResult) error {
	return nil
}
