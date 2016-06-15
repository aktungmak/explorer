package main

import (
	"fmt"
	"github.com/aktungmak/explorer"
)

func main() {
	app, err := explorer.NewApp("https://147.214.115.3:10443/rest/v0/",
		"sysadmin",
		"sysadmin123")
	fmt.Printf("err: %s\n", err)
	app.EventLoop()
}
