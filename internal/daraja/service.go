package daraja

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"renewit-go/config"
	"time"
)

type DarajaService struct {
	ConsumerKey       string
	ConsumerSecret    string
	BusinessShortCode string
	PassKey           string
	BaseURL           string
	CallbackURL       string
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

func NewDarajaService() *DarajaService {
	return &DarajaService{
		ConsumerKey:       config.DarajaConsumerKey(),
		ConsumerSecret:    config.DarajaConsumerSecret(),
		BusinessShortCode: config.DarajaShortcode(),
		PassKey:           config.DarajaPasskey(),
		BaseURL:           "https://sandbox.safaricom.co.ke",
		CallbackURL:       config.DarajaCallbackURL(),
	}
}

func (d *DarajaService) GetAccessToken() (string, error) {

	url := d.BaseURL + "/oauth/v1/generate?grant_type=client_credentials"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.SetBasicAuth(
		d.ConsumerKey,
		d.ConsumerSecret,
	)

	client := &http.Client{}

	response, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var token AccessTokenResponse

	err = json.NewDecoder(response.Body).Decode(&token)

	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func (d *DarajaService) GeneratePassword() (string, string) {

	timestamp := time.Now().Format("20060102150405")

	rawPassword := d.BusinessShortCode + d.PassKey + timestamp

	encodedPassword := base64.StdEncoding.EncodeToString([]byte(rawPassword))

	return encodedPassword, timestamp
}

func (d *DarajaService) STKPush(
	phoneNumber string,
	amount int,
	accountReference string,
	transactionDesc string,
) (*STKPushResponse, error) {

	token, err := d.GetAccessToken()

	if err != nil {
		return nil, err
	}

	password, timestamp := d.GeneratePassword()

	request := STKPushRequest{
		BusinessShortCode: d.BusinessShortCode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            amount,
		PartyA:            phoneNumber,
		PartyB:            d.BusinessShortCode,
		PhoneNumber:       phoneNumber,
		CallBackURL:       d.CallbackURL,
		AccountReference:  accountReference,
		TransactionDesc:   transactionDesc,
	}

	jsonData, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	url := d.BaseURL + "/mpesa/stkpush/v1/processrequest"

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to read Daraja response: %v",
			err,
		)
	}

	fmt.Println("=================================")
	fmt.Println("DARAJA STK RESPONSE")
	fmt.Println("HTTP Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
	fmt.Println("=================================")

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf(
			"Daraja returned HTTP %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	var response STKPushResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf(
			"failed to decode Daraja response: %v",
			err,
		)
	}

	return &response, nil
}
