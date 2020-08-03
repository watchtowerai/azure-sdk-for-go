package subscription

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

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/subscription/mgmt/2019-10-01-preview/subscription"

// AdPrincipal active Directory Principal who’ll get owner access on the new subscription.
type AdPrincipal struct {
	// ObjectID - Object id of the Principal
	ObjectID *string `json:"objectId,omitempty"`
}

// CanceledSubscriptionID the ID of the canceled subscription
type CanceledSubscriptionID struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The ID of the canceled subscription
	Value *string `json:"value,omitempty"`
}

// CreateCspSubscriptionFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type CreateCspSubscriptionFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *CreateCspSubscriptionFuture) Result(client Client) (cr CreationResult, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.CreateCspSubscriptionFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("subscription.CreateCspSubscriptionFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if cr.Response.Response, err = future.GetResult(sender); err == nil && cr.Response.Response.StatusCode != http.StatusNoContent {
		cr, err = client.CreateCspSubscriptionResponder(cr.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "subscription.CreateCspSubscriptionFuture", "Result", cr.Response.Response, "Failure responding to request")
		}
	}
	return
}

// CreateSubscriptionFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type CreateSubscriptionFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *CreateSubscriptionFuture) Result(client Client) (cr CreationResult, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.CreateSubscriptionFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("subscription.CreateSubscriptionFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if cr.Response.Response, err = future.GetResult(sender); err == nil && cr.Response.Response.StatusCode != http.StatusNoContent {
		cr, err = client.CreateSubscriptionResponder(cr.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "subscription.CreateSubscriptionFuture", "Result", cr.Response.Response, "Failure responding to request")
		}
	}
	return
}

// CreateSubscriptionInEnrollmentAccountFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type CreateSubscriptionInEnrollmentAccountFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *CreateSubscriptionInEnrollmentAccountFuture) Result(client Client) (cr CreationResult, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.CreateSubscriptionInEnrollmentAccountFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("subscription.CreateSubscriptionInEnrollmentAccountFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if cr.Response.Response, err = future.GetResult(sender); err == nil && cr.Response.Response.StatusCode != http.StatusNoContent {
		cr, err = client.CreateSubscriptionInEnrollmentAccountResponder(cr.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "subscription.CreateSubscriptionInEnrollmentAccountFuture", "Result", cr.Response.Response, "Failure responding to request")
		}
	}
	return
}

// CreationParameters subscription Creation Parameters required to create a new Azure subscription.
type CreationParameters struct {
	// DisplayName - The display name of the subscription.
	DisplayName *string `json:"displayName,omitempty"`
	// ManagementGroupID - The Management Group Id.
	ManagementGroupID *string `json:"managementGroupId,omitempty"`
	// Owners - The list of principals that should be granted Owner access on the subscription. Principals should be of type User, Service Principal or Security Group.
	Owners *[]AdPrincipal `json:"owners,omitempty"`
	// OfferType - The offer type of the subscription. For example, MS-AZR-0017P (EnterpriseAgreement) and MS-AZR-0148P (EnterpriseAgreement devTest) are available. Only valid when creating a subscription in a enrollment account scope. Possible values include: 'MSAZR0017P', 'MSAZR0148P'
	OfferType OfferType `json:"offerType,omitempty"`
	// AdditionalParameters - Additional, untyped parameters to support custom subscription creation scenarios.
	AdditionalParameters map[string]interface{} `json:"additionalParameters"`
}

// MarshalJSON is the custom marshaler for CreationParameters.
func (cp CreationParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cp.DisplayName != nil {
		objectMap["displayName"] = cp.DisplayName
	}
	if cp.ManagementGroupID != nil {
		objectMap["managementGroupId"] = cp.ManagementGroupID
	}
	if cp.Owners != nil {
		objectMap["owners"] = cp.Owners
	}
	if cp.OfferType != "" {
		objectMap["offerType"] = cp.OfferType
	}
	if cp.AdditionalParameters != nil {
		objectMap["additionalParameters"] = cp.AdditionalParameters
	}
	return json.Marshal(objectMap)
}

// CreationResult the created subscription object.
type CreationResult struct {
	autorest.Response `json:"-"`
	// SubscriptionLink - The link to the new subscription. Use this link to check the status of subscription creation operation.
	SubscriptionLink *string `json:"subscriptionLink,omitempty"`
}

// EnabledSubscriptionID the ID of the subscriptions that is being enabled
type EnabledSubscriptionID struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The ID of the subscriptions that is being enabled
	Value *string `json:"value,omitempty"`
}

