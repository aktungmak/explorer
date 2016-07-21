package explorer

import (
	"encoding/json"
	"github.com/chzyer/readline"
	"io/ioutil"
	"net/url"
	"os"
)

type Config struct {
	Root     string
	Current  string
	History  []string
	Marks    map[string]string
	AutoOpts bool
	AutoBody bool
	User     string
	Pass     string
	Insecure bool
}

func LoadConfig(filename string) (*App, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(c)
	if err != nil {
		return nil, err
	}

	a, err := NewApp(c.Root, c.User, c.Pass, c.Insecure)
	if err != nil {
		return a, err
	}

	a.Root, err = url.Parse(c.Root)
	if err != nil {
		return a, err
	}
	a.Current, _ = url.Parse(c.Current)
	if err != nil {
		a.Current = a.Root
	}

	a.History = make([]*url.URL, 0, len(c.History))
	for _, s := range c.History {
		u, err := url.Parse(s)
		if err != nil {
			continue
		}
		a.History = append(a.History, u)
	}

	a.Marks = make(map[string]*url.URL)
	for k, s := range c.Marks {
		u, err := url.Parse(s)
		if err != nil {
			continue
		}
		a.Marks[k] = u
		__m_completer.SetChildren(append(__m_completer.GetChildren(),
			readline.PcItem(k)))
	}

	a.AutoOpts = c.AutoOpts
	a.AutoBody = c.AutoBody

	return a, nil
}

func (a *App) SaveConfig(filename string) error {
	c := &Config{}
	c.Root = a.Root.String()
	c.Current = a.Current.String()

	c.History = make([]string, len(a.History))
	for i, u := range a.History {
		c.History[i] = u.String()
	}

	c.Marks = make(map[string]string)
	for k, u := range a.Marks {
		c.Marks[k] = u.String()
	}

	c.AutoOpts = a.AutoOpts
	c.AutoBody = a.AutoBody

	c.User = a.Client.Username
	c.Pass = a.Client.Password
	c.Insecure = a.Insecure

	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, []byte(b), 0644)
	return err
}
