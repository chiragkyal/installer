package relay

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// Listen ...
	Listen AccessRights = "Listen"
	// Manage ...
	Manage AccessRights = "Manage"
	// SendEnumValue ...
	SendEnumValue AccessRights = "Send"
)

// PossibleAccessRightsValues returns an array of possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{Listen, Manage, SendEnumValue}
}

// KeyType enumerates the values for key type.
type KeyType string

const (
	// PrimaryKey ...
	PrimaryKey KeyType = "PrimaryKey"
	// SecondaryKey ...
	SecondaryKey KeyType = "SecondaryKey"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{PrimaryKey, SecondaryKey}
}

// ProvisioningStateEnum enumerates the values for provisioning state enum.
type ProvisioningStateEnum string

const (
	// Created ...
	Created ProvisioningStateEnum = "Created"
	// Deleted ...
	Deleted ProvisioningStateEnum = "Deleted"
	// Failed ...
	Failed ProvisioningStateEnum = "Failed"
	// Succeeded ...
	Succeeded ProvisioningStateEnum = "Succeeded"
	// Unknown ...
	Unknown ProvisioningStateEnum = "Unknown"
	// Updating ...
	Updating ProvisioningStateEnum = "Updating"
)

// PossibleProvisioningStateEnumValues returns an array of possible values for the ProvisioningStateEnum const type.
func PossibleProvisioningStateEnumValues() []ProvisioningStateEnum {
	return []ProvisioningStateEnum{Created, Deleted, Failed, Succeeded, Unknown, Updating}
}

// RelaytypeEnum enumerates the values for relaytype enum.
type RelaytypeEnum string

const (
	// HTTP ...
	HTTP RelaytypeEnum = "Http"
	// NetTCP ...
	NetTCP RelaytypeEnum = "NetTcp"
)

// PossibleRelaytypeEnumValues returns an array of possible values for the RelaytypeEnum const type.
func PossibleRelaytypeEnumValues() []RelaytypeEnum {
	return []RelaytypeEnum{HTTP, NetTCP}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Standard ...
	Standard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{Standard}
}

// UnavailableReason enumerates the values for unavailable reason.
type UnavailableReason string

const (
	// InvalidName ...
	InvalidName UnavailableReason = "InvalidName"
	// NameInLockdown ...
	NameInLockdown UnavailableReason = "NameInLockdown"
	// NameInUse ...
	NameInUse UnavailableReason = "NameInUse"
	// None ...
	None UnavailableReason = "None"
	// SubscriptionIsDisabled ...
	SubscriptionIsDisabled UnavailableReason = "SubscriptionIsDisabled"
	// TooManyNamespaceInCurrentSubscription ...
	TooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// PossibleUnavailableReasonValues returns an array of possible values for the UnavailableReason const type.
func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{InvalidName, NameInLockdown, NameInUse, None, SubscriptionIsDisabled, TooManyNamespaceInCurrentSubscription}
}