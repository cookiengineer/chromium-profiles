package config

import "path/filepath"

func Sandboxes() string {
	return filepath.Join(Work(), "Sandboxes")
}
