package config

import "path/filepath"

func Sandbox(name string) string {
	return filepath.Join(Work(), "Sandboxes", name)
}
