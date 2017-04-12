package explorer

import (
	"encoding/json"
	"github.com/aktungmak/odata-client"
	"github.com/chzyer/readline"
	"io/ioutil"
	"net/url"
	"sort"
)

func (a *App) Back() string {
	prev := a.History[len(a.History)-2]
	a.History = a.History[:len(a.History)-2]
	return a.Goto("GET", prev, nil)
}
func (a *App) Dump(filename string) string {
	err := ioutil.WriteFile(filename, []byte(a.LastBody), 0644)
	if err != nil {
		return "Error dumping body: " + err.Error()
	}
	return "successfully dumped body in " + filename
}
func (a *App) Help() string {
	return __HELP_TEXT__
}
func (a *App) Home() string {
	return a.Goto("GET", a.Root, nil)
}
func (a *App) ShowLinks() string {
	return a.LinksString()
}
func (a *App) ShowBody() string {
	return string(a.LastBody)
}
func (a *App) Mark(name string) string {
	a.Marks[name] = a.Current
	__m_completer.SetChildren(append(__m_completer.GetChildren(),
		readline.PcItem(name)))
	return "Marked current location as " + name
}
func (a *App) Jump(name string) string {
	u, ok := a.Marks[name]
	if !ok {
		return "no mark called " + name
	}
	return a.Goto("GET", u, nil)
}

func (a *App) Set(setting, value string) string {
	switch setting {
	case "body":
		if value == "on" {
			a.AutoBody = true
			return "enabled auto printing of request body"
		} else {
			a.AutoBody = false
			return "disabled auto printing of request body"
		}
	case "opts":
		if value == "on" {
			a.AutoOpts = true
			return "enabled auto printing of link options"
		} else {
			a.AutoOpts = false
			return "disabled auto printing of link options"
		}
	}
	return "unknown option: " + value
}

// change to the (unqualified) url provided, even if it does
// not exist yet. this allows us to make a POST from the new
// location if needed.
func (a *App) Goto(method string, u *url.URL, body []byte) string {
	if u == nil {
		return "invalid URL"
	}
	// update app state and clear old links and body
	a.History = append(a.History, u)
	a.Current = u
	a.Links = make(UrlSlice, 0)
	a.LastBody = []byte{}

	// qualify the url with our service root
	fullUrl := a.Root.ResolveReference(u)

	// make the request
	res, err := a.Client.DoRaw(method, fullUrl.String(), body)
	if err != nil {
		a.LastStatus = "0 ERROR"
		return err.Error()
	}
	a.LastStatus = res.Status

	// try to extract the response body
	defer res.Body.Close()
	a.LastBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return res.Status + ": " + err.Error()
	}

	// parse json response
	var pBody map[string]interface{}
	err = json.Unmarshal(a.LastBody, &pBody)
	if err != nil {
		//return res.Status + ": " + err.Error()
		return res.Status
	}

	// extract links
	linkMap := odata.ParseLinks(pBody, "")
	a.Links = make(UrlSlice, len(linkMap))
	i := 0
	for _, v := range linkMap {
		a.Links[i] = v
		i++
	}
	sort.Sort(a.Links)
	return res.Status
}
