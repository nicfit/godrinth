package godrinth

import (
	"fmt"
	"os"
)

const (
	Version               = "0.1.0"
	modrinthUrlProdUrl    = "https://api.modrinth.com/"
	modrinthUrlStagingUrl = "https://staging-api.modrinth.com/"
)

var (
	ModrinthUrl    string
	ModrinthApiUrl string
	UserAgent      = fmt.Sprintf("github.com/nicfit/godrinth/%s", Version)
)

func init() {
	ModrinthUrl = modrinthUrlProdUrl
	if os.Getenv("GODRINTH_DEVEL") != "" {
		ModrinthUrl = modrinthUrlStagingUrl
	}
	ModrinthApiUrl = fmt.Sprintf("%sv2/", ModrinthUrl)
}
