package browser

import (
	"os/exec"

	"github.com/op/go-logging"
)

// Open browser tab
func Open(log *logging.Logger, url string) {
	exe := exec.Command("sensible-browser", "--new-tab", url)
	err := exe.Start()
	if err != nil {
		log.Error(err)
	}
	log.Noticef("\"%s\" successfully opened", url)
}
