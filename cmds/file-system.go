package cmds

import (
	"context"
	"fmt"
	"io"

	"github.com/dc-dc-dc/lambda-cli/api"
)

type FileSystemCommand struct {
	apiHandler *api.APIHandler
}

func NewFileSystemCommand(apiHandler *api.APIHandler) CommandHandler {
	return &FileSystemCommand{
		apiHandler: apiHandler,
	}
}

func (f *FileSystemCommand) HandleCommand(cmd string, args []string) error {
	switch cmd {
	case "list":
		return f.listFileSystems(args)
	}
	return fmt.Errorf("unknown cmd %s", cmd)
}

func (f *FileSystemCommand) listFileSystems(args []string) error {
	httpRes, err := f.apiHandler.Get(context.TODO(), "/file-systems")
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()
	raw, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	prettyJson(raw)

	return nil
}

func (f *FileSystemCommand) GetAvailableCommands() []string {
	return []string{"list"}
}
