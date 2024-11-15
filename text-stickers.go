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

type YGTextStickerService struct {
	Key string `json:"key"`
}

func (textStickerService *YGTextStickerService) UseKey() string {
	return fmt.Sprintf("Bearer %s", textStickerService.Key)
}

func (textStickerService *YGTextStickerService) GetStickersList() (err error, response ListResponse[TextSticker]) {
	url := "https://ru.yougile.com/api-v2/string-stickers"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", textStickerService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetStickersList StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (textStickerService *YGTextStickerService) CreateTextSticker(createTextStickerRequest TextStickerRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/string-stickers"
	payloadByte, _ := json.Marshal(createTextStickerRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", textStickerService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("CreateTextSticker StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (textStickerService *YGTextStickerService) GetTextStickerById(stickerId string) (err error, response TextStickerResponse) {
	url := "https://ru.yougile.com/api-v2/string-stickers/" + stickerId
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", textStickerService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetTextStickerById StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (textStickerService *YGTextStickerService) EditTextSticker(textStickerId string, editTextStickerRequest EditTextStickerRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/string-stickers/" + textStickerId
	payloadByte, _ := json.Marshal(editTextStickerRequest)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", textStickerService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditTextSticker StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

type TextStickerResponse struct {
	Id      string `json:"id"`
	Deleted bool   `json:"deleted"`
	Name    string `json:"name"`
	Icon    string `json:"icon"`
	States  []struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"states"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type EditTextStickerRequest struct {
	Deleted bool   `json:"deleted"`
	Name    string `json:"name"`
	Icon    string `json:"icon"`
}

type TextStickerRequest struct {
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	States []struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"states"`
}

type TextSticker struct {
	Id      string `json:"id"`
	Deleted bool   `json:"deleted"`
	Name    string `json:"name"`
	Icon    string `json:"icon"`
	States  []struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"states"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
