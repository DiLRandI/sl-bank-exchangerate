package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func ParseConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read the file, [%w]", err)
	}

	var config Config

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unable to parse the config, [%w]", err)
	}

	return &config, nil
}
