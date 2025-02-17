# Release History

## 4.1.0 (2023-07-28)
### Features Added

- New value `RegionStorageToNetworkProximityAcrossT2`, `RegionStorageToNetworkProximityT1AndAcrossT2`, `RegionStorageToNetworkProximityT1AndT2AndAcrossT2`, `RegionStorageToNetworkProximityT2AndAcrossT2` added to enum type `RegionStorageToNetworkProximity`
- New value `VolumeStorageToNetworkProximityAcrossT2` added to enum type `VolumeStorageToNetworkProximity`
- New function `*VolumesClient.BeginListGetGroupIDListForLdapUser(context.Context, string, string, string, string, GetGroupIDListForLDAPUserRequest, *VolumesClientBeginListGetGroupIDListForLdapUserOptions) (*runtime.Poller[VolumesClientListGetGroupIDListForLdapUserResponse], error)`
- New struct `GetGroupIDListForLDAPUserRequest`
- New struct `GetGroupIDListForLDAPUserResponse`
- New field `Identity` in struct `AccountPatch`
- New field `SnapshotDirectoryVisible` in struct `VolumePatchProperties`
- New field `ActualThroughputMibps`, `OriginatingResourceID` in struct `VolumeProperties`


## 4.0.0 (2023-03-24)
### Breaking Changes

- Type of `Account.Identity` has been changed from `*Identity` to `*ManagedServiceIdentity`
- Type alias `IdentityType` has been removed
- Function `NewVaultsClient` has been removed
- Function `*VaultsClient.NewListPager` has been removed
- Struct `Identity` has been removed
- Struct `Vault` has been removed
- Struct `VaultList` has been removed
- Struct `VaultProperties` has been removed
- Struct `VaultsClient` has been removed
- Field `VaultID` of struct `VolumeBackupProperties` has been removed

### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module
- New enum type `FileAccessLogs` with values `FileAccessLogsDisabled`, `FileAccessLogsEnabled`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New function `*BackupsClient.BeginRestoreFiles(context.Context, string, string, string, string, string, BackupRestoreFiles, *BackupsClientBeginRestoreFilesOptions) (*runtime.Poller[BackupsClientRestoreFilesResponse], error)`
- New function `*VolumesClient.BeginBreakFileLocks(context.Context, string, string, string, string, *VolumesClientBeginBreakFileLocksOptions) (*runtime.Poller[VolumesClientBreakFileLocksResponse], error)`
- New struct `BackupRestoreFiles`
- New struct `BreakFileLocksRequest`
- New struct `ManagedServiceIdentity`
- New struct `VolumeRelocationProperties`
- New field `PreferredServersForLdapClient` in struct `ActiveDirectory`
- New field `SystemData` in struct `Backup`
- New field `SystemData` in struct `Snapshot`
- New field `DataStoreResourceID` in struct `VolumeProperties`
- New field `FileAccessLogs` in struct `VolumeProperties`
- New field `IsLargeVolume` in struct `VolumeProperties`
- New field `ProvisionedAvailabilityZone` in struct `VolumeProperties`
- New field `VolumeRelocation` in struct `VolumePropertiesDataProtection`
- New field `Tags` in struct `VolumeQuotaRulePatch`


## 3.0.0 (2022-09-16)
### Breaking Changes

- Type of `AccountEncryption.KeySource` has been changed from `*string` to `*KeySource`
- Field `Location` of struct `Vault` has been removed

### Features Added

