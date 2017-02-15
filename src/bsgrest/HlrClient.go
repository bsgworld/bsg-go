package bsgrest

import "time"

type HlrClient struct {
	*BaseClient
}
type HlrRequest struct {
	// Phone number
	Msisdn string `json:"msisdn"`
	// External id
	Reference string `json:"reference"`
	// Tariff number. An integer from 0 to 9
	Tariff      *string `json:"tariff"`
	CallbackUrl string `json:"callback_url"`
}
type HlrData struct {
	*BsgError
	Hlrs []struct {
		*BsgError
		// External id
		Reference string `json:"reference"`
		// Tariff number. An integer from 0 to 9
		TariffCode string `json:"tariff_code"`
		// URL for service response
		CallbackURL string `json:"callback_url"`
		// Request costs
		Price float64 `json:"price"`
		// Currency name
		Currency string `json:"currency"`
		// Service internal id
		ID string `json:"id"`
	} `json:"result"`
	// Request total costs
	TotalPrice float64 `json:"total_price"`
	// Currency name
	Currency string `json:"currency"`
}
type HlrInfo struct {
	*BsgError
	// Country name in Russian
	NameRu string `json:"name_ru"`
	// Country name in English
	NameEn string `json:"name_en"`
	Brand  string `json:"brand"`
	Name   string `json:"name"`
	// Phone number
	Msisdn string `json:"msisdn"`
	// Service internal id
	ID string `json:"id"`
	// External id
	Reference string `json:"reference"`
	// Phone mccmnc
	Network string `json:"network"`
	// Request status. Possible values: sent, absent, active, unknown,
	Status string `json:"status"`
	// Request details
	Details struct {
		Imsi        string `json:"imsi"`
		LocationMsc string `json:"location_msc"`
		Ported      int `json:"ported"`
		Roaming     int `json:"roaming"`
	} `json:"details"`
	CreatedDatetime time.Time `json:"createdDatetime"`
	StatusDatetime  time.Time `json:"statusDatetime"`
}

const hlr_create_method = "/hlr/create"
const hlr_get_prices_method = "/hlr/prices"
const hlr_get_info_by_id_method = "/hlr"
const hlr_get_info_by_reference_method = "/hlr/reference"

func NewHlrClient(serviceUrl string, apiKey string) *HlrClient {
	client := new(HlrClient)
	client.BaseClient = NewBaseClient(serviceUrl, apiKey)
	return client
}

// Retrieve the prices of default HLR tariff.
func (client *HlrClient) GetHlrPrices() BsgPrices {
	var prices BsgPrices
	_DoJsonRequest(client.BaseClient, hlr_get_prices_method, &prices)
	return prices
}

// Retrieve the prices of specified HLR tariff.
func (client *HlrClient) GetHlrPricesById(id int) BsgPrices {
	var prices BsgPrices
	_DoJsonRequestByIntegerId(client.BaseClient, hlr_get_prices_method, id, &prices)
	return prices
}

// Create a new HLR.
// msisdn - phone number
// reference - external id
// tariff - tariff number. An integer from 0 to 9. Set null for default
// callbackUrl - URL for service response
func (client *HlrClient) CreateHlrWithParameters(msisdn string, reference string, tariff *string, callbackUrl string) HlrData {
	var request HlrRequest
	request.CallbackUrl = callbackUrl
	request.Msisdn = msisdn
	request.Reference = reference
	request.Tariff = tariff
	return client.CreateHlr(request)
}

// Create a new HLR.
// hlrRequest - HLR request
func (client *HlrClient) CreateHlr(request HlrRequest) HlrData {
	var hlrRequests = []HlrRequest{request}
	return client.CreateHlrs(hlrRequests)
}

// Create a new HLR.
// hlrRequest - array of HLR requests
func (client *HlrClient) CreateHlrs(requests []HlrRequest) HlrData {
	var hlr HlrData
	_DoJsonCreateRequest(client.BaseClient, hlr_create_method, requests, &hlr)
	return hlr
}

// Retrieve the information of specific HLR.
// ref - external ID
func (client *HlrClient) GetHlrInfoByReference(ref string) HlrInfo {
	var info HlrInfo
	_DoJsonRequestByStringId(client.BaseClient, hlr_get_info_by_reference_method, ref, &info)
	return info
}

// Retrieve the information of specific HLR.
// id - internal ID
func (client *HlrClient) GetHlrInfoById(id string) HlrInfo {
	var info HlrInfo
	_DoJsonRequestByStringId(client.BaseClient, hlr_get_info_by_id_method, id, &info)
	return info
}
