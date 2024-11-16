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

type YGSprintStickerService struct {
	Key string `json:"key"`
}

func (sprintStickerService *YGSprintStickerService) UseKey() string {
	return fmt.Sprintf("Bearer %s", sprintStickerService.Key)
}

func (sprintStickerService *YGSprintStickerService) GetSprintStickerList() (err error, response ListResponse[SprintSticker]) {
	url := "https://ru.yougile.com/api-v2/sprint-stickers"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", sprintStickerService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetSprintStickerList StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (sprintStickerService *YGSprintStickerService) CreateSprintSticker(createSprintStickerRequest CreateSprintStickerRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/sprint-stickers"
	payloadByte, _ := json.Marshal(createSprintStickerRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", sprintStickerService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("CreateSprintSticker StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (sprintStickerService *YGSprintStickerService) GetSprintStickerById(sprintStickerId string) (err error, response SprintSticker) {
	url := "https://ru.yougile.com/api-v2/sprint-stickers/" + sprintStickerId
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", sprintStickerService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetSprintStickerById StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (sprintStickerService *YGSprintStickerService) EditSprintSticker(sprintStickerId string, editSprintStickerRequest EditSprintStickerRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/sprint-stickers/" + sprintStickerId
	payloadByte, _ := json.Marshal(editSprintStickerRequest)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", sprintStickerService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditSprintSticker StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

type EditSprintStickerRequest struct {
	Deleted bool   `json:"deleted"`
	Name    string `json:"name"`
}

type SprintSticker struct {
	Id      string `json:"id"`
	Deleted bool   `json:"deleted"`
	Name    string `json:"name"`
	States  []struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"states"`
}

type CreateSprintStickerRequest struct {
	Name   string `json:"name"`
	States []struct {
		Name  string `json:"name"`
		Begin int    `json:"begin,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"states"`
}
