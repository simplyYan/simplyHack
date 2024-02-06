package simplyhack

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/BurntSushi/toml"
)

// SimplyHack struct represents the simplyHack library.
type SimplyHack struct {
	Name   string
	Config map[string]string
}

// New creates a new instance of SimplyHack with a specified TOML file.
func New(fileName string) (*SimplyHack, error) {
	config := make(map[string]string)
	if _, err := toml.DecodeFile(fileName, &config); err != nil {
		return nil, err
	}

	return &SimplyHack{
		Name:   fileName,
		Config: config,
	}, nil
}

// Area transpiles and executes the provided code based on the TOML configuration.
func (sh *SimplyHack) Area(code string) {
	for keyword, replacement := range sh.Config {
		code = strings.ReplaceAll(code, keyword, replacement)
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", code)
	} else {
		cmd = exec.Command("sh", "-c", code)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing code:", err)
	}
}
