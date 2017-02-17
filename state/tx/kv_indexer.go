package tx

import (
	"bytes"
	"fmt"

	db "github.com/tendermint/go-db"
	"github.com/tendermint/go-wire"
	"github.com/tendermint/tendermint/types"
)

// KVIndexer is a simpliest possible indexer, backed
// by Key-Value storage (levelDB). It could only
// index transaction by its identifier.
type KVIndexer struct {
	Store db.DB
}

// Tx gets transaction from the KV storage and returns it or nil if the
// transaction is not found.
func (indexer *KVIndexer) Tx(hash string) (*types.TxResult, error) {
	if hash == "" {
		return nil, ErrorEmptyHash
	}

	rawBytes := indexer.Store.Get([]byte(hash))
	if rawBytes == nil {
		return nil, nil
	}

	r := bytes.NewReader(rawBytes)
	var n int
	var err error
	txResult := wire.ReadBinary(&types.TxResult{}, r, 0, &n, &err).(*types.TxResult)
	if err != nil {
		return nil, fmt.Errorf("Error reading TxResult: %v", err)
	}

	return txResult, nil
}

// Index writes transactions to the KV storage.
func (indexer *KVIndexer) Index(batch []IndexerKVPair) error {
	storeBatch := indexer.Store.NewBatch()
	for _, pair := range batch {
		if pair.Hash == "" {
			return ErrorEmptyHash
		}
		rawBytes := wire.BinaryBytes(&pair.Result)
		storeBatch.Set([]byte(pair.Hash), rawBytes)
	}
	storeBatch.Write()
	return nil
}
