package explorer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aktungmak/odata-client"
	"github.com/chzyer/readline"
	"io/ioutil"
	"net/url"
	"os"
)

// client types
const (
	Basic = iota
	Manual
	Token
)

type Config struct {
	Root       string
	Current    string
	History    []string
	Marks      map[string]string
	AutoOpts   bool
	AutoBody   bool
	User       string
	Pass       string
	Token      string
	ClientType int
	Insecure   bool
}

func LoadConfig(filename string) (*App, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	c := &Config{}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(c)
	if err != nil {
		return nil, err
	}

	servroot, err := url.Parse(c.Root)
	if err != nil {
		return nil, err
	}

	var cl odata.Client
	// create correct client based on ClientType
	switch c.ClientType {
	case Basic:
		cl = odata.NewBaClient(c.User, c.Pass, c.Insecure)
	case Manual:
		cl = odata.NewManualClient(c.Token, c.Insecure)
	case Token:
		cl, err = odata.NewTokenClient(servroot.Hostname(), c.User, c.Pass, c.Insecure)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid ClientType in config")
	}

	a, err := NewApp(servroot, cl)
	if err != nil {
		return nil, err
	}

	a.Current, err = url.Parse(c.Current)
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

	// save client based on type
	switch v := a.Client.(type) {
	case *odata.BaClient:
		c.ClientType = Basic
		c.User = v.Username
		c.Pass = v.Password
	case *odata.ManualClient:
		c.ClientType = Manual
		c.Token = v.Token
	case *odata.TokenClient:
		c.ClientType = Token
		c.User = v.Username
		c.Pass = v.Password
	default:
		fmt.Printf("client type: %T", v)
		fmt.Printf("client type: %T", a.Client)
		return errors.New("app has an unimplemented client type")
	}

	// TODO work this out
	// c.Insecure = a.Insecure
	c.Insecure = true

	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, []byte(b), 0644)
	return err
}
