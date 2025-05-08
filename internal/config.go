package internal

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	TemplateDir string `toml:"template-dir"`
}

func GetConfig() (Config, error) {
	var config Config

	f, err := os.ReadFile("config.toml")
	if err != nil {
		fmt.Println("config.toml not found")
		return Config{}, fmt.Errorf("config.toml not found")
	}

	_, err = toml.Decode(string(f), &config)
	if err != nil {
		fmt.Println("error decoding config.toml:", err)
		return Config{}, fmt.Errorf("error decofing config.toml")
	}

	return config, nil
}
