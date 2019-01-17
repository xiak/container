package app

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/xiak/container/app/cmd"
)

func Run() error {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	command := cmd.NewDefaultCommand()
	return command.Execute()
}
