package cmds

import (
	"context"
	"fmt"

	"github.com/dc-dc-dc/lambdalabs-cli/api"
)

type CommandHandler interface {
	HandleCommand(ctx context.Context, cmd string, args []string) error
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
			"i":  "instance",
			"fs": "file-system",
		},
		Cmds: map[string]CommandHandler{
			"instance":    NewInstanceCommand(apiHandler),
			"ssh":         NewSSHKeyCommand(apiHandler),
			"file-system": NewFileSystemCommand(apiHandler),
		},
	}
}

func (h *Handler) HandleCommand(ctx context.Context, cmd string, args []string) error {
	if alias, ok := h.AliasMap[cmd]; ok {
		cmd = alias
	}

	if cmd == "help" {
		h.PrintHelp("sub", h.GetAvailableCommands())
		return nil
	}

	if handler, ok := h.Cmds[cmd]; ok {
		if len(args) == 0 {
			h.PrintHelp("sub", handler.GetAvailableCommands())
			return nil // or return error ?
		}
		_cmd := args[0]
		_args := args[1:]
		if !containsString(_cmd, handler.GetAvailableCommands()) {
			h.PrintHelp(cmd, handler.GetAvailableCommands())
			return nil // or error?
		}

		err := handler.HandleCommand(ctx, _cmd, _args)
		if t, ok := err.(*api.APIError); ok {
			// if the conversion is fine print the json directly
			if raw, err := t.Raw(); err == nil {
				prettyJson(raw)
				return nil
			}
		}
		return err
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
