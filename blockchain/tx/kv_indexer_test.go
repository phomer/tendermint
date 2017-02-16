package tx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	abci "github.com/tendermint/abci/types"
	db "github.com/tendermint/go-db"
	"github.com/tendermint/tendermint/types"
)

func TestKVIndexerIndex(t *testing.T) {
	indexer := &KVIndexer{db.NewMemDB()}

	tx := types.Tx("HELLO WORLD")
	txResult := &types.TxResult{tx, 1, abci.ResponseDeliverTx{Data: []byte{0}, Code: abci.CodeType_OK, Log: ""}}
	hash := string(tx.Hash())

	indexer.Index(hash, *txResult)
	loadedTxResult, err := indexer.Tx(hash)
	assert.Nil(t, err)
	assert.Equal(t, txResult, loadedTxResult)
}
