package main

import (
	"context"
	"fmt"
	godrinth "github.com/nicfit/godrinth/pkg"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Panicln(fmt.Sprintf("usage: %s PROJECT_SLUG_OR_ID", os.Args[0]))
	}

	var ctx = context.Background()

	meta, err := godrinth.GetMeta(ctx)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Metadata:", meta)

	project, err := godrinth.GetProject(ctx, os.Args[1])
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Project:", project.Slug)
}
