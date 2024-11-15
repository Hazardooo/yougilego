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

type YGUsersService struct {
	Key string `json:"token"`
}

func (userService *YGUsersService) UseKey() string {
	return fmt.Sprintf("Bearer %s", userService.Key)
}

func (userService *YGUsersService) GetUsers() (err error, users ListResponse[UsersResponse]) {
	url := "https://ru.yougile.com/api-v2/users"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", userService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetUsers StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &users)
	return
}

func (userService *YGUsersService) InviteUser(email string, isAdmin bool) (err error, userId IDResponse) {
	url := "https://ru.yougile.com/api-v2/users"
	payload := SendInviteRequest{
		Email:   email,
		IsAdmin: isAdmin,
	}
	payloadByte, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", userService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("InviteUser StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &userId)
	return
}

func (userService *YGUsersService) GetUserById(userId string) (err error, user UsersResponse) {
	url := "https://ru.yougile.com/api-v2/users/" + userId
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", userService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetUserById StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &user)
	return
}

func (userService *YGUsersService) EditUserPerm(userId string, isAdmin bool) (err error, userIdResponse IDResponse) {
	url := "https://ru.yougile.com/api-v2/users/" + userId
	payload := strings.NewReader(fmt.Sprintf("{\n  \"isAdmin\": %b\n}", isAdmin))
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", userService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditUserPerm StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &userIdResponse)
	return
}

func (userService *YGUsersService) DeleteFromCompany(userId string) (err error, response bool) {
	url := "https://ru.yougile.com/api-v2/users/" + userId
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", userService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("DeleteFromCompany StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return err, false
	}
	defer res.Body.Close()
	return err, true
}

type UsersResponse struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	IsAdmin      bool   `json:"isAdmin"`
	RealName     string `json:"realName"`
	Status       string `json:"status"`
	LastActivity string `json:"lastActivity"`
}

type SendInviteRequest struct {
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
}

type IDResponse struct {
	Id string `json:"id"`
}
