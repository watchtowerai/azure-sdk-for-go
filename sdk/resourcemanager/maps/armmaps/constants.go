//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaps

const (
	moduleName    = "armmaps"
	moduleVersion = "v1.0.0"
)

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// IdentityType - Values can be systemAssignedIdentity or userAssignedIdentity
type IdentityType string

const (
	IdentityTypeDelegatedResourceIdentity IdentityType = "delegatedResourceIdentity"
	IdentityTypeSystemAssignedIdentity    IdentityType = "systemAssignedIdentity"
	IdentityTypeUserAssignedIdentity      IdentityType = "userAssignedIdentity"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeDelegatedResourceIdentity,
		IdentityTypeSystemAssignedIdentity,
		IdentityTypeUserAssignedIdentity,
	}
}

// InfrastructureEncryption - Values are enabled and disabled.
type InfrastructureEncryption string

const (
	InfrastructureEncryptionDisabled InfrastructureEncryption = "disabled"
	InfrastructureEncryptionEnabled  InfrastructureEncryption = "enabled"
)

// PossibleInfrastructureEncryptionValues returns the possible values for the InfrastructureEncryption const type.
func PossibleInfrastructureEncryptionValues() []InfrastructureEncryption {
	return []InfrastructureEncryption{
		InfrastructureEncryptionDisabled,
		InfrastructureEncryptionEnabled,
	}
}

// KeyType - Whether the operation refers to the primary or secondary key.
type KeyType string

const (
	KeyTypePrimary   KeyType = "primary"
	KeyTypeSecondary KeyType = "secondary"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimary,
		KeyTypeSecondary,
	}
}

// Kind - The Kind of the Maps Account.
type Kind string

const (
	KindGen1 Kind = "Gen1"
	KindGen2 Kind = "Gen2"
)

// PossibleKindValues returns the possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{
		KindGen1,
		KindGen2,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned, UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// Name - The name of the SKU, in standard format (such as S0).
type Name string

const (
	NameG2 Name = "G2"
	NameS0 Name = "S0"
	NameS1 Name = "S1"
)

// PossibleNameValues returns the possible values for the Name const type.
func PossibleNameValues() []Name {
	return []Name{
		NameG2,
		NameS0,
		NameS1,
	}
}

// SigningKey - The Map account key to use for signing. Picking primaryKey or secondaryKey will use the Map account Shared
// Keys, and using managedIdentity will use the auto-renewed private key to sign the SAS.
type SigningKey string

const (
	SigningKeyManagedIdentity SigningKey = "managedIdentity"
	SigningKeyPrimaryKey      SigningKey = "primaryKey"
	SigningKeySecondaryKey    SigningKey = "secondaryKey"
)

// PossibleSigningKeyValues returns the possible values for the SigningKey const type.
func PossibleSigningKeyValues() []SigningKey {
	return []SigningKey{
		SigningKeyManagedIdentity,
		SigningKeyPrimaryKey,
		SigningKeySecondaryKey,
	}
}
