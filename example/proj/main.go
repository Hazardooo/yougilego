package main

import (
	"fmt"
	"github.com/Hazardooo/yougilego"
	"log"
	"os"
)

func main() {
	auth := yougilego.YGAuthService{
		Login:    os.Getenv("LOGIN"),
		Password: os.Getenv("PASSWORD"),
	}
	err, company := auth.GetListCompany("Hazardooo")
	if err != nil {
		log.Println(err)
		return
	}
	myCompany := company.Content[0]
	err, keysList := auth.GetKeysList(myCompany.Id)
	if err != nil {
		log.Println(err)
		return
	}
	var key string
	if len(keysList) == 0 {
		err, key = auth.CreateKey(myCompany.Id)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		key = keysList[0].Key
	}
	projectService := yougilego.YGProjectService{
		Key: key,
	}
	err, projectsList := projectService.GetProjeсts()
	if err != nil {
		log.Println(err)
		return
	}
	myProj := projectsList.Content[0]
	roleProjService := yougilego.YGRoleProjectService{Key: key}
	err, _ = roleProjService.CreateRoleProject(myProj.Id, yougilego.RoleProjectRequest{
		Name:        "Consultant",
		Description: "Sell me this pen",
		Permissions: struct {
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
			Children struct{} `json:"children"`
		}{
			EditTitle: true,
			Delete:    true,
			AddBoard:  true,
			Boards: struct {
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
			}{
				EditTitle:    true,
				Delete:       true,
				Move:         true,
				ShowStickers: true,
				EditStickers: true,
				AddColumn:    true,
				Columns: struct {
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
				}{
					EditTitle: true,
					Delete:    true,
					Move:      "no",
					AddTask:   true,
					AllTasks: struct {
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
					}{
						Show:            true,
						Delete:          true,
						EditTitle:       true,
						EditDescription: true,
						Complete:        true,
						Close:           true,
						AssignUsers:     "no",
						Connect:         true,
						EditSubtasks:    "no",
						EditStickers:    true,
						EditPins:        true,
						Move:            "no",
						SendMessages:    true,
						SendFiles:       true,
						EditWhoToNotify: "no",
					},
					WithMeTasks: struct {
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
					}{
						Show:            true,
						Delete:          true,
						EditTitle:       true,
						EditDescription: true,
						Complete:        true,
						Close:           true,
						AssignUsers:     "no",
						Connect:         true,
						EditSubtasks:    "no",
						EditStickers:    true,
						EditPins:        true,
						Move:            "no",
						SendMessages:    true,
						SendFiles:       true,
						EditWhoToNotify: "no",
					},
					MyTasks: struct {
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
					}{
						Show:            true,
						Delete:          true,
						EditTitle:       true,
						EditDescription: true,
						Complete:        true,
						Close:           true,
						AssignUsers:     "no",
						Connect:         true,
						EditSubtasks:    "no",
						EditStickers:    true,
						EditPins:        true,
						Move:            "no",
						SendMessages:    true,
						SendFiles:       true,
						EditWhoToNotify: "no",
					},
					CreatedByMeTasks: struct {
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
					}{
						Show:            true,
						Delete:          true,
						EditTitle:       true,
						EditDescription: true,
						Complete:        true,
						Close:           true,
						AssignUsers:     "no",
						Connect:         true,
						EditSubtasks:    "no",
						EditStickers:    true,
						EditPins:        true,
						Move:            "no",
						SendMessages:    true,
						SendFiles:       true,
						EditWhoToNotify: "no",
					},
				},
				Settings: true,
			},
			Children: struct{}{},
		},
	})
	if err != nil {
		log.Println(err)
		return
	}

	err, roleProjList := roleProjService.GetRoleProjectList(myProj.Id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(roleProjList)
}
