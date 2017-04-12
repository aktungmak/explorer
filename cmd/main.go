package main

import (
	"flag"
	"fmt"
	"github.com/aktungmak/explorer"
	"github.com/aktungmak/odata-client"
	"io/ioutil"
	"os"
)

func main() {
	var app *explorer.App
	var cl odata.Client
	var err error

	// command line flag definitions
	var ba = flag.Bool("ba", true, "use basic auth instead of token. true by default!")
	var cf = flag.String("cf", "", "config file created using the 'save' command")
	var is = flag.Bool("k", true, "like curl -k, ignore bad certs. true by default!")
	var ps = flag.String("p", "", "password for basic auth")
	var sr = flag.String("sr", "none", "the URL to the API service root")
	var tf = flag.String("tf", "", "filename containing a JWT to be used")
	var tk = flag.String("tk", "", "JWT to be used in token auth")
	var us = flag.String("u", "", "username for basic auth")
	flag.Parse()

	// check the service root provided is valid
	servroot, err := url.Parse(sr)
	if err != nil {
		fmt.Printf("invalid or missing service root URL: %s", err)
		fmt.Println("make sure you included the -sr option.")
		os.Exit(2)
	}

	if cf != "" {
		// we have a config file, so use that
		app, err = explorer.LoadConfig(cf)
	} else {
		// try and guess what mode we are in based on the provided args
		if us != "" && ps != "" && ba {
			// basic auth
			cl = odata.NewBaClient(us, ps, is)
		} else if us != "" && ps != "" && !ba {
			// auto auth token retrieval
			cl = odata.NewTokenClient(us, ps, is)
		} else if tk != "" {
			// use provided token string
			cl = odata.NewManualClient(tk, is)
		} else if tf != "" {
			// load token from file
			data, err := ioutil.ReadFile(tf)
			if err != nil {
				fmt.Printf("can't load token from file: %s", err)
				os.Exit(2)
			}
			cl = odata.NewManualClient(string(data), is)
		} else {
			// invalid combination of arguments
			return "invalid args"
		}
		app, err := explorer.NewApp(servroot, cl)
		if err != nil {
			fmt.Printf("can't start application: %s", err)
			os.Exit(2)
		}
	}
	app.EventLoop()
}
