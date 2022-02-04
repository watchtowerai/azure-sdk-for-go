//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SolutionsClient contains the methods for the SecuritySolutions group.
// Don't use this type directly, use NewSolutionsClient() instead.
type SolutionsClient struct {
	host           string
	subscriptionID string
	ascLocation    string
	pl             runtime.Pipeline
}

// NewSolutionsClient creates a new instance of SolutionsClient with the specified values.
// subscriptionID - Azure subscription ID
// ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSolutionsClient(subscriptionID string, ascLocation string, credential azcore.TokenCredential, options *arm.ClientOptions) *SolutionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SolutionsClient{
		subscriptionID: subscriptionID,
		ascLocation:    ascLocation,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets a specific Security Solution.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// securitySolutionName - Name of security solution.
// options - SolutionsClientGetOptions contains the optional parameters for the SolutionsClient.Get method.
func (client *SolutionsClient) Get(ctx context.Context, resourceGroupName string, securitySolutionName string, options *SolutionsClientGetOptions) (SolutionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, securitySolutionName, options)
	if err != nil {
		return SolutionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SolutionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SolutionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SolutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, securitySolutionName string, options *SolutionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/securitySolutions/{securitySolutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.ascLocation == "" {
		return nil, errors.New("parameter client.ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(client.ascLocation))
	if securitySolutionName == "" {
		return nil, errors.New("parameter securitySolutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securitySolutionName}", url.PathEscape(securitySolutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SolutionsClient) getHandleResponse(resp *http.Response) (SolutionsClientGetResponse, error) {
	result := SolutionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Solution); err != nil {
		return SolutionsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of Security Solutions for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - SolutionsClientListOptions contains the optional parameters for the SolutionsClient.List method.
func (client *SolutionsClient) List(options *SolutionsClientListOptions) *SolutionsClientListPager {
	return &SolutionsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SolutionsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SolutionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SolutionsClient) listCreateRequest(ctx context.Context, options *SolutionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securitySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SolutionsClient) listHandleResponse(resp *http.Response) (SolutionsClientListResponse, error) {
	result := SolutionsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SolutionList); err != nil {
		return SolutionsClientListResponse{}, err
	}
	return result, nil
}