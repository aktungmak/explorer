package explorer

import (
	"github.com/chzyer/readline"
)

var __completer = readline.NewPrefixCompleter(
	readline.PcItem("back"),
	readline.PcItem("body"),
	readline.PcItem("dump"),
	readline.PcItem("exit"),
	readline.PcItem("goto"),
	readline.PcItem("help"),
	readline.PcItem("home"),
	readline.PcItem("jump"),
	readline.PcItem("man",
		readline.PcItem("DELETE"),
		readline.PcItem("GET"),
		readline.PcItem("PATCH"),
		readline.PcItem("POST"),
		readline.PcItem("PUT"),
	),
	readline.PcItem("mark"),
	readline.PcItem("quit"),
	readline.PcItem("save"),
	readline.PcItem("set",
		readline.PcItem("body", readline.PcItem("on"), readline.PcItem("off")),
		readline.PcItem("opts", readline.PcItem("on"), readline.PcItem("off")),
	),
	readline.PcItem("where"),
)
