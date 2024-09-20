# \DnsZoneAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDnsZone**](DnsZoneAPI.md#CreateDnsZone) | **Post** /zones/create | Creates a new authoritative zone
[**DeleteDnsZone**](DnsZoneAPI.md#DeleteDnsZone) | **Post** /zones/delete | Deletes an authoritative zone
[**GetDnsZoneOptions**](DnsZoneAPI.md#GetDnsZoneOptions) | **Get** /zones/options/get | Gets the zone specific options
[**ListDnsZones**](DnsZoneAPI.md#ListDnsZones) | **Get** /zones/list | List all authoritative zones hosted on this DNS server. The list contains only the zones that the user has View permissions for. These API calls requires permission for both the Zones section as well as the individual permission for each zone. 
[**SetDnsZoneOptions**](DnsZoneAPI.md#SetDnsZoneOptions) | **Post** /zones/options/set | Sets the zone specific options



## CreateDnsZone

> CreateZoneResponse CreateDnsZone(ctx).Zone(zone).Type_(type_).UseSoaSerialDateScheme(useSoaSerialDateScheme).PrimaryNameServerAddresses(primaryNameServerAddresses).ZoneTransferProtocol(zoneTransferProtocol).Protocol(protocol).Forwarder(forwarder).DnssecValidation(dnssecValidation).ProxyType(proxyType).ProxyAddress(proxyAddress).ProxyPort(proxyPort).ProxyUsername(proxyUsername).ProxyPassword(proxyPassword).Execute()

Creates a new authoritative zone

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mleczkom/technitium-api-client/technitium"
)

func main() {
	zone := "zone_example" // string | The domain name for creating the new zone. The value can be valid domain name, an IP address,  or an network address in CIDR format. When value is IP address or network address, a reverse zone is created. 
	type_ := "type__example" // string | The type of zone to be created. Valid values are [Primary, Secondary, Stub, Forwarder]
	useSoaSerialDateScheme := true // bool | Set value to true to enable using date scheme for SOA serial. This optional parameter is used  only with Primary zone. Default value is false  (optional)
	primaryNameServerAddresses := []string{"Inner_example"} // []string | List of comma separated IP addresses of the primary name server. This optional parameter is used only with Secondary and Stub zones. If this parameter is not used, the DNS server will try to recursively resolve the primary name server addresses automatically  (optional)
	zoneTransferProtocol := "zoneTransferProtocol_example" // string | The zone transfer protocol to be used by secondary zones. Valid values are [Tcp, Tls, Quic]  (optional)
	protocol := "protocol_example" // string | The DNS transport protocol to be used by the conditional forwarder zone. This optional parameter is used with Conditional Forwarder zones. Valid values are [Udp, Tcp, Tls, Https, Quic]. Default Udp protocol is used when this parameter is missing  (optional)
	forwarder := "forwarder_example" // string | The address of the DNS server to be used as a forwarder. This optional parameter is required to be used with Conditional Forwarder zones. A special value this-server can be used as a forwarder which when used will forward all the requests internally to this DNS server such that you can override the zone with records and rest of the zone gets resolved via This Server   (optional)
	dnssecValidation := true // bool | Set this boolean value to indicate if DNSSEC validation must be done. This optional parameter is required  to be used with Conditional Forwarder zones  (optional)
	proxyType := "proxyType_example" // string | The type of proxy that must be used for conditional forwarding. This optional parameter is required to be used with Conditional Forwarder zones. Valid values are [NoProxy, DefaultProxy, Http, Socks5].  Default value DefaultProxy is used when this parameter is missing  (optional)
	proxyAddress := "proxyAddress_example" // string | The proxy server address to use when proxyType is configured. This optional parameter is required to be used with Conditional Forwarder zones  (optional)
	proxyPort := "proxyPort_example" // string | The proxy server port to use when proxyType is configured. This optional parameter is required to be used with Conditional Forwarder zones  (optional)
	proxyUsername := "proxyUsername_example" // string | The proxy server username to use when proxyType is configured. This optional parameter is required to be used with Conditional Forwarder zones  (optional)
	proxyPassword := "proxyPassword_example" // string | The proxy server password to use when proxyType is configured. This optional parameter is required to be used with Conditional Forwarder zones  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsZoneAPI.CreateDnsZone(context.Background()).Zone(zone).Type_(type_).UseSoaSerialDateScheme(useSoaSerialDateScheme).PrimaryNameServerAddresses(primaryNameServerAddresses).ZoneTransferProtocol(zoneTransferProtocol).Protocol(protocol).Forwarder(forwarder).DnssecValidation(dnssecValidation).ProxyType(proxyType).ProxyAddress(proxyAddress).ProxyPort(proxyPort).ProxyUsername(proxyUsername).ProxyPassword(proxyPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZoneAPI.CreateDnsZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDnsZone`: CreateZoneResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsZoneAPI.CreateDnsZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDnsZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string** | The domain name for creating the new zone. The value can be valid domain name, an IP address,  or an network address in CIDR format. When value is IP address or network address, a reverse zone is created.  | 
 **type_** | **string** | The type of zone to be created. Valid values are [Primary, Secondary, Stub, Forwarder] | 
 **useSoaSerialDateScheme** | **bool** | Set value to true to enable using date scheme for SOA serial. This optional parameter is used  only with Primary zone. Default value is false  | 
 **primaryNameServerAddresses** | **[]string** | List of comma separated IP addresses of the primary name server. This optional parameter is used only with Secondary and Stub zones. If this parameter is not used, the DNS server will try to recursively resolve the primary name server addresses automatically  | 
 **zoneTransferProtocol** | **string** | The zone transfer protocol to be used by secondary zones. Valid values are [Tcp, Tls, Quic]  | 
 **protocol** | **string** | The DNS transport protocol to be used by the conditional forwarder zone. This optional parameter is used with Conditional Forwarder zones. Valid values are [Udp, Tcp, Tls, Https, Quic]. Default Udp protocol is used when this parameter is missing  | 
 **forwarder** | **string** | The address of the DNS server to be used as a forwarder. This optional parameter is required to be used with Conditional Forwarder zones. A special value this-server can be used as a forwarder which when used will forward all the requests internally to this DNS server such that you can override the zone with records and rest of the zone gets resolved via This Server   | 
 **dnssecValidation** | **bool** | Set this boolean value to indicate if DNSSEC validation must be done. This optional parameter is required  to be used with Conditional Forwarder zones  | 
 **proxyType** | **string** | The type of proxy that must be used for conditional forwarding. This optional parameter is required to be used with Conditional Forwarder zones. Valid values are [NoProxy, DefaultProxy, Http, Socks5].  Default value DefaultProxy is used when this parameter is missing  | 
 **proxyAddress** | **string** | The proxy server address to use when proxyType is configured. This optional parameter is required to be used with Conditional Forwarder zones  | 
 **proxyPort** | **string** | The proxy server port to use when proxyType is configured. This optional parameter is required to be used with Conditional Forwarder zones  | 
 **proxyUsername** | **string** | The proxy server username to use when proxyType is configured. This optional parameter is required to be used with Conditional Forwarder zones  | 
 **proxyPassword** | **string** | The proxy server password to use when proxyType is configured. This optional parameter is required to be used with Conditional Forwarder zones  | 

### Return type

[**CreateZoneResponse**](CreateZoneResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDnsZone

> CommonResponse DeleteDnsZone(ctx).Zone(zone).Execute()

Deletes an authoritative zone

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mleczkom/technitium-api-client/technitium"
)

func main() {
	zone := "zone_example" // string | Dns zone name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsZoneAPI.DeleteDnsZone(context.Background()).Zone(zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZoneAPI.DeleteDnsZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDnsZone`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsZoneAPI.DeleteDnsZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDnsZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string** | Dns zone name | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDnsZoneOptions

> GetZoneOptionsResponse GetDnsZoneOptions(ctx).Zone(zone).IncludeAvailableTsigKeyNames(includeAvailableTsigKeyNames).Execute()

Gets the zone specific options

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mleczkom/technitium-api-client/technitium"
)

func main() {
	zone := "zone_example" // string | The domain name of the zone to get options
	includeAvailableTsigKeyNames := true // bool | Set to true to include list of available TSIG key names on the DNS server (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsZoneAPI.GetDnsZoneOptions(context.Background()).Zone(zone).IncludeAvailableTsigKeyNames(includeAvailableTsigKeyNames).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZoneAPI.GetDnsZoneOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsZoneOptions`: GetZoneOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsZoneAPI.GetDnsZoneOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDnsZoneOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string** | The domain name of the zone to get options | 
 **includeAvailableTsigKeyNames** | **bool** | Set to true to include list of available TSIG key names on the DNS server | 

### Return type

[**GetZoneOptionsResponse**](GetZoneOptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDnsZones

> ListZonesResponse ListDnsZones(ctx).PageNumber(pageNumber).ZonesPerPage(zonesPerPage).Execute()

List all authoritative zones hosted on this DNS server. The list contains only the zones that the user has View permissions for. These API calls requires permission for both the Zones section as well as the individual permission for each zone. 

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mleczkom/technitium-api-client/technitium"
)

func main() {
	pageNumber := int32(56) // int32 | When this parameter is specified, the API will return paginated results based on the page number and  zones per pages options. When not specified, the API will return a list of all zones.  (optional)
	zonesPerPage := int32(56) // int32 | The number of zones per page to be returned. This option is only used when pageNumber options is specified. The default value is 10 when not specified.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsZoneAPI.ListDnsZones(context.Background()).PageNumber(pageNumber).ZonesPerPage(zonesPerPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZoneAPI.ListDnsZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDnsZones`: ListZonesResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsZoneAPI.ListDnsZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDnsZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | When this parameter is specified, the API will return paginated results based on the page number and  zones per pages options. When not specified, the API will return a list of all zones.  | 
 **zonesPerPage** | **int32** | The number of zones per page to be returned. This option is only used when pageNumber options is specified. The default value is 10 when not specified.  | 

### Return type

[**ListZonesResponse**](ListZonesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDnsZoneOptions

> CommonResponse SetDnsZoneOptions(ctx).Zone(zone).Disabled(disabled).ZoneTransfer(zoneTransfer).ZoneTransferNameServers(zoneTransferNameServers).ZoneTransferTsigKeyNames(zoneTransferTsigKeyNames).Notify(notify).NotifyNameServers(notifyNameServers).Update(update).UpdateIpAddresses(updateIpAddresses).Execute()

Sets the zone specific options

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mleczkom/technitium-api-client/technitium"
)

func main() {
	zone := "zone_example" // string | The domain name of the zone to set options
	disabled := true // bool | Sets if the zone is enabled or disabled (optional)
	zoneTransfer := "zoneTransfer_example" // string | Sets if the zone allows zone transfer. Valid options are [Deny, Allow, AllowOnlyZoneNameServers, AllowOnlySpecifiedNameServers,  AllowBothZoneAndSpecifiedNameServers]. This option is valid only for Primary and Secondary zones.  (optional)
	zoneTransferNameServers := []string{"Inner_example"} // []string | A list of comma separated IP or network addresses which should be allowed to perform zone transfer. This list is enabled only when zoneTransfer option is set to AllowOnlySpecifiedNameServers or  AllowBothZoneAndSpecifiedNameServers. This option is valid only for Primary and Secondary zones  (optional)
	zoneTransferTsigKeyNames := []string{"Inner_example"} // []string | A list of comma separated TSIG keys names that are authorized to perform a zone transfer. Set this option to false to clear all key names. This option is valid only for Primary and Secondary zones  (optional)
	notify := "notify_example" // string | Sets if the DNS server should notify other DNS servers for zone updates. Valid options are [None, ZoneNameServers,  SpecifiedNameServers, BothZoneAndSpecifiedNameServers]. This option is valid only for Primary and Secondary zones.  (optional)
	notifyNameServers := []string{"Inner_example"} // []string | A list of comma separated IP addresses which should be notified by the DNS server for zone updates. This list is used only when notify option is set to SpecifiedNameServers or BothZoneAndSpecifiedNameServers. This option is valid only for Primary and Secondary zones.  (optional)
	update := "update_example" // string | Sets if the DNS server should allow dynamic updates (RFC 2136). This option is valid only  for Primary, Secondary and Forwarder zones. Valid options for Primary zones are [Deny, Allow,  AllowOnlyZoneNameServers, AllowOnlySpecifiedIpAddresses, AllowBothZoneNameServersAndSpecifiedIpAddresses]. Valid options for Secondary and Forwarder zones are [Deny, Allow, AllowOnlySpecifiedIpAddresses].  (optional)
	updateIpAddresses := []string{"Inner_example"} // []string | A list of comma separated IP or network addresses which should be allowed to perform dynamic updates. This list is enabled only when update option is set to AllowOnlySpecifiedIpAddresses or AllowBothZoneNameServersAndSpecifiedIpAddresses. This option is valid only for Primary,  Secondary and Forwarder zones  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsZoneAPI.SetDnsZoneOptions(context.Background()).Zone(zone).Disabled(disabled).ZoneTransfer(zoneTransfer).ZoneTransferNameServers(zoneTransferNameServers).ZoneTransferTsigKeyNames(zoneTransferTsigKeyNames).Notify(notify).NotifyNameServers(notifyNameServers).Update(update).UpdateIpAddresses(updateIpAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZoneAPI.SetDnsZoneOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDnsZoneOptions`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsZoneAPI.SetDnsZoneOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDnsZoneOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string** | The domain name of the zone to set options | 
 **disabled** | **bool** | Sets if the zone is enabled or disabled | 
 **zoneTransfer** | **string** | Sets if the zone allows zone transfer. Valid options are [Deny, Allow, AllowOnlyZoneNameServers, AllowOnlySpecifiedNameServers,  AllowBothZoneAndSpecifiedNameServers]. This option is valid only for Primary and Secondary zones.  | 
 **zoneTransferNameServers** | **[]string** | A list of comma separated IP or network addresses which should be allowed to perform zone transfer. This list is enabled only when zoneTransfer option is set to AllowOnlySpecifiedNameServers or  AllowBothZoneAndSpecifiedNameServers. This option is valid only for Primary and Secondary zones  | 
 **zoneTransferTsigKeyNames** | **[]string** | A list of comma separated TSIG keys names that are authorized to perform a zone transfer. Set this option to false to clear all key names. This option is valid only for Primary and Secondary zones  | 
 **notify** | **string** | Sets if the DNS server should notify other DNS servers for zone updates. Valid options are [None, ZoneNameServers,  SpecifiedNameServers, BothZoneAndSpecifiedNameServers]. This option is valid only for Primary and Secondary zones.  | 
 **notifyNameServers** | **[]string** | A list of comma separated IP addresses which should be notified by the DNS server for zone updates. This list is used only when notify option is set to SpecifiedNameServers or BothZoneAndSpecifiedNameServers. This option is valid only for Primary and Secondary zones.  | 
 **update** | **string** | Sets if the DNS server should allow dynamic updates (RFC 2136). This option is valid only  for Primary, Secondary and Forwarder zones. Valid options for Primary zones are [Deny, Allow,  AllowOnlyZoneNameServers, AllowOnlySpecifiedIpAddresses, AllowBothZoneNameServersAndSpecifiedIpAddresses]. Valid options for Secondary and Forwarder zones are [Deny, Allow, AllowOnlySpecifiedIpAddresses].  | 
 **updateIpAddresses** | **[]string** | A list of comma separated IP or network addresses which should be allowed to perform dynamic updates. This list is enabled only when update option is set to AllowOnlySpecifiedIpAddresses or AllowBothZoneNameServersAndSpecifiedIpAddresses. This option is valid only for Primary,  Secondary and Forwarder zones  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

