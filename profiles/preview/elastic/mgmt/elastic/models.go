//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package elastic

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/elastic/mgmt/2020-07-01/elastic"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type DeploymentStatus = original.DeploymentStatus

const (
	DeploymentStatusHealthy   DeploymentStatus = original.DeploymentStatusHealthy
	DeploymentStatusUnhealthy DeploymentStatus = original.DeploymentStatusUnhealthy
)

type LiftrResourceCategories = original.LiftrResourceCategories

const (
	LiftrResourceCategoriesMonitorLogs LiftrResourceCategories = original.LiftrResourceCategoriesMonitorLogs
	LiftrResourceCategoriesUnknown     LiftrResourceCategories = original.LiftrResourceCategoriesUnknown
)

type ManagedIdentityTypes = original.ManagedIdentityTypes

const (
	ManagedIdentityTypesSystemAssigned ManagedIdentityTypes = original.ManagedIdentityTypesSystemAssigned
)

type MonitoringStatus = original.MonitoringStatus

const (
	MonitoringStatusDisabled MonitoringStatus = original.MonitoringStatusDisabled
	MonitoringStatusEnabled  MonitoringStatus = original.MonitoringStatusEnabled
)

type OperationName = original.OperationName

const (
	OperationNameAdd    OperationName = original.OperationNameAdd
	OperationNameDelete OperationName = original.OperationNameDelete
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted     ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled     ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreating     ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleted      ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting     ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed       ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateNotSpecified ProvisioningState = original.ProvisioningStateNotSpecified
	ProvisioningStateSucceeded    ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating     ProvisioningState = original.ProvisioningStateUpdating
)

type SendingLogs = original.SendingLogs

const (
	SendingLogsFalse SendingLogs = original.SendingLogsFalse
	SendingLogsTrue  SendingLogs = original.SendingLogsTrue
)

type TagAction = original.TagAction

const (
	TagActionExclude TagAction = original.TagActionExclude
	TagActionInclude TagAction = original.TagActionInclude
)

