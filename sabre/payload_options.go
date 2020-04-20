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
	"github.com/hyperledger/transact-sdk-go/src/protobuf/sabre_pb2"
)

// SabrePayloadOption provides the functional options used for constructing
// a SabrePayload
type SabrePayloadOption func(*SabrePayloadBuilder) error

// WithAction sets the sabre payload action
// returns error if actionName is invalid
func WithAction(action sabre_pb2.SabrePayload_Action) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setAction(action)
		return nil
	}
}

// WithContractName sets the sabre payload contract name
func WithContractName(name string) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setContractName(name)
		return nil
	}
}

// WithContractVersion sets the sabre payload contract version
func WithContractVersion(version string) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setContractVersion(version)
		return nil
	}
}

// WithContract sets the sabre payload contract bytes
func WithContract(contract []byte) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setContract(contract)
		return nil
	}
}

// WithInputs sets the sabre payload inputs
func WithInputs(inputs []string) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setInputs(inputs)
		return nil
	}
}

// WithOutputs sets the sabre payload outputs
func WithOutputs(outputs []string) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setOutputs(outputs)
		return nil
	}
}

// WithExecuteContractPayload sets the sabre payload execute contract bytes
func WithExecuteContractPayload(payload []byte) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setExecuteContractPayload(payload)
		return nil
	}
}

// WithOwners sets the sabre payload owners
func WithOwners(owners []string) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setOwners(owners)
		return nil
	}
}

// WithNamespace sets the sabre payload namespace name
func WithNamespace(namespace string) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setNamespace(namespace)
		return nil
	}
}

// WithNamespaceReadPermission sets the sabre payload namespace read permission
func WithNamespaceReadPermission(r bool) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setReadNamespacePermission(r)
		return nil
	}
}

// WithNamespaceWritePermission sets the sabre payload namespace write permission
func WithNamespaceWritePermission(w bool) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setWriteNamespacePermssion(w)
		return nil
	}
}

// WithSmartPermissionName sets the sabre payload smart permission name
func WithSmartPermissionName(name string) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setSmartPermssionName(name)
		return nil
	}
}

// WithOrgID sets the sabre payload org id
func WithOrgID(id string) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setOrgID(id)
		return nil
	}
}

// WithSmartPermissionFunction sets the sabre payload smart permission function bytes
func WithSmartPermissionFunction(function []byte) SabrePayloadOption {
	return func(s *SabrePayloadBuilder) error {
		s.setSmartPermssionFunction(function)
		return nil
	}
}
