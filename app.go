package explorer

import (
	"bufio"
	"fmt"
	"github.com/aktungmak/ccm-client"
	"net/url"
	"os"
	"strings"
)

type App struct {
	Root       *url.URL
	Current    *url.URL
	History    []*url.URL
	Marks      map[string]*url.URL
	Client     *ccm.Client
	Links      []*url.URL
	LastBody   []byte
	LastStatus string
	Reader     *bufio.Reader // replace this with a readline implementation
}

func NewApp(servroot, user, pass string) (*App, error) {
	sr, err := url.Parse(servroot)
	if err != nil {
		return &App{}, err
	}

	c, err := ccm.NewClient(sr.Host, user, pass)
	if err != nil {
		return &App{}, err
	}

	a := &App{
		Root:    sr,
		Current: sr,
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
	for i, url := range a.Links {
		ret += fmt.Sprintf("%d: %s\n", i, url)
	}
	return ret
}

func (a *App) EventLoop() {
	a.Goto(a.Root)
	for {
		fmt.Println(a.LinksString())
		text := a.getLine()
		if text == "quit" || text == "exit" {
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