type BaseClient = original.BaseClient
type CloudDeployment = original.CloudDeployment
type CloudUser = original.CloudUser
type CompanyInfo = original.CompanyInfo
type DeploymentInfoClient = original.DeploymentInfoClient
type DeploymentInfoResponse = original.DeploymentInfoResponse
type ErrorResponseBody = original.ErrorResponseBody
type FilteringTag = original.FilteringTag
type IdentityProperties = original.IdentityProperties
type LogRules = original.LogRules
type MonitorProperties = original.MonitorProperties
type MonitorResource = original.MonitorResource
type MonitorResourceListResponse = original.MonitorResourceListResponse
type MonitorResourceListResponseIterator = original.MonitorResourceListResponseIterator
type MonitorResourceListResponsePage = original.MonitorResourceListResponsePage
type MonitorResourceUpdateParameters = original.MonitorResourceUpdateParameters
type MonitoredResource = original.MonitoredResource
type MonitoredResourceListResponse = original.MonitoredResourceListResponse
type MonitoredResourceListResponseIterator = original.MonitoredResourceListResponseIterator
type MonitoredResourceListResponsePage = original.MonitoredResourceListResponsePage
type MonitoredResourcesClient = original.MonitoredResourcesClient
type MonitoringTagRules = original.MonitoringTagRules
type MonitoringTagRulesListResponse = original.MonitoringTagRulesListResponse
type MonitoringTagRulesListResponseIterator = original.MonitoringTagRulesListResponseIterator
type MonitoringTagRulesListResponsePage = original.MonitoringTagRulesListResponsePage
type MonitoringTagRulesProperties = original.MonitoringTagRulesProperties
type MonitorsClient = original.MonitorsClient
type MonitorsCreateFuture = original.MonitorsCreateFuture
type MonitorsDeleteFuture = original.MonitorsDeleteFuture
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResult = original.OperationResult
type OperationsClient = original.OperationsClient
type Properties = original.Properties
type ResourceProviderDefaultErrorResponse = original.ResourceProviderDefaultErrorResponse
type ResourceSku = original.ResourceSku
type SystemData = original.SystemData
type TagRulesClient = original.TagRulesClient
type TagRulesDeleteFuture = original.TagRulesDeleteFuture
type UserInfo = original.UserInfo
type VMCollectionClient = original.VMCollectionClient
type VMCollectionUpdate = original.VMCollectionUpdate
type VMHostClient = original.VMHostClient
type VMHostListResponse = original.VMHostListResponse
type VMHostListResponseIterator = original.VMHostListResponseIterator
type VMHostListResponsePage = original.VMHostListResponsePage
type VMIngestionClient = original.VMIngestionClient
type VMIngestionDetailsResponse = original.VMIngestionDetailsResponse
type VMResources = original.VMResources

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDeploymentInfoClient(subscriptionID string) DeploymentInfoClient {
	return original.NewDeploymentInfoClient(subscriptionID)
}
func NewDeploymentInfoClientWithBaseURI(baseURI string, subscriptionID string) DeploymentInfoClient {
	return original.NewDeploymentInfoClientWithBaseURI(baseURI, subscriptionID)
}
func NewMonitorResourceListResponseIterator(page MonitorResourceListResponsePage) MonitorResourceListResponseIterator {
	return original.NewMonitorResourceListResponseIterator(page)
}
func NewMonitorResourceListResponsePage(cur MonitorResourceListResponse, getNextPage func(context.Context, MonitorResourceListResponse) (MonitorResourceListResponse, error)) MonitorResourceListResponsePage {
	return original.NewMonitorResourceListResponsePage(cur, getNextPage)
}
func NewMonitoredResourceListResponseIterator(page MonitoredResourceListResponsePage) MonitoredResourceListResponseIterator {
	return original.NewMonitoredResourceListResponseIterator(page)
}
func NewMonitoredResourceListResponsePage(cur MonitoredResourceListResponse, getNextPage func(context.Context, MonitoredResourceListResponse) (MonitoredResourceListResponse, error)) MonitoredResourceListResponsePage {
	return original.NewMonitoredResourceListResponsePage(cur, getNextPage)
}
func NewMonitoredResourcesClient(subscriptionID string) MonitoredResourcesClient {
	return original.NewMonitoredResourcesClient(subscriptionID)
}
func NewMonitoredResourcesClientWithBaseURI(baseURI string, subscriptionID string) MonitoredResourcesClient {
	return original.NewMonitoredResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewMonitoringTagRulesListResponseIterator(page MonitoringTagRulesListResponsePage) MonitoringTagRulesListResponseIterator {
	return original.NewMonitoringTagRulesListResponseIterator(page)
}
func NewMonitoringTagRulesListResponsePage(cur MonitoringTagRulesListResponse, getNextPage func(context.Context, MonitoringTagRulesListResponse) (MonitoringTagRulesListResponse, error)) MonitoringTagRulesListResponsePage {
	return original.NewMonitoringTagRulesListResponsePage(cur, getNextPage)
}
func NewMonitorsClient(subscriptionID string) MonitorsClient {
	return original.NewMonitorsClient(subscriptionID)
}
func NewMonitorsClientWithBaseURI(baseURI string, subscriptionID string) MonitorsClient {
	return original.NewMonitorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagRulesClient(subscriptionID string) TagRulesClient {
	return original.NewTagRulesClient(subscriptionID)
}
func NewTagRulesClientWithBaseURI(baseURI string, subscriptionID string) TagRulesClient {
	return original.NewTagRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVMCollectionClient(subscriptionID string) VMCollectionClient {
	return original.NewVMCollectionClient(subscriptionID)
}
func NewVMCollectionClientWithBaseURI(baseURI string, subscriptionID string) VMCollectionClient {
	return original.NewVMCollectionClientWithBaseURI(baseURI, subscriptionID)
}
func NewVMHostClient(subscriptionID string) VMHostClient {
	return original.NewVMHostClient(subscriptionID)
}
func NewVMHostClientWithBaseURI(baseURI string, subscriptionID string) VMHostClient {
	return original.NewVMHostClientWithBaseURI(baseURI, subscriptionID)
}
func NewVMHostListResponseIterator(page VMHostListResponsePage) VMHostListResponseIterator {
	return original.NewVMHostListResponseIterator(page)
}
func NewVMHostListResponsePage(cur VMHostListResponse, getNextPage func(context.Context, VMHostListResponse) (VMHostListResponse, error)) VMHostListResponsePage {
	return original.NewVMHostListResponsePage(cur, getNextPage)
}
func NewVMIngestionClient(subscriptionID string) VMIngestionClient {
	return original.NewVMIngestionClient(subscriptionID)
}
func NewVMIngestionClientWithBaseURI(baseURI string, subscriptionID string) VMIngestionClient {
	return original.NewVMIngestionClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleDeploymentStatusValues() []DeploymentStatus {
	return original.PossibleDeploymentStatusValues()
}
func PossibleLiftrResourceCategoriesValues() []LiftrResourceCategories {
	return original.PossibleLiftrResourceCategoriesValues()
}
func PossibleManagedIdentityTypesValues() []ManagedIdentityTypes {
	return original.PossibleManagedIdentityTypesValues()
}
func PossibleMonitoringStatusValues() []MonitoringStatus {
	return original.PossibleMonitoringStatusValues()
}
func PossibleOperationNameValues() []OperationName {
	return original.PossibleOperationNameValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSendingLogsValues() []SendingLogs {
	return original.PossibleSendingLogsValues()
}
func PossibleTagActionValues() []TagAction {
	return original.PossibleTagActionValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}