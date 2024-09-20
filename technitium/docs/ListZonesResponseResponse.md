# ListZonesResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**TotalZones** | Pointer to **int32** |  | [optional] 
**Zones** | Pointer to [**[]ListZonesResponseResponseZonesInner**](ListZonesResponseResponseZonesInner.md) |  | [optional] 

## Methods

### NewListZonesResponseResponse

`func NewListZonesResponseResponse() *ListZonesResponseResponse`

NewListZonesResponseResponse instantiates a new ListZonesResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListZonesResponseResponseWithDefaults

`func NewListZonesResponseResponseWithDefaults() *ListZonesResponseResponse`

NewListZonesResponseResponseWithDefaults instantiates a new ListZonesResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *ListZonesResponseResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ListZonesResponseResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ListZonesResponseResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ListZonesResponseResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetTotalPages

`func (o *ListZonesResponseResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ListZonesResponseResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ListZonesResponseResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *ListZonesResponseResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalZones

`func (o *ListZonesResponseResponse) GetTotalZones() int32`

GetTotalZones returns the TotalZones field if non-nil, zero value otherwise.

### GetTotalZonesOk

`func (o *ListZonesResponseResponse) GetTotalZonesOk() (*int32, bool)`

GetTotalZonesOk returns a tuple with the TotalZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalZones

`func (o *ListZonesResponseResponse) SetTotalZones(v int32)`

SetTotalZones sets TotalZones field to given value.

### HasTotalZones

`func (o *ListZonesResponseResponse) HasTotalZones() bool`

HasTotalZones returns a boolean if a field has been set.

### GetZones

`func (o *ListZonesResponseResponse) GetZones() []ListZonesResponseResponseZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ListZonesResponseResponse) GetZonesOk() (*[]ListZonesResponseResponseZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ListZonesResponseResponse) SetZones(v []ListZonesResponseResponseZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ListZonesResponseResponse) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


