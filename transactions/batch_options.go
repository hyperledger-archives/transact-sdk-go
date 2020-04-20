package transactions

import (
	"github.com/hyperledger/transact-sdk-go/src/protobuf/transaction_pb2"
)

// BatchOption provide the functional options for building a BatchBuilder
type BatchOption func(*BatchBuilder) error

// WithTransactions sets the Batch Option Transactions
func WithTransactions(t []*transaction_pb2.Transaction) BatchOption {
	return func(b *BatchBuilder) error {
		b.setTransactions(t)
		return nil
	}
}

// WithTrace sets the Batch Option for Trace
func WithTrace(trace bool) BatchOption {
	return func(b *BatchBuilder) error {
		b.setTrace(trace)
		return nil
	}
}
