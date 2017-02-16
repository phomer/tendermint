package types

import (
	abci "github.com/tendermint/abci/types"
)

// TxResult contains DeliverTx response and height of the block, where this
// transaction was committed.
//
// One usage is indexing transaction results (see "blockchain/tx/indexer.go").
type TxResult struct {
	Tx                Tx
	Height            int
	DeliverTxResponse abci.ResponseDeliverTx
}
