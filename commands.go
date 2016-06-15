package explorer

import (
	"net/url"
    "io/ioutil"
    "encoding/json"
    "github.com/aktungmak/ccm-client"
)

func (a *App) Back() string {
	prev := a.History[len(a.History)-1]
	a.History = a.History[:len(a.History)-1]
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
func (a *App) Goto(u *url.URL) string {
    // update state, even if url is not (yet) valid
	a.History = append(a.History, u)
    a.Current = u

    // make the request
    res, err := a.Client.DoRaw("GET", u.String())
    if err != nil {
        return err.Error()
    }
    a.Response = res

    // try to extract the response body
    defer res.Body.Close()
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return err.Error()
    }

    // parse json response
    var pdat map[string]interface{}
	err = json.Unmarshal(data, &pdat)
    if err != nil {
        return "couldn't parse JSON response: " + err.Error()
    }

    // extract links - todo should perhaps specify the "Links" key?
	a.Links = ccm.ParseLinks(pdat, "")

    return res.Status
}
