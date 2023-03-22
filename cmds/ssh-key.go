package cmds

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/dc-dc-dc/lambda-cli/api"
)

type SSHKeyCommand struct {
	apiHandler *api.APIHandler
}

func NewSSHKeyCommand(apiHandler *api.APIHandler) CommandHandler {
	return &SSHKeyCommand{
		apiHandler: apiHandler,
	}
}

func (s *SSHKeyCommand) HandleCommand(cmd string, args []string) error {
	switch cmd {
	case "list":
		return s.listSSHKeys(args)
	case "add":
		return s.addSSHKey(args)
	case "delete":
		return s.deleteSSHKey(args)
	}
	return fmt.Errorf("unknown cmd %s", cmd)
}

func (s *SSHKeyCommand) listSSHKeys(args []string) error {
	httpRes, err := s.apiHandler.Get(context.TODO(), "/ssh-keys")
	if err != nil {
		return err
	}
	raw, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	prettyJson(raw)
	return nil
}

func (s *SSHKeyCommand) addSSHKey(args []string) error {
	fs := flag.NewFlagSet("add-ssh", flag.ContinueOnError)
	name := fs.String("name", "", "name of the ssh key")
	publicKey := fs.String("public-key", "", "public key to upload")
	if err := fs.Parse(args); err != nil {
		return err
	}
	httpRes, err := s.apiHandler.Post(context.TODO(), "/ssh-keys", api.SSHAddRequest{
		Name:      *name,
		PublicKey: *publicKey,
	})

	if err != nil {
		return err
	}
	raw, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}

	prettyJson(raw)

	return nil
}

func (s *SSHKeyCommand) deleteSSHKey(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected one arg of ssh-key id got %d args", len(args))
	}
	id := args[0]
	if _, err := s.apiHandler.Delete(context.TODO(), fmt.Sprintf("/ssh-keys/%s", id)); err != nil {
		return err
	}
	return nil
}

func (s *SSHKeyCommand) GetAvailableCommands() []string {
	return []string{"list", "add", "delete"}
}
