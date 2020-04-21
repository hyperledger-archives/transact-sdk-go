// Copyright 2020 Tyson Foods, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// -----------------------------------------------------------------------------

package transactions

import (
	"github.com/golang/protobuf/proto"

	"github.com/hyperledger/transact-sdk-go/errors"
	"github.com/hyperledger/transact-sdk-go/src/protobuf/transaction_pb2"
	"github.com/hyperledger/transact-sdk-go/transactions/signing"
)

// IBatchBuilder defines the interface for building Batches of Transactions
type IBatchBuilder interface {
	GetTransactions() []*transaction_pb2.Transaction
	GetTrace() bool
	Build(signer *signing.Signer) ([]byte, error)
}

// NewBatchBuilder returns a new instance of IBatchBuilder interface
func NewBatchBuilder(opts ...BatchOption) (IBatchBuilder, error) {
	b := &BatchBuilder{}
	for _, opt := range opts {
		err := opt(b)
		if err != nil {
			return nil, err
		}
	}
	return b, nil
}

// BatchBuilder implements the builder pattern for batches of transactions
// and implements the IBatchBuilder interface
type BatchBuilder struct {
	transactions []*transaction_pb2.Transaction
	trace        bool
}

// GetTransactions returns the slice of transactions for the batch
func (b *BatchBuilder) GetTransactions() []*transaction_pb2.Transaction {
	return b.transactions
}

// GetTrace returns the trace setting for the batch
func (b *BatchBuilder) GetTrace() bool {
	return b.trace
}

func (b *BatchBuilder) setTransactions(t []*transaction_pb2.Transaction) { b.transactions = t }
func (b *BatchBuilder) setTrace(t bool)                                  { b.trace = t }

func (b *BatchBuilder) buildHeader(signer *signing.Signer) ([]byte, error) {
	if len(b.transactions) == 0 {
		return nil, errors.NewMissingFieldError("Transactions")
	}

	transactionIds := make([]string, len(b.transactions))
	for _, txn := range b.transactions {
		transactionIds = append(transactionIds, txn.GetHeaderSignature())
	}

	header := &transaction_pb2.BatchHeader{
		SignerPublicKey: signer.GetPublicKey().AsHex(),
		TransactionIds:  transactionIds,
	}

	headerBytes, err := proto.Marshal(header)
	if err != nil {
		return nil, errors.NewProtobufEncodingError(err)
	}

	return headerBytes, nil
}

// Build returns a built batch as a byte slice or an error
// if fields were missing or proto failed to marshal
func (b *BatchBuilder) Build(signer *signing.Signer) ([]byte, error) {
	header, err := b.buildHeader(signer)
	if err != nil {
		return nil, err
	}

	headerSignature := signer.Sign(header)
	batch := &transaction_pb2.Batch{
		Header:          header,
		HeaderSignature: headerSignature,
		Transactions:    b.transactions,
		Trace:           b.trace,
	}
	batchList := &transaction_pb2.BatchList{
		Batches: []*transaction_pb2.Batch{
			batch,
		},
	}

	batchListBytes, err := proto.Marshal(batchList)
	if err != nil {
		return nil, errors.NewProtobufEncodingError(err)
	}

	return batchListBytes, nil
}
