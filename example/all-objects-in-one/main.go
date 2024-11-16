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
	fmt.Println(company)
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
	fmt.Println(key)
	usersService := yougilego.YGUsersService{Key: key}
	err, userList := usersService.GetUsers()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(userList)
	projectService := yougilego.YGProjectService{Key: key}
	err, projList := projectService.GetProjeсtList()
	if err != nil {
		log.Println(err)
		return
	}
	roleProjectService := yougilego.YGRoleProjectService{Key: key}
	err, roleProjList := roleProjectService.GetRoleProjectList(projList.Content[0].Id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(roleProjList)

	departmentsService := yougilego.YGDepartmentsService{Key: key}
	err, departList := departmentsService.GetDepartList()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(departList)
	boardsService := yougilego.YGBoardsService{Key: key}
	err, boardList := boardsService.GetBoards()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(boardList)
	columnService := yougilego.YGColumnService{Key: key}
	err, columnList := columnService.GetColumnList()
	if err != nil {
		log.Println(err)
		return
	}
	myColumn := columnList.Content[0]
	taskService := yougilego.YGTaskService{Key: key}
	err, taskList := taskService.GetTasks(myColumn.Id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(taskList)
	textStickerService := yougilego.YGTextStickerService{Key: key}
	err, stickerList := textStickerService.GetStickersList()
	if err != nil {
		log.Println(err)
		return
	}
	mySticker := stickerList.Content[0]
	statusTextStickerService := yougilego.YGStatusTextStickerService{Key: key}
	err, statusStickerList := statusTextStickerService.GetStatusTextStickerById(mySticker.Id, mySticker.States[0].Id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(statusStickerList)
	sprintStickerService := yougilego.YGSprintStickerService{Key: key}
	err, sprintStickers := sprintStickerService.GetSprintStickerList()
	if err != nil {
		log.Println(err)
		return
	}
	mySprintStickers := sprintStickers.Content[0]
	statusStickerSprintService := yougilego.YGStatusStickerSprintService{Key: key}
	err, statusStickerSprint := statusStickerSprintService.GetStatusStickerSprint(mySprintStickers.Id, mySprintStickers.States[0].Id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(statusStickerSprint)
	groupChatService := yougilego.YGGroupChatService{Key: key}
	err, groupChatList := groupChatService.GetGroupChatList()
	if err != nil {
		log.Println(err)
		return
	}
	myChat := groupChatList.Content[0]
	messageService := yougilego.YGChatMessageService{Key: key}
	err, history := messageService.GetHistoryChat(myChat.Id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(history)
}
