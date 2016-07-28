package explorer

import (
	"github.com/chzyer/readline"
)

// this will be populated with the mark names
var __m_completer = readline.PcItem("jump")

var __completer = readline.NewPrefixCompleter(
	__m_completer,
	readline.PcItem("back"),
	readline.PcItem("body"),
	readline.PcItem("dump"),
	readline.PcItem("exit"),
	readline.PcItem("freq",
		readline.PcItem("DELETE"),
		readline.PcItem("GET"),
		readline.PcItem("PATCH"),
		readline.PcItem("POST"),
		readline.PcItem("PUT"),
	),
	readline.PcItem("goto"),
	readline.PcItem("help"),
	readline.PcItem("home"),
	readline.PcItem("last"),
	readline.PcItem("mark"),
	readline.PcItem("opts"),
	readline.PcItem("quit"),
	readline.PcItem("req",
		readline.PcItem("DELETE"),
		readline.PcItem("GET"),
		readline.PcItem("PATCH"),
		readline.PcItem("POST"),
		readline.PcItem("PUT"),
	),
	readline.PcItem("save"),
	readline.PcItem("set",
		readline.PcItem("body", readline.PcItem("on"), readline.PcItem("off")),
		readline.PcItem("opts", readline.PcItem("on"), readline.PcItem("off")),
	),
	readline.PcItem("up"),
	readline.PcItem("where"),
)
