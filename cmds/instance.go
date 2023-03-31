package cmds

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dc-dc-dc/lambdalabs-cli/api"
)

type InstanceCommand struct {
	apiHandler *api.APIHandler
}

func NewInstanceCommand(apiHandler *api.APIHandler) CommandHandler {
	return &InstanceCommand{
		apiHandler: apiHandler,
	}
}

func (t *InstanceCommand) HandleCommand(ctx context.Context, cmd string, args []string) error {
	if os.Getenv("DEBUG") == "1" {
		fmt.Printf("got cmd %s, args - %v\n", cmd, args)
	}
	switch cmd {
	case "types":
		return t.handleInstanceTypeListCommand(ctx, args)
	case "create":
		return t.handleInstanceCreateCommand(ctx, args)
	case "delete":
		return t.handleInstanceDeleteCommand(ctx, args)
	case "list":
		return t.handleInstanceListCommand(ctx, args)
	case "get":
		return t.handleInstanceGetCommand(ctx, args)
	}
	return fmt.Errorf("unknown cmd %s", cmd)
}

func (t *InstanceCommand) handleInstanceTypeListCommand(ctx context.Context, args []string) error {
	httpRes, err := t.apiHandler.Get(ctx, "/instance-types")
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()
	rawData, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	prettyJson(rawData)
	return nil
}

func (t *InstanceCommand) handleInstanceCreateCommand(ctx context.Context, args []string) error {
	// region, instance_type_name, ssh_key_names - required
	// quantity, name, file_system_names
	fs := flag.NewFlagSet("create", flag.ContinueOnError)
	region := fs.String("region", "", "the region of the instance")
	instanceType := fs.String("type", "", "the instance type")
	sshKeys := fs.String("ssh-keys", "", "comma seperated name of ssh-keys to install")
	fileSystemNames := fs.String("file-systems", "", "comma seperated names of file systems to add")
	quantity := fs.Int("q", 1, "number of instances to spin up")
	name := fs.String("name", "", "the name of the instance")

	if err := fs.Parse(args); err != nil {
		return err
	}
	req := api.InstanceCreateAPIRequest{
		RegionName:       *region,
		InstanceTypeName: *instanceType,
		SSHKeyNames:      strings.Split(*sshKeys, ","),
		Quantity:         *quantity,
		Name:             *name,
	}
	if *fileSystemNames != "" {
		req.FileSystemNames = strings.Split(*fileSystemNames, ",")

	}
	// use default region if not provided
	if req.RegionName == "" {
		req.RegionName = t.apiHandler.GetDefaultRegion()
	}

	httpRes, err := t.apiHandler.Post(ctx, "/instance-operations/launch", req)
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()
	rawData, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	prettyJson(rawData)
	return nil
}

func (t *InstanceCommand) handleInstanceDeleteCommand(ctx context.Context, args []string) error {
	// TODO: validate the ids
	httpRes, err := t.apiHandler.Post(ctx, "/instance-operations/terminate", api.InstanceDeleteApiRequest{InstanceIds: args})
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()
	rawData, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	prettyJson(rawData)
	return nil
}

func (t *InstanceCommand) handleInstanceGetCommand(ctx context.Context, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("instance get expected id")
	}
	id := args[0]
	httpRes, err := t.apiHandler.Get(ctx, fmt.Sprintf("/instances/%s", id))
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()
	rawData, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	prettyJson(rawData)
	return nil
}

func (t *InstanceCommand) handleInstanceListCommand(ctx context.Context, args []string) error {
	httpRes, err := t.apiHandler.Get(ctx, "/instances")
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()
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