// ErrorResponse describes the format of Error response.
type ErrorResponse struct {
	// Code - Error code
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// ListResult subscription list operation response.
type ListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of subscriptions.
	Value *[]Model `json:"value,omitempty"`
	// NextLink - The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ListResultIterator provides access to a complete listing of Model values.
type ListResultIterator struct {
	i    int
	page ListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListResultIterator) Response() ListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListResultIterator) Value() Model {
	if !iter.page.NotDone() {
		return Model{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ListResultIterator type.
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return ListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (lr ListResult) IsEmpty() bool {
	return lr.Value == nil || len(*lr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (lr ListResult) hasNextLink() bool {
	return lr.NextLink != nil && len(*lr.NextLink) != 0
}

// listResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lr ListResult) listResultPreparer(ctx context.Context) (*http.Request, error) {
	if !lr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lr.NextLink)))
}

// ListResultPage contains a page of Model values.
type ListResultPage struct {
	fn func(context.Context, ListResult) (ListResult, error)
	lr ListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.lr)
		if err != nil {
			return err
		}
		page.lr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListResultPage) NotDone() bool {
	return !page.lr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListResultPage) Response() ListResult {
	return page.lr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListResultPage) Values() []Model {
	if page.lr.IsEmpty() {
		return nil
	}
	return *page.lr.Value
}

// Creates a new instance of the ListResultPage type.
func NewListResultPage(getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return ListResultPage{fn: getNextPage}
}

// Location location information.
type Location struct {
	// ID - READ-ONLY; The fully qualified ID of the location. For example, /subscriptions/00000000-0000-0000-0000-000000000000/locations/westus.
	ID *string `json:"id,omitempty"`
	// SubscriptionID - READ-ONLY; The subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// Name - READ-ONLY; The location name.
	Name *string `json:"name,omitempty"`
	// DisplayName - READ-ONLY; The display name of the location.
	DisplayName *string `json:"displayName,omitempty"`
	// Latitude - READ-ONLY; The latitude of the location.
	Latitude *string `json:"latitude,omitempty"`
	// Longitude - READ-ONLY; The longitude of the location.
	Longitude *string `json:"longitude,omitempty"`
}

// LocationListResult location list operation response.
type LocationListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of locations.
	Value *[]Location `json:"value,omitempty"`
}

// Model subscription information.
type Model struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The fully qualified ID for the subscription. For example, /subscriptions/00000000-0000-0000-0000-000000000000.
	ID *string `json:"id,omitempty"`
	// SubscriptionID - READ-ONLY; The subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// DisplayName - READ-ONLY; The subscription display name.
	DisplayName *string `json:"displayName,omitempty"`
	// State - READ-ONLY; The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted. Possible values include: 'Enabled', 'Warned', 'PastDue', 'Disabled', 'Deleted'
	State State `json:"state,omitempty"`
	// SubscriptionPolicies - The subscription policies.
	SubscriptionPolicies *Policies `json:"subscriptionPolicies,omitempty"`
	// AuthorizationSource - The authorization source of the request. Valid values are one or more combinations of Legacy, RoleBased, Bypassed, Direct and Management. For example, 'Legacy, RoleBased'.
	AuthorizationSource *string `json:"authorizationSource,omitempty"`
}

// MarshalJSON is the custom marshaler for Model.
func (mVar Model) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mVar.SubscriptionPolicies != nil {
		objectMap["subscriptionPolicies"] = mVar.SubscriptionPolicies
	}
	if mVar.AuthorizationSource != nil {
		objectMap["authorizationSource"] = mVar.AuthorizationSource
	}
	return json.Marshal(objectMap)
}

// ModernCspSubscriptionCreationParameters the parameters required to create a new CSP subscription.
type ModernCspSubscriptionCreationParameters struct {
	// DisplayName - The friendly name of the subscription.
	DisplayName *string `json:"displayName,omitempty"`
	// SkuID - The SKU ID of the Azure plan. Azure plan determines the pricing and service-level agreement of the subscription.  Use 001 for Microsoft Azure Plan and 002 for Microsoft Azure Plan for DevTest.
	SkuID *string `json:"skuId,omitempty"`
	// ResellerID - Reseller ID, basically MPN Id.
	ResellerID *string `json:"resellerId,omitempty"`
}

// ModernSubscriptionCreationParameters the parameters required to create a new subscription.
type ModernSubscriptionCreationParameters struct {
	// DisplayName - The friendly name of the subscription.
	DisplayName *string `json:"displayName,omitempty"`
	// SkuID - The SKU ID of the Azure plan. Azure plan determines the pricing and service-level agreement of the subscription.  Use 001 for Microsoft Azure Plan and 002 for Microsoft Azure Plan for DevTest.
	SkuID *string `json:"skuId,omitempty"`
	// CostCenter - If set, the cost center will show up on the Azure usage and charges file.
	CostCenter *string `json:"costCenter,omitempty"`
	// Owner - If specified, the AD principal will get owner access to the subscription, along with the user who is performing the create subscription operation
	Owner *AdPrincipal `json:"owner,omitempty"`
	// ManagementGroupID - The identifier of the management group to which this subscription will be associated.
	ManagementGroupID *string `json:"managementGroupId,omitempty"`
	// AdditionalParameters - Additional, untyped parameters to support custom subscription creation scenarios.
	AdditionalParameters map[string]interface{} `json:"additionalParameters"`
}

