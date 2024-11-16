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

type YGEventSubscribeService struct {
	Key string `json:"key"`
}

func (eventSubscribeService *YGEventSubscribeService) UseKey() string {
	return fmt.Sprintf("Bearer %s", eventSubscribeService.Key)
}

func (eventSubscribeService *YGEventSubscribeService) CreateSubscribe(createSubsRequest CreateSubs) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/webhooks"
	payloadByte, _ := json.Marshal(createSubsRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", eventSubscribeService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("CreateSubscribe StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (eventSubscribeService *YGEventSubscribeService) GetSubscribeList(includeDeleted bool) (err error, response SubscribeResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/webhooks?includeDeleted=%b", includeDeleted)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", eventSubscribeService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetSubscribeList StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (eventSubscribeService *YGEventSubscribeService) EditSubscribe(subsId string, editSubsRequest EditSubsRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/webhooks/" + subsId
	payloadByte, _ := json.Marshal(editSubsRequest)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", eventSubscribeService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditSubscribe StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

type CreateSubs struct {
	Url   string `json:"url"`
	Event string `json:"event"`
}

type SubscribeResponse struct {
	Id                       string `json:"id"`
	Deleted                  bool   `json:"deleted"`
	Url                      string `json:"url"`
	Event                    string `json:"event"`
	Disabled                 bool   `json:"disabled"`
	LastSuccess              int    `json:"lastSuccess"`
	FailuresSinceLastSuccess int    `json:"failuresSinceLastSuccess"`
}

type EditSubsRequest struct {
	Deleted  bool   `json:"deleted"`
	Url      string `json:"url"`
	Event    string `json:"event"`
	Disabled bool   `json:"disabled"`
}
