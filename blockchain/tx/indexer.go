package tx

import (
	"errors"

	"github.com/tendermint/tendermint/types"
)

// Indexer interface defines methods to index and search transactions. It is
// designed that way so it is easy to swap default KVIndexer with more
// advanced tools like Lucene, Solr, bleve
// https://github.com/blevesearch/bleve, ElasticSearch, etc.
type Indexer interface {
	// Index analyzes, indexes or stores mapped result fields. Supplied
	// identifier is bound to analyzed result and will be retrieved by search
	// requests.
	Index(hash string, result types.TxResult) error
	// Tx returns specified transaction or nil if the transaction is not indexed
	// or stored.
	Tx(hash string) (*types.TxResult, error)
}

// ErrorEmptyHash indicates empty hash
var ErrorEmptyHash = errors.New("Transaction hash cannot be empty")
