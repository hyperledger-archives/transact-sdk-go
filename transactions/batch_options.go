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
