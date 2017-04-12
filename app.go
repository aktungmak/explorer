package explorer

import (
	"fmt"
	"github.com/aktungmak/odata-client"
	"github.com/chzyer/readline"
	"net/url"
	"strings"
)

type App struct {
	Root       *url.URL
	Current    *url.URL
	History    []*url.URL
	Marks      map[string]*url.URL
	Client     odata.Client
	Links      UrlSlice
	LastBody   []byte
	LastStatus string
	Reader     *readline.Instance `json:"-"`
	AutoOpts   bool
	AutoBody   bool
}

func NewApp(servroot *url.URL, client odata.Client) (*App, error) {
	rdr, err := readline.NewEx(&readline.Config{
		Prompt:       ": ",
		AutoComplete: __completer,
	})
	if err != nil {
		return nil, err
	}

	a := &App{
		Root:     servroot,
		Current:  servroot,
		History:  make([]*url.URL, 0, 10),
		Marks:    make(map[string]*url.URL),
		Client:   client,
		Reader:   rdr,
		AutoOpts: true,
		AutoBody: false,
	}

	return a, err
}

// implements sort.Interface so the links are predictably ordered
type UrlSlice []*url.URL

func (a UrlSlice) Len() int           { return len(a) }
func (a UrlSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a UrlSlice) Less(i, j int) bool { return a[i].Path < a[j].Path }

// format a string of the options and their indexes
func (a *App) LinksString() string {
	if len(a.Links) == 0 {
		return "no links in the response\n"
	}
	var ret string
	for i, url := range a.Links {
		ret += fmt.Sprintf("%d: %s\n", i, url)
	}
	return ret
}

func (a *App) EventLoop() {
	defer a.Reader.Close()
	fmt.Println("-> " + a.Goto("GET", a.Current, nil))
	for {
		if a.AutoOpts {
			fmt.Println(a.LinksString())
		}

		text, err := a.Reader.Readline()
		if err != nil {
			fmt.Printf("Error reading command: %s\n", err)
			break
		}
		text = strings.TrimSpace(text)
		if text == "quit" || text == "exit" {
			break
		}
		res := a.ParseCommand(text)
		fmt.Println("-> " + res)
		if a.AutoBody {
			fmt.Println(a.ShowBody())
		}
	}
}
