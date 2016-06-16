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
func (a *App) Home() string {
	a.History = append(a.History, a.Root)
	return a.Goto(a.Root)
}
func (a *App) Help() string {
	return "one day this will be useful help"
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
	ru := a.Root.ResolveReference(u)

	// make the request
	res, err := a.Client.DoRaw("GET", ru.String())
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
	var pdat map[string]interface{}
	err = json.Unmarshal(a.LastBody, &pdat)
	if err != nil {
		return "couldn't parse JSON response: " + err.Error()
	}

	// extract links - todo should perhaps specify the "Links" key?
	lmap := ccm.ParseLinks(pdat, "")
	a.Links = make([]*url.URL, len(lmap))
	i := 0
	for _, v := range lmap {
		a.Links[i] = v
		i++
	}
	return res.Status
}
