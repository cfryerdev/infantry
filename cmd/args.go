package cmd

import (
	"github.com/urfave/cli/v2"
)

// SetupArgs Sets up the args for the cli interface
func SetupArgs(plan string, output string) *cli.App {

	args := cli.NewApp()
	args.Name = "infantry"
	args.Usage = "Performance and Load testing."
	args.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "plan",
			Aliases:     []string{"p"},
			Usage:       "The yaml/json file used to load test.",
			Destination: &plan,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "The output file used for reporting.",
			Destination: &output,
		},
	}

	return args
}
