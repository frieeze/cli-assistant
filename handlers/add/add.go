package add

import (
	"bufio"
	"cli-assistant/cmdlist"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/op/go-logging"
)

// Add Command
func Add(log *logging.Logger, cmds cmdlist.RepoManager) {
	fmt.Printf("Adding/Updating command: \n")

	cmdName := getInput("	Command Name: ")
	input := getInput("	Command type ([A]pplication | [b]rowser): ")
	cmdType := getCommandType(input, cmds)
	if cmdType == "unknown" {
		log.Warningf("Command type unknown: %s\n", input)
		return
	}
	cmdAlias := getInput("	Command alias: ")
	if cmdAlias == "" {
		log.Errorf("No alias \n")
		return
	}
	cmdArgs := strings.Fields(getInput("	Command Args: "))
	cmdDescription := getInput("	Command description: ")
	command := cmdlist.Command{
		Alias:       cmdAlias,
		Type:        cmdType,
		Args:        cmdArgs,
		Description: cmdDescription,
	}
	cmds.Add(cmdName, command)
	log.Noticef("Command '%s' successfully added/updated", cmdName)
	go saveCommands(log, cmds)
}

func saveCommands(log *logging.Logger, cmds cmdlist.RepoManager) {
	file, err := json.MarshalIndent(cmds, "", " ")
	if err != nil {
		log.Error(err)
		return
	}
	err = ioutil.WriteFile(cmdlist.CommandFileName, file, 0644)
	if err != nil {
		log.Error(err)
		return
	}
	log.Debug("Commands saved successfully")

}

func getInput(display string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(display)
	input, _ := reader.ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}

func getCommandType(input string, cmds cmdlist.RepoManager) string {
	for key, accepted := range cmds.GetTypes() {
		for _, i := range accepted {
			if i == input {
				return key
			}
		}

	}
	return "unknown"
}
