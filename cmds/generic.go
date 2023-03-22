package cmds

import (
	"fmt"

	"github.com/dc-dc-dc/lambda-cli/api"
)

type CommandHandler interface {
	HandleCommand(cmd string, args []string) error
	GetAvailableCommands() []string
}

type Handler struct {
	Version  string
	AliasMap map[string]string
	Cmds     map[string]CommandHandler
}

func NewHandler(version string, apiHandler *api.APIHandler) *Handler {
	return &Handler{
		Version: version,
		AliasMap: map[string]string{
			"i": "instance",
		},
		Cmds: map[string]CommandHandler{
			"instance": NewInstanceCommand(apiHandler),
		},
	}
}

func (h *Handler) HandleCommand(cmd string, args []string) error {
	if alias, ok := h.AliasMap[cmd]; ok {
		cmd = alias
	}

	if cmd == "help" {
		h.PrintHelp("", h.GetAvailableCommands())
		return nil
	}

	if handler, ok := h.Cmds[cmd]; ok {
		if len(args) == 0 {
			h.PrintHelp("", handler.GetAvailableCommands())
			return nil // or return error ?
		}
		_cmd := args[0]
		_args := args[1:]
		if !containsString(_cmd, handler.GetAvailableCommands()) {
			h.PrintHelp(cmd, handler.GetAvailableCommands())
			return nil // or error?
		}

		return handler.HandleCommand(_cmd, _args)
	}

	return fmt.Errorf("unknown command %s", cmd)
}

func (h *Handler) GetAvailableCommands() []string {
	res := make([]string, len(h.Cmds))
	var i int
	for s := range h.Cmds {
		res[i] = s
		i += 1
	}
	return res
}

func (h *Handler) PrintHelp(handler string, cmds []string) {
	fmt.Println("*********************************")
	fmt.Printf("*  LambdaLabs cli version %s *\n", h.Version)
	fmt.Println("*********************************")
	fmt.Printf("%s commands:\n", handler)
	for _, s := range cmds {
		fmt.Printf("\t %s\n", s)
	}
}

func containsString(needle string, hay []string) bool {
	for _, s := range hay {
		if s == needle {
			return true
		}
	}
	return false
}
