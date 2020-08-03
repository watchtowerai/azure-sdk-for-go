package blockchain

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// MembersClient is the REST API for Azure Blockchain Service
type MembersClient struct {
	BaseClient
}

// NewMembersClient creates an instance of the MembersClient client.
func NewMembersClient(subscriptionID string) MembersClient {
	return NewMembersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMembersClientWithBaseURI creates an instance of the MembersClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMembersClientWithBaseURI(baseURI string, subscriptionID string) MembersClient {
	return MembersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a blockchain member.
// Parameters:
// blockchainMemberName - blockchain member name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// blockchainMember - payload to create a blockchain member.
func (client MembersClient) Create(ctx context.Context, blockchainMemberName string, resourceGroupName string, blockchainMember *Member) (result MembersCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MembersClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, blockchainMemberName, resourceGroupName, blockchainMember)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client MembersClient) CreatePreparer(ctx context.Context, blockchainMemberName string, resourceGroupName string, blockchainMember *Member) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if blockchainMember != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(blockchainMember))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client MembersClient) CreateSender(req *http.Request) (future MembersCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client MembersClient) CreateResponder(resp *http.Response) (result Member, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a blockchain member.
// Parameters:
// blockchainMemberName - blockchain member name
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
func (client MembersClient) Delete(ctx context.Context, blockchainMemberName string, resourceGroupName string) (result MembersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MembersClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, blockchainMemberName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MembersClient) DeletePreparer(ctx context.Context, blockchainMemberName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MembersClient) DeleteSender(req *http.Request) (future MembersDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MembersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get details about a blockchain member.
// Parameters:
// blockchainMemberName - blockchain member name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
func (client MembersClient) Get(ctx context.Context, blockchainMemberName string, resourceGroupName string) (result Member, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MembersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, blockchainMemberName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MembersClient) GetPreparer(ctx context.Context, blockchainMemberName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MembersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MembersClient) GetResponder(resp *http.Response) (result Member, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the blockchain members for a resource group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
func (client MembersClient) List(ctx context.Context, resourceGroupName string) (result MemberCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MembersClient.List")
		defer func() {
			sc := -1
			if result.mc.Response.Response != nil {
				sc = result.mc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.mc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "List", resp, "Failure sending request")
		return
	}

	result.mc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "List", resp, "Failure responding to request")
	}
	if result.mc.hasNextLink() && result.mc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client MembersClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MembersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MembersClient) ListResponder(resp *http.Response) (result MemberCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client MembersClient) listNextResults(ctx context.Context, lastResults MemberCollection) (result MemberCollection, err error) {
	req, err := lastResults.memberCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "blockchain.MembersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "blockchain.MembersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client MembersClient) ListComplete(ctx context.Context, resourceGroupName string) (result MemberCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MembersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// ListAll lists the blockchain members for a subscription.
func (client MembersClient) ListAll(ctx context.Context) (result MemberCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MembersClient.ListAll")
		defer func() {
			sc := -1
			if result.mc.Response.Response != nil {
				sc = result.mc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.mc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.mc, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "ListAll", resp, "Failure responding to request")
	}
	if result.mc.hasNextLink() && result.mc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client MembersClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Blockchain/blockchainMembers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client MembersClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client MembersClient) ListAllResponder(resp *http.Response) (result MemberCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client MembersClient) listAllNextResults(ctx context.Context, lastResults MemberCollection) (result MemberCollection, err error) {
	req, err := lastResults.memberCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "blockchain.MembersClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "blockchain.MembersClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client MembersClient) ListAllComplete(ctx context.Context) (result MemberCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MembersClient.ListAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAll(ctx)
	return
}

// ListAPIKeys lists the API keys for a blockchain member.
// Parameters:
// blockchainMemberName - blockchain member name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
func (client MembersClient) ListAPIKeys(ctx context.Context, blockchainMemberName string, resourceGroupName string) (result APIKeyCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MembersClient.ListAPIKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListAPIKeysPreparer(ctx, blockchainMemberName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "ListAPIKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAPIKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "ListAPIKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListAPIKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "ListAPIKeys", resp, "Failure responding to request")
	}

	return
}

// ListAPIKeysPreparer prepares the ListAPIKeys request.
func (client MembersClient) ListAPIKeysPreparer(ctx context.Context, blockchainMemberName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/listApiKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAPIKeysSender sends the ListAPIKeys request. The method will close the
// http.Response Body if it receives an error.
func (client MembersClient) ListAPIKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAPIKeysResponder handles the response to the ListAPIKeys request. The method always
// closes the http.Response Body.
func (client MembersClient) ListAPIKeysResponder(resp *http.Response) (result APIKeyCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListConsortiumMembers lists the consortium members for a blockchain member.
// Parameters:
// blockchainMemberName - blockchain member name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
func (client MembersClient) ListConsortiumMembers(ctx context.Context, blockchainMemberName string, resourceGroupName string) (result ConsortiumMemberCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MembersClient.ListConsortiumMembers")
		defer func() {
			sc := -1
			if result.cmc.Response.Response != nil {
				sc = result.cmc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listConsortiumMembersNextResults
	req, err := client.ListConsortiumMembersPreparer(ctx, blockchainMemberName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "ListConsortiumMembers", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListConsortiumMembersSender(req)
	if err != nil {
		result.cmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "ListConsortiumMembers", resp, "Failure sending request")
		return
	}

	result.cmc, err = client.ListConsortiumMembersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "ListConsortiumMembers", resp, "Failure responding to request")
	}
	if result.cmc.hasNextLink() && result.cmc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListConsortiumMembersPreparer prepares the ListConsortiumMembers request.
func (client MembersClient) ListConsortiumMembersPreparer(ctx context.Context, blockchainMemberName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/consortiumMembers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListConsortiumMembersSender sends the ListConsortiumMembers request. The method will close the
// http.Response Body if it receives an error.
func (client MembersClient) ListConsortiumMembersSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListConsortiumMembersResponder handles the response to the ListConsortiumMembers request. The method always
// closes the http.Response Body.
func (client MembersClient) ListConsortiumMembersResponder(resp *http.Response) (result ConsortiumMemberCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listConsortiumMembersNextResults retrieves the next set of results, if any.
func (client MembersClient) listConsortiumMembersNextResults(ctx context.Context, lastResults ConsortiumMemberCollection) (result ConsortiumMemberCollection, err error) {
	req, err := lastResults.consortiumMemberCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "blockchain.MembersClient", "listConsortiumMembersNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListConsortiumMembersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "blockchain.MembersClient", "listConsortiumMembersNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListConsortiumMembersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "listConsortiumMembersNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListConsortiumMembersComplete enumerates all values, automatically crossing page boundaries as required.
func (client MembersClient) ListConsortiumMembersComplete(ctx context.Context, blockchainMemberName string, resourceGroupName string) (result ConsortiumMemberCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MembersClient.ListConsortiumMembers")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListConsortiumMembers(ctx, blockchainMemberName, resourceGroupName)
	return
}

// ListRegenerateAPIKeys regenerate the API keys for a blockchain member.
// Parameters:
// blockchainMemberName - blockchain member name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// APIKey - api key to be regenerate
func (client MembersClient) ListRegenerateAPIKeys(ctx context.Context, blockchainMemberName string, resourceGroupName string, APIKey *APIKey) (result APIKeyCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MembersClient.ListRegenerateAPIKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListRegenerateAPIKeysPreparer(ctx, blockchainMemberName, resourceGroupName, APIKey)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "ListRegenerateAPIKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRegenerateAPIKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "ListRegenerateAPIKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListRegenerateAPIKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "ListRegenerateAPIKeys", resp, "Failure responding to request")
	}

	return
}

// ListRegenerateAPIKeysPreparer prepares the ListRegenerateAPIKeys request.
func (client MembersClient) ListRegenerateAPIKeysPreparer(ctx context.Context, blockchainMemberName string, resourceGroupName string, APIKey *APIKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/regenerateApiKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if APIKey != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(APIKey))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRegenerateAPIKeysSender sends the ListRegenerateAPIKeys request. The method will close the
// http.Response Body if it receives an error.
func (client MembersClient) ListRegenerateAPIKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListRegenerateAPIKeysResponder handles the response to the ListRegenerateAPIKeys request. The method always
// closes the http.Response Body.
func (client MembersClient) ListRegenerateAPIKeysResponder(resp *http.Response) (result APIKeyCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update a blockchain member.
// Parameters:
// blockchainMemberName - blockchain member name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// blockchainMember - payload to update the blockchain member.
func (client MembersClient) Update(ctx context.Context, blockchainMemberName string, resourceGroupName string, blockchainMember *MemberUpdate) (result Member, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MembersClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, blockchainMemberName, resourceGroupName, blockchainMember)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.MembersClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client MembersClient) UpdatePreparer(ctx context.Context, blockchainMemberName string, resourceGroupName string, blockchainMember *MemberUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if blockchainMember != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(blockchainMember))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client MembersClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client MembersClient) UpdateResponder(resp *http.Response) (result Member, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