// MarshalJSON is the custom marshaler for ModernSubscriptionCreationParameters.
func (mscp ModernSubscriptionCreationParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mscp.DisplayName != nil {
		objectMap["displayName"] = mscp.DisplayName
	}
	if mscp.SkuID != nil {
		objectMap["skuId"] = mscp.SkuID
	}
	if mscp.CostCenter != nil {
		objectMap["costCenter"] = mscp.CostCenter
	}
	if mscp.Owner != nil {
		objectMap["owner"] = mscp.Owner
	}
	if mscp.ManagementGroupID != nil {
		objectMap["managementGroupId"] = mscp.ManagementGroupID
	}
	if mscp.AdditionalParameters != nil {
		objectMap["additionalParameters"] = mscp.AdditionalParameters
	}
	return json.Marshal(objectMap)
}

// Name the new name of the subscription.
type Name struct {
	// SubscriptionName - New subscription name
	SubscriptionName *string `json:"subscriptionName,omitempty"`
}

// Operation REST API operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Subscription
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of the request to list operations. It contains a list of operations and a URL
// link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of operations.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// Policies subscription policies.
type Policies struct {
	// LocationPlacementID - READ-ONLY; The subscription location placement ID. The ID indicates which regions are visible for a subscription. For example, a subscription with a location placement Id of Public_2014-09-01 has access to Azure public regions.
	LocationPlacementID *string `json:"locationPlacementId,omitempty"`
	// QuotaID - READ-ONLY; The subscription quota ID.
	QuotaID *string `json:"quotaId,omitempty"`
	// SpendingLimit - READ-ONLY; The subscription spending limit. Possible values include: 'On', 'Off', 'CurrentPeriodOff'
	SpendingLimit SpendingLimit `json:"spendingLimit,omitempty"`
}

// RenamedSubscriptionID the ID of the subscriptions that is being renamed
type RenamedSubscriptionID struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The ID of the subscriptions that is being renamed
	Value *string `json:"value,omitempty"`
}

// TenantIDDescription tenant Id information.
type TenantIDDescription struct {
	// ID - READ-ONLY; The fully qualified ID of the tenant. For example, /tenants/00000000-0000-0000-0000-000000000000.
	ID *string `json:"id,omitempty"`
	// TenantID - READ-ONLY; The tenant ID. For example, 00000000-0000-0000-0000-000000000000.
	TenantID *string `json:"tenantId,omitempty"`
}

// TenantListResult tenant Ids information.
type TenantListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of tenants.
	Value *[]TenantIDDescription `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// TenantListResultIterator provides access to a complete listing of TenantIDDescription values.
type TenantListResultIterator struct {
	i    int
	page TenantListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *TenantListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *TenantListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter TenantListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter TenantListResultIterator) Response() TenantListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter TenantListResultIterator) Value() TenantIDDescription {
	if !iter.page.NotDone() {
		return TenantIDDescription{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the TenantListResultIterator type.
func NewTenantListResultIterator(page TenantListResultPage) TenantListResultIterator {
	return TenantListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (tlr TenantListResult) IsEmpty() bool {
	return tlr.Value == nil || len(*tlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (tlr TenantListResult) hasNextLink() bool {
	return tlr.NextLink != nil && len(*tlr.NextLink) != 0
}

// tenantListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (tlr TenantListResult) tenantListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !tlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(tlr.NextLink)))
}

// TenantListResultPage contains a page of TenantIDDescription values.
type TenantListResultPage struct {
	fn  func(context.Context, TenantListResult) (TenantListResult, error)
	tlr TenantListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *TenantListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.tlr)
		if err != nil {
			return err
		}
		page.tlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *TenantListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page TenantListResultPage) NotDone() bool {
	return !page.tlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page TenantListResultPage) Response() TenantListResult {
	return page.tlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page TenantListResultPage) Values() []TenantIDDescription {
	if page.tlr.IsEmpty() {
		return nil
	}
	return *page.tlr.Value
}

// Creates a new instance of the TenantListResultPage type.
func NewTenantListResultPage(getNextPage func(context.Context, TenantListResult) (TenantListResult, error)) TenantListResultPage {
	return TenantListResultPage{fn: getNextPage}
}
