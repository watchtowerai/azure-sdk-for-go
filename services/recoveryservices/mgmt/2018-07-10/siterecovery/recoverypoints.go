package siterecovery

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

// RecoveryPointsClient is the client for the RecoveryPoints methods of the Siterecovery service.
type RecoveryPointsClient struct {
	BaseClient
}

// NewRecoveryPointsClient creates an instance of the RecoveryPointsClient client.
func NewRecoveryPointsClient(subscriptionID string, resourceGroupName string, resourceName string) RecoveryPointsClient {
	return NewRecoveryPointsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewRecoveryPointsClientWithBaseURI creates an instance of the RecoveryPointsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRecoveryPointsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) RecoveryPointsClient {
	return RecoveryPointsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Get get the details of specified recovery point.
// Parameters:
// fabricName - the fabric name.
// protectionContainerName - the protection container name.
// replicatedProtectedItemName - the replication protected item's name.
// recoveryPointName - the recovery point name.
func (client RecoveryPointsClient) Get(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string, recoveryPointName string) (result RecoveryPoint, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoveryPointsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, fabricName, protectionContainerName, replicatedProtectedItemName, recoveryPointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.RecoveryPointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.RecoveryPointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.RecoveryPointsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecoveryPointsClient) GetPreparer(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string, recoveryPointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                  autorest.Encode("path", fabricName),
		"protectionContainerName":     autorest.Encode("path", protectionContainerName),
		"recoveryPointName":           autorest.Encode("path", recoveryPointName),
		"replicatedProtectedItemName": autorest.Encode("path", replicatedProtectedItemName),
		"resourceGroupName":           autorest.Encode("path", client.ResourceGroupName),
		"resourceName":                autorest.Encode("path", client.ResourceName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/recoveryPoints/{recoveryPointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecoveryPointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecoveryPointsClient) GetResponder(resp *http.Response) (result RecoveryPoint, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByReplicationProtectedItems lists the available recovery points for a replication protected item.
// Parameters:
// fabricName - the fabric name.
// protectionContainerName - the protection container name.
// replicatedProtectedItemName - the replication protected item's name.
func (client RecoveryPointsClient) ListByReplicationProtectedItems(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string) (result RecoveryPointCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoveryPointsClient.ListByReplicationProtectedItems")
		defer func() {
			sc := -1
			if result.RPCVar.Response.Response != nil {
				sc = result.RPCVar.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByReplicationProtectedItemsNextResults
	req, err := client.ListByReplicationProtectedItemsPreparer(ctx, fabricName, protectionContainerName, replicatedProtectedItemName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.RecoveryPointsClient", "ListByReplicationProtectedItems", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationProtectedItemsSender(req)
	if err != nil {
		result.RPCVar.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.RecoveryPointsClient", "ListByReplicationProtectedItems", resp, "Failure sending request")
		return
	}

	result.RPCVar, err = client.ListByReplicationProtectedItemsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.RecoveryPointsClient", "ListByReplicationProtectedItems", resp, "Failure responding to request")
	}
	if result.RPCVar.hasNextLink() && result.RPCVar.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByReplicationProtectedItemsPreparer prepares the ListByReplicationProtectedItems request.
func (client RecoveryPointsClient) ListByReplicationProtectedItemsPreparer(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                  autorest.Encode("path", fabricName),
		"protectionContainerName":     autorest.Encode("path", protectionContainerName),
		"replicatedProtectedItemName": autorest.Encode("path", replicatedProtectedItemName),
		"resourceGroupName":           autorest.Encode("path", client.ResourceGroupName),
		"resourceName":                autorest.Encode("path", client.ResourceName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/recoveryPoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByReplicationProtectedItemsSender sends the ListByReplicationProtectedItems request. The method will close the
// http.Response Body if it receives an error.
func (client RecoveryPointsClient) ListByReplicationProtectedItemsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByReplicationProtectedItemsResponder handles the response to the ListByReplicationProtectedItems request. The method always
// closes the http.Response Body.
func (client RecoveryPointsClient) ListByReplicationProtectedItemsResponder(resp *http.Response) (result RecoveryPointCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByReplicationProtectedItemsNextResults retrieves the next set of results, if any.
func (client RecoveryPointsClient) listByReplicationProtectedItemsNextResults(ctx context.Context, lastResults RecoveryPointCollection) (result RecoveryPointCollection, err error) {
	req, err := lastResults.recoveryPointCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.RecoveryPointsClient", "listByReplicationProtectedItemsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByReplicationProtectedItemsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.RecoveryPointsClient", "listByReplicationProtectedItemsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByReplicationProtectedItemsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.RecoveryPointsClient", "listByReplicationProtectedItemsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByReplicationProtectedItemsComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecoveryPointsClient) ListByReplicationProtectedItemsComplete(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string) (result RecoveryPointCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoveryPointsClient.ListByReplicationProtectedItems")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByReplicationProtectedItems(ctx, fabricName, protectionContainerName, replicatedProtectedItemName)
	return
}
