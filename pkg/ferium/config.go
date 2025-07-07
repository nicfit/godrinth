package ferium

import (
	"encoding/json"
	"github.com/adrg/xdg"
	"os"
	"path"

	godrinth "github.com/nicfit/godrinth/pkg"
)

var DEFAULT_CONFIG_FILE = path.Join(xdg.ConfigHome, "ferium", "config.json")

type config struct {
	activeProfile int       `json:"active_profile"`
	profiles      []profile `json:"profiles"`
}

type profile struct {
	Name        string `json:"name"`
	OutputDir   string `json:"output_dir"`
	GameVersion string `json:"game_version"`
	ModLoader   string `json:"mod_loader"`
	Mods        []mod  `json:"mods"`
}

type mod struct {
	Name       string `json:"name"`
	Identifier struct {
		ModrinthProject  string   `json:"ModrinthProject,omitempty"`
		GitHubRepository []string `json:"GitHubRepository,omitempty"`
	} `json:"identifier"`
	CheckGameVersion bool `json:"check_game_version,omitempty"`
}

func (c *config) Profiles() []string {
	profiles := make([]string, len(c.profiles))
	for i, p := range c.profiles {
		profiles[i] = p.Name
	}
	return profiles
}

func LoadConfig(path string) (godrinth.Config, error) {
	if path == "" {
		path = DEFAULT_CONFIG_FILE
	}
	var config config

	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = fp.Close() }()

	decoder := json.NewDecoder(fp)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
