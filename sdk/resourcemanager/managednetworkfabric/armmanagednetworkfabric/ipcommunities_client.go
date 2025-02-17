//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagednetworkfabric

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// IPCommunitiesClient contains the methods for the IPCommunities group.
// Don't use this type directly, use NewIPCommunitiesClient() instead.
type IPCommunitiesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIPCommunitiesClient creates a new instance of IPCommunitiesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIPCommunitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IPCommunitiesClient, error) {
	cl, err := arm.NewClient(moduleName+".IPCommunitiesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IPCommunitiesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Implements an IP Community PUT method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ipCommunityName - Name of the IP Community.
//   - body - Request payload.
//   - options - IPCommunitiesClientBeginCreateOptions contains the optional parameters for the IPCommunitiesClient.BeginCreate
//     method.
func (client *IPCommunitiesClient) BeginCreate(ctx context.Context, resourceGroupName string, ipCommunityName string, body IPCommunity, options *IPCommunitiesClientBeginCreateOptions) (*runtime.Poller[IPCommunitiesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, ipCommunityName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IPCommunitiesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[IPCommunitiesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Implements an IP Community PUT method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *IPCommunitiesClient) create(ctx context.Context, resourceGroupName string, ipCommunityName string, body IPCommunity, options *IPCommunitiesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, ipCommunityName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *IPCommunitiesClient) createCreateRequest(ctx context.Context, resourceGroupName string, ipCommunityName string, body IPCommunity, options *IPCommunitiesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/ipCommunities/{ipCommunityName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipCommunityName == "" {
		return nil, errors.New("parameter ipCommunityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipCommunityName}", url.PathEscape(ipCommunityName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Implements IP Community DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ipCommunityName - Name of the IP Community.
//   - options - IPCommunitiesClientBeginDeleteOptions contains the optional parameters for the IPCommunitiesClient.BeginDelete
//     method.
func (client *IPCommunitiesClient) BeginDelete(ctx context.Context, resourceGroupName string, ipCommunityName string, options *IPCommunitiesClientBeginDeleteOptions) (*runtime.Poller[IPCommunitiesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, ipCommunityName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IPCommunitiesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[IPCommunitiesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Implements IP Community DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *IPCommunitiesClient) deleteOperation(ctx context.Context, resourceGroupName string, ipCommunityName string, options *IPCommunitiesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ipCommunityName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IPCommunitiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ipCommunityName string, options *IPCommunitiesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/ipCommunities/{ipCommunityName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipCommunityName == "" {
		return nil, errors.New("parameter ipCommunityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipCommunityName}", url.PathEscape(ipCommunityName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements an IP Community GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ipCommunityName - Name of the IP Community.
//   - options - IPCommunitiesClientGetOptions contains the optional parameters for the IPCommunitiesClient.Get method.
func (client *IPCommunitiesClient) Get(ctx context.Context, resourceGroupName string, ipCommunityName string, options *IPCommunitiesClientGetOptions) (IPCommunitiesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ipCommunityName, options)
	if err != nil {
		return IPCommunitiesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IPCommunitiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IPCommunitiesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IPCommunitiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, ipCommunityName string, options *IPCommunitiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/ipCommunities/{ipCommunityName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipCommunityName == "" {
		return nil, errors.New("parameter ipCommunityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipCommunityName}", url.PathEscape(ipCommunityName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IPCommunitiesClient) getHandleResponse(resp *http.Response) (IPCommunitiesClientGetResponse, error) {
	result := IPCommunitiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPCommunity); err != nil {
		return IPCommunitiesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Implements IP Communities list by resource group GET method.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - IPCommunitiesClientListByResourceGroupOptions contains the optional parameters for the IPCommunitiesClient.NewListByResourceGroupPager
//     method.
func (client *IPCommunitiesClient) NewListByResourceGroupPager(resourceGroupName string, options *IPCommunitiesClientListByResourceGroupOptions) *runtime.Pager[IPCommunitiesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[IPCommunitiesClientListByResourceGroupResponse]{
		More: func(page IPCommunitiesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IPCommunitiesClientListByResourceGroupResponse) (IPCommunitiesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IPCommunitiesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return IPCommunitiesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IPCommunitiesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IPCommunitiesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *IPCommunitiesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/ipCommunities"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IPCommunitiesClient) listByResourceGroupHandleResponse(resp *http.Response) (IPCommunitiesClientListByResourceGroupResponse, error) {
	result := IPCommunitiesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPCommunitiesListResult); err != nil {
		return IPCommunitiesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Implements IP Communities list by subscription GET method.
//
// Generated from API version 2023-06-15
//   - options - IPCommunitiesClientListBySubscriptionOptions contains the optional parameters for the IPCommunitiesClient.NewListBySubscriptionPager
//     method.
func (client *IPCommunitiesClient) NewListBySubscriptionPager(options *IPCommunitiesClientListBySubscriptionOptions) *runtime.Pager[IPCommunitiesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[IPCommunitiesClientListBySubscriptionResponse]{
		More: func(page IPCommunitiesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IPCommunitiesClientListBySubscriptionResponse) (IPCommunitiesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IPCommunitiesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return IPCommunitiesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IPCommunitiesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *IPCommunitiesClient) listBySubscriptionCreateRequest(ctx context.Context, options *IPCommunitiesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetworkFabric/ipCommunities"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *IPCommunitiesClient) listBySubscriptionHandleResponse(resp *http.Response) (IPCommunitiesClientListBySubscriptionResponse, error) {
	result := IPCommunitiesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPCommunitiesListResult); err != nil {
		return IPCommunitiesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - API to update certain properties of the IP Community resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ipCommunityName - Name of the IP Community.
//   - body - IP Community properties to update.
//   - options - IPCommunitiesClientBeginUpdateOptions contains the optional parameters for the IPCommunitiesClient.BeginUpdate
//     method.
func (client *IPCommunitiesClient) BeginUpdate(ctx context.Context, resourceGroupName string, ipCommunityName string, body IPCommunityPatch, options *IPCommunitiesClientBeginUpdateOptions) (*runtime.Poller[IPCommunitiesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, ipCommunityName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IPCommunitiesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[IPCommunitiesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - API to update certain properties of the IP Community resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *IPCommunitiesClient) update(ctx context.Context, resourceGroupName string, ipCommunityName string, body IPCommunityPatch, options *IPCommunitiesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, ipCommunityName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *IPCommunitiesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, ipCommunityName string, body IPCommunityPatch, options *IPCommunitiesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/ipCommunities/{ipCommunityName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipCommunityName == "" {
		return nil, errors.New("parameter ipCommunityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipCommunityName}", url.PathEscape(ipCommunityName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
