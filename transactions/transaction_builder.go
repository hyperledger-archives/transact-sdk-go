// Copyright 2020 txross
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

	"github.com/hyperledger/transact-sdk-go/crypto"
	"github.com/hyperledger/transact-sdk-go/errors"
	"github.com/hyperledger/transact-sdk-go/src/protobuf/transaction_pb2"
	"github.com/hyperledger/transact-sdk-go/transactions/signing"
)

// ITransactionBuilder provides the interface for building transactions
type ITransactionBuilder interface {
	GetBatcherPublicKey() signing.PublicKey
	GetDependencies() []string
	GetFamilyName() string
	GetFamilyVersion() string
	GetInputs() []string
	GetOutputs() []string
	GetNonce() string
	GetPayload() []byte
	Build(*signing.Signer) (*transaction_pb2.Transaction, error)
}

// NewTransactionBuilder creates a TransactionBuilder from provided TransactionBuilderOptions
func NewTransactionBuilder(opts ...TransactionBuilderOption) (ITransactionBuilder, error) {
	t := &TransactionBuilder{}
	for _, opt := range opts {
		if err := opt(t); err != nil {
			return nil, err
		}
	}
	return t, nil
}

// TransactionBuilder is the struct for creating a new Transaction
// and implements the ITransactionBuilder interface
type TransactionBuilder struct {
	batcherPublicKey signing.PublicKey
	dependencies     []string
	familyName       string
	familyVersion    string
	inputs           []string
	outputs          []string
	nonce            string
	payload          []byte
}

// GetBatcherPublicKey returns the transaction builder's batcher public key
func (t *TransactionBuilder) GetBatcherPublicKey() signing.PublicKey { return t.batcherPublicKey }

// GetDependencies returns the transaction builder's dependencies as a slice of string
func (t *TransactionBuilder) GetDependencies() []string { return t.dependencies }

// GetFamilyName returns the transaction builder's family name string
func (t *TransactionBuilder) GetFamilyName() string { return t.familyName }

// GetFamilyVersion returns the transaction builder's family version string
func (t *TransactionBuilder) GetFamilyVersion() string { return t.familyVersion }

// GetInputs returns the transaction builder's input addresses as a slice of string
func (t *TransactionBuilder) GetInputs() []string { return t.inputs }

// GetOutputs returns the transaction builder's output addresses as a slice of string
func (t *TransactionBuilder) GetOutputs() []string { return t.outputs }

// GetNonce returns the transaction builder's nonce string value
func (t *TransactionBuilder) GetNonce() string { return t.nonce }

// GetPayload returns the transaction builder's payload bytes
func (t *TransactionBuilder) GetPayload() []byte { return t.payload }

func (t *TransactionBuilder) buildTransactionHeader(signer *signing.Signer) ([]byte, error) {
	if t.familyName == "" {
		return nil, errors.NewMissingFieldError("family name")
	}
	if t.familyVersion == "" {
		return nil, errors.NewMissingFieldError("family version")
	}
	if len(t.inputs) == 0 {
		return nil, errors.NewMissingFieldError("inputs")
	}
	if len(t.outputs) == 0 {
		return nil, errors.NewMissingFieldError("outputs")
	}
	if len(t.payload) == 0 {
		return nil, errors.NewMissingFieldError("payload")
	}

	signingKey := t.batcherPublicKey
	if signingKey == nil {
		signingKey = signer.GetPublicKey()
	}

	header := &transaction_pb2.TransactionHeader{
		FamilyName:      t.familyName,
		FamilyVersion:   t.familyVersion,
		Inputs:          t.inputs,
		Outputs:         t.outputs,
		SignerPublicKey: signingKey.AsHex(),
		Dependencies:    t.dependencies,
		Nonce:           t.nonce,
		PayloadSha512:   crypto.NewSha512Hash(t.payload),
	}

	headerBytes, err := proto.Marshal(header)
	if err != nil {
		return nil, errors.NewProtobufEncodingError(err)
	}

	return headerBytes, nil
}

// Build creates a Transaction provided signer of the transaction.
// Returns an Transaction and an error indicating missing fields
// or proto marshalling errors, if any
func (t *TransactionBuilder) Build(signer *signing.Signer) (*transaction_pb2.Transaction, error) {
	header, err := t.buildTransactionHeader(signer)
	if err != nil {
		return nil, err
	}

	headerSignature := signer.Sign(header)

	return &transaction_pb2.Transaction{
		Header:          header,
		HeaderSignature: headerSignature,
		Payload:         t.payload,
	}, nil
}

func (t *TransactionBuilder) setBatcherPublicKey(key signing.PublicKey) { t.batcherPublicKey = key }
func (t *TransactionBuilder) setDependencies(deps []string)             { t.dependencies = deps }
func (t *TransactionBuilder) setFamilyName(name string)                 { t.familyName = name }
func (t *TransactionBuilder) setFamilyVersion(version string)           { t.familyVersion = version }
func (t *TransactionBuilder) setInputs(inputs []string)                 { t.inputs = inputs }
func (t *TransactionBuilder) setOutputs(outputs []string)               { t.outputs = outputs }
func (t *TransactionBuilder) setNonce(nonce string)                     { t.nonce = nonce }
func (t *TransactionBuilder) setPayload(payload []byte)                 { t.payload = payload }
