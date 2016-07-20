package main

import (
	"fmt"
	"github.com/aktungmak/explorer"
	"os"
)

func main() {
	var app *explorer.App
	var err error
	if len(os.Args) > 3 {
		app, err = explorer.NewApp(os.Args[1], os.Args[2], os.Args[3], true)
	} else if len(os.Args) > 1 {
		app, err = explorer.LoadConfig(os.Args[1])
	} else {
		fmt.Println("usage: explore.exe (<servRoot> <uname> <passwd> | <configFile>)")
		os.Exit(2)
	}

	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		app.EventLoop()
	}
}