- New const `KeyVaultStatusCreated`
- New const `IdentityTypeNone`
- New const `IdentityTypeUserAssigned`
- New const `IdentityTypeSystemAssigned`
- New const `RegionStorageToNetworkProximityDefault`
- New const `KeyVaultStatusUpdating`
- New const `KeyVaultStatusError`
- New const `RegionStorageToNetworkProximityT1AndT2`
- New const `KeySourceMicrosoftKeyVault`
- New const `SmbNonBrowsableDisabled`
- New const `SmbNonBrowsableEnabled`
- New const `RegionStorageToNetworkProximityT1`
- New const `SmbAccessBasedEnumerationDisabled`
- New const `SmbAccessBasedEnumerationEnabled`
- New const `KeySourceMicrosoftNetApp`
- New const `RegionStorageToNetworkProximityT2`
- New const `KeyVaultStatusDeleted`
- New const `KeyVaultStatusInUse`
- New const `IdentityTypeSystemAssignedUserAssigned`
- New type alias `SmbAccessBasedEnumeration`
- New type alias `SmbNonBrowsable`
- New type alias `KeySource`
- New type alias `KeyVaultStatus`
- New type alias `IdentityType`
- New type alias `RegionStorageToNetworkProximity`
- New function `PossibleIdentityTypeValues() []IdentityType`
- New function `PossibleKeySourceValues() []KeySource`
- New function `*ResourceClient.QueryRegionInfo(context.Context, string, *ResourceClientQueryRegionInfoOptions) (ResourceClientQueryRegionInfoResponse, error)`
- New function `PossibleSmbAccessBasedEnumerationValues() []SmbAccessBasedEnumeration`
- New function `PossibleSmbNonBrowsableValues() []SmbNonBrowsable`
- New function `*AccountsClient.BeginRenewCredentials(context.Context, string, string, *AccountsClientBeginRenewCredentialsOptions) (*runtime.Poller[AccountsClientRenewCredentialsResponse], error)`
- New function `PossibleRegionStorageToNetworkProximityValues() []RegionStorageToNetworkProximity`
- New function `PossibleKeyVaultStatusValues() []KeyVaultStatus`
- New struct `AccountsClientBeginRenewCredentialsOptions`
- New struct `AccountsClientRenewCredentialsResponse`
- New struct `EncryptionIdentity`
- New struct `Identity`
- New struct `KeyVaultProperties`
- New struct `RegionInfo`
- New struct `RegionInfoAvailabilityZoneMappingsItem`
- New struct `RelocateVolumeRequest`
- New struct `ResourceClientQueryRegionInfoOptions`
- New struct `ResourceClientQueryRegionInfoResponse`
- New struct `UserAssignedIdentity`
- New field `Body` in struct `VolumesClientBeginRelocateOptions`
- New field `SmbAccessBasedEnumeration` in struct `VolumeProperties`
- New field `SmbNonBrowsable` in struct `VolumeProperties`
- New field `DeleteBaseSnapshot` in struct `VolumeProperties`
- New field `Identity` in struct `Account`
- New field `Identity` in struct `AccountEncryption`
- New field `KeyVaultProperties` in struct `AccountEncryption`
- New field `DisableShowmount` in struct `AccountProperties`


## 2.1.0 (2022-07-21)
### Features Added

- New const `EncryptionKeySourceMicrosoftKeyVault`
- New function `*VolumesClient.BeginReestablishReplication(context.Context, string, string, string, string, ReestablishReplicationRequest, *VolumesClientBeginReestablishReplicationOptions) (*runtime.Poller[VolumesClientReestablishReplicationResponse], error)`
- New struct `ReestablishReplicationRequest`
- New struct `VolumesClientBeginReestablishReplicationOptions`
- New struct `VolumesClientReestablishReplicationResponse`
- New field `CoolnessPeriod` in struct `VolumePatchProperties`
- New field `CoolAccess` in struct `VolumePatchProperties`
- New field `KeyVaultPrivateEndpointResourceID` in struct `VolumeProperties`
- New field `CoolAccess` in struct `PoolPatchProperties`


## 2.0.0 (2022-06-22)
### Breaking Changes

- Function `*BackupsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, string, *BackupsClientBeginUpdateOptions)` to `(context.Context, string, string, string, string, string, BackupPatch, *BackupsClientBeginUpdateOptions)`
- Type of `VolumeProperties.EncryptionKeySource` has been changed from `*string` to `*EncryptionKeySource`
- Field `Body` of struct `BackupsClientBeginUpdateOptions` has been removed
- Field `Tags` of struct `VolumeGroupDetails` has been removed
- Field `Tags` of struct `VolumeGroup` has been removed

### Features Added

