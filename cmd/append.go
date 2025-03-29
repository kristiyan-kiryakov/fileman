package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewAppendCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "append <filename> <content>",
		Short: "append content to file",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			content := args[1]

			if content == "" {
				fmt.Println("content is empty")
			}

			file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("Error opening file: ", err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					fmt.Printf("Error closing file: ", err)
				}
			}(file)

			_, err = file.WriteString(content)
			if err != nil {
				fmt.Printf("Error appending to file: ", err)
			}

			fmt.Printf("file appended successfully")
			return
		},
	}
}
