package explorer

import (
	"net/url"
	"strconv"
	"strings"
)

// parse the command line and do what is required
// return a message (error or success) for the user
func (a *App) ParseCommand(line string) string {
	hp := strings.HasPrefix
	cmds := strings.Split(line, " ")
	switch {
	case hp(line, "back"):
		return a.Back()
	case hp(line, "body"):
		return a.ShowBody()
	case hp(line, "dump"):
		if len(cmds) < 2 {
			return "need to specify the filename"
		}
		return a.Dump(cmds[1])
	case hp(line, "goto"):
		if len(cmds) < 2 {
			return "need to specify the URI to goto"
		}
		u, err := url.Parse(cmds[1])
		if err != nil {
			return "that's not a valid uri: " + err.Error()
		}
		return a.Goto("GET", u, nil)
	case hp(line, "help"):
		return a.Help()
	case hp(line, "home"):
		return a.Home()
	case hp(line, "jump"):
		if len(cmds) < 2 {
			return "please specify the name to jump to"
		}
		return a.Jump(cmds[1])
	case hp(line, "opts"):
		return a.ShowLinks()
	case hp(line, "mark"):
		if len(cmds) < 2 {
			return "you must specify a name for the mark"
		}
		return a.Mark(cmds[1])
	case hp(line, "man"):
		if len(cmds) < 3 {
			return "must specify at least method and URI"
		}
		u, err := url.Parse(cmds[2])
		if err != nil {
			return "that's not a valid uri: " + err.Error()
		}
		body := []byte{}
		if len(cmds) > 3 {
			body = []byte(cmds[3])
		}
		return a.Goto(cmds[1], u, body)
	case hp(line, "save"):
		if len(cmds) < 2 {
			return "need filename to save config in"
		}
		err := a.SaveConfig(cmds[1])
		if err != nil {
			return "error saving config: " + err.Error()
		}
		return "saved config in " + cmds[1]
	case hp(line, "set"):
		if len(cmds) < 3 {
			return "must specify setting and on or off"
		}
		return a.Set(cmds[1], cmds[2])
	case hp(line, "where"):
		return a.Root.ResolveReference(a.Current).String()
	default:
		// by default assume it is a link index
		i, err := strconv.ParseInt(cmds[0], 10, 0)
		if err != nil {
			break
		}
		if int(i) < 0 || int(i) >= len(a.Links) {
			return "no link with index " + cmds[0]
		}
		return a.Goto("GET", a.Links[i], nil)
	}
	return "didn't understand the command... try help"
}
