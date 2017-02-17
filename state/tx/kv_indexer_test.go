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

	err := indexer.Index([]IndexerKVPair{
		{hash, *txResult},
	})
	require.Nil(t, err)

	loadedTxResult, err := indexer.Tx(hash)
	require.Nil(t, err)
	assert.Equal(t, txResult, loadedTxResult)
}

func benchmarkKVIndexerIndex(txsCount int, b *testing.B) {
	tx := types.Tx("HELLO WORLD")
	txResult := &types.TxResult{tx, 1, abci.ResponseDeliverTx{Data: []byte{0}, Code: abci.CodeType_OK, Log: ""}}

	dir, err := ioutil.TempDir("", "tx_indexer_db")
	if err != nil {
		b.Fatal(err)
	}
	defer os.RemoveAll(dir)

	store := db.NewDB("tx_indexer", "leveldb", dir)
	indexer := &KVIndexer{store}

	batch := make([]IndexerKVPair, txsCount)
	for i := 0; i < txsCount; i++ {
		batch[i] = IndexerKVPair{fmt.Sprintf("hash%v", i), *txResult}
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		err = indexer.Index(batch)
	}
}

func BenchmarkKVIndexerIndex1(b *testing.B)     { benchmarkKVIndexerIndex(1, b) }
func BenchmarkKVIndexerIndex500(b *testing.B)   { benchmarkKVIndexerIndex(500, b) }
func BenchmarkKVIndexerIndex1000(b *testing.B)  { benchmarkKVIndexerIndex(1000, b) }
func BenchmarkKVIndexerIndex2000(b *testing.B)  { benchmarkKVIndexerIndex(2000, b) }
func BenchmarkKVIndexerIndex10000(b *testing.B) { benchmarkKVIndexerIndex(10000, b) }
