package cmds

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
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
	if os.Getenv("DEBUG") == "1" {
		fmt.Printf("got cmd %s, args - %v\n", cmd, args)
	}
	switch cmd {
	case "types":
		return t.handleInstanceTypeListCommand(args)
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

func (t *InstanceCommand) handleInstanceTypeListCommand(args []string) error {
	httpRes, err := t.apiHandler.Get(context.TODO(), "/instance-types")
	if err != nil {
		return err
	}
	rawData, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	prettyJson(rawData)
	return nil
}

func (t *InstanceCommand) handleInstanceCreateCommand(args []string) error {
	// region, instance_type_name, ssh_key_names - required
	// quantity, name, file_system_names
	fs := flag.NewFlagSet("create", flag.ContinueOnError)
	region := fs.String("region", "", "the region of the instance")
	instanceType := fs.String("type", "", "the instance type")
	sshKeys := fs.String("ssh-keys", "", "comma seperated name of ssh-keys to install")
	fileSystemNames := fs.String("file-systems", "", "comma seperated names of file systems to add")
	quantity := fs.Int("q", 1, "number of instances to spin up")
	name := fs.String("name", "", "the name of the instance")

	fs.Parse(args)
	req := api.InstanceCreateAPIRequest{
		RegionName:       *region,
		InstanceTypeName: *instanceType,
		SSHKeyNames:      strings.Split(*sshKeys, ","),
		Quantity:         *quantity,
	}
	if *fileSystemNames != "" {
		req.FileSystemNames = strings.Split(*fileSystemNames, ",")

	}
	if *name != "" {
		req.Name = name
	}
	httpRes, err := t.apiHandler.Post(context.TODO(), "/instance-operations/launch", req)
	if err != nil {
		return err
	}
	rawData, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	prettyJson(rawData)
	return nil
}

func (t *InstanceCommand) handleInstanceDeleteCommand(args []string) error {
	// TODO: validate the ids
	httpRes, err := t.apiHandler.Post(context.TODO(), "/instance-operations/terminate", api.InstanceDeleteApiRequest{InstanceIds: args})
	if err != nil {
		return err
	}
	rawData, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	prettyJson(rawData)
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
	rawData, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	prettyJson(rawData)
	return nil
}

func (t *InstanceCommand) handleInstanceListCommand(args []string) error {
	httpRes, err := t.apiHandler.Get(context.TODO(), "/instances")
	if err != nil {
		return err
	}
	rawData, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	prettyJson(rawData)
	// defer httpRes.Body.Close()
	// data := &api.ListInstanceResponse{}
	// if err := json.NewDecoder(httpRes.Body).Decode(&data); err != nil {
	// 	return err
	// }

	// if len(data.Data) == 0 {
	// 	fmt.Println("no instances running")
	// }

	// for _, s := range data.Data {
	// 	fmt.Printf("%+v\n", s)
	// }
	return nil
}

func (t *InstanceCommand) GetAvailableCommands() []string {
	return []string{"types", "create", "delete", "list", "get"}
}
