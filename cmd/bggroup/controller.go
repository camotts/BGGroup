package main

import (
	"github.com/camotts/bggroup/controller"
	"github.com/camotts/bggroup/controller/config"
	"github.com/michaelquigley/cf"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var controllerCmd *controllerCommand

type controllerCommand struct {
	cmd *cobra.Command
}

func init() {
	controllerCmd = newControllerCommand()
	rootCmd.AddCommand(controllerCmd.cmd)
}

func newControllerCommand() *controllerCommand {
	cmd := &cobra.Command{
		Use:  "controller <configPath>",
		Args: cobra.ExactArgs(1),
	}
	command := &controllerCommand{cmd: cmd}
	cmd.Run = command.run
	return command
}

func (cmd *controllerCommand) run(_ *cobra.Command, args []string) {
	cfg, err := config.LoadConfig(args[0])
	if err != nil {
		panic(err)
	}
	logrus.Infof(cf.Dump(cfg, cf.DefaultOptions()))
	if err := controller.Run(cfg); err != nil {
		panic(err)
	}
}
