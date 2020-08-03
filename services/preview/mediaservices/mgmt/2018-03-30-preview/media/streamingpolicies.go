package media

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// StreamingPoliciesClient is the client for the StreamingPolicies methods of the Media service.
type StreamingPoliciesClient struct {
	BaseClient
}

// NewStreamingPoliciesClient creates an instance of the StreamingPoliciesClient client.
func NewStreamingPoliciesClient(subscriptionID string) StreamingPoliciesClient {
	return NewStreamingPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStreamingPoliciesClientWithBaseURI creates an instance of the StreamingPoliciesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewStreamingPoliciesClientWithBaseURI(baseURI string, subscriptionID string) StreamingPoliciesClient {
	return StreamingPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a Streaming Policy in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingPolicyName - the Streaming Policy name.
// parameters - the request parameters
func (client StreamingPoliciesClient) Create(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string, parameters StreamingPolicy) (result StreamingPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingPoliciesClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.StreamingPolicyProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.StreamingPolicyProperties.EnvelopeEncryption", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.StreamingPolicyProperties.EnvelopeEncryption.EnabledProtocols", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.StreamingPolicyProperties.EnvelopeEncryption.EnabledProtocols.Download", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.StreamingPolicyProperties.EnvelopeEncryption.EnabledProtocols.Dash", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.StreamingPolicyProperties.EnvelopeEncryption.EnabledProtocols.Hls", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.StreamingPolicyProperties.EnvelopeEncryption.EnabledProtocols.SmoothStreaming", Name: validation.Null, Rule: true, Chain: nil},
						}},
					}},
					{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCenc", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCenc.EnabledProtocols", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCenc.EnabledProtocols.Download", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCenc.EnabledProtocols.Dash", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCenc.EnabledProtocols.Hls", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCenc.EnabledProtocols.SmoothStreaming", Name: validation.Null, Rule: true, Chain: nil},
							}},
						}},
					{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCbcs", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCbcs.EnabledProtocols", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCbcs.EnabledProtocols.Download", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCbcs.EnabledProtocols.Dash", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCbcs.EnabledProtocols.Hls", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCbcs.EnabledProtocols.SmoothStreaming", Name: validation.Null, Rule: true, Chain: nil},
							}},
							{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCbcs.Drm", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCbcs.Drm.FairPlay", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.StreamingPolicyProperties.CommonEncryptionCbcs.Drm.FairPlay.AllowPersistentLicense", Name: validation.Null, Rule: true, Chain: nil}}},
								}},
						}},
					{Target: "parameters.StreamingPolicyProperties.NoEncryption", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.StreamingPolicyProperties.NoEncryption.EnabledProtocols", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.StreamingPolicyProperties.NoEncryption.EnabledProtocols.Download", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.StreamingPolicyProperties.NoEncryption.EnabledProtocols.Dash", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.StreamingPolicyProperties.NoEncryption.EnabledProtocols.Hls", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.StreamingPolicyProperties.NoEncryption.EnabledProtocols.SmoothStreaming", Name: validation.Null, Rule: true, Chain: nil},
							}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("media.StreamingPoliciesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, streamingPolicyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client StreamingPoliciesClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string, parameters StreamingPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":         autorest.Encode("path", accountName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"streamingPolicyName": autorest.Encode("path", streamingPolicyName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingPoliciesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client StreamingPoliciesClient) CreateResponder(resp *http.Response) (result StreamingPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Streaming Policy in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingPolicyName - the Streaming Policy name.
func (client StreamingPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingPoliciesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, streamingPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client StreamingPoliciesClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":         autorest.Encode("path", accountName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"streamingPolicyName": autorest.Encode("path", streamingPolicyName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingPoliciesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client StreamingPoliciesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the details of a Streaming Policy in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// streamingPolicyName - the Streaming Policy name.
func (client StreamingPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string) (result StreamingPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, streamingPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client StreamingPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":         autorest.Encode("path", accountName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"streamingPolicyName": autorest.Encode("path", streamingPolicyName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StreamingPoliciesClient) GetResponder(resp *http.Response) (result StreamingPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the Streaming Policies in the account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// filter - restricts the set of items returned.
// top - specifies a non-negative integer n that limits the number of items returned from a collection. The
// service returns the number of available items up to but not greater than the specified value n.
// orderby - specifies the key by which the result collection should be ordered.
func (client StreamingPoliciesClient) List(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (result StreamingPolicyCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingPoliciesClient.List")
		defer func() {
			sc := -1
			if result.spc.Response.Response != nil {
				sc = result.spc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, accountName, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.spc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "List", resp, "Failure sending request")
		return
	}

	result.spc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "List", resp, "Failure responding to request")
	}
	if result.spc.hasNextLink() && result.spc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client StreamingPoliciesClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client StreamingPoliciesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client StreamingPoliciesClient) ListResponder(resp *http.Response) (result StreamingPolicyCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client StreamingPoliciesClient) listNextResults(ctx context.Context, lastResults StreamingPolicyCollection) (result StreamingPolicyCollection, err error) {
	req, err := lastResults.streamingPolicyCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.StreamingPoliciesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client StreamingPoliciesClient) ListComplete(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (result StreamingPolicyCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StreamingPoliciesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, accountName, filter, top, orderby)
	return
}
