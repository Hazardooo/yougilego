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
