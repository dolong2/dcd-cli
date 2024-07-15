package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <resourceId>",
	Short: "command to update a resource",
	Long:  `this command is used to update a resource`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("file", "f", "", "file path where a resource format is defined")
	updateCmd.Flags().StringP("template", "", "", "resource template to json")
}
