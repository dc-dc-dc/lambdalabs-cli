package main

import (
	"fmt"
	"os"

	"github.com/dc-dc-dc/lambda-cli/api"
	"github.com/dc-dc-dc/lambda-cli/cmds"
)

var (
	version = "0.0.1"
)

func main() {
	apiKey := os.Getenv("LAMBDA_API_KEY")
	if apiKey == "" {
		panic("env LAMBDA_API_KEY not set")
	}
	apiHandler := api.NewAPIHandler(apiKey)
	h := cmds.NewHandler(version, apiHandler)
	if len(os.Args) == 1 {
		// print help
		h.PrintHelp("", h.GetAvailableCommands())
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
