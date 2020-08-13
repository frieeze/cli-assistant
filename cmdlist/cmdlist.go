package cmdlist

import (
	"encoding/json"
	"io/ioutil"

	"github.com/op/go-logging"
)

// Command file name
var commandFileName = "commands.json"

// RepoManager interface
type RepoManager interface {
	Get(string) Command
	GetAll() map[string]Command
	GetTypes() map[string][]string
	Add(string, Command)
	Del(string)
	Save(*logging.Logger)
}

// Command struct
type Command struct {
	Alias       string   `json:"Alias"`
	Type        string   `json:"Type"`
	Args        []string `json:"Args"`
	Description string   `json:"Description"`
}

// Repo struct
type Repo struct {
	Repo  map[string]Command  `json:"Repo"`
	Types map[string][]string `json:"Types"`
}

// New Repo
func New(log *logging.Logger) *Repo {
	file, err := ioutil.ReadFile(commandFileName)
	if err != nil {
		log.Error(err)
		return &Repo{Repo: map[string]Command{
			"help": Command{
				Alias:       "",
				Type:        "help",
				Args:        []string{},
				Description: "Display every recorded commands",
			},
			"add": Command{
				Alias:       "",
				Type:        "add",
				Args:        []string{},
				Description: "Add or update a command. To update an existing command enter the name when asked",
			},
			"del": Command{
				Alias:       "",
				Type:        "del",
				Args:        []string{},
				Description: "Remove a recorded command",
			},
		},
			Types: map[string][]string{
				"application": {"", "a", "A", "app", "application", "Application"},
				"browser":     {"b", "B", "browser", "Browser"}},
		}
	}
	data := Repo{}
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Error(err)
		return &Repo{Repo: map[string]Command{
			"help": Command{
				Alias:       "",
				Type:        "help",
				Args:        []string{},
				Description: "Display every recorded commands",
			},
			"add": Command{
				Alias:       "",
				Type:        "add",
				Args:        []string{},
				Description: "Add or update a command. To update an existing command enter the name when asked",
			},
			"del": Command{
				Alias:       "",
				Type:        "del",
				Args:        []string{},
				Description: "Remove a recorded command",
			},
		},
			Types: map[string][]string{
				"application": {"", "a", "A", "app", "application", "Application"},
				"browser":     {"b", "B", "browser", "Browser"}},
		}
	}
	return &data
}

// Get Command
func (r *Repo) Get(alias string) Command {
	return r.Repo[alias]
}

// GetAll commands
func (r *Repo) GetAll() map[string]Command {
	return r.Repo
}

// GetTypes commands
func (r *Repo) GetTypes() map[string][]string {
	return r.Types
}

// Add Command
func (r *Repo) Add(name string, cmd Command) {
	r.Repo[name] = cmd
}

// Del Command
func (r *Repo) Del(name string) {
	delete(r.Repo, name)
}

// Save repo to file
func (r *Repo) Save(log *logging.Logger) {
	file, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		log.Error(err)
		return
	}
	err = ioutil.WriteFile(commandFileName, file, 0644)
	if err != nil {
		log.Error(err)
		return
	}
	log.Debug("Commands saved successfully")

}