- New const `ProvisioningStateDeleting`
- New const `TypeIndividualGroupQuota`
- New const `TypeDefaultUserQuota`
- New const `TypeIndividualUserQuota`
- New const `ProvisioningStateSucceeded`
- New const `ProvisioningStateAccepted`
- New const `ProvisioningStateCreating`
- New const `ProvisioningStateMoving`
- New const `ProvisioningStatePatching`
- New const `ProvisioningStateFailed`
- New const `EncryptionKeySourceMicrosoftNetApp`
- New const `TypeDefaultGroupQuota`
- New function `*VolumesClient.BeginRelocate(context.Context, string, string, string, string, *VolumesClientBeginRelocateOptions) (*runtime.Poller[VolumesClientRelocateResponse], error)`
- New function `*VolumeQuotaRulesClient.BeginDelete(context.Context, string, string, string, string, string, *VolumeQuotaRulesClientBeginDeleteOptions) (*runtime.Poller[VolumeQuotaRulesClientDeleteResponse], error)`
- New function `*VolumeQuotaRulesClient.Get(context.Context, string, string, string, string, string, *VolumeQuotaRulesClientGetOptions) (VolumeQuotaRulesClientGetResponse, error)`
- New function `*VolumesClient.NewListReplicationsPager(string, string, string, string, *VolumesClientListReplicationsOptions) *runtime.Pager[VolumesClientListReplicationsResponse]`
- New function `PossibleTypeValues() []Type`
- New function `*VolumesClient.BeginResetCifsPassword(context.Context, string, string, string, string, *VolumesClientBeginResetCifsPasswordOptions) (*runtime.Poller[VolumesClientResetCifsPasswordResponse], error)`
- New function `*VolumeQuotaRulesClient.BeginCreate(context.Context, string, string, string, string, string, VolumeQuotaRule, *VolumeQuotaRulesClientBeginCreateOptions) (*runtime.Poller[VolumeQuotaRulesClientCreateResponse], error)`
- New function `*VolumesClient.BeginFinalizeRelocation(context.Context, string, string, string, string, *VolumesClientBeginFinalizeRelocationOptions) (*runtime.Poller[VolumesClientFinalizeRelocationResponse], error)`
- New function `*VolumeQuotaRulesClient.NewListByVolumePager(string, string, string, string, *VolumeQuotaRulesClientListByVolumeOptions) *runtime.Pager[VolumeQuotaRulesClientListByVolumeResponse]`
- New function `PossibleProvisioningStateValues() []ProvisioningState`
- New function `*VolumesClient.BeginRevertRelocation(context.Context, string, string, string, string, *VolumesClientBeginRevertRelocationOptions) (*runtime.Poller[VolumesClientRevertRelocationResponse], error)`
- New function `*VolumeQuotaRulesClient.BeginUpdate(context.Context, string, string, string, string, string, VolumeQuotaRulePatch, *VolumeQuotaRulesClientBeginUpdateOptions) (*runtime.Poller[VolumeQuotaRulesClientUpdateResponse], error)`
- New function `PossibleEncryptionKeySourceValues() []EncryptionKeySource`
- New function `NewVolumeQuotaRulesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*VolumeQuotaRulesClient, error)`
- New struct `ListReplications`
- New struct `Replication`
- New struct `TrackedResource`
- New struct `VolumeQuotaRule`
- New struct `VolumeQuotaRulePatch`
- New struct `VolumeQuotaRulesClient`
- New struct `VolumeQuotaRulesClientBeginCreateOptions`
- New struct `VolumeQuotaRulesClientBeginDeleteOptions`
- New struct `VolumeQuotaRulesClientBeginUpdateOptions`
- New struct `VolumeQuotaRulesClientCreateResponse`
- New struct `VolumeQuotaRulesClientDeleteResponse`
- New struct `VolumeQuotaRulesClientGetOptions`
- New struct `VolumeQuotaRulesClientGetResponse`
- New struct `VolumeQuotaRulesClientListByVolumeOptions`
- New struct `VolumeQuotaRulesClientListByVolumeResponse`
- New struct `VolumeQuotaRulesClientUpdateResponse`
- New struct `VolumeQuotaRulesList`
- New struct `VolumeQuotaRulesProperties`
- New struct `VolumeRelocationProperties`
- New struct `VolumesClientBeginFinalizeRelocationOptions`
- New struct `VolumesClientBeginRelocateOptions`
- New struct `VolumesClientBeginResetCifsPasswordOptions`
- New struct `VolumesClientBeginRevertRelocationOptions`
- New struct `VolumesClientFinalizeRelocationResponse`
- New struct `VolumesClientListReplicationsOptions`
- New struct `VolumesClientListReplicationsResponse`
- New struct `VolumesClientRelocateResponse`
- New struct `VolumesClientResetCifsPasswordResponse`
- New struct `VolumesClientRevertRelocationResponse`
- New field `SystemData` in struct `ProxyResource`
- New field `Zones` in struct `Volume`
- New field `SystemData` in struct `Resource`
- New field `Encrypted` in struct `VolumeProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
