package azblob

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"context"
	"encoding/xml"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// serviceClient is the client for the Service methods of the Azblob service.
type serviceClient struct {
	managementClient
}

// newServiceClient creates an instance of the serviceClient client.
func newServiceClient(url url.URL, p pipeline.Pipeline) serviceClient {
	return serviceClient{newManagementClient(url, p)}
}

// GetAccountInfo returns the sku name and account kind
func (client serviceClient) GetAccountInfo(ctx context.Context) (*ServiceGetAccountInfoResponse, error) {
	req, err := client.getAccountInfoPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getAccountInfoResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ServiceGetAccountInfoResponse), err
}

// getAccountInfoPreparer prepares the GetAccountInfo request.
func (client serviceClient) getAccountInfoPreparer() (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("restype", "account")
	params.Set("comp", "properties")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getAccountInfoResponder handles the response to the GetAccountInfo request.
func (client serviceClient) getAccountInfoResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &ServiceGetAccountInfoResponse{rawResponse: resp.Response()}, err
}

// GetProperties gets the properties of a storage account's Blob service, including properties for Storage Analytics
// and CORS (Cross-Origin Resource Sharing) rules.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> requestID is provides a client-generated, opaque value with a 1 KB
// character limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client serviceClient) GetProperties(ctx context.Context, timeout *int32, requestID *string) (*StorageServiceProperties, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPropertiesPreparer(timeout, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getPropertiesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*StorageServiceProperties), err
}

// getPropertiesPreparer prepares the GetProperties request.
func (client serviceClient) getPropertiesPreparer(timeout *int32, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "service")
	params.Set("comp", "properties")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	return req, nil
}

// getPropertiesResponder handles the response to the GetProperties request.
func (client serviceClient) getPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &StorageServiceProperties{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, err
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetStatistics retrieves statistics related to replication for the Blob service. It is only available on the
// secondary location endpoint when read-access geo-redundant replication is enabled for the storage account.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> requestID is provides a client-generated, opaque value with a 1 KB
// character limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client serviceClient) GetStatistics(ctx context.Context, timeout *int32, requestID *string) (*StorageServiceStats, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getStatisticsPreparer(timeout, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getStatisticsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*StorageServiceStats), err
}

// getStatisticsPreparer prepares the GetStatistics request.
func (client serviceClient) getStatisticsPreparer(timeout *int32, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "service")
	params.Set("comp", "stats")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	return req, nil
}

// getStatisticsResponder handles the response to the GetStatistics request.
func (client serviceClient) getStatisticsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &StorageServiceStats{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, err
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetUserDelegationKey retrieves a user delegation key for the Blob service. This is only a valid operation when using
// bearer token authentication.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> requestID is provides a client-generated, opaque value with a 1 KB
// character limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client serviceClient) GetUserDelegationKey(ctx context.Context, keyInfo KeyInfo, timeout *int32, requestID *string) (*UserDelegationKey, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getUserDelegationKeyPreparer(keyInfo, timeout, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getUserDelegationKeyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*UserDelegationKey), err
}

