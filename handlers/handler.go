package handlers

import (
	"cli-assistant/cmdlist"
	"cli-assistant/handlers/add"
	"cli-assistant/handlers/application"
	"cli-assistant/handlers/browser"
	"cli-assistant/handlers/help"
	"cli-assistant/handlers/remove"

	"github.com/op/go-logging"
)

// Handle inputs
func Handle(log *logging.Logger, input string, cmds cmdlist.RepoManager) {
	ask := cmds.Get(input)
	switch ask.Type {
	case "help":
		help.Display(log, cmds)
	case "add":
		add.Add(log, cmds)
	case "del":
		remove.Del(log, cmds)
	case "browser":
		browser.Open(log, ask.Alias)
	case "application":
		application.Open(log, ask, input)
	default:
		log.Warningf("Unknown command: %s \n", input)
	}

}
