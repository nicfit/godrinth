package main

import (
	"context"
	"fmt"
	"github.com/nicfit/godrinth/pkg/ferium"
	"os"

	godrinth "github.com/nicfit/godrinth/pkg"
	"github.com/urfave/cli/v3"
)

func getConfig(cmd *cli.Command) *godrinth.Config {
	for _, c := range cmd.Lineage() {
		if cfg, exists := c.Metadata["config"]; exists {
			if config, ok := cfg.(*godrinth.Config); ok {
				return config
			}
		}
	}
	return nil
}

var mainCommand = &cli.Command{
	Name: "godrinth",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Load configuration from `FILE`",
			Value:   ferium.DEFAULT_CONFIG_FILE,
		},
	},
	Commands: []*cli.Command{
		getCommand,
		searchCommand,
	},
	Before: func(ctx context.Context, command *cli.Command) (context.Context, error) {
		// Load configuration
		config, err := ferium.LoadConfig(command.String("config"))
		if err != nil {
			return ctx, fmt.Errorf("failed to load configuration: %w", err)
		}
		command.Metadata["config"] = config

		return ctx, nil
	},
	Action: func(ctx context.Context, command *cli.Command) error {
		fmt.Println("FIXME: TUI", command.Metadata["config"])
		return nil
	},
}

var getCommand = &cli.Command{
	Name: "get",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		fmt.Println("FIXME: GET", getConfig(cmd))

		project, err := godrinth.GetProject(ctx, cmd.Args().First())
		if err != nil {
			return err
		}
		fmt.Println("Project:", godrinth.IndentedJson(project))

		return nil
	},
}

var searchCommand = &cli.Command{
	Name: "search",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		fmt.Println("FIXME: SEARCH", getConfig(cmd))

		projects, err := godrinth.SearchProject(ctx, cmd.Args().First(), &godrinth.SearchOptions{
			Facets: "[[\"project_type=mod\"]]",
			Index:  "relevance",
			Offset: 0,
			Limit:  20,
		})
		if err != nil {
			return err
		}
		fmt.Println("Projects:", godrinth.IndentedJson(projects))

		if len(projects.Hits) == 0 {
			return fmt.Errorf("no projects found for query: %s", cmd.Args().First())
		}

		return nil
	},
}

func main() {
	var ctx = context.Background()

	meta, err := godrinth.GetMeta(ctx)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Metadata Error:", err)
		os.Exit(1)
	}
	fmt.Println("Metadata:", godrinth.IndentedJson(meta))

	if err := mainCommand.Run(ctx, os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}
}
