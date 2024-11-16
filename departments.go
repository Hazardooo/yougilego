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

type YGDepartmentsService struct {
	Key string `json:"key"`
}

func (departService *YGDepartmentsService) UseKey() string {
	return fmt.Sprintf("Bearer %s", departService.Key)
}

func (departService *YGDepartmentsService) GetDepartmentsList() (err error, response ListResponse[DepartResponse]) {
	url := "https://ru.yougile.com/api-v2/departments"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", departService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetDepartmentsList StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (departService *YGDepartmentsService) CreateDepart(createDepartRequest DepartRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/departments"
	payloadByte, _ := json.Marshal(createDepartRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", departService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("CreateDepart StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (departService *YGDepartmentsService) GetDepartById(departId string) (err error, response DepartResponse) {
	url := "https://ru.yougile.com/api-v2/departments/" + departId
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", departService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetDepartById StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (departService *YGDepartmentsService) EditDepart(departId string, editDepartRequest EditDepartRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/departments/" + departId
	payloadByte, _ := json.Marshal(editDepartRequest)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", departService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditDepart StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

type DepartRequest struct {
	Title    string            `json:"title"`
	ParentId string            `json:"parentId"`
	Users    map[string]string `json:"users"`
}

type DepartResponse struct {
	Deleted  bool              `json:"deleted"`
	Id       string            `json:"id"`
	Title    string            `json:"title"`
	ParentId string            `json:"parentId"`
	Users    map[string]string `json:"users"`
}

type EditDepartRequest struct {
	Deleted  bool              `json:"deleted"`
	Title    string            `json:"title"`
	ParentId string            `json:"parentId"`
	Users    map[string]string `json:"users"`
}
