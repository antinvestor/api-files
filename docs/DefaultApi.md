# \DefaultApi

All URIs are relative to *http://api.antinvestor.com/files*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFile**](DefaultApi.md#AddFile) | **Post** / | 
[**DeleteFile**](DefaultApi.md#DeleteFile) | **Delete** /{id} | 
[**FindFileById**](DefaultApi.md#FindFileById) | **Get** /{id} | 
[**FindFiles**](DefaultApi.md#FindFiles) | **Get** / | 



## AddFile

> File AddFile(ctx, optional)



Stores a new file. Duplicates are going to be ignored if the hash of the file and meta data both match an existing file

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

### Return type

[**File**](File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFile

> DeleteFile(ctx, id)



deletes a single file based on the ID supplied

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID of file to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindFileById

> *os.File FindFileById(ctx, id)



Returns a file based on supplied ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID of file to fetch | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindFiles

> []File FindFiles(ctx, optional)



Returns all files from the system that the user has access to 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindFilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **optional.String**| filters the files by the subscription by | 
 **groupId** | **optional.String**| filters the files by the grouping id | 
 **limit** | **optional.Int32**| maximum number of results to return | 

### Return type

[**[]File**](File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

