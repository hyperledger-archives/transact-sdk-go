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

package sabre

import (
	"github.com/hyperledger/transact-sdk-go/transactions"
)

// SabreTransactionOption provides functional options
// for customizing a SabreTransactionBuilder
type SabreTransactionOption func(*SabreTransactionBuilder) error

// WithPayloadBuilder provides the option for setting the ISabrePayloadBuilder for a Sabre Transaction
func WithPayloadBuilder(p ISabrePayloadBuilder) SabreTransactionOption {
	return func(s *SabreTransactionBuilder) error {
		_, err := p.Build()
		if err != nil {
			return err
		}
		s.setPayloadBuilder(p)
		return nil
	}
}

// WithTransactionBuilder provides the option for setting the ITransactionBuilder for a Sabre Transaction
func WithTransactionBuilder(t transactions.ITransactionBuilder) SabreTransactionOption {
	return func(s *SabreTransactionBuilder) error {
		s.setTransactionBuilder(t)
		return nil
	}
}