// getUserDelegationKeyPreparer prepares the GetUserDelegationKey request.
func (client serviceClient) getUserDelegationKeyPreparer(keyInfo KeyInfo, timeout *int32, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("POST", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "service")
	params.Set("comp", "userdelegationkey")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	b, err := xml.Marshal(keyInfo)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/xml")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// getUserDelegationKeyResponder handles the response to the GetUserDelegationKey request.
func (client serviceClient) getUserDelegationKeyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &UserDelegationKey{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, err
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListContainersSegment the List Containers Segment operation returns a list of the containers under the specified
// account
//
// prefix is filters the results to return only containers whose name begins with the specified prefix. marker is a
// string value that identifies the portion of the list of containers to be returned with the next listing operation.
// The operation returns the NextMarker value within the response body if the listing operation did not return all
// containers remaining to be listed with the current page. The NextMarker value can be used as the value for the
// marker parameter in a subsequent call to request the next page of list items. The marker value is opaque to the
// client. maxresults is specifies the maximum number of containers to return. If the request does not specify
// maxresults, or specifies a value greater than 5000, the server will return up to 5000 items. Note that if the
// listing operation crosses a partition boundary, then the service will return a continuation token for retrieving the
// remainder of the results. For this reason, it is possible that the service will return fewer results than specified
// by maxresults, or than the default of 5000. include is include this parameter to specify that the container's
// metadata be returned as part of the response body. timeout is the timeout parameter is expressed in seconds. For
// more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> requestID is provides a client-generated, opaque value with a 1 KB
// character limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client serviceClient) ListContainersSegment(ctx context.Context, prefix *string, marker *string, maxresults *int32, include ListContainersIncludeType, timeout *int32, requestID *string) (*ListContainersSegmentResponse, error) {
	if err := validate([]validation{
		{targetValue: maxresults,
			constraints: []constraint{{target: "maxresults", name: null, rule: false,
				chain: []constraint{{target: "maxresults", name: inclusiveMinimum, rule: 1, chain: nil}}}}},
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.listContainersSegmentPreparer(prefix, marker, maxresults, include, timeout, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listContainersSegmentResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ListContainersSegmentResponse), err
}

// listContainersSegmentPreparer prepares the ListContainersSegment request.
func (client serviceClient) listContainersSegmentPreparer(prefix *string, marker *string, maxresults *int32, include ListContainersIncludeType, timeout *int32, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if prefix != nil && len(*prefix) > 0 {
		params.Set("prefix", *prefix)
	}
	if marker != nil && len(*marker) > 0 {
		params.Set("marker", *marker)
	}
	if maxresults != nil {
		params.Set("maxresults", strconv.FormatInt(int64(*maxresults), 10))
	}
	if include != ListContainersIncludeNone {
		params.Set("include", string(include))
	}
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("comp", "list")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	return req, nil
}

// listContainersSegmentResponder handles the response to the ListContainersSegment request.
func (client serviceClient) listContainersSegmentResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ListContainersSegmentResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, err
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// SetProperties sets properties for a storage account's Blob service endpoint, including properties for Storage
// Analytics and CORS (Cross-Origin Resource Sharing) rules
//
// storageServiceProperties is the StorageService properties. timeout is the timeout parameter is expressed in seconds.
// For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> requestID is provides a client-generated, opaque value with a 1 KB
// character limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client serviceClient) SetProperties(ctx context.Context, storageServiceProperties StorageServiceProperties, timeout *int32, requestID *string) (*ServiceSetPropertiesResponse, error) {
	if err := validate([]validation{
		{targetValue: storageServiceProperties,
			constraints: []constraint{{target: "storageServiceProperties.Logging", name: null, rule: false,
				chain: []constraint{{target: "storageServiceProperties.Logging.RetentionPolicy", name: null, rule: true,
					chain: []constraint{{target: "storageServiceProperties.Logging.RetentionPolicy.Days", name: null, rule: false,
						chain: []constraint{{target: "storageServiceProperties.Logging.RetentionPolicy.Days", name: inclusiveMinimum, rule: 1, chain: nil}}},
					}},
				}},
				{target: "storageServiceProperties.HourMetrics", name: null, rule: false,
					chain: []constraint{{target: "storageServiceProperties.HourMetrics.RetentionPolicy", name: null, rule: false,
						chain: []constraint{{target: "storageServiceProperties.HourMetrics.RetentionPolicy.Days", name: null, rule: false,
							chain: []constraint{{target: "storageServiceProperties.HourMetrics.RetentionPolicy.Days", name: inclusiveMinimum, rule: 1, chain: nil}}},
						}},
					}},
				{target: "storageServiceProperties.MinuteMetrics", name: null, rule: false,
					chain: []constraint{{target: "storageServiceProperties.MinuteMetrics.RetentionPolicy", name: null, rule: false,
						chain: []constraint{{target: "storageServiceProperties.MinuteMetrics.RetentionPolicy.Days", name: null, rule: false,
							chain: []constraint{{target: "storageServiceProperties.MinuteMetrics.RetentionPolicy.Days", name: inclusiveMinimum, rule: 1, chain: nil}}},
						}},
					}},
				{target: "storageServiceProperties.DeleteRetentionPolicy", name: null, rule: false,
					chain: []constraint{{target: "storageServiceProperties.DeleteRetentionPolicy.Days", name: null, rule: false,
						chain: []constraint{{target: "storageServiceProperties.DeleteRetentionPolicy.Days", name: inclusiveMinimum, rule: 1, chain: nil}}},
					}}}},
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setPropertiesPreparer(storageServiceProperties, timeout, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setPropertiesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ServiceSetPropertiesResponse), err
}

// setPropertiesPreparer prepares the SetProperties request.
func (client serviceClient) setPropertiesPreparer(storageServiceProperties StorageServiceProperties, timeout *int32, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "service")
	params.Set("comp", "properties")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	b, err := xml.Marshal(storageServiceProperties)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/xml")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// setPropertiesResponder handles the response to the SetProperties request.
func (client serviceClient) setPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &ServiceSetPropertiesResponse{rawResponse: resp.Response()}, err
}

// SubmitBatch the Batch operation allows multiple API calls to be embedded into a single HTTP request.
//
// body is initial data body will be closed upon successful return. Callers should ensure closure when receiving an
// error.contentLength is the length of the request. multipartContentType is required. The value of this header must be
// multipart/mixed with a batch boundary. Example header value: multipart/mixed; boundary=batch_<GUID> timeout is the
// timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> requestID is provides a client-generated, opaque value with a 1 KB
// character limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client serviceClient) SubmitBatch(ctx context.Context, body io.ReadSeeker, contentLength int64, multipartContentType string, timeout *int32, requestID *string) (*SubmitBatchResponse, error) {
	if err := validate([]validation{
		{targetValue: body,
			constraints: []constraint{{target: "body", name: null, rule: true, chain: nil}}},
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.submitBatchPreparer(body, contentLength, multipartContentType, timeout, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.submitBatchResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*SubmitBatchResponse), err
}

// submitBatchPreparer prepares the SubmitBatch request.
func (client serviceClient) submitBatchPreparer(body io.ReadSeeker, contentLength int64, multipartContentType string, timeout *int32, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("POST", client.url, body)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("comp", "batch")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	req.Header.Set("Content-Type", multipartContentType)
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	return req, nil
}

// submitBatchResponder handles the response to the SubmitBatch request.
func (client serviceClient) submitBatchResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	return &SubmitBatchResponse{rawResponse: resp.Response()}, err
}
