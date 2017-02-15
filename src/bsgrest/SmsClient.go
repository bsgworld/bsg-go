package bsgrest

type SmsClient struct {
	*BaseClient
}
type SingleSmsRequest struct {
	// Destination. "phone"/"phones"
	Destination string `json:"destination"`
	// Sender name
	Originator string `json:"originator"`
	// Message text
	Body string `json:"body"`
	*SmsPhone
	// Validity period. The period must be between 0 and 72 hours
	Validity string `json:"validity"`
	// Tariff number. An integer from 0 to 9
	Tariff *string `json:"tariff"`
}
type MultipleSmsRequest struct {
	// Validity period. The period must be between 0 and 72 hours
	Validity string `json:"validity"`
	// Tariff number. An integer from 0 to 9
	Tariff *string `json:"tariff"`
	// Destination. "phone"/"phones"
	Destination string `json:"destination"`
	// Sender name
	Originator string `json:"originator"`
	// Message text
	Body string `json:"body"`
	// Recipients
	Phones []SmsPhone `json:"phones"`
}
type SmsPhone struct {
	// Phone number
	Msisdn string `json:"msisdn"`
	// External id
	Reference string `json:"reference"`
}
type MultipleSmsData struct {
	*BsgError
	// created task ID
	TaskID string `json:"task_id"`
	// Created SMS
	Smses []Message `json:"result"`
	// Request total costs
	TotalPrice float64 `json:"total_price"`
	// Currency name
	Currency string `json:"currency"`
}
type SingleSmsData struct {
	*BsgError
	// Created SMS
	Sms Message `json:"result"`
}
type SmsTaskInfo struct {
	*BsgError
	// Sender name
	Originator string `json:"originator"`
	// Message text
	Body string `json:"body"`
	// Validity period. The period must be between 0 and 72 hours
	Validity int `json:"validity"`
	// Request total costs
	Totalprice float64 `json:"totalprice"`
	// Currency name
	Currency string `json:"currency"`
	// Sent count
	Sent int `json:"sent"`
	// Delivered count
	Delivered int `json:"delivered"`
	// Expired count
	Expired int `json:"expired"`
	// Undeliverable count
	Undeliverable int `json:"undeliverable"`
	// Unknown count
	Unknown int `json:"unknown"`
}

const sms_create_method = "/sms/create"
const sms_get_prices_method = "/sms/prices"
const sms_get_info_by_id_method = "/sms"
const sms_get_task_info_by_id_method = "/sms/task"
const sms_get_info_by_reference_method = "/sms/reference"

func NewSmsClient(serviceUrl string, apiKey string) *SmsClient {
	client := new(SmsClient)
	client.BaseClient = NewBaseClient(serviceUrl, apiKey)
	return client
}

// Retrieve the prices of default Sms tariff.
func (client *SmsClient) GetSmsPrices() BsgPrices {
	var prices BsgPrices
	_DoJsonRequest(client.BaseClient, sms_get_prices_method, &prices)
	return prices
}

// Retrieve the prices of specified Sms tariff.
func (client *SmsClient) GetSmsPricesById(id int) BsgPrices {
	var prices BsgPrices
	_DoJsonRequestByIntegerId(client.BaseClient, sms_get_prices_method, id, &prices)
	return prices
}

// Create a new SMS.
func (client *SmsClient) CreateSingleSms(request SingleSmsRequest) SingleSmsData {
	var sms SingleSmsData
	request.Destination = "phone"
	_DoJsonCreateRequest(client.BaseClient, sms_create_method, request, &sms)
	return sms
}

// Create a new SMS.
// originator - sender name
// body - SMS text
// phone - phone number with reference
// validity - validity period. The period must be between 0 and 72 hours
// tariff - tariff number. An integer from 0 to 9. Set null for default
func (client *SmsClient) CreateSms(validity string, tariff *string, originator string, body string, phone *SmsPhone) SingleSmsData {
	return client.CreateSingleSms(SingleSmsRequest{Validity:validity, Tariff:tariff, Originator:originator, Body:body, SmsPhone:phone})
}

// Create a new SMSes.
func (client *SmsClient) CreateMultipleSms(request MultipleSmsRequest) MultipleSmsData {
	var sms MultipleSmsData
	request.Destination = "phones"
	_DoJsonCreateRequest(client.BaseClient, sms_create_method, request, &sms)
	return sms
}

// Retrieve the information of specific SMS package.
// taskId - task ID
func (client *SmsClient) GetSmsInfoByTaskId(taskId string) SmsTaskInfo {
	var sms SmsTaskInfo
	_DoJsonRequestByStringId(client.BaseClient, sms_get_task_info_by_id_method, taskId, &sms)
	return sms
}

// Retrieve the information of specific SMS.
// ref - external ID
func (client *SmsClient) GetSmsInfoByReference(ref string) MessageInfo {
	var sms MessageInfo
	_DoJsonRequestByStringId(client.BaseClient, sms_get_info_by_reference_method, ref, &sms)
	return sms
}

// Retrieve the information of specific SMS.
// id - internal ID
func (client *SmsClient) GetSmsInfoById(id string) MessageInfo {
	var sms MessageInfo
	_DoJsonRequestByStringId(client.BaseClient, sms_get_info_by_id_method, id, &sms)
	return sms
}
