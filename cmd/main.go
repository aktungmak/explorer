package main

import (
	"fmt"
	"github.com/aktungmak/explorer"
)

func main() {
	//app, err := explorer.NewApp("https://147.214.115.3:10443/rest/v0/",
	//    "sysadmin",
	//    "sysadmin123")
	app, err := explorer.NewApp("https://127.0.0.1:9001/rest/v0/",
		"localhost\\sysadmin",
		"Sysadmin123@")
	if err != nil {
		fmt.Printf("err: %s\n", err)
	} else {
		app.EventLoop()
	}
}
