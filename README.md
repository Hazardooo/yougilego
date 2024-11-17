[![en](https://img.shields.io/badge/lang-rus-red.svg)](https://github.com/Hazardooo/yougilego/blob/main/README.ru.md)
# YougileGo

YougileGo is a Go package that provides a comprehensive set of methods and structures for connecting to and retrieving data from the Yougile platform.

## Features

- Authentication with Yougile.
- Retrieve lists of companies, users, projects, departments, boards, columns, tasks, stickers, and more.
- Manage and interact with group chats and messages.
- Work with sprint stickers and text stickers.
- Subscribe to webhook events.

## Installation

Install YougileGo using `go get`:

```sh
go get github.com/Hazardooo/yougilego
```
## Usage
To make it clearer to use the library, check out the [Yougile API](https://ru.yougile.com/api-v2#/) documentation.

### Getting the company key

[An example of using all objects](https://github.com/Hazardooo/yougilego/blob/main/example/all-objects-in-one/main.go)

```go
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
}
```

In the future, the company key will be needed to work with the rest of the YouGile data

| Object                | Status |
|-----------------------|--------|
| users                 | Support|
| projects              | Support|
| project-roles         | Support|
| departments           | Support|
| boards                | Support|
| columns               | Support|
| tasks                 | Support|
| text-stickers         | Support|
| status-text-stickers  | Support|
| sprint-sticker        | Support|
| status-sprint-sticker | Support|
| group-chats           | Support|
| chat-messages         | Support|
| event-subs            | Support|

### Getting the tasks from column
```go
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
```

### Example of creating an event subscription

[Example of creating an event subscription](https://github.com/Hazardooo/yougilego/blob/main/example/sub-event/main.go)

To create a subscription to an event, you need to get your webhook address via [webhook.site](https://webhook.site/).

```go
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
	eventSubscribeService := yougilego.YGEventSubscribeService{Key: key}
	err, eventList := eventSubscribeService.GetSubscribeList(false)
	if err != nil {
		log.Println(err)
		return
	}
	var event = yougilego.SubscribeResponse{}
	if len(eventList) != 0 {
		event = eventList[0]
	} else {
		err, _ = eventSubscribeService.CreateSubscribe(yougilego.CreateSubs{
			Url:   "https://webhook.site/d19ddaf8-36d5-4dcd-acf1-51b91b6dd375",
			Event: "task-created",
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
	fmt.Println(event)
}
```
After creating a subscription, go to the Webhook site.site will instantly display all the events that took place in your YouGile

![изображение](https://github.com/user-attachments/assets/2b8fd496-9a38-4a54-9a4e-4e86100293ca)
![изображение](https://github.com/user-attachments/assets/78cdff49-326b-408d-8172-307142d639eb)


Variations of event subscriptions
| Object          | Events                                                     |
|------------------|------------------------------------------------------------|
| project         | created, deleted, restored, moved, renamed, updated         |
| board           | created, deleted, restored, moved, renamed, updated         |
| column          | created, deleted, restored, moved, renamed, updated         |
| task            | created, deleted, restored, moved, renamed, updated         |
| sticker         | created, deleted, restored, moved, renamed, updated         |
| department      | created, deleted, restored, moved, renamed, updated         |
| group_chat      | created, deleted, restored, moved, renamed, updated         |
| chat_message    | created, deleted, restored, moved, renamed, updated         |
| user            | added, removed                                              |

To subscribe to all events of a certain object, you can use the <object>-* format. For example:

- task-* — all events for tasks.
- .* — all events for all objects.

## Contributing
Contributions are welcome! Feel free to submit a pull request or open an issue to report bugs or suggest new features.

## License
This project is licensed under the [MIT License](https://github.com/Hazardooo/yougilego/blob/main/LICENSE).