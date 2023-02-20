package profile

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Load 加载并序列化配置
func Load(file string, v any) error {
	open, err := os.Open(file)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer open.Close()
	stat, err := os.Stat(file)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		file = filepath.Join(file, "manager.yaml")
	}

	ext := strings.ToLower(filepath.Ext(file))
	switch ext {
	case ".yml", ".yaml":
		err = yaml.NewDecoder(open).Decode(v)
	default:
		err = json.NewDecoder(open).Decode(v)
	}

	return err
}
