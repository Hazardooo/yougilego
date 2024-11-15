package yougilego

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type YGColumnService struct {
	Key string `json:"key"`
}

func (columnService *YGColumnService) UseKey() string {
	return fmt.Sprintf("Bearer %s", columnService.Key)
}

func (columnService *YGColumnService) GetColumn() (err error, columns ListResponse[ColumnResponse]) {
	url := "https://ru.yougile.com/api-v2/columns"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", columnService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetColumn StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &columns)
	return
}

func (columnService *YGColumnService) CreateColumn(createColumnRequest CreateColumn) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/columns"
	payloadByte, _ := json.Marshal(createColumnRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", columnService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("CreateColumn StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (columnService *YGColumnService) GetColumnById(columnId string) (err error, response ColumnResponse) {
	url := "https://ru.yougile.com/api-v2/columns/" + columnId
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", columnService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetColumnById StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (columnService *YGColumnService) EditColumn(columnId string, editColumnRequest EditColumnRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/columns/" + columnId
	payloadByte, _ := json.Marshal(editColumnRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", columnService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetProjectById StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

type CreateColumn struct {
	Title   string `json:"title"`
	Color   int    `json:"color"`
	BoardId string `json:"boardId"`
}

type EditColumnRequest struct {
	Deleted bool   `json:"deleted"`
	Title   string `json:"title"`
	Color   int    `json:"color"`
	BoardId string `json:"boardId"`
}

type ColumnResponse struct {
	Id      string `json:"id"`
	Deleted bool   `json:"deleted"`
	Title   string `json:"title"`
	Color   int    `json:"color"`
	BoardId string `json:"boardId"`
}
