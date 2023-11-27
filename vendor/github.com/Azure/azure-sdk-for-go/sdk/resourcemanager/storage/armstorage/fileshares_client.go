//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

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

// FileSharesClient contains the methods for the FileShares group.
// Don't use this type directly, use NewFileSharesClient() instead.
type FileSharesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFileSharesClient creates a new instance of FileSharesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFileSharesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FileSharesClient, error) {
	cl, err := arm.NewClient(moduleName+".FileSharesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FileSharesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates a new share under the specified account as described by request body. The share resource includes metadata
// and properties for that share. It does not include a list of the files contained by
// the share.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - shareName - The name of the file share within the specified storage account. File share names must be between 3 and 63
//     characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-)
//     character must be immediately preceded and followed by a letter or number.
//   - fileShare - Properties of the file share to create.
//   - options - FileSharesClientCreateOptions contains the optional parameters for the FileSharesClient.Create method.
func (client *FileSharesClient) Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare FileShare, options *FileSharesClientCreateOptions) (FileSharesClientCreateResponse, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, shareName, fileShare, options)
	if err != nil {
		return FileSharesClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FileSharesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return FileSharesClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *FileSharesClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare FileShare, options *FileSharesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, fileShare); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *FileSharesClient) createHandleResponse(resp *http.Response) (FileSharesClientCreateResponse, error) {
	result := FileSharesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileShare); err != nil {
		return FileSharesClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specified share under its account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - shareName - The name of the file share within the specified storage account. File share names must be between 3 and 63
//     characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-)
//     character must be immediately preceded and followed by a letter or number.
//   - options - FileSharesClientDeleteOptions contains the optional parameters for the FileSharesClient.Delete method.
func (client *FileSharesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *FileSharesClientDeleteOptions) (FileSharesClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
	if err != nil {
		return FileSharesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FileSharesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return FileSharesClientDeleteResponse{}, err
	}
	return FileSharesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FileSharesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *FileSharesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	if options != nil && options.Include != nil {
		reqQP.Set("$include", *options.Include)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.XMSSnapshot != nil {
		req.Raw().Header["x-ms-snapshot"] = []string{*options.XMSSnapshot}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets properties of a specified share.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - shareName - The name of the file share within the specified storage account. File share names must be between 3 and 63
//     characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-)
//     character must be immediately preceded and followed by a letter or number.
//   - options - FileSharesClientGetOptions contains the optional parameters for the FileSharesClient.Get method.
func (client *FileSharesClient) Get(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *FileSharesClientGetOptions) (FileSharesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
	if err != nil {
		return FileSharesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FileSharesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FileSharesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FileSharesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *FileSharesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.XMSSnapshot != nil {
		req.Raw().Header["x-ms-snapshot"] = []string{*options.XMSSnapshot}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FileSharesClient) getHandleResponse(resp *http.Response) (FileSharesClientGetResponse, error) {
	result := FileSharesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileShare); err != nil {
		return FileSharesClientGetResponse{}, err
	}
	return result, nil
}

// Lease - The Lease Share operation establishes and manages a lock on a share for delete operations. The lock duration can
// be 15 to 60 seconds, or can be infinite.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - shareName - The name of the file share within the specified storage account. File share names must be between 3 and 63
//     characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-)
//     character must be immediately preceded and followed by a letter or number.
//   - options - FileSharesClientLeaseOptions contains the optional parameters for the FileSharesClient.Lease method.
func (client *FileSharesClient) Lease(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *FileSharesClientLeaseOptions) (FileSharesClientLeaseResponse, error) {
	var err error
	req, err := client.leaseCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
	if err != nil {
		return FileSharesClientLeaseResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FileSharesClientLeaseResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FileSharesClientLeaseResponse{}, err
	}
	resp, err := client.leaseHandleResponse(httpResp)
	return resp, err
}

// leaseCreateRequest creates the Lease request.
func (client *FileSharesClient) leaseCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *FileSharesClientLeaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}/lease"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.XMSSnapshot != nil {
		req.Raw().Header["x-ms-snapshot"] = []string{*options.XMSSnapshot}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.Parameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// leaseHandleResponse handles the Lease response.
func (client *FileSharesClient) leaseHandleResponse(resp *http.Response) (FileSharesClientLeaseResponse, error) {
	result := FileSharesClientLeaseResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.LeaseShareResponse); err != nil {
		return FileSharesClientLeaseResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all shares.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - options - FileSharesClientListOptions contains the optional parameters for the FileSharesClient.NewListPager method.
func (client *FileSharesClient) NewListPager(resourceGroupName string, accountName string, options *FileSharesClientListOptions) *runtime.Pager[FileSharesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FileSharesClientListResponse]{
		More: func(page FileSharesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FileSharesClientListResponse) (FileSharesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FileSharesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return FileSharesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FileSharesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FileSharesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *FileSharesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	if options != nil && options.Maxpagesize != nil {
		reqQP.Set("$maxpagesize", *options.Maxpagesize)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FileSharesClient) listHandleResponse(resp *http.Response) (FileSharesClientListResponse, error) {
	result := FileSharesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileShareItems); err != nil {
		return FileSharesClientListResponse{}, err
	}
	return result, nil
}

// Restore - Restore a file share within a valid retention days if share soft delete is enabled
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - shareName - The name of the file share within the specified storage account. File share names must be between 3 and 63
//     characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-)
//     character must be immediately preceded and followed by a letter or number.
//   - options - FileSharesClientRestoreOptions contains the optional parameters for the FileSharesClient.Restore method.
func (client *FileSharesClient) Restore(ctx context.Context, resourceGroupName string, accountName string, shareName string, deletedShare DeletedShare, options *FileSharesClientRestoreOptions) (FileSharesClientRestoreResponse, error) {
	var err error
	req, err := client.restoreCreateRequest(ctx, resourceGroupName, accountName, shareName, deletedShare, options)
	if err != nil {
		return FileSharesClientRestoreResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FileSharesClientRestoreResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FileSharesClientRestoreResponse{}, err
	}
	return FileSharesClientRestoreResponse{}, nil
}

// restoreCreateRequest creates the Restore request.
func (client *FileSharesClient) restoreCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, deletedShare DeletedShare, options *FileSharesClientRestoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}/restore"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, deletedShare); err != nil {
		return nil, err
	}
	return req, nil
}

// Update - Updates share properties as specified in request body. Properties not mentioned in the request will not be changed.
// Update fails if the specified share does not already exist.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - shareName - The name of the file share within the specified storage account. File share names must be between 3 and 63
//     characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-)
//     character must be immediately preceded and followed by a letter or number.
//   - fileShare - Properties to update for the file share.
//   - options - FileSharesClientUpdateOptions contains the optional parameters for the FileSharesClient.Update method.
func (client *FileSharesClient) Update(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare FileShare, options *FileSharesClientUpdateOptions) (FileSharesClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, shareName, fileShare, options)
	if err != nil {
		return FileSharesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FileSharesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FileSharesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *FileSharesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare FileShare, options *FileSharesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, fileShare); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *FileSharesClient) updateHandleResponse(resp *http.Response) (FileSharesClientUpdateResponse, error) {
	result := FileSharesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileShare); err != nil {
		return FileSharesClientUpdateResponse{}, err
	}
	return result, nil
}
