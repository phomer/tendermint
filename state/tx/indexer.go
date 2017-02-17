package tx

import (
	"errors"

	"github.com/tendermint/tendermint/types"
)

// Indexer interface defines methods to index and search transactions.
//
// It is designed that way so it is easy to swap default KVIndexer with more
// advanced tools like Lucene, Solr, bleve
// https://github.com/blevesearch/bleve, ElasticSearch, etc. We are talking
// about string key here, which we could get from `result.Tx`.
type Indexer interface {

	// Index analyzes, indexes or stores mapped result fields. Supplied
	// hash is bound to analyzed result and will be retrieved by search
	// requests.
	//
	// Index takes an batch of transactions because handling transactions one by
	// one would be slow.
	Index(batch []IndexerKVPair) error

	// Tx returns specified transaction or nil if the transaction is not indexed
	// or stored.
	Tx(hash string) (*types.TxResult, error)
}

// IndexerKVPair is a key-value tuple, used by Index function.
type IndexerKVPair struct {
	Hash   string
	Result types.TxResult
}

// ErrorEmptyHash indicates empty hash
var ErrorEmptyHash = errors.New("Transaction hash cannot be empty")
