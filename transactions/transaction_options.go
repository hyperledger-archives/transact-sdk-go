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
	"github.com/hyperledger/transact-sdk-go/transactions/signing"
)

// TransactionBuilderOption provides the functional option for creating a new Transaction Builder
type TransactionBuilderOption func(t *TransactionBuilder) error

// WithBatcherPublicKey provides the TransactionBuilderOption for
// defining a public key to sign batches
func WithBatcherPublicKey(key signing.PublicKey) TransactionBuilderOption {
	return func(t *TransactionBuilder) error {
		t.setBatcherPublicKey(key)
		return nil
	}
}

// WithDependencies provides the TransactionBuilderOption for
// defining transaction dependencies
func WithDependencies(deps []string) TransactionBuilderOption {
	return func(t *TransactionBuilder) error {
		t.setDependencies(deps)
		return nil
	}
}

// WithFamilyName provides the TransactionBuilderOption for
// defining transaction family name
func WithFamilyName(name string) TransactionBuilderOption {
	return func(t *TransactionBuilder) error {
		t.setFamilyName(name)
		return nil
	}
}

// WithFamilyVersion provides the TransactionBuilderOption for
// defining transaction family version
func WithFamilyVersion(version string) TransactionBuilderOption {
	return func(t *TransactionBuilder) error {
		t.setFamilyVersion(version)
		return nil
	}
}

// WithInputs provides the TransactionBuilderOption for
// defining transaction input addresses
func WithInputs(inputs []string) TransactionBuilderOption {
	return func(t *TransactionBuilder) error {
		t.setInputs(inputs)
		return nil
	}
}

// WithOutputs provides the TransactionBuilderOption for
// defining transaction output addresses
func WithOutputs(outputs []string) TransactionBuilderOption {
	return func(t *TransactionBuilder) error {
		t.setOutputs(outputs)
		return nil
	}
}

// WithNonce provides the TransactionBuilderOption for
// defining transaction nonce
func WithNonce(nonce string) TransactionBuilderOption {
	return func(t *TransactionBuilder) error {
		t.setNonce(nonce)
		return nil
	}
}

// WithPayload provides the TransactionBuilderOption for
// defining transaction payload bytes
func WithPayload(payload []byte) TransactionBuilderOption {
	return func(t *TransactionBuilder) error {
		t.setPayload(payload)
		return nil
	}
}
