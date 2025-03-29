package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewReadCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "read <filename>",
		Short: "Read and display the contents of a file",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]

			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading file: ", err)
				return
			}

			fileContent := string(data)

			fmt.Println(fileContent)
			return
		},
	}
}
