package application

import (
	"cli-assistant/cmdlist"
	"os/exec"

	"github.com/op/go-logging"
)

// Open application
func Open(log *logging.Logger, cmd cmdlist.Command, name string) {
	exe := exec.Command(cmd.Alias, cmd.Args...)
	err := exe.Start()
	if err != nil {
		log.Error(err)
		return
	}
	log.Noticef("\"%s\" successfully executed", name)

}
