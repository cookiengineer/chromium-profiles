package actions

import "github.com/cookiengineer/chromium-profiles/config"
import "fmt"
import "os"
import "path/filepath"
import "sort"
import "strings"

func List() error {

	entries, err := os.ReadDir(config.Sandboxes())

	if err == nil {

		lines := make([]string, 0)

		for _, entry := range entries {

			if entry.IsDir() == true {

				profile_name    := entry.Name()
				profile_variant := ""

				sandbox := config.Sandbox(profile_name)
				bytes, err2 := os.ReadFile(filepath.Join(sandbox, ".variant"))

				if err2 == nil {
					profile_variant = strings.TrimSpace(string(bytes))
				}

				if profile_variant != "" {
					lines = append(lines, fmt.Sprintf("- %s (%s)", profile_name, profile_variant))
				} else {
					lines = append(lines, fmt.Sprintf("- %s", profile_name))
				}

			}

		}

		sort.Strings(lines)

		fmt.Fprint(os.Stdout, "Available Chromium Profiles:\n\n")

		for _, line := range lines {
			fmt.Fprintf(os.Stdout, "%s\n", line)
		}

		return nil

	} else {
		return err
	}

}
