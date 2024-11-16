package yougilego

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type YGAuthService struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// GetListCompany This method is designed to get a list of companies
func (authService *YGAuthService) GetListCompany(companyName string) (err error, response ListResponse[getCompanyResponse]) {
	url := "https://ru.yougile.com/api-v2/auth/companies"
	getListCompanyRequest := getCompanyListRequest{
		Login:    authService.Login,
		Password: authService.Password,
		Name:     companyName,
	}
	payloadByte, err := json.Marshal(getListCompanyRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetListCompany StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return err, response
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

// GetKeysList This method is designed to get a list of keys
func (authService *YGAuthService) GetKeysList(companyId string) (err error, response []getKeysListResponse) {
	url := "https://ru.yougile.com/api-v2/auth/keys/get"
	payload := AuthRequest{
		Login:     authService.Login,
		Password:  authService.Password,
		CompanyId: companyId,
	}
	payloadByte, err := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetKeysList StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return err, response
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

// CreateKey This method is designed to create a company Key
func (authService *YGAuthService) CreateKey(companyId string) (err error, response string) {
	url := "https://ru.yougile.com/api-v2/auth/keys"
	payload := AuthRequest{
		Login:     authService.Login,
		Password:  authService.Password,
		CompanyId: companyId,
	}
	payloadByte, err := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("CreateKey StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return err, response
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	keyResponse := createKeyResponse{}
	json.Unmarshal(body, &keyResponse)
	response = keyResponse.Key
	return
}

// DeleteKey This method is designed to delete the company Key
func (authService *YGAuthService) DeleteKey(response string) {
	url := "https://ru.yougile.com/api-v2/auth/keys/" + response
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
}

type AuthRequest struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	CompanyId string `json:"companyId"`
}

type getCompanyListRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type getCompanyResponse struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"isAdmin"`
}

type getKeysListResponse struct {
	Key       string `json:"Key"`
	CompanyId string `json:"companyId"`
	Timestamp string `json:"timestamp"`
	Deleted   bool   `json:"deleted"`
}

type createKeyResponse struct {
	Key string `json:"key"`
}
