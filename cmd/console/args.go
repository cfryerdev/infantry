package console

import (
	"github.com/urfave/cli/v2"
)

// SetupArgs Sets up the args for the cli interface
func SetupArgs(plan *string, output *string) *cli.App {

	args := cli.NewApp()
	args.HideHelp = true
	args.Name = "Infantry"
	args.Usage = "Performance and Load testing."
	args.Version = "0.0.1"
	args.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "plan",
			Aliases:     []string{"p"},
			Usage:       "The yaml file (plan) used to orchestrate how to perform load testing.",
			Destination: plan,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "The output file used for reporting.",
			Destination: output,
		},
	}

	return args
}
