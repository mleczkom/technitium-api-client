# GetZoneOptionsResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Internal** | Pointer to **bool** |  | [optional] 
**DnssecStatus** | Pointer to **string** |  | [optional] 
**NotifyFailed** | Pointer to **bool** |  | [optional] 
**NotifyFailedFor** | Pointer to **[]string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**ZoneTransfer** | Pointer to **string** |  | [optional] 
**ZoneTransferNameServers** | Pointer to **[]string** |  | [optional] 
**ZoneTransferTsigKeyNames** | Pointer to **[]string** |  | [optional] 
**Notify** | Pointer to **string** |  | [optional] 
**NotifyNameServers** | Pointer to **[]string** |  | [optional] 
**Update** | Pointer to **string** |  | [optional] 
**UpdateIpAddresses** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetZoneOptionsResponseResponse

`func NewGetZoneOptionsResponseResponse() *GetZoneOptionsResponseResponse`

NewGetZoneOptionsResponseResponse instantiates a new GetZoneOptionsResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneOptionsResponseResponseWithDefaults

`func NewGetZoneOptionsResponseResponseWithDefaults() *GetZoneOptionsResponseResponse`

NewGetZoneOptionsResponseResponseWithDefaults instantiates a new GetZoneOptionsResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetZoneOptionsResponseResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetZoneOptionsResponseResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetZoneOptionsResponseResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetZoneOptionsResponseResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetZoneOptionsResponseResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetZoneOptionsResponseResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetZoneOptionsResponseResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetZoneOptionsResponseResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInternal

