package cmds

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/dc-dc-dc/lambda-cli/api"
)

type InstanceCommand struct {
	apiHandler *api.APIHandler
}

func NewInstanceCommand(apiHandler *api.APIHandler) CommandHandler {
	return &InstanceCommand{
		apiHandler: apiHandler,
	}
}

func (t *InstanceCommand) HandleCommand(cmd string, args []string) error {
	fmt.Printf("got cmd %s, args - %v\n", cmd, args)
	switch cmd {
	case "create":
		return t.handleInstanceCreateCommand(args)
	case "delete":
		return t.handleInstanceDeleteCommand(args)
	case "list":
		return t.handleInstanceListCommand(args)
	case "get":
		return t.handleInstanceGetCommand(args)
	}
	return fmt.Errorf("unknown cmd %s", cmd)
}

func (t *InstanceCommand) handleInstanceCreateCommand(args []string) error {
	// region, instance_type_name, ssh_key_names - required
	// quantity, name, file_system_names
	region := flag.String("region", "", "the region of the instance")
	instanceType := flag.String("instance-type", "", "the instance type")
	sshKeys := flag.String("ssh-keys", "", "comma seperated name of ssh-keys to install")
	fileSystems := flag.String("file-systems", "", "comma seperated names of file systems to add")
	quantity := flag.Int("q", 1, "number of instances to spin up")
	name := flag.String("name", "", "the name of the instance")
	flag.Parse()
	httpRes, err := t.apiHandler.Post(context.TODO(), "/instance-operations/launch", api.InstanceCreateAPIRequest{
		RegionName:       *region,
		InstanceTypeName: *instanceType,
		SSHKeyNames:      strings.Split(*sshKeys, ","),
		FileSystemNames:  strings.Split(*fileSystems, ","),
		Name:             name,
		Quantity:         *quantity,
	})
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()
	data := &api.InstanceCreateAPIResponse{}
	if err := json.NewDecoder(httpRes.Body).Decode(&data); err != nil {
		return err
	}
	fmt.Printf("%s\n", strings.Join(data.Data.InstanceIds, ", "))
	return nil
}

func (t *InstanceCommand) handleInstanceDeleteCommand(args []string) error {
	// TODO: validate the ids
	if _, err := t.apiHandler.Post(context.TODO(), "/instance-operations/terminate", api.InstanceDeleteApiRequest{InstanceIds: args}); err != nil {
		return err
	}

	return nil
}

func (t *InstanceCommand) handleInstanceGetCommand(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("instance get expected id")
	}
	id := args[0]
	httpRes, err := t.apiHandler.Get(context.TODO(), fmt.Sprintf("/instances/%s", id))
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()
	data := &api.InstanceGetAPIResponse{}
	if err := json.NewDecoder(httpRes.Body).Decode(&data); err != nil {
		return err
	}
	fmt.Printf("%+v", data.Data)
	return nil
}

func (t *InstanceCommand) handleInstanceListCommand(args []string) error {
	httpRes, err := t.apiHandler.Get(context.TODO(), "/instances")
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()
	data := &api.ListInstanceResponse{}
	if err := json.NewDecoder(httpRes.Body).Decode(&data); err != nil {
		return err
	}

	if len(data.Data) == 0 {
		fmt.Println("no instances running")
	}

	for _, s := range data.Data {
		fmt.Printf("%+v\n", s)
	}
	return nil
}

func (t *InstanceCommand) GetAvailableCommands() []string {
	return []string{"create", "delete", "list", "get"}
}
