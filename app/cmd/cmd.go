package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xiak/container/app/cmd/container"
	"github.com/xiak/container/app/cmd/templates"
	"github.com/xiak/container/app/cmd/version"
)

func NewDefaultCommand() *cobra.Command {
	return NewCommand()
}

func NewCommand() *cobra.Command {
	cmds := cobra.Command{
		Use:				"Simple code to build a sanbox",
		Short:  			"Linux namespace, cgroup, rootfs",
		Long:   			templates.Long(`
			Using linux namespace, cgroups, rootfs to controll applications
		
			Find more information at:
			https://github.com/xiak/container
		`),
	}
	cmds.AddCommand(version.NewCmdVersion())
	cmds.AddCommand(container.NewCmd())
	return &cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}