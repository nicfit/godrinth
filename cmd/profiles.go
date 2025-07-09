package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
)

func newProfilesCommand() *cli.Command {
	return &cli.Command{
		Name:  "profiles",
		Usage: "Manage profiles",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List all profiles",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					config := getConfig(cmd)
					fmt.Println("Available profiles:")
					for _, profile := range config.Profiles() {
						fmt.Println(profile.GetName())
					}
					return nil
				},
			},
		},
	}
}
