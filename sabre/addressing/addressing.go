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

package addressing

import (
	"github.com/hyperledger/transact-sdk-go/crypto"
)

const (
	// NamespaceRegistryPrefix is the namespace registry prefix for global state
	NamespaceRegistryPrefix = "00ec00"

	// ContractPrefix is the contract prefix for global state
	ContractPrefix = "00ec02"

	// ContractRegistryPrefix is the contract registry prefix for global state
	ContractRegistryPrefix = "00ec01"

	// SmartPermissionPrefix is the smart permission prefix for global state (00ec03)
	SmartPermissionPrefix = "00ec03"

	// PikePrefix is the global address for pike
	PikePrefix = "cad11d"
)

// CalculateNamespaceRegistryAddress calculates the registry address for a namespace.
// It accepts a namespace {string}
// Returns the Contract Address {string}
func CalculateNamespaceRegistryAddress(namespace string) string {
	prefix := namespace[:6]
	hash := crypto.NewSha512Hash([]byte(prefix))[:64]
	return concat(NamespaceRegistryPrefix, hash)
}

// ComputeContractAddress calculates the contract address for a version of a contract.
// Accepts the name {string} and version {string} of a contract
// Returns the Contract Address {string}
func ComputeContractAddress(name, version string) string {
	input := concat(concat(name, ","), version)
	hash := crypto.NewSha512Hash([]byte(input))[:64]
	return concat(ContractPrefix, hash)
}

// ComputeContractRegistryAddress calculates the registry address for a contract.
// Accepts the name {string} of the smart contract
// Returns the Contract Registry Address {string}
func ComputeContractRegistryAddress(name string) string {
	hash := crypto.NewSha512Hash([]byte(name))[:64]
	return concat(ContractRegistryPrefix, hash)
}

// ComputeContractPrefix calculates the state address prefix for a given
// contract name {string}
// Returns the address prefix { string } length of 6
func ComputeContractPrefix(contractName string) string {
	return crypto.NewSha512Hash([]byte(contractName))[:6]
}

// concat uses a preallocated copy method for efficient string concatenation
func concat(s1, s2 string) string {
	s := make([]byte, len(s1)+len(s2))
	bl := 0
	bl += copy(s[bl:], s1)
	bl += copy(s[bl:], s2)
	return string(s)
}
