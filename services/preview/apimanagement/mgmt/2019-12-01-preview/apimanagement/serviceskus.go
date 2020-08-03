package apimanagement

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

// ServiceSkusClient is the apiManagement Client
type ServiceSkusClient struct {
	BaseClient
}

// NewServiceSkusClient creates an instance of the ServiceSkusClient client.
func NewServiceSkusClient(subscriptionID string) ServiceSkusClient {
	return NewServiceSkusClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceSkusClientWithBaseURI creates an instance of the ServiceSkusClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewServiceSkusClientWithBaseURI(baseURI string, subscriptionID string) ServiceSkusClient {
	return ServiceSkusClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListAvailableServiceSkus gets all available SKU for a given API Management service
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
func (client ServiceSkusClient) ListAvailableServiceSkus(ctx context.Context, resourceGroupName string, serviceName string) (result ResourceSkuResultsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceSkusClient.ListAvailableServiceSkus")
		defer func() {
			sc := -1
			if result.rsr.Response.Response != nil {
				sc = result.rsr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.ServiceSkusClient", "ListAvailableServiceSkus", err.Error())
	}

	result.fn = client.listAvailableServiceSkusNextResults
	req, err := client.ListAvailableServiceSkusPreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ServiceSkusClient", "ListAvailableServiceSkus", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAvailableServiceSkusSender(req)
	if err != nil {
		result.rsr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.ServiceSkusClient", "ListAvailableServiceSkus", resp, "Failure sending request")
		return
	}

	result.rsr, err = client.ListAvailableServiceSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ServiceSkusClient", "ListAvailableServiceSkus", resp, "Failure responding to request")
	}
	if result.rsr.hasNextLink() && result.rsr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListAvailableServiceSkusPreparer prepares the ListAvailableServiceSkus request.
func (client ServiceSkusClient) ListAvailableServiceSkusPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/skus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAvailableServiceSkusSender sends the ListAvailableServiceSkus request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceSkusClient) ListAvailableServiceSkusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAvailableServiceSkusResponder handles the response to the ListAvailableServiceSkus request. The method always
// closes the http.Response Body.
func (client ServiceSkusClient) ListAvailableServiceSkusResponder(resp *http.Response) (result ResourceSkuResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAvailableServiceSkusNextResults retrieves the next set of results, if any.
func (client ServiceSkusClient) listAvailableServiceSkusNextResults(ctx context.Context, lastResults ResourceSkuResults) (result ResourceSkuResults, err error) {
	req, err := lastResults.resourceSkuResultsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ServiceSkusClient", "listAvailableServiceSkusNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAvailableServiceSkusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ServiceSkusClient", "listAvailableServiceSkusNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAvailableServiceSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ServiceSkusClient", "listAvailableServiceSkusNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAvailableServiceSkusComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServiceSkusClient) ListAvailableServiceSkusComplete(ctx context.Context, resourceGroupName string, serviceName string) (result ResourceSkuResultsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceSkusClient.ListAvailableServiceSkus")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAvailableServiceSkus(ctx, resourceGroupName, serviceName)
	return
}
