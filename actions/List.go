package actions

import "github.com/cookiengineer/chromium-profiles/config"
import "fmt"
import "os"
import "sort"

func List() error {

	entries, err := os.ReadDir(config.ProfileRoot())

	if err == nil {

		lines := make([]string, 0)

		for _, entry := range entries {

			if entry.IsDir() == true {
				lines = append(lines, fmt.Sprintf("- %s", entry.Name()))
			}

		}

		sort.Strings(lines)

		fmt.Fprint(os.Stdout, "Available Profiles:\n")

		for _, line := range lines {
			fmt.Fprintf(os.Stdout, "%s\n", line)
		}

		return nil

	} else {
		return err
	}

}
