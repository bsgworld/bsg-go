package bsgrest

type ViberClient struct {
	*BaseClient
}
type ViberRequest struct {
	// Tariff number. An integer from 0 to 9
	Tariff *string `json:"tariff"`
	// Validity period. The period must be between 0 and 72 hours
	Validity string `json:"validity"`
	// Request messages
	ViberMessages []ViberMessage `json:"messages"`
}
type ViberMessage struct {
	// Recipients
	Recipients []Recipient `json:"to"`
	// Message text
	Text string `json:"text"`
	// Sender name
	AlphaName string `json:"alpha_name"`
	// Promotional attribute
	IsPromotional bool `json:"is_promotional"`
	// Additional options
	Options ViberOptions `json:"options"`
}
type ViberOptions struct {
	Viber ViberOptionsContainer `json:"viber"`
}
type Recipient struct {
	// Phone number
	Msisdn string `json:"msisdn"`
	// External id
	Reference string `json:"reference"`
}
type ViberOptionsContainer struct {
	Img     string `json:"img"`
	Caption string `json:"caption"`
	Action  string `json:"action"`
}
type ViberData struct {
	*BsgError
	// Created messages
	Messages []Message `json:"result"`
	// Currency name
	Currency string `json:"currency"`
	// Request total costs
	TotalPrice float64 `json:"total_price"`
}

const viber_create_method = "/viber/create"
const viber_get_prices_method = "/viber/prices"
const viber_get_info_by_id_method = "/viber"
const viber_get_info_by_reference_method = "/viber/reference"

func NewViberClient(serviceUrl string, apiKey string) *ViberClient {
	client := new(ViberClient)
	client.BaseClient = NewBaseClient(serviceUrl, apiKey)
	return client
}

// Retrieve the prices of default Viber tariff.
func (client *ViberClient) GetViberPrices() BsgPrices {
	var prices BsgPrices
	_DoJsonRequest(client.BaseClient, viber_get_prices_method, &prices)
	return prices
}

// Retrieve the prices of specified Viber tariff.
func (client *ViberClient) GetViberPricesById(id int) BsgPrices {
	var prices BsgPrices
	_DoJsonRequestByIntegerId(client.BaseClient, viber_get_prices_method, id, &prices)
	return prices
}

// Create a new Viber message.
func (client *ViberClient) CreateViber(request ViberRequest) ViberData {
	var viber ViberData
	_DoJsonCreateRequest(client.BaseClient, viber_create_method, request, &viber)
	return viber
}

// Retrieve the information of specific Viber message.
// ref - external ID
func (client *ViberClient) GetViberInfoByReference(ref string) MessageInfo {
	var viberInfo MessageInfo
	_DoJsonRequestByStringId(client.BaseClient, viber_get_info_by_reference_method, ref, &viberInfo)
	return viberInfo
}

// Retrieve the information of specific Viber message.
// id - internal ID
func (client *ViberClient) GetViberInfoById(id string) MessageInfo {
	var viberInfo MessageInfo
	_DoJsonRequestByStringId(client.BaseClient, viber_get_info_by_id_method, id, &viberInfo)
	return viberInfo
}
