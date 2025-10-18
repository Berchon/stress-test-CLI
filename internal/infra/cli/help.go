package cli

import (
	"github.com/spf13/cobra"
)

func NewHelpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "Displays help information about the CLI",
		Long:  `Shows usage examples and explanation of flags for the stress-test CLI tool.`,
		Run: func(cmd *cobra.Command, args []string) {
			PrintHelp()
		},
	}
}
