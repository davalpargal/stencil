package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// ProtocCmd calls protoc
func ProtocCmd() *cobra.Command {
	var configFile string

	cmd := &cobra.Command{
		Use:     "protoc",
		Aliases: []string{"p"},
		Short:   "Execute a protoc command",
		Example: "stencil protoc -- --descriptor_set_out \"<abs-path for op>\" --include_imports <abs-dir-path for import proto>/*.proto --proto_path \"<abs-path for ip>\"",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := "protoc"
			command := exec.Command(app, args...)
			protoPath, err := getProtoPathFromArgs(args)
			if err != nil {
				return err
			}
			//execute the protoc command in specified proto_path
			if protoPath != "" {
				command.Dir = protoPath
			}
			stdout, err := command.Output()
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			fmt.Println(string(stdout))
			fmt.Println("Successfully executed Protoc operation.")
			return nil
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "./config.yaml", "Config file path")
	return cmd
}

func getProtoPathFromArgs(args []string) (string, error) {
	for i, arg := range args {
		if arg == "--proto_path" {
			if i+1 == len(args) {
				return "", errors.New("--proto_path value not provided")
			}
			return args[i+1], nil
		}
	}
	return "", nil
}
