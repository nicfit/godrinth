package godrinth

import (
	"encoding/json"
	"github.com/adrg/xdg"
	"os"
	"path"
)

type Config struct {
	ActiveProfile int       `json:"active_profile"`
	Profiles      []Profile `json:"profiles"`
}

type Profile struct {
	Name        string `json:"name"`
	OutputDir   string `json:"output_dir"`
	GameVersion string `json:"game_version"`
	ModLoader   string `json:"mod_loader"`
	Mods        []Mod  `json:"mods"`
}

type Mod struct {
	Name       string `json:"name"`
	Identifier struct {
		ModrinthProject  string   `json:"ModrinthProject,omitempty"`
		GitHubRepository []string `json:"GitHubRepository,omitempty"`
	} `json:"identifier"`
	CheckGameVersion bool `json:"check_game_version,omitempty"`
}

var DEFAULT_CONFIG_FILE = path.Join(xdg.ConfigHome, "ferium", "config.json")

func LoadConfig(path string) (*Config, error) {
	var config Config

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
