package main

import (
	"bufio"
	"cli-assistant/cmdlist"
	"cli-assistant/handlers"
	"fmt"
	"os"
	"strings"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("cli-assistant")

func main() {
	initLogger()
	cmds := cmdlist.New(log)
	input, end := "", "q"
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Type '%s' to quit\n", end)
	for true {
		fmt.Printf(">>> ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if input == end {
			return
		}
		handlers.Handle(log, input, cmds)
	}

}

func initLogger() {
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{level} ▶ %{message} %{color:reset}`,
	)
	level := logging.NOTICE

	// format := logging.MustStringFormatter(
	// 	`%{color}%{time:15:04:05.000} %{shortfile} ▶ %{level} %{color:reset} %{message}`,
	// )
	// level := logging.DEBUG

	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatedBackend := logging.NewBackendFormatter(backend, format)
	leveledBackend := logging.AddModuleLevel(formatedBackend)
	leveledBackend.SetLevel(level, "")

	// Set the backends to be used.
	logging.SetBackend(leveledBackend)

	log.Debug("testing this shit")
}
