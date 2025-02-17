//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmosforpostgresql

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

// ConfigurationsClient contains the methods for the Configurations group.
// Don't use this type directly, use NewConfigurationsClient() instead.
type ConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConfigurationsClient creates a new instance of ConfigurationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName+".ConfigurationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets information of a configuration for coordinator and nodes.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - configurationName - The name of the cluster configuration.
//   - options - ConfigurationsClientGetOptions contains the optional parameters for the ConfigurationsClient.Get method.
func (client *ConfigurationsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, options *ConfigurationsClientGetOptions) (ConfigurationsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, configurationName, options)
	if err != nil {
		return ConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, options *ConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/{clusterName}/configurations/{configurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationsClient) getHandleResponse(resp *http.Response) (ConfigurationsClientGetResponse, error) {
	result := ConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Configuration); err != nil {
		return ConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// GetCoordinator - Gets information of a configuration for coordinator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - configurationName - The name of the cluster configuration.
//   - options - ConfigurationsClientGetCoordinatorOptions contains the optional parameters for the ConfigurationsClient.GetCoordinator
//     method.
func (client *ConfigurationsClient) GetCoordinator(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, options *ConfigurationsClientGetCoordinatorOptions) (ConfigurationsClientGetCoordinatorResponse, error) {
	var err error
	req, err := client.getCoordinatorCreateRequest(ctx, resourceGroupName, clusterName, configurationName, options)
	if err != nil {
		return ConfigurationsClientGetCoordinatorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationsClientGetCoordinatorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationsClientGetCoordinatorResponse{}, err
	}
	resp, err := client.getCoordinatorHandleResponse(httpResp)
	return resp, err
}

// getCoordinatorCreateRequest creates the GetCoordinator request.
func (client *ConfigurationsClient) getCoordinatorCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, options *ConfigurationsClientGetCoordinatorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/{clusterName}/coordinatorConfigurations/{configurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getCoordinatorHandleResponse handles the GetCoordinator response.
func (client *ConfigurationsClient) getCoordinatorHandleResponse(resp *http.Response) (ConfigurationsClientGetCoordinatorResponse, error) {
	result := ConfigurationsClientGetCoordinatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerConfiguration); err != nil {
		return ConfigurationsClientGetCoordinatorResponse{}, err
	}
	return result, nil
}

// GetNode - Gets information of a configuration for worker nodes.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - configurationName - The name of the cluster configuration.
//   - options - ConfigurationsClientGetNodeOptions contains the optional parameters for the ConfigurationsClient.GetNode method.
func (client *ConfigurationsClient) GetNode(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, options *ConfigurationsClientGetNodeOptions) (ConfigurationsClientGetNodeResponse, error) {
	var err error
	req, err := client.getNodeCreateRequest(ctx, resourceGroupName, clusterName, configurationName, options)
	if err != nil {
		return ConfigurationsClientGetNodeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationsClientGetNodeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationsClientGetNodeResponse{}, err
	}
	resp, err := client.getNodeHandleResponse(httpResp)
	return resp, err
}

// getNodeCreateRequest creates the GetNode request.
func (client *ConfigurationsClient) getNodeCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, options *ConfigurationsClientGetNodeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/{clusterName}/nodeConfigurations/{configurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNodeHandleResponse handles the GetNode response.
func (client *ConfigurationsClient) getNodeHandleResponse(resp *http.Response) (ConfigurationsClientGetNodeResponse, error) {
	result := ConfigurationsClientGetNodeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerConfiguration); err != nil {
		return ConfigurationsClientGetNodeResponse{}, err
	}
	return result, nil
}

// NewListByClusterPager - List all the configurations of a cluster.
//
// Generated from API version 2022-11-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - options - ConfigurationsClientListByClusterOptions contains the optional parameters for the ConfigurationsClient.NewListByClusterPager
//     method.
func (client *ConfigurationsClient) NewListByClusterPager(resourceGroupName string, clusterName string, options *ConfigurationsClientListByClusterOptions) *runtime.Pager[ConfigurationsClientListByClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationsClientListByClusterResponse]{
		More: func(page ConfigurationsClientListByClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConfigurationsClientListByClusterResponse) (ConfigurationsClientListByClusterResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConfigurationsClientListByClusterResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationsClientListByClusterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationsClientListByClusterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByClusterHandleResponse(resp)
		},
	})
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *ConfigurationsClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ConfigurationsClientListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/{clusterName}/configurations"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *ConfigurationsClient) listByClusterHandleResponse(resp *http.Response) (ConfigurationsClientListByClusterResponse, error) {
	result := ConfigurationsClientListByClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterConfigurationListResult); err != nil {
		return ConfigurationsClientListByClusterResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - List all the configurations of a server in cluster.
//
// Generated from API version 2022-11-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - serverName - The name of the server.
//   - options - ConfigurationsClientListByServerOptions contains the optional parameters for the ConfigurationsClient.NewListByServerPager
//     method.
func (client *ConfigurationsClient) NewListByServerPager(resourceGroupName string, clusterName string, serverName string, options *ConfigurationsClientListByServerOptions) *runtime.Pager[ConfigurationsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationsClientListByServerResponse]{
		More: func(page ConfigurationsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConfigurationsClientListByServerResponse) (ConfigurationsClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, clusterName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConfigurationsClientListByServerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationsClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationsClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ConfigurationsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, serverName string, options *ConfigurationsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/{clusterName}/servers/{serverName}/configurations"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ConfigurationsClient) listByServerHandleResponse(resp *http.Response) (ConfigurationsClientListByServerResponse, error) {
	result := ConfigurationsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerConfigurationListResult); err != nil {
		return ConfigurationsClientListByServerResponse{}, err
	}
	return result, nil
}

// BeginUpdateOnCoordinator - Updates configuration of coordinator in a cluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - configurationName - The name of the cluster configuration.
//   - parameters - The required parameters for updating a cluster configuration.
//   - options - ConfigurationsClientBeginUpdateOnCoordinatorOptions contains the optional parameters for the ConfigurationsClient.BeginUpdateOnCoordinator
//     method.
func (client *ConfigurationsClient) BeginUpdateOnCoordinator(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, parameters ServerConfiguration, options *ConfigurationsClientBeginUpdateOnCoordinatorOptions) (*runtime.Poller[ConfigurationsClientUpdateOnCoordinatorResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateOnCoordinator(ctx, resourceGroupName, clusterName, configurationName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConfigurationsClientUpdateOnCoordinatorResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ConfigurationsClientUpdateOnCoordinatorResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateOnCoordinator - Updates configuration of coordinator in a cluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-08
func (client *ConfigurationsClient) updateOnCoordinator(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, parameters ServerConfiguration, options *ConfigurationsClientBeginUpdateOnCoordinatorOptions) (*http.Response, error) {
	var err error
	req, err := client.updateOnCoordinatorCreateRequest(ctx, resourceGroupName, clusterName, configurationName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateOnCoordinatorCreateRequest creates the UpdateOnCoordinator request.
func (client *ConfigurationsClient) updateOnCoordinatorCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, parameters ServerConfiguration, options *ConfigurationsClientBeginUpdateOnCoordinatorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/{clusterName}/coordinatorConfigurations/{configurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdateOnNode - Updates configuration of worker nodes in a cluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - configurationName - The name of the cluster configuration.
//   - parameters - The required parameters for updating a cluster configuration.
//   - options - ConfigurationsClientBeginUpdateOnNodeOptions contains the optional parameters for the ConfigurationsClient.BeginUpdateOnNode
//     method.
func (client *ConfigurationsClient) BeginUpdateOnNode(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, parameters ServerConfiguration, options *ConfigurationsClientBeginUpdateOnNodeOptions) (*runtime.Poller[ConfigurationsClientUpdateOnNodeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateOnNode(ctx, resourceGroupName, clusterName, configurationName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConfigurationsClientUpdateOnNodeResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ConfigurationsClientUpdateOnNodeResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateOnNode - Updates configuration of worker nodes in a cluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-08
func (client *ConfigurationsClient) updateOnNode(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, parameters ServerConfiguration, options *ConfigurationsClientBeginUpdateOnNodeOptions) (*http.Response, error) {
	var err error
	req, err := client.updateOnNodeCreateRequest(ctx, resourceGroupName, clusterName, configurationName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateOnNodeCreateRequest creates the UpdateOnNode request.
func (client *ConfigurationsClient) updateOnNodeCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, parameters ServerConfiguration, options *ConfigurationsClientBeginUpdateOnNodeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/{clusterName}/nodeConfigurations/{configurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
