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

type YGStatusTextStickerService struct {
	Key string `json:"key"`
}

func (statusTextStickerService *YGStatusTextStickerService) UseKey() string {
	return fmt.Sprintf("Bearer %s", statusTextStickerService.Key)
}

func (statusTextStickerService *YGStatusTextStickerService) GetStatusTextStickerById(stickerId string) (err error, response StatusTextStickerById) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/string-stickers/%s/states/stickerStateId", stickerId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", statusTextStickerService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetStatusTextStickerById StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (statusTextStickerService *YGStatusTextStickerService) EditStatusTextSticker(stickerId string, editStatusTextStickerRequest EditStatusTextStickerRequest) (err error, response IDResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/string-stickers/%s/states/stickerStateId", stickerId)
	payloadByte, _ := json.Marshal(editStatusTextStickerRequest)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", statusTextStickerService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditStatusTextSticker StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (statusTextStickerService *YGStatusTextStickerService) CreateStatusTextSticker(stickerId string, createStatusTextStickerRequest CreateStatusTextStickerRequest) (err error, response IDResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/string-stickers/%s/states", stickerId)
	payloadByte, _ := json.Marshal(createStatusTextStickerRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", statusTextStickerService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("CreateStatusTextSticker StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

type StatusTextStickerById struct {
	Deleted bool   `json:"deleted"`
	Id      string `json:"id"`
	Name    string `json:"name"`
	Color   string `json:"color"`
}

type EditStatusTextStickerRequest struct {
	Deleted bool   `json:"deleted"`
	Name    string `json:"name"`
	Color   string `json:"color"`
}

type CreateStatusTextStickerRequest struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}
