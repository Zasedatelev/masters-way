/*
Masters way general API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// FavoriteUserAPIService FavoriteUserAPI service
type FavoriteUserAPIService service

type ApiCreateFavoriteUserRequest struct {
	ctx context.Context
	ApiService *FavoriteUserAPIService
	request *SchemasCreateFavoriteUserPayload
}

// query params
func (r ApiCreateFavoriteUserRequest) Request(request SchemasCreateFavoriteUserPayload) ApiCreateFavoriteUserRequest {
	r.request = &request
	return r
}

func (r ApiCreateFavoriteUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateFavoriteUserExecute(r)
}

/*
CreateFavoriteUser Create a new favorite user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFavoriteUserRequest
*/
func (a *FavoriteUserAPIService) CreateFavoriteUser(ctx context.Context) ApiCreateFavoriteUserRequest {
	return ApiCreateFavoriteUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *FavoriteUserAPIService) CreateFavoriteUserExecute(r ApiCreateFavoriteUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FavoriteUserAPIService.CreateFavoriteUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/favoriteUsers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.request == nil {
		return nil, reportError("request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.request
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteFavoriteUserRequest struct {
	ctx context.Context
	ApiService *FavoriteUserAPIService
	donorUserUuid string
	acceptorUserUuid string
}

func (r ApiDeleteFavoriteUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteFavoriteUserExecute(r)
}

/*
DeleteFavoriteUser Delete favoriteUser by UUID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param donorUserUuid donorUser UUID
 @param acceptorUserUuid acceptorUser UUID
 @return ApiDeleteFavoriteUserRequest
*/
func (a *FavoriteUserAPIService) DeleteFavoriteUser(ctx context.Context, donorUserUuid string, acceptorUserUuid string) ApiDeleteFavoriteUserRequest {
	return ApiDeleteFavoriteUserRequest{
		ApiService: a,
		ctx: ctx,
		donorUserUuid: donorUserUuid,
		acceptorUserUuid: acceptorUserUuid,
	}
}

// Execute executes the request
func (a *FavoriteUserAPIService) DeleteFavoriteUserExecute(r ApiDeleteFavoriteUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FavoriteUserAPIService.DeleteFavoriteUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/favoriteUsers/{donorUserUuid}/{acceptorUserUuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"donorUserUuid"+"}", url.PathEscape(parameterValueToString(r.donorUserUuid, "donorUserUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"acceptorUserUuid"+"}", url.PathEscape(parameterValueToString(r.acceptorUserUuid, "acceptorUserUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
