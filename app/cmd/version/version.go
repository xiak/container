package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

const Version string = "v0.0.1"

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use: "version",
		Short: "Print Short version",
		Long:  "Print Long version",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	}
	return cmd
}

func printVersion() {
	fmt.Printf("Version: %s\n", Version)
}