`func (o *GetZoneOptionsResponseResponse) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *GetZoneOptionsResponseResponse) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *GetZoneOptionsResponseResponse) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *GetZoneOptionsResponseResponse) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetDnssecStatus

`func (o *GetZoneOptionsResponseResponse) GetDnssecStatus() string`

GetDnssecStatus returns the DnssecStatus field if non-nil, zero value otherwise.

### GetDnssecStatusOk

`func (o *GetZoneOptionsResponseResponse) GetDnssecStatusOk() (*string, bool)`

GetDnssecStatusOk returns a tuple with the DnssecStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecStatus

`func (o *GetZoneOptionsResponseResponse) SetDnssecStatus(v string)`

SetDnssecStatus sets DnssecStatus field to given value.

### HasDnssecStatus

`func (o *GetZoneOptionsResponseResponse) HasDnssecStatus() bool`

HasDnssecStatus returns a boolean if a field has been set.

### GetNotifyFailed

`func (o *GetZoneOptionsResponseResponse) GetNotifyFailed() bool`

GetNotifyFailed returns the NotifyFailed field if non-nil, zero value otherwise.

### GetNotifyFailedOk

`func (o *GetZoneOptionsResponseResponse) GetNotifyFailedOk() (*bool, bool)`

GetNotifyFailedOk returns a tuple with the NotifyFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFailed

`func (o *GetZoneOptionsResponseResponse) SetNotifyFailed(v bool)`

SetNotifyFailed sets NotifyFailed field to given value.

### HasNotifyFailed

`func (o *GetZoneOptionsResponseResponse) HasNotifyFailed() bool`

HasNotifyFailed returns a boolean if a field has been set.

### GetNotifyFailedFor

`func (o *GetZoneOptionsResponseResponse) GetNotifyFailedFor() []string`

GetNotifyFailedFor returns the NotifyFailedFor field if non-nil, zero value otherwise.

### GetNotifyFailedForOk

`func (o *GetZoneOptionsResponseResponse) GetNotifyFailedForOk() (*[]string, bool)`

GetNotifyFailedForOk returns a tuple with the NotifyFailedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFailedFor

`func (o *GetZoneOptionsResponseResponse) SetNotifyFailedFor(v []string)`

SetNotifyFailedFor sets NotifyFailedFor field to given value.

### HasNotifyFailedFor

`func (o *GetZoneOptionsResponseResponse) HasNotifyFailedFor() bool`

HasNotifyFailedFor returns a boolean if a field has been set.

### GetDisabled

`func (o *GetZoneOptionsResponseResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetZoneOptionsResponseResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetZoneOptionsResponseResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetZoneOptionsResponseResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetZoneTransfer

`func (o *GetZoneOptionsResponseResponse) GetZoneTransfer() string`

GetZoneTransfer returns the ZoneTransfer field if non-nil, zero value otherwise.

### GetZoneTransferOk

`func (o *GetZoneOptionsResponseResponse) GetZoneTransferOk() (*string, bool)`

GetZoneTransferOk returns a tuple with the ZoneTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTransfer

`func (o *GetZoneOptionsResponseResponse) SetZoneTransfer(v string)`

SetZoneTransfer sets ZoneTransfer field to given value.

### HasZoneTransfer

`func (o *GetZoneOptionsResponseResponse) HasZoneTransfer() bool`

HasZoneTransfer returns a boolean if a field has been set.

### GetZoneTransferNameServers

`func (o *GetZoneOptionsResponseResponse) GetZoneTransferNameServers() []string`

GetZoneTransferNameServers returns the ZoneTransferNameServers field if non-nil, zero value otherwise.

### GetZoneTransferNameServersOk

`func (o *GetZoneOptionsResponseResponse) GetZoneTransferNameServersOk() (*[]string, bool)`

GetZoneTransferNameServersOk returns a tuple with the ZoneTransferNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTransferNameServers

`func (o *GetZoneOptionsResponseResponse) SetZoneTransferNameServers(v []string)`

SetZoneTransferNameServers sets ZoneTransferNameServers field to given value.

### HasZoneTransferNameServers

`func (o *GetZoneOptionsResponseResponse) HasZoneTransferNameServers() bool`

HasZoneTransferNameServers returns a boolean if a field has been set.

### GetZoneTransferTsigKeyNames

`func (o *GetZoneOptionsResponseResponse) GetZoneTransferTsigKeyNames() []string`

GetZoneTransferTsigKeyNames returns the ZoneTransferTsigKeyNames field if non-nil, zero value otherwise.

### GetZoneTransferTsigKeyNamesOk

`func (o *GetZoneOptionsResponseResponse) GetZoneTransferTsigKeyNamesOk() (*[]string, bool)`

GetZoneTransferTsigKeyNamesOk returns a tuple with the ZoneTransferTsigKeyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTransferTsigKeyNames

`func (o *GetZoneOptionsResponseResponse) SetZoneTransferTsigKeyNames(v []string)`

SetZoneTransferTsigKeyNames sets ZoneTransferTsigKeyNames field to given value.

### HasZoneTransferTsigKeyNames

`func (o *GetZoneOptionsResponseResponse) HasZoneTransferTsigKeyNames() bool`

HasZoneTransferTsigKeyNames returns a boolean if a field has been set.

### GetNotify

`func (o *GetZoneOptionsResponseResponse) GetNotify() string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *GetZoneOptionsResponseResponse) GetNotifyOk() (*string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *GetZoneOptionsResponseResponse) SetNotify(v string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *GetZoneOptionsResponseResponse) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyNameServers

`func (o *GetZoneOptionsResponseResponse) GetNotifyNameServers() []string`

GetNotifyNameServers returns the NotifyNameServers field if non-nil, zero value otherwise.

### GetNotifyNameServersOk

`func (o *GetZoneOptionsResponseResponse) GetNotifyNameServersOk() (*[]string, bool)`

GetNotifyNameServersOk returns a tuple with the NotifyNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyNameServers

`func (o *GetZoneOptionsResponseResponse) SetNotifyNameServers(v []string)`

SetNotifyNameServers sets NotifyNameServers field to given value.

### HasNotifyNameServers

`func (o *GetZoneOptionsResponseResponse) HasNotifyNameServers() bool`

HasNotifyNameServers returns a boolean if a field has been set.

### GetUpdate

`func (o *GetZoneOptionsResponseResponse) GetUpdate() string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *GetZoneOptionsResponseResponse) GetUpdateOk() (*string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *GetZoneOptionsResponseResponse) SetUpdate(v string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *GetZoneOptionsResponseResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUpdateIpAddresses

`func (o *GetZoneOptionsResponseResponse) GetUpdateIpAddresses() []string`

GetUpdateIpAddresses returns the UpdateIpAddresses field if non-nil, zero value otherwise.

### GetUpdateIpAddressesOk

`func (o *GetZoneOptionsResponseResponse) GetUpdateIpAddressesOk() (*[]string, bool)`

GetUpdateIpAddressesOk returns a tuple with the UpdateIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateIpAddresses

`func (o *GetZoneOptionsResponseResponse) SetUpdateIpAddresses(v []string)`

SetUpdateIpAddresses sets UpdateIpAddresses field to given value.

### HasUpdateIpAddresses

`func (o *GetZoneOptionsResponseResponse) HasUpdateIpAddresses() bool`

HasUpdateIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


