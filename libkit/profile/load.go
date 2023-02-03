package profile

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func Load(file string, v any) error {
	open, err := os.Open(file)
	if err != nil {
		return err
	}

	ext := strings.ToLower(filepath.Ext(file))
	switch ext {
	case ".yml", ".yaml":
		err = yaml.NewDecoder(open).Decode(v)
	default:
		err = json.NewDecoder(open).Decode(v)
	}

	_ = open.Close()

	return err
}
