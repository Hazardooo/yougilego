package main

import (
	"fmt"
	"github.com/Hazardooo/yougilego"
)

func main() {
	auth := yougilego.YGAuthService{
		Login:    "Hazdoo@outlook.com",
		Password: "Pav.03.07.2004!",
	}

	err, company := auth.GetListCompany("Hazardooo")
	if err != nil {
		return
	}
	fmt.Println(company)

	err, keys := auth.GetKeysList(company.Content[0].Id)
	if err != nil {
		return
	}
	fmt.Println(keys)

	err, key := auth.CreateKey(company.Content[0].Id)
	if err != nil {
		return
	}
	fmt.Println(key)

	auth.DeleteKey(keys[0].Key)
}
