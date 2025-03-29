package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewWriteCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "create <filename> <content>",
		Short: "create  a new file",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			content := args[1]

			if content == "" {
				fmt.Printf(
					"empty content",
				)
			}

			err := os.WriteFile(path, []byte(content), 0644)
			if err != nil {
				fmt.Printf("Error writing file: ", err)
				return
			}

			fmt.Print("file written successfully")
			return
		},
	}
}
