//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice

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

// GlobalClient contains the methods for the Global group.
// Don't use this type directly, use NewGlobalClient() instead.
type GlobalClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGlobalClient creates a new instance of GlobalClient with the specified values.
// subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGlobalClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *GlobalClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &GlobalClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetDeletedWebApp - Description for Get deleted app for a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// deletedSiteID - The numeric ID of the deleted app, e.g. 12345
// options - GlobalClientGetDeletedWebAppOptions contains the optional parameters for the GlobalClient.GetDeletedWebApp method.
func (client *GlobalClient) GetDeletedWebApp(ctx context.Context, deletedSiteID string, options *GlobalClientGetDeletedWebAppOptions) (GlobalClientGetDeletedWebAppResponse, error) {
	req, err := client.getDeletedWebAppCreateRequest(ctx, deletedSiteID, options)
	if err != nil {
		return GlobalClientGetDeletedWebAppResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GlobalClientGetDeletedWebAppResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GlobalClientGetDeletedWebAppResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDeletedWebAppHandleResponse(resp)
}

// getDeletedWebAppCreateRequest creates the GetDeletedWebApp request.
func (client *GlobalClient) getDeletedWebAppCreateRequest(ctx context.Context, deletedSiteID string, options *GlobalClientGetDeletedWebAppOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites/{deletedSiteId}"
	if deletedSiteID == "" {
		return nil, errors.New("parameter deletedSiteID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deletedSiteId}", url.PathEscape(deletedSiteID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDeletedWebAppHandleResponse handles the GetDeletedWebApp response.
func (client *GlobalClient) getDeletedWebAppHandleResponse(resp *http.Response) (GlobalClientGetDeletedWebAppResponse, error) {
	result := GlobalClientGetDeletedWebAppResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedSite); err != nil {
		return GlobalClientGetDeletedWebAppResponse{}, err
	}
	return result, nil
}

// GetDeletedWebAppSnapshots - Description for Get all deleted apps for a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// deletedSiteID - The numeric ID of the deleted app, e.g. 12345
// options - GlobalClientGetDeletedWebAppSnapshotsOptions contains the optional parameters for the GlobalClient.GetDeletedWebAppSnapshots
// method.
func (client *GlobalClient) GetDeletedWebAppSnapshots(ctx context.Context, deletedSiteID string, options *GlobalClientGetDeletedWebAppSnapshotsOptions) (GlobalClientGetDeletedWebAppSnapshotsResponse, error) {
	req, err := client.getDeletedWebAppSnapshotsCreateRequest(ctx, deletedSiteID, options)
	if err != nil {
		return GlobalClientGetDeletedWebAppSnapshotsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GlobalClientGetDeletedWebAppSnapshotsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GlobalClientGetDeletedWebAppSnapshotsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDeletedWebAppSnapshotsHandleResponse(resp)
}

// getDeletedWebAppSnapshotsCreateRequest creates the GetDeletedWebAppSnapshots request.
func (client *GlobalClient) getDeletedWebAppSnapshotsCreateRequest(ctx context.Context, deletedSiteID string, options *GlobalClientGetDeletedWebAppSnapshotsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites/{deletedSiteId}/snapshots"
	if deletedSiteID == "" {
		return nil, errors.New("parameter deletedSiteID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deletedSiteId}", url.PathEscape(deletedSiteID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDeletedWebAppSnapshotsHandleResponse handles the GetDeletedWebAppSnapshots response.
func (client *GlobalClient) getDeletedWebAppSnapshotsHandleResponse(resp *http.Response) (GlobalClientGetDeletedWebAppSnapshotsResponse, error) {
	result := GlobalClientGetDeletedWebAppSnapshotsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotArray); err != nil {
		return GlobalClientGetDeletedWebAppSnapshotsResponse{}, err
	}
	return result, nil
}

// GetSubscriptionOperationWithAsyncResponse - Description for Gets an operation in a subscription and given region
// If the operation fails it returns an *azcore.ResponseError type.
// location - Location name
// operationID - Operation Id
// options - GlobalClientGetSubscriptionOperationWithAsyncResponseOptions contains the optional parameters for the GlobalClient.GetSubscriptionOperationWithAsyncResponse
// method.
func (client *GlobalClient) GetSubscriptionOperationWithAsyncResponse(ctx context.Context, location string, operationID string, options *GlobalClientGetSubscriptionOperationWithAsyncResponseOptions) (GlobalClientGetSubscriptionOperationWithAsyncResponseResponse, error) {
	req, err := client.getSubscriptionOperationWithAsyncResponseCreateRequest(ctx, location, operationID, options)
	if err != nil {
		return GlobalClientGetSubscriptionOperationWithAsyncResponseResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GlobalClientGetSubscriptionOperationWithAsyncResponseResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return GlobalClientGetSubscriptionOperationWithAsyncResponseResponse{}, runtime.NewResponseError(resp)
	}
	return GlobalClientGetSubscriptionOperationWithAsyncResponseResponse{RawResponse: resp}, nil
}

// getSubscriptionOperationWithAsyncResponseCreateRequest creates the GetSubscriptionOperationWithAsyncResponse request.
func (client *GlobalClient) getSubscriptionOperationWithAsyncResponseCreateRequest(ctx context.Context, location string, operationID string, options *GlobalClientGetSubscriptionOperationWithAsyncResponseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/operations/{operationId}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}