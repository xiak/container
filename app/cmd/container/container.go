package container

import (
	"github.com/spf13/cobra"
	"github.com/xiak/container/app/cmd/templates"
)

func NewCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "run",
		Short: "Run a container",
		Long:  templates.Long("Create a container"),
		Example: templates.Example(`
  			app run
		`),
	}
	return &cmd
}
