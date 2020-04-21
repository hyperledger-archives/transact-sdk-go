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

package sabre

import (
	"github.com/golang/protobuf/proto"

	"github.com/hyperledger/transact-sdk-go/crypto"
	"github.com/hyperledger/transact-sdk-go/errors"
	"github.com/hyperledger/transact-sdk-go/src/protobuf/transaction_pb2"
	"github.com/hyperledger/transact-sdk-go/transactions"
	"github.com/hyperledger/transact-sdk-go/transactions/signing"
)

const (
	// SabreFamilyName is the Sawtooth Sabre transaction family name (sabre)
	SabreFamilyName = "sabre"
	// SabreFamilyVersion is the Sawtooth Sabre transaction family version (0.4)
	SabreFamilyVersion = "0.4"
	// SmartPermissionPrefix is the smart permission prefix for global state (00ec03)
	SmartPermissionPrefix string = "00ec03"
	// PikePrefix is the global address for pike
	PikePrefix string = "cad11d"
)

// NewSabreTransactionBuilder  provides the builder for creating a new Sabre Transaction
func NewSabreTransactionBuilder(opts ...SabreTransactionOption) (*SabreTransactionBuilder, error) {
	s := &SabreTransactionBuilder{}
	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// SabreTransactionBuilder implements ITransactionBuilder
// interface to build a Sabre Transaction
type SabreTransactionBuilder struct {
	transactionBuilder transactions.ITransactionBuilder
	payloadBuilder     ISabrePayloadBuilder
}

func (s *SabreTransactionBuilder) setPayloadBuilder(p ISabrePayloadBuilder) {
	s.payloadBuilder = p
}

func (s *SabreTransactionBuilder) setTransactionBuilder(t transactions.ITransactionBuilder) {
	s.transactionBuilder = t
}

// Build creates a Sabre Transaction provided a signer of the transaction.
// Returns an Transaction and an error indicating missing fields
// or proto marshalling errors, if any
func (s *SabreTransactionBuilder) Build(signer *signing.Signer) (*transaction_pb2.Transaction, error) {

	if s.payloadBuilder == nil {
		return nil, errors.NewMissingFieldError("payload builder")
	}

	if s.transactionBuilder == nil {
		return nil, errors.NewMissingFieldError("transaction builder")
	}

	sabrePayload, err := s.payloadBuilder.Build()
	if err != nil {
		return nil, err
	}

	signingKey := s.transactionBuilder.GetBatcherPublicKey()
	if signingKey == nil {
		signingKey = signer.GetPublicKey()
	}

	payloadBytes, err := proto.Marshal(sabrePayload)
	if err != nil {
		return nil, errors.NewProtobufEncodingError(err)
	}

	header := &transaction_pb2.TransactionHeader{
		FamilyName:      SabreFamilyName,
		FamilyVersion:   SabreFamilyVersion,
		Inputs:          s.transactionBuilder.GetInputs(),
		Outputs:         s.transactionBuilder.GetOutputs(),
		SignerPublicKey: signingKey.AsHex(),
		Dependencies:    s.transactionBuilder.GetDependencies(),
		Nonce:           s.transactionBuilder.GetNonce(),
		PayloadSha512:   crypto.NewSha512Hash(payloadBytes),
	}

	headerBytes, err := proto.Marshal(header)
	if err != nil {
		return nil, errors.NewProtobufEncodingError(err)
	}

	headerSignature := signer.Sign(headerBytes)

	return &transaction_pb2.Transaction{
		Header:          headerBytes,
		HeaderSignature: headerSignature,
		Payload:         payloadBytes,
	}, nil
}
