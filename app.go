package explorer

import (
	"bufio"
	"fmt"
	"github.com/aktungmak/ccm-client"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type App struct {
	Root     *url.URL
	Current  *url.URL
	History  []*url.URL
	Marks    map[string]*url.URL
	Client   *ccm.Client
	Links    map[string]*url.URL
	Response *http.Response
	Reader   *bufio.Reader // replace this with a readline implementation
}

func NewApp(servroot, user, pass string) (*App, error) {
	s, err := url.Parse(servroot)
	if err != nil {
		return &App{}, err
	}
	c, err := ccm.NewClient(s.Host, user, pass)
	if err != nil {
		return &App{}, err
	}
	a := &App{
		Root:    s,
		Current: s,
		History: make([]*url.URL, 10),
		Marks:   make(map[string]*url.URL),
		Client:  c,
		Reader:  bufio.NewReader(os.Stdin),
	}

	return a, err
}

// format a string of the options and their indexes
func (a *App) LinksString() string {
	var ret string
	for name, url := range a.Links {
		ret += fmt.Sprintf("%s: %s\n", name, url)
	}
	return ret
}

func (a *App) EventLoop() {
    a.Goto(a.Root)
	for {
		fmt.Println(a.LinksString())
		text := a.getLine()
		if text == "quit" {
			break
		}
		res := a.ParseCommand(text)
		fmt.Println("-> " + res)
	}
}

// temporary function until I use readline
func (a *App) getLine() string {
	fmt.Print(": ")
	text, err := a.Reader.ReadString('\n')
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return ""
	}
	return strings.TrimSpace(text)

}