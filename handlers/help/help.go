package help

import (
	"cli-assistant/cmdlist"
	"fmt"

	"github.com/op/go-logging"
)

// Display help
func Display(log *logging.Logger, cmds cmdlist.RepoManager) {
	commands := cmds.GetAll()
	fmt.Printf("Application commands:\n")
	for key, val := range commands {
		if val.Alias == "" {
			fmt.Printf("	%s				%s\n", key, val.Description)
		}

	}

	fmt.Printf("\nRegistered commands:\n")
	for key, val := range commands {
		if key != "add" && key != "help" {
			fmt.Printf("	%s				%s\n", key, val.Description)
		}

	}
	fmt.Printf("\n\n\n")
}
