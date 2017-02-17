package tx

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/abci/types"
	db "github.com/tendermint/go-db"
	"github.com/tendermint/tendermint/types"
)

func TestKVIndexerIndex(t *testing.T) {
	indexer := &KVIndexer{db.NewMemDB()}

	tx := types.Tx("HELLO WORLD")
	txResult := &types.TxResult{tx, 1, abci.ResponseDeliverTx{Data: []byte{0}, Code: abci.CodeType_OK, Log: ""}}
	hash := string(tx.Hash())

	err := indexer.Index(hash, *txResult)
	require.Nil(t, err)

	loadedTxResult, err := indexer.Tx(hash)
	require.Nil(t, err)
	assert.Equal(t, txResult, loadedTxResult)
}

func BenchmarkKVIndexerIndex(b *testing.B) {
	tx := types.Tx("HELLO WORLD")
	txResult := &types.TxResult{tx, 1, abci.ResponseDeliverTx{Data: []byte{0}, Code: abci.CodeType_OK, Log: ""}}

	dir, err := ioutil.TempDir("", "tx_indexer_db")
	if err != nil {
		b.Fatal(err)
	}
	defer os.RemoveAll(dir)

	store := db.NewDB("tx_indexer", "leveldb", dir)
	indexer := &KVIndexer{store}

	for n := 0; n < b.N; n++ {
		hash := fmt.Sprintf("hash%v", n)
		err = indexer.Index(hash, *txResult)
	}
}
