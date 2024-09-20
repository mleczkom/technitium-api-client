# \DnsRecordAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDnsRecord**](DnsRecordAPI.md#CreateDnsRecord) | **Post** /zones/records/add | Adds an resource record for an authoritative zone
[**DeleteDnsRecord**](DnsRecordAPI.md#DeleteDnsRecord) | **Post** /zones/records/delete | Deletes a record from an authoritative zone
[**UpdateDnsRecord**](DnsRecordAPI.md#UpdateDnsRecord) | **Post** /zones/records/update | Updates an existing record in an authoritative zone



## CreateDnsRecord

> CreateRecordResponse CreateDnsRecord(ctx).Domain(domain).Zone(zone).Type_(type_).Ttl(ttl).Overwrite(overwrite).Comments(comments).IpAddress(ipAddress).Ptr(ptr).CreatePtrZone(createPtrZone).UpdateSvcbHints(updateSvcbHints).NameServer(nameServer).Glue(glue).Cname(cname).PtrName(ptrName).Exchange(exchange).Preference(preference).Text(text).SplitText(splitText).Priority(priority).Weight(weight).Port(port).Target(target).NaptrOrder(naptrOrder).NaptrPreference(naptrPreference).NaptrFlags(naptrFlags).NaptrServices(naptrServices).NaptrRegexp(naptrRegexp).NaptrReplacement(naptrReplacement).Dname(dname).KeyTag(keyTag).Algorithm(algorithm).DigestType(digestType).Digest(digest).SshfpAlgorithm(sshfpAlgorithm).SshfpFingerprintType(sshfpFingerprintType).SshfpFingerprint(sshfpFingerprint).TlsaCertificateUsage(tlsaCertificateUsage).TlsaSelector(tlsaSelector).TlsaMatchingType(tlsaMatchingType).TlsaCertificateAssociationData(tlsaCertificateAssociationData).SvcPriority(svcPriority).SvcTargetName(svcTargetName).SvcParams(svcParams).AutoIpv4Hint(autoIpv4Hint).AutoIpv6Hint(autoIpv6Hint).UriPriority(uriPriority).UriWeight(uriWeight).Uri(uri).Flags(flags).Tag(tag).Value(value).Aname(aname).Protocol(protocol).Forwarder(forwarder).DnssecValidation(dnssecValidation).ProxyType(proxyType).ProxyAddress(proxyAddress).ProxyPort(proxyPort).ProxyUsername(proxyUsername).ProxyPassword(proxyPassword).AppName(appName).ClassPath(classPath).RecordData(recordData).Rdata(rdata).Execute()

Adds an resource record for an authoritative zone

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
	domain := "domain_example" // string | The domain name of the zone to add record
	zone := "zone_example" // string | The name of the authoritative zone into which the domain exists. When unspecified, the closest authoritative zone will be used.  (optional)
	type_ := "type__example" // string | he DNS resource record type. Supported record types are [A, AAAA, NS, CNAME, PTR, MX, TXT, SRV, DNAME, DS, SSHFP, TLSA, SVCB, HTTPS, URI, CAA] and proprietary types [ANAME, FWD, APP].  Unknown record types are also supported since v11.2.  (optional)
	ttl := int32(56) // int32 | The DNS resource record TTL value. This is the value in seconds that the DNS resolvers can cache the record for. When not specified the default TTL value from settings will be used  (optional)
	overwrite := true // bool | This option when set to true will overwrite existing resource record set for the selected type with the new record. Default value of false will add the new record into existing resource record set.  (optional)
	comments := "comments_example" // string | Sets comments for the added resource record (optional)
	ipAddress := "ipAddress_example" // string | The IP address for adding A or AAAA record. A special value of request-ip-address  can be used to set the record with the IP address of the API HTTP request to help with  dynamic DNS update applications. This option is required and used only for A and AAAA records.  (optional)
	ptr := true // bool | Set this option to true to add a reverse PTR record for the IP address in the A or AAAA record. This option is used only for A and AAAA records  (optional)
	createPtrZone := true // bool | Set this option to true to create a reverse zone for PTR record. This option is used for A and AAAA records  (optional)
	updateSvcbHints := true // bool | Set this option to true to update any SVCB/HTTPS records in the zone that has Automatic Hints option enabled  and matches its target name with the current record's domain name. This option is used for A and AAAA records  (optional)
	nameServer := "nameServer_example" // string | The name server domain name. This option is required for adding NS record.  (optional)
	glue := "glue_example" // string | This is the glue address for the name server in the NS record. This optional  parameter is used for adding NS record  (optional)
	cname := "cname_example" // string | The CNAME domain name. This option is required for adding CNAME record.  (optional)
	ptrName := "ptrName_example" // string | The PTR domain name. This option is required for adding PTR record.  (optional)
	exchange := "exchange_example" // string | The exchange domain name. This option is required for adding MX record  (optional)
	preference := "preference_example" // string | This is the preference value for MX record type. This option is required for adding MX record  (optional)
	text := "text_example" // string | The text data for TXT record. This option is required for adding TXT record  (optional)
	splitText := true // bool | Set to true for using new line char to split text into multiple character-strings for adding TXT record.  (optional)
	priority := int32(56) // int32 | This parameter is required for adding the SRV record  (optional)
	weight := int32(56) // int32 | This parameter is required for adding the SRV record  (optional)
	port := int32(56) // int32 | This parameter is required for adding the SRV record  (optional)
	target := "target_example" // string | This parameter is required for adding the SRV record  (optional)
	naptrOrder := int32(56) // int32 | This parameter is required for adding the NAPTR record  (optional)
	naptrPreference := int32(56) // int32 | This parameter is required for adding the NAPTR record  (optional)
	naptrFlags := "naptrFlags_example" // string | This parameter is required for adding the NAPTR record  (optional)
	naptrServices := "naptrServices_example" // string | This parameter is required for adding the NAPTR record  (optional)
	naptrRegexp := "naptrRegexp_example" // string | This parameter is required for adding the NAPTR record  (optional)
	naptrReplacement := "naptrReplacement_example" // string | This parameter is required for adding the NAPTR record  (optional)
	dname := "dname_example" // string | The DNAME domain name. This option is required for adding DNAME record  (optional)
	keyTag := "keyTag_example" // string | This parameter is required for adding DS record  (optional)
	algorithm := "algorithm_example" // string | Valid values are [RSAMD5, DSA, RSASHA1, DSA-NSEC3-SHA1, RSASHA1-NSEC3-SHA1, RSASHA256, RSASHA512, ECC-GOST, ECDSAP256SHA256, ECDSAP384SHA384, ED25519, ED448]. This parameter is required for adding DS record.  (optional)
	digestType := "digestType_example" // string | Valid values are [SHA1, SHA256, GOST-R-34-11-94, SHA384]. This parameter is required for adding DS record  (optional)
	digest := "digest_example" // string | A hex string value. This parameter is required for adding DS record  (optional)
	sshfpAlgorithm := "sshfpAlgorithm_example" // string | Valid values are [RSA, DSA, ECDSA, Ed25519, Ed448]. This parameter is required for adding SSHFP record.  (optional)
	sshfpFingerprintType := "sshfpFingerprintType_example" // string | Valid values are [SHA1, SHA256]. This parameter is required for adding SSHFP record.  (optional)
	sshfpFingerprint := "sshfpFingerprint_example" // string | A hex string value. This parameter is required for adding SSHFP record  (optional)
	tlsaCertificateUsage := "tlsaCertificateUsage_example" // string | Valid values are [PKIX-TA, PKIX-EE, DANE-TA, DANE-EE]. This parameter is required for adding TLSA record  (optional)
	tlsaSelector := "tlsaSelector_example" // string | Valid values are [Cert, SPKI]. This parameter is required for adding TLSA record  (optional)
	tlsaMatchingType := "tlsaMatchingType_example" // string | Valid value are [Full, SHA2-256, SHA2-512]. This parameter is required for adding TLSA record  (optional)
	tlsaCertificateAssociationData := "tlsaCertificateAssociationData_example" // string | A X509 certificate in PEM format or a hex string value. This parameter  is required for adding TLSA record  (optional)
	svcPriority := "svcPriority_example" // string | The priority value for SVCB or HTTPS record. This parameter is required for adding SCVB or HTTPS record  (optional)
	svcTargetName := "svcTargetName_example" // string | The target domain name for SVCB or HTTPS record. This parameter is required for adding SCVB or HTTPS record  (optional)
	svcParams := "svcParams_example" // string | The service parameters for SVCB or HTTPS record which is a pipe separated list of key and value. For example, alpn|h2,h3|port|53443. To clear existing values, set it to false. This parameter is required for adding SCVB or HTTPS record  (optional)
	autoIpv4Hint := true // bool | Set this option to true to enable Automatic Hints for the ipv4hint parameter in the svcParams. This option is valid only for SVCB and HTTPS records  (optional)
	autoIpv6Hint := true // bool | Set this option to true to enable Automatic Hints for the ipv6hint parameter in the svcParams. This option is valid only for SVCB and HTTPS records.  (optional)
	uriPriority := int32(56) // int32 | The priority value for adding the URI record  (optional)
	uriWeight := int32(56) // int32 | The weight value for adding the URI record  (optional)
	uri := "uri_example" // string | The URI value for adding the URI record  (optional)
	flags := int32(56) // int32 | This parameter is required for adding the CAA record  (optional)
	tag := "tag_example" // string | This parameter is required for adding the CAA record  (optional)
	value := "value_example" // string | This parameter is required for adding the CAA record  (optional)
	aname := "aname_example" // string | The ANAME domain name. This option is required for adding ANAME record  (optional)
	protocol := "protocol_example" // string | This parameter is required for adding the FWD record. Valid values are [Udp, Tcp, Tls, Https, Quic]  (optional)
	forwarder := "forwarder_example" // string | The forwarder address. A special value of this-server can be used to directly forward requests internally to the DNS server. This parameter is required for adding the FWD record  (optional)
	dnssecValidation := true // bool | Set this boolean value to indicate if DNSSEC validation must be done. This optional parameter is to be used with FWD records. Default value is false  (optional)
	proxyType := "proxyType_example" // string | The type of proxy that must be used for conditional forwarding. This optional parameter is to be used with FWD records. Valid values are [NoProxy, DefaultProxy, Http, Socks5]. Default value  DefaultProxy is used when this parameter is missing  (optional)
	proxyAddress := "proxyAddress_example" // string | The proxy server address to use when proxyType is configured. This optional parameter is  to be used with FWD records  (optional)
	proxyPort := "proxyPort_example" // string | The proxy server port to use when proxyType is configured. This optional parameter is to be used with FWD records  (optional)
	proxyUsername := "proxyUsername_example" // string | The proxy server username to use when proxyType is configured. This optional parameter is to be used with FWD records  (optional)
	proxyPassword := "proxyPassword_example" // string | The proxy server password to use when proxyType is configured. This optional parameter is to be used with FWD records.  (optional)
	appName := "appName_example" // string | The name of the DNS app. This parameter is required for adding the APP record  (optional)
	classPath := "classPath_example" // string | This parameter is required for adding the APP record  (optional)
	recordData := "recordData_example" // string | This parameter is used for adding the APP record as per the DNS app requirements  (optional)
	rdata := "rdata_example" // string | This parameter is used for adding unknown i.e. unsupported record types. The value must be formatted as a hex string or a colon separated hex string.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsRecordAPI.CreateDnsRecord(context.Background()).Domain(domain).Zone(zone).Type_(type_).Ttl(ttl).Overwrite(overwrite).Comments(comments).IpAddress(ipAddress).Ptr(ptr).CreatePtrZone(createPtrZone).UpdateSvcbHints(updateSvcbHints).NameServer(nameServer).Glue(glue).Cname(cname).PtrName(ptrName).Exchange(exchange).Preference(preference).Text(text).SplitText(splitText).Priority(priority).Weight(weight).Port(port).Target(target).NaptrOrder(naptrOrder).NaptrPreference(naptrPreference).NaptrFlags(naptrFlags).NaptrServices(naptrServices).NaptrRegexp(naptrRegexp).NaptrReplacement(naptrReplacement).Dname(dname).KeyTag(keyTag).Algorithm(algorithm).DigestType(digestType).Digest(digest).SshfpAlgorithm(sshfpAlgorithm).SshfpFingerprintType(sshfpFingerprintType).SshfpFingerprint(sshfpFingerprint).TlsaCertificateUsage(tlsaCertificateUsage).TlsaSelector(tlsaSelector).TlsaMatchingType(tlsaMatchingType).TlsaCertificateAssociationData(tlsaCertificateAssociationData).SvcPriority(svcPriority).SvcTargetName(svcTargetName).SvcParams(svcParams).AutoIpv4Hint(autoIpv4Hint).AutoIpv6Hint(autoIpv6Hint).UriPriority(uriPriority).UriWeight(uriWeight).Uri(uri).Flags(flags).Tag(tag).Value(value).Aname(aname).Protocol(protocol).Forwarder(forwarder).DnssecValidation(dnssecValidation).ProxyType(proxyType).ProxyAddress(proxyAddress).ProxyPort(proxyPort).ProxyUsername(proxyUsername).ProxyPassword(proxyPassword).AppName(appName).ClassPath(classPath).RecordData(recordData).Rdata(rdata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsRecordAPI.CreateDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDnsRecord`: CreateRecordResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsRecordAPI.CreateDnsRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | The domain name of the zone to add record | 
 **zone** | **string** | The name of the authoritative zone into which the domain exists. When unspecified, the closest authoritative zone will be used.  | 
 **type_** | **string** | he DNS resource record type. Supported record types are [A, AAAA, NS, CNAME, PTR, MX, TXT, SRV, DNAME, DS, SSHFP, TLSA, SVCB, HTTPS, URI, CAA] and proprietary types [ANAME, FWD, APP].  Unknown record types are also supported since v11.2.  | 
 **ttl** | **int32** | The DNS resource record TTL value. This is the value in seconds that the DNS resolvers can cache the record for. When not specified the default TTL value from settings will be used  | 
 **overwrite** | **bool** | This option when set to true will overwrite existing resource record set for the selected type with the new record. Default value of false will add the new record into existing resource record set.  | 
 **comments** | **string** | Sets comments for the added resource record | 
 **ipAddress** | **string** | The IP address for adding A or AAAA record. A special value of request-ip-address  can be used to set the record with the IP address of the API HTTP request to help with  dynamic DNS update applications. This option is required and used only for A and AAAA records.  | 
 **ptr** | **bool** | Set this option to true to add a reverse PTR record for the IP address in the A or AAAA record. This option is used only for A and AAAA records  | 
 **createPtrZone** | **bool** | Set this option to true to create a reverse zone for PTR record. This option is used for A and AAAA records  | 
 **updateSvcbHints** | **bool** | Set this option to true to update any SVCB/HTTPS records in the zone that has Automatic Hints option enabled  and matches its target name with the current record&#39;s domain name. This option is used for A and AAAA records  | 
 **nameServer** | **string** | The name server domain name. This option is required for adding NS record.  | 
 **glue** | **string** | This is the glue address for the name server in the NS record. This optional  parameter is used for adding NS record  | 
 **cname** | **string** | The CNAME domain name. This option is required for adding CNAME record.  | 
 **ptrName** | **string** | The PTR domain name. This option is required for adding PTR record.  | 
 **exchange** | **string** | The exchange domain name. This option is required for adding MX record  | 
 **preference** | **string** | This is the preference value for MX record type. This option is required for adding MX record  | 
 **text** | **string** | The text data for TXT record. This option is required for adding TXT record  | 
 **splitText** | **bool** | Set to true for using new line char to split text into multiple character-strings for adding TXT record.  | 
 **priority** | **int32** | This parameter is required for adding the SRV record  | 
 **weight** | **int32** | This parameter is required for adding the SRV record  | 
 **port** | **int32** | This parameter is required for adding the SRV record  | 
 **target** | **string** | This parameter is required for adding the SRV record  | 
 **naptrOrder** | **int32** | This parameter is required for adding the NAPTR record  | 
 **naptrPreference** | **int32** | This parameter is required for adding the NAPTR record  | 
 **naptrFlags** | **string** | This parameter is required for adding the NAPTR record  | 
 **naptrServices** | **string** | This parameter is required for adding the NAPTR record  | 
 **naptrRegexp** | **string** | This parameter is required for adding the NAPTR record  | 
 **naptrReplacement** | **string** | This parameter is required for adding the NAPTR record  | 
 **dname** | **string** | The DNAME domain name. This option is required for adding DNAME record  | 
 **keyTag** | **string** | This parameter is required for adding DS record  | 
 **algorithm** | **string** | Valid values are [RSAMD5, DSA, RSASHA1, DSA-NSEC3-SHA1, RSASHA1-NSEC3-SHA1, RSASHA256, RSASHA512, ECC-GOST, ECDSAP256SHA256, ECDSAP384SHA384, ED25519, ED448]. This parameter is required for adding DS record.  | 
 **digestType** | **string** | Valid values are [SHA1, SHA256, GOST-R-34-11-94, SHA384]. This parameter is required for adding DS record  | 
 **digest** | **string** | A hex string value. This parameter is required for adding DS record  | 
 **sshfpAlgorithm** | **string** | Valid values are [RSA, DSA, ECDSA, Ed25519, Ed448]. This parameter is required for adding SSHFP record.  | 
 **sshfpFingerprintType** | **string** | Valid values are [SHA1, SHA256]. This parameter is required for adding SSHFP record.  | 
 **sshfpFingerprint** | **string** | A hex string value. This parameter is required for adding SSHFP record  | 
 **tlsaCertificateUsage** | **string** | Valid values are [PKIX-TA, PKIX-EE, DANE-TA, DANE-EE]. This parameter is required for adding TLSA record  | 
 **tlsaSelector** | **string** | Valid values are [Cert, SPKI]. This parameter is required for adding TLSA record  | 
 **tlsaMatchingType** | **string** | Valid value are [Full, SHA2-256, SHA2-512]. This parameter is required for adding TLSA record  | 
 **tlsaCertificateAssociationData** | **string** | A X509 certificate in PEM format or a hex string value. This parameter  is required for adding TLSA record  | 
 **svcPriority** | **string** | The priority value for SVCB or HTTPS record. This parameter is required for adding SCVB or HTTPS record  | 
 **svcTargetName** | **string** | The target domain name for SVCB or HTTPS record. This parameter is required for adding SCVB or HTTPS record  | 
 **svcParams** | **string** | The service parameters for SVCB or HTTPS record which is a pipe separated list of key and value. For example, alpn|h2,h3|port|53443. To clear existing values, set it to false. This parameter is required for adding SCVB or HTTPS record  | 
 **autoIpv4Hint** | **bool** | Set this option to true to enable Automatic Hints for the ipv4hint parameter in the svcParams. This option is valid only for SVCB and HTTPS records  | 
 **autoIpv6Hint** | **bool** | Set this option to true to enable Automatic Hints for the ipv6hint parameter in the svcParams. This option is valid only for SVCB and HTTPS records.  | 
 **uriPriority** | **int32** | The priority value for adding the URI record  | 
 **uriWeight** | **int32** | The weight value for adding the URI record  | 
 **uri** | **string** | The URI value for adding the URI record  | 
 **flags** | **int32** | This parameter is required for adding the CAA record  | 
 **tag** | **string** | This parameter is required for adding the CAA record  | 
 **value** | **string** | This parameter is required for adding the CAA record  | 
 **aname** | **string** | The ANAME domain name. This option is required for adding ANAME record  | 
 **protocol** | **string** | This parameter is required for adding the FWD record. Valid values are [Udp, Tcp, Tls, Https, Quic]  | 
 **forwarder** | **string** | The forwarder address. A special value of this-server can be used to directly forward requests internally to the DNS server. This parameter is required for adding the FWD record  | 
 **dnssecValidation** | **bool** | Set this boolean value to indicate if DNSSEC validation must be done. This optional parameter is to be used with FWD records. Default value is false  | 
 **proxyType** | **string** | The type of proxy that must be used for conditional forwarding. This optional parameter is to be used with FWD records. Valid values are [NoProxy, DefaultProxy, Http, Socks5]. Default value  DefaultProxy is used when this parameter is missing  | 
 **proxyAddress** | **string** | The proxy server address to use when proxyType is configured. This optional parameter is  to be used with FWD records  | 
 **proxyPort** | **string** | The proxy server port to use when proxyType is configured. This optional parameter is to be used with FWD records  | 
 **proxyUsername** | **string** | The proxy server username to use when proxyType is configured. This optional parameter is to be used with FWD records  | 
 **proxyPassword** | **string** | The proxy server password to use when proxyType is configured. This optional parameter is to be used with FWD records.  | 
 **appName** | **string** | The name of the DNS app. This parameter is required for adding the APP record  | 
 **classPath** | **string** | This parameter is required for adding the APP record  | 
 **recordData** | **string** | This parameter is used for adding the APP record as per the DNS app requirements  | 
 **rdata** | **string** | This parameter is used for adding unknown i.e. unsupported record types. The value must be formatted as a hex string or a colon separated hex string.  | 

### Return type

[**CreateRecordResponse**](CreateRecordResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDnsRecord

> CommonResponse DeleteDnsRecord(ctx).Domain(domain).Zone(zone).Type_(type_).IpAddress(ipAddress).UpdateSvcbHints(updateSvcbHints).NameServer(nameServer).PtrName(ptrName).Preference(preference).Exchange(exchange).Text(text).SplitText(splitText).Priority(priority).Weight(weight).Port(port).Target(target).NaptrOrder(naptrOrder).NaptrPreference(naptrPreference).NaptrFlags(naptrFlags).NaptrServices(naptrServices).NaptrRegexp(naptrRegexp).NaptrReplacement(naptrReplacement).KeyTag(keyTag).Algorithm(algorithm).DigestType(digestType).Digest(digest).SshfpAlgorithm(sshfpAlgorithm).SshfpFingerprintType(sshfpFingerprintType).SshfpFingerprint(sshfpFingerprint).TlsaCertificateUsage(tlsaCertificateUsage).TlsaSelector(tlsaSelector).TlsaMatchingType(tlsaMatchingType).TlsaCertificateAssociationData(tlsaCertificateAssociationData).SvcPriority(svcPriority).SvcTargetName(svcTargetName).SvcParams(svcParams).UriPriority(uriPriority).UriWeight(uriWeight).Uri(uri).Flags(flags).Tag(tag).Value(value).Aname(aname).Protocol(protocol).Forwarder(forwarder).Rdata(rdata).Execute()

Deletes a record from an authoritative zone

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
	domain := "domain_example" // string | The domain name of the zone to delete the record
	zone := "zone_example" // string | The name of the authoritative zone into which the domain exists. When unspecified, the closest authoritative zone will be used  (optional)
	type_ := "type__example" // string | The type of the resource record to delete  (optional)
	ipAddress := "ipAddress_example" // string | This parameter is required when deleting A or AAAA record  (optional)
	updateSvcbHints := true // bool | Set this option to true to update any SVCB/HTTPS records in the zone that has Automatic Hints option enabled and matches its target name with the current record's domain name. This option is used for A and AAAA records.  (optional)
	nameServer := "nameServer_example" // string | This parameter is required when deleting NS record  (optional)
	ptrName := "ptrName_example" // string | This parameter is required when deleting PTR record  (optional)
	preference := "preference_example" // string | This parameter is required when deleting MX record  (optional)
	exchange := "exchange_example" // string | This parameter is required when deleting MX record  (optional)
	text := "text_example" // string | This parameter is required when deleting TXT record  (optional)
	splitText := true // bool | This parameter is used when deleting TXT record. Default value is set to false when unspecified  (optional)
	priority := int32(56) // int32 | This parameter is required when deleting the SRV record  (optional)
	weight := int32(56) // int32 | This parameter is required when deleting the SRV record  (optional)
	port := int32(56) // int32 | This parameter is required when deleting the SRV record  (optional)
	target := "target_example" // string | This parameter is required when deleting the SRV record  (optional)
	naptrOrder := int32(56) // int32 | This parameter is required when deleting the NAPTR record  (optional)
	naptrPreference := int32(56) // int32 | This parameter is required for deleting the NAPTR record  (optional)
	naptrFlags := "naptrFlags_example" // string | This parameter is required for deleting the NAPTR record  (optional)
	naptrServices := "naptrServices_example" // string | This parameter is required for deleting the NAPTR record  (optional)
	naptrRegexp := "naptrRegexp_example" // string | This parameter is required for deleting the NAPTR record  (optional)
	naptrReplacement := "naptrReplacement_example" // string | This parameter is required for deleting the NAPTR record  (optional)
	keyTag := "keyTag_example" // string | This parameter is required when deleting DS record  (optional)
	algorithm := "algorithm_example" // string | This parameter is required when deleting DS record  (optional)
	digestType := "digestType_example" // string | This parameter is required when deleting DS record  (optional)
	digest := "digest_example" // string | This parameter is required when deleting DS record  (optional)
	sshfpAlgorithm := "sshfpAlgorithm_example" // string | This parameter is required when deleting SSHFP record  (optional)
	sshfpFingerprintType := "sshfpFingerprintType_example" // string | This parameter is required when deleting SSHFP record  (optional)
	sshfpFingerprint := "sshfpFingerprint_example" // string | This parameter is required when deleting SSHFP record  (optional)
	tlsaCertificateUsage := "tlsaCertificateUsage_example" // string | This parameter is required when deleting TLSA record  (optional)
	tlsaSelector := "tlsaSelector_example" // string | This parameter is required when deleting TLSA record  (optional)
	tlsaMatchingType := "tlsaMatchingType_example" // string | This parameter is required when deleting TLSA record  (optional)
	tlsaCertificateAssociationData := "tlsaCertificateAssociationData_example" // string | This parameter is required when deleting TLSA record  (optional)
	svcPriority := "svcPriority_example" // string | The priority value for SVCB or HTTPS record. This parameter is required for deleting SCVB or HTTPS record  (optional)
	svcTargetName := "svcTargetName_example" // string | TThe priority value for SVCB or HTTPS record. This parameter is required for deleting SCVB or HTTPS record  (optional)
	svcParams := "svcParams_example" // string | The service parameters for SVCB or HTTPS record which is a pipe separated list of key and value. For example, alpn|h2,h3|port|53443. To clear existing values, set it to false. This parameter  is required for deleting SCVB or HTTPS record.  (optional)
	uriPriority := int32(56) // int32 | The priority value in the URI record. This parameter is required when deleting the URI record  (optional)
	uriWeight := int32(56) // int32 | The weight value in the URI record. This parameter is required when deleting the URI record  (optional)
	uri := "uri_example" // string | The URI value in the URI record. This parameter is required when deleting the URI record  (optional)
	flags := int32(56) // int32 | This is the flags parameter in the CAA record. This parameter is required when deleting  the CAA record  (optional)
	tag := "tag_example" // string | This is the tag parameter in the CAA record. This parameter is required when deleting the CAA record.  (optional)
	value := "value_example" // string | This parameter is required when deleting the CAA record  (optional)
	aname := "aname_example" // string | This parameter is required when deleting the ANAME record  (optional)
	protocol := "protocol_example" // string | This is the protocol parameter in the FWD record. Valid values are [Udp, Tcp, Tls, Https, Quic].  This parameter is optional and default value Udp will be used when deleting the FWD record  (optional)
	forwarder := "forwarder_example" // string | This parameter is required when deleting the FWD record  (optional)
	rdata := "rdata_example" // string | This parameter is used for deleting unknown i.e. unsupported record types.  The value must be formatted as a hex string or a colon separated hex string.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsRecordAPI.DeleteDnsRecord(context.Background()).Domain(domain).Zone(zone).Type_(type_).IpAddress(ipAddress).UpdateSvcbHints(updateSvcbHints).NameServer(nameServer).PtrName(ptrName).Preference(preference).Exchange(exchange).Text(text).SplitText(splitText).Priority(priority).Weight(weight).Port(port).Target(target).NaptrOrder(naptrOrder).NaptrPreference(naptrPreference).NaptrFlags(naptrFlags).NaptrServices(naptrServices).NaptrRegexp(naptrRegexp).NaptrReplacement(naptrReplacement).KeyTag(keyTag).Algorithm(algorithm).DigestType(digestType).Digest(digest).SshfpAlgorithm(sshfpAlgorithm).SshfpFingerprintType(sshfpFingerprintType).SshfpFingerprint(sshfpFingerprint).TlsaCertificateUsage(tlsaCertificateUsage).TlsaSelector(tlsaSelector).TlsaMatchingType(tlsaMatchingType).TlsaCertificateAssociationData(tlsaCertificateAssociationData).SvcPriority(svcPriority).SvcTargetName(svcTargetName).SvcParams(svcParams).UriPriority(uriPriority).UriWeight(uriWeight).Uri(uri).Flags(flags).Tag(tag).Value(value).Aname(aname).Protocol(protocol).Forwarder(forwarder).Rdata(rdata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsRecordAPI.DeleteDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDnsRecord`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsRecordAPI.DeleteDnsRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | The domain name of the zone to delete the record | 
 **zone** | **string** | The name of the authoritative zone into which the domain exists. When unspecified, the closest authoritative zone will be used  | 
 **type_** | **string** | The type of the resource record to delete  | 
 **ipAddress** | **string** | This parameter is required when deleting A or AAAA record  | 
 **updateSvcbHints** | **bool** | Set this option to true to update any SVCB/HTTPS records in the zone that has Automatic Hints option enabled and matches its target name with the current record&#39;s domain name. This option is used for A and AAAA records.  | 
 **nameServer** | **string** | This parameter is required when deleting NS record  | 
 **ptrName** | **string** | This parameter is required when deleting PTR record  | 
 **preference** | **string** | This parameter is required when deleting MX record  | 
 **exchange** | **string** | This parameter is required when deleting MX record  | 
 **text** | **string** | This parameter is required when deleting TXT record  | 
 **splitText** | **bool** | This parameter is used when deleting TXT record. Default value is set to false when unspecified  | 
 **priority** | **int32** | This parameter is required when deleting the SRV record  | 
 **weight** | **int32** | This parameter is required when deleting the SRV record  | 
 **port** | **int32** | This parameter is required when deleting the SRV record  | 
 **target** | **string** | This parameter is required when deleting the SRV record  | 
 **naptrOrder** | **int32** | This parameter is required when deleting the NAPTR record  | 
 **naptrPreference** | **int32** | This parameter is required for deleting the NAPTR record  | 
 **naptrFlags** | **string** | This parameter is required for deleting the NAPTR record  | 
 **naptrServices** | **string** | This parameter is required for deleting the NAPTR record  | 
 **naptrRegexp** | **string** | This parameter is required for deleting the NAPTR record  | 
 **naptrReplacement** | **string** | This parameter is required for deleting the NAPTR record  | 
 **keyTag** | **string** | This parameter is required when deleting DS record  | 
 **algorithm** | **string** | This parameter is required when deleting DS record  | 
 **digestType** | **string** | This parameter is required when deleting DS record  | 
 **digest** | **string** | This parameter is required when deleting DS record  | 
 **sshfpAlgorithm** | **string** | This parameter is required when deleting SSHFP record  | 
 **sshfpFingerprintType** | **string** | This parameter is required when deleting SSHFP record  | 
 **sshfpFingerprint** | **string** | This parameter is required when deleting SSHFP record  | 
 **tlsaCertificateUsage** | **string** | This parameter is required when deleting TLSA record  | 
 **tlsaSelector** | **string** | This parameter is required when deleting TLSA record  | 
 **tlsaMatchingType** | **string** | This parameter is required when deleting TLSA record  | 
 **tlsaCertificateAssociationData** | **string** | This parameter is required when deleting TLSA record  | 
 **svcPriority** | **string** | The priority value for SVCB or HTTPS record. This parameter is required for deleting SCVB or HTTPS record  | 
 **svcTargetName** | **string** | TThe priority value for SVCB or HTTPS record. This parameter is required for deleting SCVB or HTTPS record  | 
 **svcParams** | **string** | The service parameters for SVCB or HTTPS record which is a pipe separated list of key and value. For example, alpn|h2,h3|port|53443. To clear existing values, set it to false. This parameter  is required for deleting SCVB or HTTPS record.  | 
 **uriPriority** | **int32** | The priority value in the URI record. This parameter is required when deleting the URI record  | 
 **uriWeight** | **int32** | The weight value in the URI record. This parameter is required when deleting the URI record  | 
 **uri** | **string** | The URI value in the URI record. This parameter is required when deleting the URI record  | 
 **flags** | **int32** | This is the flags parameter in the CAA record. This parameter is required when deleting  the CAA record  | 
 **tag** | **string** | This is the tag parameter in the CAA record. This parameter is required when deleting the CAA record.  | 
 **value** | **string** | This parameter is required when deleting the CAA record  | 
 **aname** | **string** | This parameter is required when deleting the ANAME record  | 
 **protocol** | **string** | This is the protocol parameter in the FWD record. Valid values are [Udp, Tcp, Tls, Https, Quic].  This parameter is optional and default value Udp will be used when deleting the FWD record  | 
 **forwarder** | **string** | This parameter is required when deleting the FWD record  | 
 **rdata** | **string** | This parameter is used for deleting unknown i.e. unsupported record types.  The value must be formatted as a hex string or a colon separated hex string.  | 

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


## UpdateDnsRecord

> UpdateRecordResponse UpdateDnsRecord(ctx).Domain(domain).NewDomain(newDomain).Zone(zone).Type_(type_).Ttl(ttl).Disable(disable).Comments(comments).IpAddress(ipAddress).NewIpAddress(newIpAddress).Ptr(ptr).CreatePtrZone(createPtrZone).UpdateSvcbHints(updateSvcbHints).NameServer(nameServer).NewNameServer(newNameServer).Glue(glue).Cname(cname).PrimaryNameServer(primaryNameServer).ResponsiblePerson(responsiblePerson).Serial(serial).Refresh(refresh).Retry(retry).Expire(expire).Minimum(minimum).PrimaryAddresses(primaryAddresses).ZoneTransferProtocol(zoneTransferProtocol).TsigKeyName(tsigKeyName).PtrName(ptrName).NewPtrName(newPtrName).Preference(preference).NewPreference(newPreference).Exchange(exchange).NewExchange(newExchange).Text(text).NewText(newText).SplitText(splitText).NewSplitText(newSplitText).Priority(priority).NewPriority(newPriority).Weight(weight).NewWeight(newWeight).Port(port).NewPort(newPort).Target(target).NewTarget(newTarget).NaptrOrder(naptrOrder).NaptrNewOrder(naptrNewOrder).NaptrPreference(naptrPreference).NaptrNewPreference(naptrNewPreference).NaptrFlags(naptrFlags).NaptrNewFlags(naptrNewFlags).NaptrServices(naptrServices).NaptrNewServices(naptrNewServices).NaptrRegexp(naptrRegexp).NaptrNewRegexp(naptrNewRegexp).NaptrReplacement(naptrReplacement).NaptrNewReplacement(naptrNewReplacement).Dname(dname).KeyTag(keyTag).NewKeyTag(newKeyTag).Algorithm(algorithm).NewAlgorithm(newAlgorithm).DigestType(digestType).NewDigestType(newDigestType).Digest(digest).NewDigest(newDigest).SshfpAlgorithm(sshfpAlgorithm).NewSshfpAlgorithm(newSshfpAlgorithm).SshfpFingerprintType(sshfpFingerprintType).NewSshfpFingerprintType(newSshfpFingerprintType).SshfpFingerprint(sshfpFingerprint).NewSshfpFingerprint(newSshfpFingerprint).TlsaCertificateUsage(tlsaCertificateUsage).NewTlsaCertificateUsage(newTlsaCertificateUsage).TlsaSelector(tlsaSelector).NewTlsaSelector(newTlsaSelector).TlsaMatchingType(tlsaMatchingType).NewTlsaMatchingType(newTlsaMatchingType).TlsaCertificateAssociationData(tlsaCertificateAssociationData).NewTlsaCertificateAssociationData(newTlsaCertificateAssociationData).SvcPriority(svcPriority).NewSvcPriority(newSvcPriority).SvcTargetName(svcTargetName).NewSvcTargetName(newSvcTargetName).SvcParams(svcParams).NewSvcParams(newSvcParams).AutoIpv4Hint(autoIpv4Hint).AutoIpv6Hint(autoIpv6Hint).UriPriority(uriPriority).NewUriPriority(newUriPriority).UriWeight(uriWeight).NewUriWeight(newUriWeight).Uri(uri).NewUri(newUri).Flags(flags).NewFlags(newFlags).Tag(tag).NewTag(newTag).Value(value).NewValue(newValue).Aname(aname).NewAName(newAName).Protocol(protocol).NewProtocol(newProtocol).Forwarder(forwarder).NewForwarder(newForwarder).DnssecValidation(dnssecValidation).ProxyType(proxyType).ProxyAddress(proxyAddress).ProxyPort(proxyPort).ProxyUsername(proxyUsername).ProxyPassword(proxyPassword).AppName(appName).ClassPath(classPath).RecordData(recordData).Rdata(rdata).NewRData(newRData).Execute()

Updates an existing record in an authoritative zone

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
	domain := "domain_example" // string | The domain name of the zone to update the record
	newDomain := "newDomain_example" // string | The new domain name to be set for the record. To be used to rename sub domain name of the record 
	zone := "zone_example" // string | The name of the authoritative zone into which the domain exists. When unspecified, the closest authoritative zone will be used  (optional)
	type_ := "type__example" // string | The type of the resource record to update  (optional)
	ttl := int32(56) // int32 | The TTL value of the resource record. Default value of 3600 is used when  parameter is missing  (optional)
	disable := true // bool | Specifies if the record should be disabled. The default value is false when this parameter is missing  (optional)
	comments := "comments_example" // string | Sets comments for the updated resource record (optional)
	ipAddress := "ipAddress_example" // string | The current IP address in the A or AAAA record. This parameter is required when updating A or AAAA record  (optional)
	newIpAddress := "newIpAddress_example" // string | The new IP address in the A or AAAA record. This parameter when missing will use the current value in the record.  (optional)
	ptr := true // bool | Set this option to true to specify if the PTR record associated with the A or AAAA record must also be updated. This option is used only for A and AAAA records  (optional)
	createPtrZone := true // bool | Set this option to true to create a reverse zone for PTR record. This option is used  only for A and AAAA records.  (optional)
	updateSvcbHints := true // bool | Set this option to true to update any SVCB/HTTPS records in the zone that has Automatic Hints  option enabled and matches its target name with the current record's domain name. This option is used for A and AAAA records  (optional)
	nameServer := "nameServer_example" // string | The current name server domain name. This option is required for updating NS record  (optional)
	newNameServer := "newNameServer_example" // string | The new server domain name. This option is used for updating NS record  (optional)
	glue := "glue_example" // string | The comma separated list of IP addresses set as glue for the NS record. This parameter  is used only when updating NS record  (optional)
	cname := "cname_example" // string | The CNAME domain name to update in the existing CNAME record  (optional)
	primaryNameServer := "primaryNameServer_example" // string | This is the primary name server parameter in the SOA record. This parameter  is required when updating the SOA record.  (optional)
	responsiblePerson := "responsiblePerson_example" // string | This is the responsible person parameter in the SOA record. This parameter  is required when updating the SOA record  (optional)
	serial := int32(56) // int32 | This is the serial parameter in the SOA record. This parameter is required  when updating the SOA record  (optional)
	refresh := int32(56) // int32 | This is the refresh parameter in the SOA record. This parameter is required  when updating the SOA record  (optional)
	retry := int32(56) // int32 | This is the retry parameter in the SOA record. This parameter is required  when updating the SOA record  (optional)
	expire := int32(56) // int32 | This is the expire parameter in the SOA record. This parameter is required when updating the SOA record  (optional)
	minimum := int32(56) // int32 | This is the minimum parameter in the SOA record. This parameter is required when updating the SOA record  (optional)
	primaryAddresses := []string{"Inner_example"} // []string | This is a comma separated list of IP addresses of the primary name server. This parameter is to be used with secondary and stub zones where the primary name server address is not directly resolvable  (optional)
	zoneTransferProtocol := "zoneTransferProtocol_example" // string | The zone transfer protocol to be used by the secondary zone. Valid values are [Tcp, Tls, Quic]. This parameter is used with SOA record  (optional)
	tsigKeyName := "tsigKeyName_example" // string | The TSIG key name to be used by the secondary zone. This parameter is used with SOA record  (optional)
	ptrName := "ptrName_example" // string | The current PTR domain name. This option is required for updating PTR record  (optional)
	newPtrName := "newPtrName_example" // string | The new PTR domain name. This option is required for updating PTR record  (optional)
	preference := "preference_example" // string | The current preference value in an MX record. This parameter when missing will default to 1 value. This parameter is used only when updating MX record  (optional)
	newPreference := "newPreference_example" // string | The new preference value in an MX record. This parameter when missing will use the old value. This parameter is used only when updating MX record  (optional)
	exchange := "exchange_example" // string | The current exchange domain name. This option is required for updating MX record  (optional)
	newExchange := "newExchange_example" // string | The new exchange domain name. This option is required for updating MX record  (optional)
	text := "text_example" // string | The current text value. This option is required for updating TXT record  (optional)
	newText := "newText_example" // string | The new text value. This option is required for updating TXT record  (optional)
	splitText := true // bool | The current split text value. This option is used for updating TXT record and is set to false when unspecified  (optional)
	newSplitText := true // bool | The new split text value. This option is used for updating TXT record and is set to current split text value when unspecified  (optional)
	priority := int32(56) // int32 | This is the current priority in the SRV record. This parameter is required when updating the SRV record  (optional)
	newPriority := int32(56) // int32 | This is the new priority in the SRV record. This parameter when missing will use the old value. This parameter is used when updating the SRV record  (optional)
	weight := int32(56) // int32 | This is the current weight in the SRV record. This parameter is required when updating the SRV record  (optional)
	newWeight := int32(56) // int32 | This is the new weight in the SRV record. This parameter when missing will use the old value. This parameter is used when updating the SRV record  (optional)
	port := int32(56) // int32 | This is the port parameter in the SRV record. This parameter is required when updating the SRV record  (optional)
	newPort := int32(56) // int32 | This is the port parameter in the SRV record. This parameter is required when updating the SRV record  (optional)
	target := "target_example" // string | The current target value. This parameter is required when updating the SRV record  (optional)
	newTarget := "newTarget_example" // string | The new target value. This parameter when missing will use the old value. This parameter is required when updating the SRV record  (optional)
	naptrOrder := int32(56) // int32 | The current value in the NAPTR record. This parameter is required when  updating the NAPTR record  (optional)
	naptrNewOrder := int32(56) // int32 | The new value in the NAPTR record. This parameter when missing will use the old value. This parameter is used when updating the NAPTR record  (optional)
	naptrPreference := int32(56) // int32 | The current value in the NAPTR record. This parameter is required when updating the NAPTR record  (optional)
	naptrNewPreference := int32(56) // int32 | The new value in the NAPTR record. This parameter when missing will use the old value. This parameter is used when updating the NAPTR record  (optional)
	naptrFlags := "naptrFlags_example" // string | The current value in the NAPTR record. This parameter is required when updating the NAPTR record  (optional)
	naptrNewFlags := "naptrNewFlags_example" // string | The new value in the NAPTR record. This parameter when missing will use the old value. This parameter is used when updating the NAPTR record  (optional)
	naptrServices := "naptrServices_example" // string | The current value in the NAPTR record. This parameter is required when updating the NAPTR record  (optional)
	naptrNewServices := "naptrNewServices_example" // string | The new value in the NAPTR record. This parameter when missing will use the old value. This parameter is used when updating the NAPTR record  (optional)
	naptrRegexp := "naptrRegexp_example" // string | The current value in the NAPTR record. This parameter is required when updating the NAPTR record  (optional)
	naptrNewRegexp := "naptrNewRegexp_example" // string | The new value in the NAPTR record. This parameter when missing will use the old value. This parameter is used when updating the NAPTR record  (optional)
	naptrReplacement := "naptrReplacement_example" // string | The current value in the NAPTR record. This parameter is required when updating the NAPTR record  (optional)
	naptrNewReplacement := "naptrNewReplacement_example" // string | The new value in the NAPTR record. This parameter when missing will use the old value. This parameter is used when updating the NAPTR record  (optional)
	dname := "dname_example" // string | The DNAME domain name. This parameter is required when updating the DNAME record  (optional)
	keyTag := "keyTag_example" // string | This parameter is required when updating DS record  (optional)
	newKeyTag := "newKeyTag_example" // string | This parameter is required when updating DS record  (optional)
	algorithm := "algorithm_example" // string | This parameter is required when updating DS record  (optional)
	newAlgorithm := "newAlgorithm_example" // string | This parameter is required when updating DS record  (optional)
	digestType := "digestType_example" // string | This parameter is required when updating DS record  (optional)
	newDigestType := "newDigestType_example" // string | This parameter is required when updating DS record  (optional)
	digest := "digest_example" // string | This parameter is required when updating DS record  (optional)
	newDigest := "newDigest_example" // string | This parameter is required when updating DS record  (optional)
	sshfpAlgorithm := "sshfpAlgorithm_example" // string | This parameter is required when updating SSHFP record  (optional)
	newSshfpAlgorithm := "newSshfpAlgorithm_example" // string | This parameter is required when updating SSHFP record  (optional)
	sshfpFingerprintType := "sshfpFingerprintType_example" // string | This parameter is required when updating SSHFP record  (optional)
	newSshfpFingerprintType := "newSshfpFingerprintType_example" // string | This parameter is required when updating SSHFP record  (optional)
	sshfpFingerprint := "sshfpFingerprint_example" // string | This parameter is required when updating SSHFP record  (optional)
	newSshfpFingerprint := "newSshfpFingerprint_example" // string | This parameter is required when updating SSHFP record  (optional)
	tlsaCertificateUsage := "tlsaCertificateUsage_example" // string | This parameter is required when updating TLSA record  (optional)
	newTlsaCertificateUsage := "newTlsaCertificateUsage_example" // string | This parameter is required when updating TLSA record  (optional)
	tlsaSelector := "tlsaSelector_example" // string | This parameter is required when updating TLSA record  (optional)
	newTlsaSelector := "newTlsaSelector_example" // string | This parameter is required when updating TLSA record  (optional)
	tlsaMatchingType := "tlsaMatchingType_example" // string | This parameter is required when updating TLSA record  (optional)
	newTlsaMatchingType := "newTlsaMatchingType_example" // string | This parameter is required when updating TLSA record  (optional)
	tlsaCertificateAssociationData := "tlsaCertificateAssociationData_example" // string | This parameter is required when updating TLSA record  (optional)
	newTlsaCertificateAssociationData := "newTlsaCertificateAssociationData_example" // string | This parameter is required when updating TLSA record  (optional)
	svcPriority := "svcPriority_example" // string | The priority value for SVCB or HTTPS record. This parameter is required for updating SCVB or HTTPS record  (optional)
	newSvcPriority := "newSvcPriority_example" // string | The new priority value for SVCB or HTTPS record. This parameter when  missing will use the old value.  (optional)
	svcTargetName := "svcTargetName_example" // string | The target domain name for SVCB or HTTPS record. This parameter is required for updating SCVB or HTTPS record  (optional)
	newSvcTargetName := "newSvcTargetName_example" // string | The new target domain name for SVCB or HTTPS record. This parameter when missing will use the old value  (optional)
	svcParams := "svcParams_example" // string | The service parameters for SVCB or HTTPS record which is a pipe separated list of key and value. For example, alpn|h2,h3|port|53443. To clear existing values,  set it to false. This parameter is required for updating SCVB or HTTPS record  (optional)
	newSvcParams := "newSvcParams_example" // string | The new service parameters for SVCB or HTTPS record which is a pipe separated list of key and value. To clear existing values, set it to false. This parameter when missing will use the old value  (optional)
	autoIpv4Hint := true // bool | Set this option to true to enable Automatic Hints for the ipv4hint parameter  in the newSvcParams. This option is valid only for SVCB and HTTPS records  (optional)
	autoIpv6Hint := true // bool | Set this option to true to enable Automatic Hints for the ipv6hint parameter in the newSvcParams. This option is valid only for SVCB and HTTPS records  (optional)
	uriPriority := int32(56) // int32 | The priority value for the URI record. This parameter is required for updating the URI record  (optional)
	newUriPriority := int32(56) // int32 | The new priority value for the URI record. This parameter when missing will use the old value  (optional)
	uriWeight := int32(56) // int32 | The weight value for the URI record. This parameter is required for updating the URI record  (optional)
	newUriWeight := int32(56) // int32 | The new weight value for the URI record. This parameter when  missing will use the old value  (optional)
	uri := "uri_example" // string | The URI value for the URI record. This parameter is required for updating the URI record  (optional)
	newUri := "newUri_example" // string | The new URI value for the URI record. This parameter when missing  will use the old value  (optional)
	flags := int32(56) // int32 | This is the flags parameter in the CAA record. This parameter is required when updating the CAA record  (optional)
	newFlags := int32(56) // int32 | This is the new value of the flags parameter in the CAA record. This parameter  is used to update the flags parameter in the CAA record  (optional)
	tag := "tag_example" // string | This is the tag parameter in the CAA record. This parameter is required  when updating the CAA record.  (optional)
	newTag := "newTag_example" // string | This is the new value of the tag parameter in the CAA record. This parameter  is used to update the tag parameter in the CAA record  (optional)
	value := "value_example" // string | The current value in CAA record. This parameter is required when updating the CAA record  (optional)
	newValue := "newValue_example" // string | The new value in CAA record. This parameter is required when updating the CAA record  (optional)
	aname := "aname_example" // string | The current ANAME domain name. This parameter is required when updating  the ANAME record  (optional)
	newAName := "newAName_example" // string | The new ANAME domain name. This parameter is required when updating the ANAME record  (optional)
	protocol := "protocol_example" // string | This is the current protocol value in the FWD record. Valid values are [Udp, Tcp, Tls, Https, Quic]. This parameter is optional and default value Udp will be used when updating the FWD record  (optional)
	newProtocol := "newProtocol_example" // string | This is the new protocol value in the FWD record. Valid values are [Udp, Tcp, Tls, Https, Quic].  This parameter is optional and default value Udp will be used when updating the FWD record.  (optional)
	forwarder := "forwarder_example" // string | The current forwarder address. This parameter is required when updating the FWD record  (optional)
	newForwarder := "newForwarder_example" // string | The new forwarder address. This parameter is required when updating the FWD record  (optional)
	dnssecValidation := true // bool | Set this boolean value to indicate if DNSSEC validation must be done.  This optional parameter is to be used with FWD records. Default value is false  (optional)
	proxyType := "proxyType_example" // string | The type of proxy that must be used for conditional forwarding. This optional parameter is to be used with FWD records. Valid values are [NoProxy, DefaultProxy, Http, Socks5]. Default value DefaultProxy is used when this parameter is missing  (optional)
	proxyAddress := "proxyAddress_example" // string | The proxy server address to use when proxyType is configured. This optional parameter is to be used with FWD records  (optional)
	proxyPort := "proxyPort_example" // string | The proxy server port to use when proxyType is configured. This optional parameter  is to be used with FWD records.  (optional)
	proxyUsername := "proxyUsername_example" // string | The proxy server username to use when proxyType is configured. This optional  parameter is to be used with FWD records.  (optional)
	proxyPassword := "proxyPassword_example" // string | The proxy server password to use when proxyType is configured. This optional parameter is to be used with FWD records.  (optional)
	appName := "appName_example" // string | This parameter is required for updating the APP record  (optional)
	classPath := "classPath_example" // string | This parameter is required for updating the APP record  (optional)
	recordData := "recordData_example" // string | This parameter is used for updating the APP record as per the DNS app requirements  (optional)
	rdata := "rdata_example" // string | This parameter is used for updating unknown i.e. unsupported record types.  The value must be formatted as a hex string or a colon separated hex string  (optional)
	newRData := "newRData_example" // string | This parameter is used for updating unknown i.e. unsupported record types. The new value that must be formatted as a hex string or a colon separated hex string  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsRecordAPI.UpdateDnsRecord(context.Background()).Domain(domain).NewDomain(newDomain).Zone(zone).Type_(type_).Ttl(ttl).Disable(disable).Comments(comments).IpAddress(ipAddress).NewIpAddress(newIpAddress).Ptr(ptr).CreatePtrZone(createPtrZone).UpdateSvcbHints(updateSvcbHints).NameServer(nameServer).NewNameServer(newNameServer).Glue(glue).Cname(cname).PrimaryNameServer(primaryNameServer).ResponsiblePerson(responsiblePerson).Serial(serial).Refresh(refresh).Retry(retry).Expire(expire).Minimum(minimum).PrimaryAddresses(primaryAddresses).ZoneTransferProtocol(zoneTransferProtocol).TsigKeyName(tsigKeyName).PtrName(ptrName).NewPtrName(newPtrName).Preference(preference).NewPreference(newPreference).Exchange(exchange).NewExchange(newExchange).Text(text).NewText(newText).SplitText(splitText).NewSplitText(newSplitText).Priority(priority).NewPriority(newPriority).Weight(weight).NewWeight(newWeight).Port(port).NewPort(newPort).Target(target).NewTarget(newTarget).NaptrOrder(naptrOrder).NaptrNewOrder(naptrNewOrder).NaptrPreference(naptrPreference).NaptrNewPreference(naptrNewPreference).NaptrFlags(naptrFlags).NaptrNewFlags(naptrNewFlags).NaptrServices(naptrServices).NaptrNewServices(naptrNewServices).NaptrRegexp(naptrRegexp).NaptrNewRegexp(naptrNewRegexp).NaptrReplacement(naptrReplacement).NaptrNewReplacement(naptrNewReplacement).Dname(dname).KeyTag(keyTag).NewKeyTag(newKeyTag).Algorithm(algorithm).NewAlgorithm(newAlgorithm).DigestType(digestType).NewDigestType(newDigestType).Digest(digest).NewDigest(newDigest).SshfpAlgorithm(sshfpAlgorithm).NewSshfpAlgorithm(newSshfpAlgorithm).SshfpFingerprintType(sshfpFingerprintType).NewSshfpFingerprintType(newSshfpFingerprintType).SshfpFingerprint(sshfpFingerprint).NewSshfpFingerprint(newSshfpFingerprint).TlsaCertificateUsage(tlsaCertificateUsage).NewTlsaCertificateUsage(newTlsaCertificateUsage).TlsaSelector(tlsaSelector).NewTlsaSelector(newTlsaSelector).TlsaMatchingType(tlsaMatchingType).NewTlsaMatchingType(newTlsaMatchingType).TlsaCertificateAssociationData(tlsaCertificateAssociationData).NewTlsaCertificateAssociationData(newTlsaCertificateAssociationData).SvcPriority(svcPriority).NewSvcPriority(newSvcPriority).SvcTargetName(svcTargetName).NewSvcTargetName(newSvcTargetName).SvcParams(svcParams).NewSvcParams(newSvcParams).AutoIpv4Hint(autoIpv4Hint).AutoIpv6Hint(autoIpv6Hint).UriPriority(uriPriority).NewUriPriority(newUriPriority).UriWeight(uriWeight).NewUriWeight(newUriWeight).Uri(uri).NewUri(newUri).Flags(flags).NewFlags(newFlags).Tag(tag).NewTag(newTag).Value(value).NewValue(newValue).Aname(aname).NewAName(newAName).Protocol(protocol).NewProtocol(newProtocol).Forwarder(forwarder).NewForwarder(newForwarder).DnssecValidation(dnssecValidation).ProxyType(proxyType).ProxyAddress(proxyAddress).ProxyPort(proxyPort).ProxyUsername(proxyUsername).ProxyPassword(proxyPassword).AppName(appName).ClassPath(classPath).RecordData(recordData).Rdata(rdata).NewRData(newRData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsRecordAPI.UpdateDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDnsRecord`: UpdateRecordResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsRecordAPI.UpdateDnsRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | The domain name of the zone to update the record | 
 **newDomain** | **string** | The new domain name to be set for the record. To be used to rename sub domain name of the record  | 
 **zone** | **string** | The name of the authoritative zone into which the domain exists. When unspecified, the closest authoritative zone will be used  | 
 **type_** | **string** | The type of the resource record to update  | 
 **ttl** | **int32** | The TTL value of the resource record. Default value of 3600 is used when  parameter is missing  | 
 **disable** | **bool** | Specifies if the record should be disabled. The default value is false when this parameter is missing  | 
 **comments** | **string** | Sets comments for the updated resource record | 
 **ipAddress** | **string** | The current IP address in the A or AAAA record. This parameter is required when updating A or AAAA record  | 
 **newIpAddress** | **string** | The new IP address in the A or AAAA record. This parameter when missing will use the current value in the record.  | 
 **ptr** | **bool** | Set this option to true to specify if the PTR record associated with the A or AAAA record must also be updated. This option is used only for A and AAAA records  | 
 **createPtrZone** | **bool** | Set this option to true to create a reverse zone for PTR record. This option is used  only for A and AAAA records.  | 
 **updateSvcbHints** | **bool** | Set this option to true to update any SVCB/HTTPS records in the zone that has Automatic Hints  option enabled and matches its target name with the current record&#39;s domain name. This option is used for A and AAAA records  | 
 **nameServer** | **string** | The current name server domain name. This option is required for updating NS record  | 
 **newNameServer** | **string** | The new server domain name. This option is used for updating NS record  | 
 **glue** | **string** | The comma separated list of IP addresses set as glue for the NS record. This parameter  is used only when updating NS record  | 
 **cname** | **string** | The CNAME domain name to update in the existing CNAME record  | 
 **primaryNameServer** | **string** | This is the primary name server parameter in the SOA record. This parameter  is required when updating the SOA record.  | 
 **responsiblePerson** | **string** | This is the responsible person parameter in the SOA record. This parameter  is required when updating the SOA record  | 
 **serial** | **int32** | This is the serial parameter in the SOA record. This parameter is required  when updating the SOA record  | 
 **refresh** | **int32** | This is the refresh parameter in the SOA record. This parameter is required  when updating the SOA record  | 
 **retry** | **int32** | This is the retry parameter in the SOA record. This parameter is required  when updating the SOA record  | 
 **expire** | **int32** | This is the expire parameter in the SOA record. This parameter is required when updating the SOA record  | 
 **minimum** | **int32** | This is the minimum parameter in the SOA record. This parameter is required when updating the SOA record  | 
 **primaryAddresses** | **[]string** | This is a comma separated list of IP addresses of the primary name server. This parameter is to be used with secondary and stub zones where the primary name server address is not directly resolvable  | 
 **zoneTransferProtocol** | **string** | The zone transfer protocol to be used by the secondary zone. Valid values are [Tcp, Tls, Quic]. This parameter is used with SOA record  | 
 **tsigKeyName** | **string** | The TSIG key name to be used by the secondary zone. This parameter is used with SOA record  | 
 **ptrName** | **string** | The current PTR domain name. This option is required for updating PTR record  | 
 **newPtrName** | **string** | The new PTR domain name. This option is required for updating PTR record  | 
 **preference** | **string** | The current preference value in an MX record. This parameter when missing will default to 1 value. This parameter is used only when updating MX record  | 
 **newPreference** | **string** | The new preference value in an MX record. This parameter when missing will use the old value. This parameter is used only when updating MX record  | 
 **exchange** | **string** | The current exchange domain name. This option is required for updating MX record  | 
 **newExchange** | **string** | The new exchange domain name. This option is required for updating MX record  | 
 **text** | **string** | The current text value. This option is required for updating TXT record  | 
 **newText** | **string** | The new text value. This option is required for updating TXT record  | 
 **splitText** | **bool** | The current split text value. This option is used for updating TXT record and is set to false when unspecified  | 
 **newSplitText** | **bool** | The new split text value. This option is used for updating TXT record and is set to current split text value when unspecified  | 
 **priority** | **int32** | This is the current priority in the SRV record. This parameter is required when updating the SRV record  | 
 **newPriority** | **int32** | This is the new priority in the SRV record. This parameter when missing will use the old value. This parameter is used when updating the SRV record  | 
 **weight** | **int32** | This is the current weight in the SRV record. This parameter is required when updating the SRV record  | 
 **newWeight** | **int32** | This is the new weight in the SRV record. This parameter when missing will use the old value. This parameter is used when updating the SRV record  | 
 **port** | **int32** | This is the port parameter in the SRV record. This parameter is required when updating the SRV record  | 
 **newPort** | **int32** | This is the port parameter in the SRV record. This parameter is required when updating the SRV record  | 
 **target** | **string** | The current target value. This parameter is required when updating the SRV record  | 
 **newTarget** | **string** | The new target value. This parameter when missing will use the old value. This parameter is required when updating the SRV record  | 
 **naptrOrder** | **int32** | The current value in the NAPTR record. This parameter is required when  updating the NAPTR record  | 
 **naptrNewOrder** | **int32** | The new value in the NAPTR record. This parameter when missing will use the old value. This parameter is used when updating the NAPTR record  | 
 **naptrPreference** | **int32** | The current value in the NAPTR record. This parameter is required when updating the NAPTR record  | 
 **naptrNewPreference** | **int32** | The new value in the NAPTR record. This parameter when missing will use the old value. This parameter is used when updating the NAPTR record  | 
 **naptrFlags** | **string** | The current value in the NAPTR record. This parameter is required when updating the NAPTR record  | 
 **naptrNewFlags** | **string** | The new value in the NAPTR record. This parameter when missing will use the old value. This parameter is used when updating the NAPTR record  | 
 **naptrServices** | **string** | The current value in the NAPTR record. This parameter is required when updating the NAPTR record  | 
 **naptrNewServices** | **string** | The new value in the NAPTR record. This parameter when missing will use the old value. This parameter is used when updating the NAPTR record  | 
 **naptrRegexp** | **string** | The current value in the NAPTR record. This parameter is required when updating the NAPTR record  | 
 **naptrNewRegexp** | **string** | The new value in the NAPTR record. This parameter when missing will use the old value. This parameter is used when updating the NAPTR record  | 
 **naptrReplacement** | **string** | The current value in the NAPTR record. This parameter is required when updating the NAPTR record  | 
 **naptrNewReplacement** | **string** | The new value in the NAPTR record. This parameter when missing will use the old value. This parameter is used when updating the NAPTR record  | 
 **dname** | **string** | The DNAME domain name. This parameter is required when updating the DNAME record  | 
 **keyTag** | **string** | This parameter is required when updating DS record  | 
 **newKeyTag** | **string** | This parameter is required when updating DS record  | 
 **algorithm** | **string** | This parameter is required when updating DS record  | 
 **newAlgorithm** | **string** | This parameter is required when updating DS record  | 
 **digestType** | **string** | This parameter is required when updating DS record  | 
 **newDigestType** | **string** | This parameter is required when updating DS record  | 
 **digest** | **string** | This parameter is required when updating DS record  | 
 **newDigest** | **string** | This parameter is required when updating DS record  | 
 **sshfpAlgorithm** | **string** | This parameter is required when updating SSHFP record  | 
 **newSshfpAlgorithm** | **string** | This parameter is required when updating SSHFP record  | 
 **sshfpFingerprintType** | **string** | This parameter is required when updating SSHFP record  | 
 **newSshfpFingerprintType** | **string** | This parameter is required when updating SSHFP record  | 
 **sshfpFingerprint** | **string** | This parameter is required when updating SSHFP record  | 
 **newSshfpFingerprint** | **string** | This parameter is required when updating SSHFP record  | 
 **tlsaCertificateUsage** | **string** | This parameter is required when updating TLSA record  | 
 **newTlsaCertificateUsage** | **string** | This parameter is required when updating TLSA record  | 
 **tlsaSelector** | **string** | This parameter is required when updating TLSA record  | 
 **newTlsaSelector** | **string** | This parameter is required when updating TLSA record  | 
 **tlsaMatchingType** | **string** | This parameter is required when updating TLSA record  | 
 **newTlsaMatchingType** | **string** | This parameter is required when updating TLSA record  | 
 **tlsaCertificateAssociationData** | **string** | This parameter is required when updating TLSA record  | 
 **newTlsaCertificateAssociationData** | **string** | This parameter is required when updating TLSA record  | 
 **svcPriority** | **string** | The priority value for SVCB or HTTPS record. This parameter is required for updating SCVB or HTTPS record  | 
 **newSvcPriority** | **string** | The new priority value for SVCB or HTTPS record. This parameter when  missing will use the old value.  | 
 **svcTargetName** | **string** | The target domain name for SVCB or HTTPS record. This parameter is required for updating SCVB or HTTPS record  | 
 **newSvcTargetName** | **string** | The new target domain name for SVCB or HTTPS record. This parameter when missing will use the old value  | 
 **svcParams** | **string** | The service parameters for SVCB or HTTPS record which is a pipe separated list of key and value. For example, alpn|h2,h3|port|53443. To clear existing values,  set it to false. This parameter is required for updating SCVB or HTTPS record  | 
 **newSvcParams** | **string** | The new service parameters for SVCB or HTTPS record which is a pipe separated list of key and value. To clear existing values, set it to false. This parameter when missing will use the old value  | 
 **autoIpv4Hint** | **bool** | Set this option to true to enable Automatic Hints for the ipv4hint parameter  in the newSvcParams. This option is valid only for SVCB and HTTPS records  | 
 **autoIpv6Hint** | **bool** | Set this option to true to enable Automatic Hints for the ipv6hint parameter in the newSvcParams. This option is valid only for SVCB and HTTPS records  | 
 **uriPriority** | **int32** | The priority value for the URI record. This parameter is required for updating the URI record  | 
 **newUriPriority** | **int32** | The new priority value for the URI record. This parameter when missing will use the old value  | 
 **uriWeight** | **int32** | The weight value for the URI record. This parameter is required for updating the URI record  | 
 **newUriWeight** | **int32** | The new weight value for the URI record. This parameter when  missing will use the old value  | 
 **uri** | **string** | The URI value for the URI record. This parameter is required for updating the URI record  | 
 **newUri** | **string** | The new URI value for the URI record. This parameter when missing  will use the old value  | 
 **flags** | **int32** | This is the flags parameter in the CAA record. This parameter is required when updating the CAA record  | 
 **newFlags** | **int32** | This is the new value of the flags parameter in the CAA record. This parameter  is used to update the flags parameter in the CAA record  | 
 **tag** | **string** | This is the tag parameter in the CAA record. This parameter is required  when updating the CAA record.  | 
 **newTag** | **string** | This is the new value of the tag parameter in the CAA record. This parameter  is used to update the tag parameter in the CAA record  | 
 **value** | **string** | The current value in CAA record. This parameter is required when updating the CAA record  | 
 **newValue** | **string** | The new value in CAA record. This parameter is required when updating the CAA record  | 
 **aname** | **string** | The current ANAME domain name. This parameter is required when updating  the ANAME record  | 
 **newAName** | **string** | The new ANAME domain name. This parameter is required when updating the ANAME record  | 
 **protocol** | **string** | This is the current protocol value in the FWD record. Valid values are [Udp, Tcp, Tls, Https, Quic]. This parameter is optional and default value Udp will be used when updating the FWD record  | 
 **newProtocol** | **string** | This is the new protocol value in the FWD record. Valid values are [Udp, Tcp, Tls, Https, Quic].  This parameter is optional and default value Udp will be used when updating the FWD record.  | 
 **forwarder** | **string** | The current forwarder address. This parameter is required when updating the FWD record  | 
 **newForwarder** | **string** | The new forwarder address. This parameter is required when updating the FWD record  | 
 **dnssecValidation** | **bool** | Set this boolean value to indicate if DNSSEC validation must be done.  This optional parameter is to be used with FWD records. Default value is false  | 
 **proxyType** | **string** | The type of proxy that must be used for conditional forwarding. This optional parameter is to be used with FWD records. Valid values are [NoProxy, DefaultProxy, Http, Socks5]. Default value DefaultProxy is used when this parameter is missing  | 
 **proxyAddress** | **string** | The proxy server address to use when proxyType is configured. This optional parameter is to be used with FWD records  | 
 **proxyPort** | **string** | The proxy server port to use when proxyType is configured. This optional parameter  is to be used with FWD records.  | 
 **proxyUsername** | **string** | The proxy server username to use when proxyType is configured. This optional  parameter is to be used with FWD records.  | 
 **proxyPassword** | **string** | The proxy server password to use when proxyType is configured. This optional parameter is to be used with FWD records.  | 
 **appName** | **string** | This parameter is required for updating the APP record  | 
 **classPath** | **string** | This parameter is required for updating the APP record  | 
 **recordData** | **string** | This parameter is used for updating the APP record as per the DNS app requirements  | 
 **rdata** | **string** | This parameter is used for updating unknown i.e. unsupported record types.  The value must be formatted as a hex string or a colon separated hex string  | 
 **newRData** | **string** | This parameter is used for updating unknown i.e. unsupported record types. The new value that must be formatted as a hex string or a colon separated hex string  | 

### Return type

[**UpdateRecordResponse**](UpdateRecordResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

