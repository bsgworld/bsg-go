// Package bsgrest provides basic access for BSG Rest Service
package bsgrest

import (
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"bytes"
)

// Default error
type BsgError struct {
	// Error Code
	Error int    `json:"error"`
	// Error description
	ErrorDescription string `json:"errorDescription"`
}

type BsgPrices struct {
	*BsgError
	Prices []struct {
		// Price type
		Type string `json:"type"`
		// Country A2 code
		Country string `json:"country"`
		// Country name
		CountryName string `json:"country_name"`
		Mcc         string `json:"mcc"`
		// Request costs
		Price string `json:"price"`
		// Currency name
		Currency string `json:"currency"`
	} `json:"prices"`
}
type Balance struct {
	*BsgError
	// Balance in the current time
	Amount string `json:"amount"`
	// Currency name
	Currency string `json:"currency"`
	// Credit limit
	Limit string `json:"limit"`
}
type BaseClient struct {
	// Service URL
	ServiceURL string
	// X-API-KEY
	ApiKey string
}
type MessageInfo struct {
	*BsgError
	// Service internal id
	ID string `json:"id"`
	// Phone number
	Msisdn string `json:"msisdn"`
	// External id
	Reference string `json:"reference"`
	TimeIn    string `json:"time_in"`
	TimeSent  string `json:"time_sent"`
	TimeDr    string `json:"time_dr"`
	// Request status
	Status string `json:"status"`
	// Request costs
	Price float64 `json:"price"`
	// Currency name
	Currency string `json:"currency"`
}
type Message struct {
	*BsgError
	// External id
	Reference string `json:"reference"`
	// Service internal id
	ID string `json:"id"`
	// Request costs
	Price float64 `json:"price"`
	// Currency name
	Currency string `json:"currency"`
}

const get_balance_method = "/common/balance"
const content_application_json = "application/json"
const content_multipart_form_data = "multipart/form-data"

func NewBaseClient(serviceUrl string, apiKey string) *BaseClient {
	var client = new(BaseClient)
	client.ServiceURL = serviceUrl
	client.ApiKey = apiKey
	return client
}

func _DoJsonRequestByStringId(baseClient *BaseClient, method string, id string, object interface{}) {
	_DoJsonRequest(baseClient, method+"/"+id, object)
}
func _DoJsonRequestByIntegerId(baseClient *BaseClient, method string, id int, object interface{}) {
	_DoJsonRequest(baseClient, method+"/"+strconv.Itoa(id), object)
}
func _DoJsonRequest(baseClient *BaseClient, method string, object interface{}) {
	req, _error :=
		http.NewRequest("GET", baseClient.ServiceURL+method, nil)
	if _error != nil {
		log.Fatal("NewRequest: ", _error)
	}

	req.Header.Set("X-API-KEY", baseClient.ApiKey)
	req.Header.Set("Content-Type", content_application_json)
	client := &http.Client{}
	resp, _error := client.Do(req)
	if _error != nil {
		log.Fatal("Do: ", _error)
	}

	if err := json.NewDecoder(resp.Body).Decode(&object); err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}
func _DoJsonCreateRequest(baseClient *BaseClient, method string, requestObject interface{}, object interface{}) {
	var requestBody, _ = json.Marshal(requestObject)
	req, _error :=
		http.NewRequest("PUT", baseClient.ServiceURL+method, bytes.NewBuffer(requestBody))
	if _error != nil {
		log.Fatal("NewRequest: ", _error)
	}

	req.Header.Set("X-API-KEY", baseClient.ApiKey)
	req.Header.Set("Content-Type", content_multipart_form_data)
	client := &http.Client{}
	resp, _error := client.Do(req)
	if _error != nil {
		log.Fatal("Do: ", _error)
	}

	if err := json.NewDecoder(resp.Body).Decode(&object); err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}

//Retrieve balance amount.
func (baseClient *BaseClient) GetBalance() Balance {
	var balance Balance
	_DoJsonRequest(baseClient, get_balance_method, &balance)
	return balance
}
