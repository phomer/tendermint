package tx

import (
	"github.com/tendermint/tendermint/types"
)

// NullIndexer acts as a /dev/null.
type NullIndexer struct{}

// Tx panics.
func (indexer *NullIndexer) Tx(hash string) (*types.TxResult, error) {
	panic("You are trying to get the transaction from a null indexer")
}

// Index returns nil.
func (indexer *NullIndexer) Index(batch []IndexerKVPair) error {
	return nil
}
