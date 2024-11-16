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

type YGStatusStickerSprintService struct {
	Key string `json:"key"`
}

func (statusStickerSprintService *YGStatusStickerSprintService) UseKey() string {
	return fmt.Sprintf("Bearer %s", statusStickerSprintService.Key)
}

func (statusStickerSprintService *YGStatusStickerSprintService) GetStatusStickerSprint(stickerId string, stickerStateId string) (err error, response StatusStickerSprint) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/sprint-stickers/%s/states/%s", stickerId, stickerStateId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", statusStickerSprintService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetStatusStickerSprint StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (statusStickerSprintService *YGStatusStickerSprintService) EditStatusStickerSprint(stickerId string, editStatusStickerRequest EditStatusStickerRequest) (err error, response IDResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/sprint-stickers/%s/states/stickerStateId", stickerId)
	payloadByte, _ := json.Marshal(editStatusStickerRequest)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", statusStickerSprintService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditStatusStickerSprint StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (statusStickerSprintService *YGStatusStickerSprintService) CreateStatusStickerSprint(stickerId string, createStatusStickerRequest CreateStatusStickerRequest) (err error, response IDResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/sprint-stickers/%s/states", stickerId)
	payloadByte, _ := json.Marshal(createStatusStickerRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", statusStickerSprintService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("CreateStatusStickerSprint StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

type CreateStatusStickerRequest struct {
	Name  string `json:"name"`
	Begin string `json:"begin"`
	End   string `json:"end"`
}

type EditStatusStickerRequest struct {
	Deleted bool   `json:"deleted"`
	Name    string `json:"name"`
	Begin   string `json:"begin"`
	End     string `json:"end"`
}

type StatusStickerSprint struct {
	Deleted bool   `json:"deleted"`
	Id      string `json:"id"`
	Name    string `json:"name"`
	Begin   string `json:"begin"`
	End     string `json:"end"`
}
