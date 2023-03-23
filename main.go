package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dc-dc-dc/lambda-cli/api"
	"github.com/dc-dc-dc/lambda-cli/cmds"
)

var (
	version = "0.0.1"
)

func main() {
	apiKey, defaultRegion := GetDefaults()
	if apiKey == "" {
		fmt.Println("env LAMBDA_API_KEY not set")
		os.Exit(1)
	}
	apiHandler := api.NewAPIHandler(apiKey, defaultRegion)
	h := cmds.NewHandler(version, apiHandler)
	if len(os.Args) == 1 {
		// print help
		h.PrintHelp("sub", h.GetAvailableCommands())
		os.Exit(1)
		return
	}
	_cmd := os.Args[1]

	_args := os.Args[2:]

	if err := h.HandleCommand(_cmd, _args); err != nil {
		fmt.Println("something went wrong :F")
		fmt.Printf("error trying to execute the command %s with args %v, err - %s", _cmd, _args, err.Error())
		os.Exit(1)
	}
}

func GetDefaults() (string, string) {
	apiKey := os.Getenv("LAMBDA_API_KEY")
	defaultRegion := os.Getenv("DEFAULT_REGION")
	if apiKey != "" && defaultRegion != "" {
		return apiKey, defaultRegion
	}

	home, err := os.UserHomeDir()
	if err == nil {
		f, err := os.Open(fmt.Sprintf("%s/.lambda", home))
		if err == nil {
			defer f.Close()
			raw, err := io.ReadAll(f)
			if err == nil {
				lines := strings.Split(string(raw), "\n")
				for _, line := range lines {
					if apiKey == "" && strings.HasPrefix(line, "LAMBDA_API_KEY=") {
						apiKey = strings.TrimPrefix(line, "LAMBDA_API_KEY=")
					}
					if defaultRegion == "" && strings.HasPrefix(line, "DEFAULT_REGION=") {
						defaultRegion = strings.TrimPrefix(line, "DEFAULT_REGION=")
					}
				}
			}
		}
	}
	return apiKey, defaultRegion
}
