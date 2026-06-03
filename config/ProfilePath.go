package config

import "os"
import "path/filepath"

func WorkRoot() string {

	home, _ := os.UserHomeDir()

	return filepath.Join(home, "Work")

}

func ProfileRoot() string {
	return filepath.Join(WorkRoot(), "Sandboxes")
}

func ProfilePath(name string) string {
	return filepath.Join(WorkRoot(), "Sandboxes", name)
}
