//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// KustoPoolDataConnectionsClient contains the methods for the KustoPoolDataConnections group.
// Don't use this type directly, use NewKustoPoolDataConnectionsClient() instead.
type KustoPoolDataConnectionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewKustoPoolDataConnectionsClient creates a new instance of KustoPoolDataConnectionsClient with the specified values.
func NewKustoPoolDataConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *KustoPoolDataConnectionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &KustoPoolDataConnectionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CheckNameAvailability - Checks that the data connection name is valid and is not already in use.
// If the operation fails it returns the *ErrorResponse error type.
func (client *KustoPoolDataConnectionsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName DataConnectionCheckNameRequest, options *KustoPoolDataConnectionsCheckNameAvailabilityOptions) (KustoPoolDataConnectionsCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, options)
	if err != nil {
		return KustoPoolDataConnectionsCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return KustoPoolDataConnectionsCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KustoPoolDataConnectionsCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *KustoPoolDataConnectionsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName DataConnectionCheckNameRequest, options *KustoPoolDataConnectionsCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dataConnectionName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *KustoPoolDataConnectionsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (KustoPoolDataConnectionsCheckNameAvailabilityResponse, error) {
	result := KustoPoolDataConnectionsCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return KustoPoolDataConnectionsCheckNameAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *KustoPoolDataConnectionsClient) checkNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreateOrUpdate - Creates or updates a data connection.
// If the operation fails it returns the *ErrorResponse error type.
func (client *KustoPoolDataConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *KustoPoolDataConnectionsBeginCreateOrUpdateOptions) (KustoPoolDataConnectionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, parameters, options)
	if err != nil {
		return KustoPoolDataConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result := KustoPoolDataConnectionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("KustoPoolDataConnectionsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return KustoPoolDataConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &KustoPoolDataConnectionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a data connection.
// If the operation fails it returns the *ErrorResponse error type.
func (client *KustoPoolDataConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *KustoPoolDataConnectionsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *KustoPoolDataConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *KustoPoolDataConnectionsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/dataConnections/{dataConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataConnectionName == "" {
		return nil, errors.New("parameter dataConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectionName}", url.PathEscape(dataConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *KustoPoolDataConnectionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDataConnectionValidation - Checks that the data connection parameters are valid.
// If the operation fails it returns the *ErrorResponse error type.
func (client *KustoPoolDataConnectionsClient) BeginDataConnectionValidation(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, parameters DataConnectionValidation, options *KustoPoolDataConnectionsBeginDataConnectionValidationOptions) (KustoPoolDataConnectionsDataConnectionValidationPollerResponse, error) {
	resp, err := client.dataConnectionValidation(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, parameters, options)
	if err != nil {
		return KustoPoolDataConnectionsDataConnectionValidationPollerResponse{}, err
	}
	result := KustoPoolDataConnectionsDataConnectionValidationPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("KustoPoolDataConnectionsClient.DataConnectionValidation", "location", resp, client.pl, client.dataConnectionValidationHandleError)
	if err != nil {
		return KustoPoolDataConnectionsDataConnectionValidationPollerResponse{}, err
	}
	result.Poller = &KustoPoolDataConnectionsDataConnectionValidationPoller{
		pt: pt,
	}
	return result, nil
}

// DataConnectionValidation - Checks that the data connection parameters are valid.
// If the operation fails it returns the *ErrorResponse error type.
func (client *KustoPoolDataConnectionsClient) dataConnectionValidation(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, parameters DataConnectionValidation, options *KustoPoolDataConnectionsBeginDataConnectionValidationOptions) (*http.Response, error) {
	req, err := client.dataConnectionValidationCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.dataConnectionValidationHandleError(resp)
	}
	return resp, nil
}

// dataConnectionValidationCreateRequest creates the DataConnectionValidation request.
func (client *KustoPoolDataConnectionsClient) dataConnectionValidationCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, parameters DataConnectionValidation, options *KustoPoolDataConnectionsBeginDataConnectionValidationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/dataConnectionValidation"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// dataConnectionValidationHandleError handles the DataConnectionValidation error response.
func (client *KustoPoolDataConnectionsClient) dataConnectionValidationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the data connection with the given name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *KustoPoolDataConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, options *KustoPoolDataConnectionsBeginDeleteOptions) (KustoPoolDataConnectionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, options)
	if err != nil {
		return KustoPoolDataConnectionsDeletePollerResponse{}, err
	}
	result := KustoPoolDataConnectionsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("KustoPoolDataConnectionsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return KustoPoolDataConnectionsDeletePollerResponse{}, err
	}
	result.Poller = &KustoPoolDataConnectionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the data connection with the given name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *KustoPoolDataConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, options *KustoPoolDataConnectionsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *KustoPoolDataConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, options *KustoPoolDataConnectionsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/dataConnections/{dataConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataConnectionName == "" {
		return nil, errors.New("parameter dataConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectionName}", url.PathEscape(dataConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *KustoPoolDataConnectionsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Returns a data connection.
// If the operation fails it returns the *ErrorResponse error type.
func (client *KustoPoolDataConnectionsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, options *KustoPoolDataConnectionsGetOptions) (KustoPoolDataConnectionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, options)
	if err != nil {
		return KustoPoolDataConnectionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return KustoPoolDataConnectionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KustoPoolDataConnectionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *KustoPoolDataConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, options *KustoPoolDataConnectionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/dataConnections/{dataConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataConnectionName == "" {
		return nil, errors.New("parameter dataConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectionName}", url.PathEscape(dataConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *KustoPoolDataConnectionsClient) getHandleResponse(resp *http.Response) (KustoPoolDataConnectionsGetResponse, error) {
	result := KustoPoolDataConnectionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return KustoPoolDataConnectionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *KustoPoolDataConnectionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByDatabase - Returns the list of data connections of the given Kusto pool database.
// If the operation fails it returns the *ErrorResponse error type.
func (client *KustoPoolDataConnectionsClient) ListByDatabase(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, options *KustoPoolDataConnectionsListByDatabaseOptions) (KustoPoolDataConnectionsListByDatabaseResponse, error) {
	req, err := client.listByDatabaseCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, options)
	if err != nil {
		return KustoPoolDataConnectionsListByDatabaseResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return KustoPoolDataConnectionsListByDatabaseResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KustoPoolDataConnectionsListByDatabaseResponse{}, client.listByDatabaseHandleError(resp)
	}
	return client.listByDatabaseHandleResponse(resp)
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *KustoPoolDataConnectionsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, options *KustoPoolDataConnectionsListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/dataConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *KustoPoolDataConnectionsClient) listByDatabaseHandleResponse(resp *http.Response) (KustoPoolDataConnectionsListByDatabaseResponse, error) {
	result := KustoPoolDataConnectionsListByDatabaseResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataConnectionListResult); err != nil {
		return KustoPoolDataConnectionsListByDatabaseResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByDatabaseHandleError handles the ListByDatabase error response.
func (client *KustoPoolDataConnectionsClient) listByDatabaseHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates a data connection.
// If the operation fails it returns the *ErrorResponse error type.
func (client *KustoPoolDataConnectionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *KustoPoolDataConnectionsBeginUpdateOptions) (KustoPoolDataConnectionsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, parameters, options)
	if err != nil {
		return KustoPoolDataConnectionsUpdatePollerResponse{}, err
	}
	result := KustoPoolDataConnectionsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("KustoPoolDataConnectionsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return KustoPoolDataConnectionsUpdatePollerResponse{}, err
	}
	result.Poller = &KustoPoolDataConnectionsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a data connection.
// If the operation fails it returns the *ErrorResponse error type.
func (client *KustoPoolDataConnectionsClient) update(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *KustoPoolDataConnectionsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *KustoPoolDataConnectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *KustoPoolDataConnectionsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/dataConnections/{dataConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataConnectionName == "" {
		return nil, errors.New("parameter dataConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectionName}", url.PathEscape(dataConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *KustoPoolDataConnectionsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}