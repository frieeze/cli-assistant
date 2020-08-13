package remove

import (
	"bufio"
	"cli-assistant/cmdlist"
	"fmt"
	"os"
	"strings"

	"github.com/op/go-logging"
)

// Del command
func Del(log *logging.Logger, cmds cmdlist.RepoManager) {
	fmt.Printf("Remove command: \n")

	cmdName := getInput("	Command Name: ")
	command := cmds.Get(cmdName)
	if command.Alias == "" {
		log.Errorf("Cannot remove command: %s \n", cmdName)
		return
	}
	fmt.Printf("The following command will be removed: \n")
	fmt.Printf("	%s				%s\n", cmdName, command.Description)
	confirmation := getInput("Are you sure ? [y/N]")
	if confirmation == "y" {
		cmds.Del(cmdName)
		log.Noticef("Command %s has been removed", cmdName)
	}
}

func getInput(display string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(display)
	input, _ := reader.ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}
