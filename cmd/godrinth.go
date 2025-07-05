package main

import (
	"context"
	"fmt"
	"os"

	godrinth "github.com/nicfit/godrinth/pkg"
	"github.com/urfave/cli/v3"
)

var mainCommand = &cli.Command{
	Name: "godrinth",
	Commands: []*cli.Command{
		getCommand,
		searchCommand,
	},
}

var getCommand = &cli.Command{
	Name: "get",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		project, err := godrinth.GetProject(ctx, cmd.Args().First())
		if err != nil {
			return err
		}
		fmt.Println("Project:", project)

		return nil
	},
}

var searchCommand = &cli.Command{
	Name: "search",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		projects, err := godrinth.SearchProject(ctx, cmd.Args().First(), &godrinth.SearchOptions{
			Facets: "[[\"project_type=mod\"]]",
			Index:  "relevance",
			Offset: 0,
			Limit:  20,
		})
		if err != nil {
			return err
		}
		fmt.Println("Projects:", projects)

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
	fmt.Println("Metadata:", meta)

	if err := mainCommand.Run(ctx, os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}
}
