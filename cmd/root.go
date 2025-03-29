package cmd

import (
	"github.com/spf13/cobra"
)

const rootCmdName = "fileman"

func NewRootCmd() *cobra.Command {
	cmdInstance := &cobra.Command{
		Use:   rootCmdName,
		Short: rootCmdName,
		Long:  rootCmdName,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmdInstance.AddCommand(
		NewReadCommand(),
		NewWriteCommand(),
		NewAppendCommand(),
	)

	return cmdInstance
}
