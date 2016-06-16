package explorer

import (
	"strconv"
	"strings"
)

// parse the command line and do what is required
// return some kind of message (error or success)
func (a *App) ParseCommand(line string) string {
	hp := strings.HasPrefix
	cmds := strings.Split(line, " ")
	switch {
	case hp(line, "back"):
		return a.Back()
	case hp(line, "home"):
		return a.Home()
	case hp(line, "help"):
		return a.Help()
	case hp(line, "link"):
		return a.ShowLinks()
	case hp(line, "body"):
		return a.ShowBody()
	case hp(line, "mark"):
		if len(cmds) < 2 {
			return "not enough args"
		}
		return a.Mark(cmds[1])
	case hp(line, "jump"):
		if len(cmds) < 2 {
			return "not enough args"
		}
		return a.Jump(cmds[1])
	default:
		i, err := strconv.ParseInt(cmds[0], 10, 0)
		if err != nil {
			break
		}
		if 0 > int(i) || int(i) >= len(a.Links) {
			return "no link with that index"
		}
		return a.Goto(a.Links[i])
	}
	return "didn't understand the command... try help"
}
