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

type YGRoleProjectService struct {
	Key string `json:"key"`
}

func (roleProjService *YGRoleProjectService) UseKey() string {
	return fmt.Sprintf("Bearer %s", roleProjService.Key)
}

func (roleProjService *YGRoleProjectService) GetRoleProjectList(projectId string) (err error, response ListResponse[RoleProjectResponse]) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/projects/%s/roles", projectId)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", roleProjService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetRoleProjectList StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (roleProjService *YGRoleProjectService) CreateRoleProject(projectId string, parameters RoleProjectRequest) (err error, response IDResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/projects/%s/roles", projectId)
	payloadByte, _ := json.Marshal(parameters)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", roleProjService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("CreateRoleProject StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (roleProjService *YGRoleProjectService) GetRoleProjById(roleProjId string, projId string) (err error, response RoleProjectResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/projects/%s/roles/%s", roleProjId, projId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", roleProjService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetRoleProjById StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (roleProjService *YGRoleProjectService) EditRoleProj(roleProjId string, projId string, parameters RoleProjectResponse) (err error, response IDResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/projects/%s/roles/%s", roleProjId, projId)
	payloadByte, _ := json.Marshal(parameters)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", roleProjService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditRoleProj StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (roleProjService *YGRoleProjectService) DeleteRoleProj(roleProjId string, projId string) (err error, response RoleProjectResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/projects/%s/roles/%s", roleProjId, projId)
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", roleProjService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("DeleteRoleProj StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

type RoleProjectResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Permissions struct {
		EditTitle bool `json:"editTitle"`
		Delete    bool `json:"delete"`
		AddBoard  bool `json:"addBoard"`
		Boards    struct {
			EditTitle    bool `json:"editTitle"`
			Delete       bool `json:"delete"`
			Move         bool `json:"move"`
			ShowStickers bool `json:"showStickers"`
			EditStickers bool `json:"editStickers"`
			AddColumn    bool `json:"addColumn"`
			Columns      struct {
				EditTitle bool   `json:"editTitle"`
				Delete    bool   `json:"delete"`
				Move      string `json:"move"`
				AddTask   bool   `json:"addTask"`
				AllTasks  struct {
					Show            bool   `json:"show"`
					Delete          bool   `json:"delete"`
					EditTitle       bool   `json:"editTitle"`
					EditDescription bool   `json:"editDescription"`
					Complete        bool   `json:"complete"`
					Close           bool   `json:"close"`
					AssignUsers     string `json:"assignUsers"`
					Connect         bool   `json:"connect"`
					EditSubtasks    string `json:"editSubtasks"`
					EditStickers    bool   `json:"editStickers"`
					EditPins        bool   `json:"editPins"`
					Move            string `json:"move"`
					SendMessages    bool   `json:"sendMessages"`
					SendFiles       bool   `json:"sendFiles"`
					EditWhoToNotify string `json:"editWhoToNotify"`
				} `json:"allTasks"`
				WithMeTasks struct {
					Show            bool   `json:"show"`
					Delete          bool   `json:"delete"`
					EditTitle       bool   `json:"editTitle"`
					EditDescription bool   `json:"editDescription"`
					Complete        bool   `json:"complete"`
					Close           bool   `json:"close"`
					AssignUsers     string `json:"assignUsers"`
					Connect         bool   `json:"connect"`
					EditSubtasks    string `json:"editSubtasks"`
					EditStickers    bool   `json:"editStickers"`
					EditPins        bool   `json:"editPins"`
					Move            string `json:"move"`
					SendMessages    bool   `json:"sendMessages"`
					SendFiles       bool   `json:"sendFiles"`
					EditWhoToNotify string `json:"editWhoToNotify"`
				} `json:"withMeTasks"`
				MyTasks struct {
					Show            bool   `json:"show"`
					Delete          bool   `json:"delete"`
					EditTitle       bool   `json:"editTitle"`
					EditDescription bool   `json:"editDescription"`
					Complete        bool   `json:"complete"`
					Close           bool   `json:"close"`
					AssignUsers     string `json:"assignUsers"`
					Connect         bool   `json:"connect"`
					EditSubtasks    string `json:"editSubtasks"`
					EditStickers    bool   `json:"editStickers"`
					EditPins        bool   `json:"editPins"`
					Move            string `json:"move"`
					SendMessages    bool   `json:"sendMessages"`
					SendFiles       bool   `json:"sendFiles"`
					EditWhoToNotify string `json:"editWhoToNotify"`
				} `json:"myTasks"`
				CreatedByMeTasks struct {
					Show            bool   `json:"show"`
					Delete          bool   `json:"delete"`
					EditTitle       bool   `json:"editTitle"`
					EditDescription bool   `json:"editDescription"`
					Complete        bool   `json:"complete"`
					Close           bool   `json:"close"`
					AssignUsers     string `json:"assignUsers"`
					Connect         bool   `json:"connect"`
					EditSubtasks    string `json:"editSubtasks"`
					EditStickers    bool   `json:"editStickers"`
					EditPins        bool   `json:"editPins"`
					Move            string `json:"move"`
					SendMessages    bool   `json:"sendMessages"`
					SendFiles       bool   `json:"sendFiles"`
					EditWhoToNotify string `json:"editWhoToNotify"`
				} `json:"createdByMeTasks"`
			} `json:"columns"`
			Settings bool `json:"settings"`
		} `json:"boards"`
		Children struct {
		} `json:"children"`
	} `json:"permissions"`
}

type RoleProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permissions struct {
		EditTitle bool `json:"editTitle"`
		Delete    bool `json:"delete"`
		AddBoard  bool `json:"addBoard"`
		Boards    struct {
			EditTitle    bool `json:"editTitle"`
			Delete       bool `json:"delete"`
			Move         bool `json:"move"`
			ShowStickers bool `json:"showStickers"`
			EditStickers bool `json:"editStickers"`
			AddColumn    bool `json:"addColumn"`
			Columns      struct {
				EditTitle bool   `json:"editTitle"`
				Delete    bool   `json:"delete"`
				Move      string `json:"move"`
				AddTask   bool   `json:"addTask"`
				AllTasks  struct {
					Show            bool   `json:"show"`
					Delete          bool   `json:"delete"`
					EditTitle       bool   `json:"editTitle"`
					EditDescription bool   `json:"editDescription"`
					Complete        bool   `json:"complete"`
					Close           bool   `json:"close"`
					AssignUsers     string `json:"assignUsers"`
					Connect         bool   `json:"connect"`
					EditSubtasks    string `json:"editSubtasks"`
					EditStickers    bool   `json:"editStickers"`
					EditPins        bool   `json:"editPins"`
					Move            string `json:"move"`
					SendMessages    bool   `json:"sendMessages"`
					SendFiles       bool   `json:"sendFiles"`
					EditWhoToNotify string `json:"editWhoToNotify"`
				} `json:"allTasks"`
				WithMeTasks struct {
					Show            bool   `json:"show"`
					Delete          bool   `json:"delete"`
					EditTitle       bool   `json:"editTitle"`
					EditDescription bool   `json:"editDescription"`
					Complete        bool   `json:"complete"`
					Close           bool   `json:"close"`
					AssignUsers     string `json:"assignUsers"`
					Connect         bool   `json:"connect"`
					EditSubtasks    string `json:"editSubtasks"`
					EditStickers    bool   `json:"editStickers"`
					EditPins        bool   `json:"editPins"`
					Move            string `json:"move"`
					SendMessages    bool   `json:"sendMessages"`
					SendFiles       bool   `json:"sendFiles"`
					EditWhoToNotify string `json:"editWhoToNotify"`
				} `json:"withMeTasks"`
				MyTasks struct {
					Show            bool   `json:"show"`
					Delete          bool   `json:"delete"`
					EditTitle       bool   `json:"editTitle"`
					EditDescription bool   `json:"editDescription"`
					Complete        bool   `json:"complete"`
					Close           bool   `json:"close"`
					AssignUsers     string `json:"assignUsers"`
					Connect         bool   `json:"connect"`
					EditSubtasks    string `json:"editSubtasks"`
					EditStickers    bool   `json:"editStickers"`
					EditPins        bool   `json:"editPins"`
					Move            string `json:"move"`
					SendMessages    bool   `json:"sendMessages"`
					SendFiles       bool   `json:"sendFiles"`
					EditWhoToNotify string `json:"editWhoToNotify"`
				} `json:"myTasks"`
				CreatedByMeTasks struct {
					Show            bool   `json:"show"`
					Delete          bool   `json:"delete"`
					EditTitle       bool   `json:"editTitle"`
					EditDescription bool   `json:"editDescription"`
					Complete        bool   `json:"complete"`
					Close           bool   `json:"close"`
					AssignUsers     string `json:"assignUsers"`
					Connect         bool   `json:"connect"`
					EditSubtasks    string `json:"editSubtasks"`
					EditStickers    bool   `json:"editStickers"`
					EditPins        bool   `json:"editPins"`
					Move            string `json:"move"`
					SendMessages    bool   `json:"sendMessages"`
					SendFiles       bool   `json:"sendFiles"`
					EditWhoToNotify string `json:"editWhoToNotify"`
				} `json:"createdByMeTasks"`
			} `json:"columns"`
			Settings bool `json:"settings"`
		} `json:"boards"`
		Children struct {
		} `json:"children"`
	} `json:"permissions"`
}
