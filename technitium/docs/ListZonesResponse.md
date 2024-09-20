# ListZonesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**ListZonesResponseResponse**](ListZonesResponseResponse.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewListZonesResponse

`func NewListZonesResponse() *ListZonesResponse`

NewListZonesResponse instantiates a new ListZonesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListZonesResponseWithDefaults

`func NewListZonesResponseWithDefaults() *ListZonesResponse`

NewListZonesResponseWithDefaults instantiates a new ListZonesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *ListZonesResponse) GetResponse() ListZonesResponseResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ListZonesResponse) GetResponseOk() (*ListZonesResponseResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ListZonesResponse) SetResponse(v ListZonesResponseResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ListZonesResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatus

`func (o *ListZonesResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListZonesResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListZonesResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListZonesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ListZonesResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListZonesResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListZonesResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListZonesResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


