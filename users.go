package yougilego

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type YGUsersService struct {
	YGEngine `json:"YGEngine"`
}

type UsersResponse struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	IsAdmin      bool   `json:"isAdmin"`
	RealName     string `json:"realName"`
	Status       string `json:"status"`
	LastActivity string `json:"lastActivity"`
}

func (userService *YGUsersService) GetAdmins() (err error, adminsIDList []string) {
	err, userList := userService.GetUsers()
	if err != nil {
		return errors.New("ошибка при получении списка пользователей"), adminsIDList
	}
	for _, usr := range userList.Content {
		if usr.IsAdmin {
			adminsIDList = append(adminsIDList, usr.Id)
		}
	}
	return
}

func (userService *YGUsersService) GetUsers() (err error, users ListResponse[UsersResponse]) {
	url := "https://ru.yougile.com/api-v2/users"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", userService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &users)
	return
}
