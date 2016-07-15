package main

import (
	"fmt"
	"github.com/aktungmak/explorer"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("usage: explore.exe <servRoot> <uname> <passwd>")
		os.Exit(2)
	}
	app, err := explorer.NewApp(os.Args[1], os.Args[2], os.Args[3])

	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		app.EventLoop()
	}
}
