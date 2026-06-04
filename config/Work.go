package config

import "os"
import "path/filepath"

func Work() string {

	home, _ := os.UserHomeDir()

	return filepath.Join(home, "Work")

}

