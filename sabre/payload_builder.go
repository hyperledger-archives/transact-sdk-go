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
	"github.com/hyperledger/transact-sdk-go/errors"
	"github.com/hyperledger/transact-sdk-go/src/protobuf/sabre_pb2"
)

// ISabrePayloadBuilder provides the builder interface for building new sabre payloads
type ISabrePayloadBuilder interface {
	GetAction() sabre_pb2.SabrePayload_Action
	GetContractName() string
	GetContractVersion() string
	GetContract() []byte
	GetExecuteContractPaylod() []byte
	GetOwners() []string
	GetNamespace() string
	GetNamespaceReadPermssion() bool
	GetNamespaceWritePermssion() bool
	GetSmartPermissionName() string
	GetOrgID() string
	GetSmartPermssionFunction() []byte
	Build() (*sabre_pb2.SabrePayload, error)
}

// NewSabrePayloadBuilder provides the builder for creating a new Sabre Payload
func NewSabrePayloadBuilder(opts ...SabrePayloadOption) (ISabrePayloadBuilder, error) {
	p := &SabrePayloadBuilder{}
	for _, opt := range opts {
		err := opt(p)
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}

// SabrePayloadBuilder abstracts from sabre_pb2.SabrePayload to implement
// the builder pattern for sabre transaction payloads
type SabrePayloadBuilder struct {
	action sabre_pb2.SabrePayload_Action
	// for contract and namespace registry permission actions
	contractName           string
	contractVersion        string
	contract               []byte
	inputs                 []string
	outputs                []string
	executeContractPayload []byte
	// for contract registry actions
	owners []string
	// for namespace, namespace registry, namespace registry permission actions
	namespace string
	read      bool
	write     bool
	// for smart permission actions
	smartPermissionName     string
	orgID                   string
	smartPermissionFunction []byte
}

func (s *SabrePayloadBuilder) GetAction() sabre_pb2.SabrePayload_Action { return s.action }
func (s *SabrePayloadBuilder) GetContractName() string                  { return s.contractName }
func (s *SabrePayloadBuilder) GetContractVersion() string               { return s.contractVersion }
func (s *SabrePayloadBuilder) GetContract() []byte                      { return s.contract }
func (s *SabrePayloadBuilder) GetExecuteContractPaylod() []byte         { return s.executeContractPayload }
func (s *SabrePayloadBuilder) GetOwners() []string                      { return s.owners }
func (s *SabrePayloadBuilder) GetNamespace() string                     { return s.namespace }
func (s *SabrePayloadBuilder) GetNamespaceReadPermssion() bool          { return s.read }
func (s *SabrePayloadBuilder) GetNamespaceWritePermssion() bool         { return s.write }
func (s *SabrePayloadBuilder) GetSmartPermissionName() string           { return s.smartPermissionName }
func (s *SabrePayloadBuilder) GetOrgID() string                         { return s.orgID }
func (s *SabrePayloadBuilder) GetSmartPermssionFunction() []byte        { return s.smartPermissionFunction }

func (s *SabrePayloadBuilder) setAction(action sabre_pb2.SabrePayload_Action) {
	s.action = action
}

func (s *SabrePayloadBuilder) setContractName(name string) {
	s.contractName = name
}

func (s *SabrePayloadBuilder) setContractVersion(version string) {
	s.contractVersion = version
}

func (s *SabrePayloadBuilder) setContract(contract []byte) {
	s.contract = contract
}

func (s *SabrePayloadBuilder) setInputs(inputs []string) {
	s.inputs = inputs
}

func (s *SabrePayloadBuilder) setOutputs(outputs []string) {
	s.outputs = outputs
}

func (s *SabrePayloadBuilder) setExecuteContractPayload(payload []byte) {
	s.executeContractPayload = payload
}

func (s *SabrePayloadBuilder) setOwners(owners []string) {
	s.owners = owners
}

func (s *SabrePayloadBuilder) setNamespace(namespace string) {
	s.namespace = namespace
}

func (s *SabrePayloadBuilder) setReadNamespacePermission(read bool) {
	s.read = read
}

func (s *SabrePayloadBuilder) setWriteNamespacePermssion(write bool) {
	s.write = write
}

func (s *SabrePayloadBuilder) setSmartPermssionName(smartPermName string) {
	s.smartPermissionName = smartPermName
}

func (s *SabrePayloadBuilder) setOrgID(orgID string) {
	s.orgID = orgID
}

func (s *SabrePayloadBuilder) setSmartPermssionFunction(function []byte) {
	s.smartPermissionFunction = function
}

func (s *SabrePayloadBuilder) Build() (*sabre_pb2.SabrePayload, error) {
	if s.action == sabre_pb2.SabrePayload_ACTION_UNSET {
		return nil, errors.NewMissingFieldError("action")
	}

	var sabrePayload = &sabre_pb2.SabrePayload{}
	var err error

	switch s.action {
	case sabre_pb2.SabrePayload_CREATE_CONTRACT:
		sabrePayload, err = s.buildCreateContractAction()
	case sabre_pb2.SabrePayload_DELETE_CONTRACT:
		sabrePayload, err = s.buildDeleteContractAction()
	case sabre_pb2.SabrePayload_EXECUTE_CONTRACT:
		sabrePayload, err = s.buildExecuteContractAction()
	case sabre_pb2.SabrePayload_CREATE_CONTRACT_REGISTRY:
		sabrePayload, err = s.buildCreateContractRegistryAction()
	case sabre_pb2.SabrePayload_DELETE_CONTRACT_REGISTRY:
		sabrePayload, err = s.buildDeleteContractRegistryAction()
	case sabre_pb2.SabrePayload_UPDATE_CONTRACT_REGISTRY_OWNERS:
		sabrePayload, err = s.buildUpdateContractRegistryOwnersAction()
	case sabre_pb2.SabrePayload_CREATE_NAMESPACE_REGISTRY:
		sabrePayload, err = s.buildCreateNamespaceRegistryAction()
	case sabre_pb2.SabrePayload_DELETE_NAMESPACE_REGISTRY:
		sabrePayload, err = s.buildDeleteNamespaceRegistryAction()
	case sabre_pb2.SabrePayload_UPDATE_NAMESPACE_REGISTRY_OWNERS:
		sabrePayload, err = s.buildUpdateNamespaceRegistryOwnersAction()
	case sabre_pb2.SabrePayload_CREATE_NAMESPACE_REGISTRY_PERMISSION:
		sabrePayload, err = s.buildCreateNamespaceRegistryPermissionAction()
	case sabre_pb2.SabrePayload_DELETE_NAMESPACE_REGISTRY_PERMISSION:
		sabrePayload, err = s.buildDeleteNamespaceRegistryPermissionAction()
	case sabre_pb2.SabrePayload_CREATE_SMART_PERMISSION:
		sabrePayload, err = s.buildCreateSmartPermissionAction()
	case sabre_pb2.SabrePayload_UPDATE_SMART_PERMISSION:
		sabrePayload, err = s.buildUpdateSmartPermissionAction()
	case sabre_pb2.SabrePayload_DELETE_SMART_PERMISSION:
		sabrePayload, err = s.buildDeleteSmartPermissionAction()
	}

	return sabrePayload, err
}

func (s *SabrePayloadBuilder) buildCreateContractAction() (*sabre_pb2.SabrePayload, error) {
	if s.contractName == "" {
		return nil, errors.NewMissingFieldError("contract name")
	}
	if s.contractVersion == "" {
		return nil, errors.NewMissingFieldError("contract version")
	}
	if len(s.contract) == 0 {
		return nil, errors.NewMissingFieldError("contract")
	}

	return &sabre_pb2.SabrePayload{
		Action: s.action,
		CreateContract: &sabre_pb2.CreateContractAction{
			Name:     s.contractName,
			Version:  s.contractVersion,
			Inputs:   s.inputs,
			Outputs:  s.outputs,
			Contract: s.contract,
		},
	}, nil
}

func (s *SabrePayloadBuilder) buildDeleteContractAction() (*sabre_pb2.SabrePayload, error) {
	if s.contractName == "" {
		return nil, errors.NewMissingFieldError("contract name")
	}
	if s.contractVersion == "" {
		return nil, errors.NewMissingFieldError("contract version")
	}

	return &sabre_pb2.SabrePayload{
		Action: s.action,
		DeleteContract: &sabre_pb2.DeleteContractAction{
			Name:    s.contractName,
			Version: s.contractVersion,
		},
	}, nil
}

func (s *SabrePayloadBuilder) buildExecuteContractAction() (*sabre_pb2.SabrePayload, error) {
	if s.contractName == "" {
		return nil, errors.NewMissingFieldError("contract name")
	}
	if s.contractVersion == "" {
		return nil, errors.NewMissingFieldError("contract version")
	}
	if len(s.executeContractPayload) == 0 {
		return nil, errors.NewMissingFieldError("execute contract payload")
	}

	return &sabre_pb2.SabrePayload{
		Action: s.action,
		ExecuteContract: &sabre_pb2.ExecuteContractAction{
			Name:    s.contractName,
			Version: s.contractVersion,
			Inputs:  s.inputs,
			Outputs: s.outputs,
			Payload: s.executeContractPayload,
		},
	}, nil
}

func (s *SabrePayloadBuilder) buildCreateContractRegistryAction() (*sabre_pb2.SabrePayload, error) {
	if s.contractName == "" {
		return nil, errors.NewMissingFieldError("contract name")
	}
	if len(s.owners) == 0 {
		return nil, errors.NewMissingFieldError("owners")
	}
	return &sabre_pb2.SabrePayload{
		Action: s.action,
		CreateContractRegistry: &sabre_pb2.CreateContractRegistryAction{
			Name:   s.contractName,
			Owners: s.owners,
		},
	}, nil
}

func (s *SabrePayloadBuilder) buildDeleteContractRegistryAction() (*sabre_pb2.SabrePayload, error) {
	if s.contractName == "" {
		return nil, errors.NewMissingFieldError("contract name")
	}
	return &sabre_pb2.SabrePayload{
		Action: s.action,
		DeleteContractRegistry: &sabre_pb2.DeleteContractRegistryAction{
			Name: s.contractName,
		},
	}, nil
}

func (s *SabrePayloadBuilder) buildUpdateContractRegistryOwnersAction() (*sabre_pb2.SabrePayload, error) {
	if s.contractName == "" {
		return nil, errors.NewMissingFieldError("contract name")
	}
	if len(s.owners) == 0 {
		return nil, errors.NewMissingFieldError("owners")
	}
	return &sabre_pb2.SabrePayload{
		Action: s.action,
		UpdateContractRegistryOwners: &sabre_pb2.UpdateContractRegistryOwnersAction{
			Name:   s.contractName,
			Owners: s.owners,
		},
	}, nil
}

func (s *SabrePayloadBuilder) buildCreateNamespaceRegistryAction() (*sabre_pb2.SabrePayload, error) {
	if s.namespace == "" {
		return nil, errors.NewMissingFieldError("namespace")
	}
	if len(s.owners) == 0 {
		return nil, errors.NewMissingFieldError("owners")
	}
	return &sabre_pb2.SabrePayload{
		Action: s.action,
		CreateNamespaceRegistry: &sabre_pb2.CreateNamespaceRegistryAction{
			Namespace: s.namespace,
			Owners:    s.owners,
		},
	}, nil
}

func (s *SabrePayloadBuilder) buildDeleteNamespaceRegistryAction() (*sabre_pb2.SabrePayload, error) {
	if s.namespace == "" {
		return nil, errors.NewMissingFieldError("namespace")
	}

	return &sabre_pb2.SabrePayload{
		Action: s.action,
		DeleteNamespaceRegistry: &sabre_pb2.DeleteNamespaceRegistryAction{
			Namespace: s.namespace,
		},
	}, nil

}

func (s *SabrePayloadBuilder) buildUpdateNamespaceRegistryOwnersAction() (*sabre_pb2.SabrePayload, error) {
	if s.namespace == "" {
		return nil, errors.NewMissingFieldError("namespace")
	}
	if len(s.owners) == 0 {
		return nil, errors.NewMissingFieldError("owners")
	}

	return &sabre_pb2.SabrePayload{
		Action: s.action,
		UpdateNamespaceRegistryOwners: &sabre_pb2.UpdateNamespaceRegistryOwnersAction{
			Namespace: s.namespace,
			Owners:    s.owners,
		},
	}, nil
}

func (s *SabrePayloadBuilder) buildCreateNamespaceRegistryPermissionAction() (*sabre_pb2.SabrePayload, error) {
	if s.namespace == "" {
		return nil, errors.NewMissingFieldError("namespace")
	}
	if s.contractName == "" {
		return nil, errors.NewMissingFieldError("contract name")
	}

	return &sabre_pb2.SabrePayload{
		Action: s.action,
		CreateNamespaceRegistryPermission: &sabre_pb2.CreateNamespaceRegistryPermissionAction{
			Namespace:    s.namespace,
			ContractName: s.contractName,
			Read:         s.read,
			Write:        s.write,
		},
	}, nil
}

func (s *SabrePayloadBuilder) buildDeleteNamespaceRegistryPermissionAction() (*sabre_pb2.SabrePayload, error) {
	if s.namespace == "" {
		return nil, errors.NewMissingFieldError("namespace")
	}
	if s.contractName == "" {
		return nil, errors.NewMissingFieldError("contract name")
	}

	return &sabre_pb2.SabrePayload{
		Action: s.action,
		DeleteNamespaceRegistryPermission: &sabre_pb2.DeleteNamespaceRegistryPermissionAction{
			Namespace:    s.namespace,
			ContractName: s.contractName,
		},
	}, nil
}

func (s *SabrePayloadBuilder) buildCreateSmartPermissionAction() (*sabre_pb2.SabrePayload, error) {
	if s.smartPermissionName == "" {
		return nil, errors.NewMissingFieldError("smart permission name")
	}
	if s.orgID == "" {
		return nil, errors.NewMissingFieldError("org id")
	}
	if len(s.smartPermissionFunction) == 0 {
		return nil, errors.NewMissingFieldError("function")
	}

	return &sabre_pb2.SabrePayload{
		Action: s.action,
		CreateSmartPermission: &sabre_pb2.CreateSmartPermissionAction{
			Name:     s.smartPermissionName,
			OrgId:    s.orgID,
			Function: s.smartPermissionFunction,
		},
	}, nil
}

func (s *SabrePayloadBuilder) buildUpdateSmartPermissionAction() (*sabre_pb2.SabrePayload, error) {
	if s.smartPermissionName == "" {
		return nil, errors.NewMissingFieldError("smart permission name")
	}
	if s.orgID == "" {
		return nil, errors.NewMissingFieldError("org id")
	}
	if len(s.smartPermissionFunction) == 0 {
		return nil, errors.NewMissingFieldError("function")
	}

	return &sabre_pb2.SabrePayload{
		Action: s.action,
		UpdateSmartPermission: &sabre_pb2.UpdateSmartPermissionAction{
			Name:     s.smartPermissionName,
			OrgId:    s.orgID,
			Function: s.smartPermissionFunction,
		},
	}, nil
}

func (s *SabrePayloadBuilder) buildDeleteSmartPermissionAction() (*sabre_pb2.SabrePayload, error) {
	if s.smartPermissionName == "" {
		return nil, errors.NewMissingFieldError("smart permission name")
	}
	if s.orgID == "" {
		return nil, errors.NewMissingFieldError("org id")
	}

	return &sabre_pb2.SabrePayload{
		Action: s.action,
		DeleteSmartPermission: &sabre_pb2.DeleteSmartPermissionAction{
			Name:  s.smartPermissionName,
			OrgId: s.orgID,
		},
	}, nil
}
