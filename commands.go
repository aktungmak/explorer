package explorer

import (
	"encoding/json"
	"github.com/aktungmak/ccm-client"
	"io/ioutil"
	"net/url"
)

func (a *App) Back() string {
	prev := a.History[len(a.History)-2]
	a.History = a.History[:len(a.History)-2]
	return a.Goto(prev)
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
	a.History = append(a.History, a.Root)
	return a.Goto(a.Root)
}
func (a *App) ShowLinks() string {
	return a.LinksString()
}
func (a *App) ShowBody() string {
	return string(a.LastBody)
}
func (a *App) Mark(name string) string {
	a.Marks[name] = a.Current
	return "Marked current location as " + name
}
func (a *App) Jump(name string) string {
	u, ok := a.Marks[name]
	if !ok {
		return "no mark called " + name
	}
	return a.Goto(u)
}

// change to the (unqualified) url provided, even if it does
// not exist yet. this allows us to make a POST from the new
// location if needed.
func (a *App) Goto(u *url.URL) string {
	// update app state
	a.History = append(a.History, u)
	a.Current = u

	// qualify the url with our service root
	fullUrl := a.Root.ResolveReference(u)

	// make the request
	res, err := a.Client.DoRaw("GET", fullUrl.String())
	if err != nil {
		a.LastStatus = "0 ERROR"
		return err.Error()
	}
	a.LastStatus = res.Status

	// try to extract the response body
	defer res.Body.Close()
	a.LastBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		a.LastBody = []byte{}
		return err.Error()
	}

	// parse json response
	var pBody map[string]interface{}
	err = json.Unmarshal(a.LastBody, &pBody)
	if err != nil {
		return "couldn't parse JSON response: " + err.Error()
	}

	// extract links - todo should perhaps specify the "Links" key?
	linkMap := ccm.ParseLinks(pBody, "")
	a.Links = make([]*url.URL, len(linkMap))
	i := 0
	for _, v := range linkMap {
		a.Links[i] = v
		i++
	}
	return res.Status
}

func (a *App) Manual(method string, uri *url.URL, body []byte) string {

	return "to be implemented"
}
