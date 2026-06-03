package actions

import "github.com/cookiengineer/chromium-profiles/config"
import "github.com/cookiengineer/chromium-profiles/extensions"
import "fmt"
import "os"
import "path/filepath"

func InstallExtensions(name string) error {

	root := config.ProfilePath(name)
	extension_root := filepath.Join(root, "chromium-extensions")

	for _, name := range bundled_extensions {

		target := filepath.Join(extension_root, name)

		err := extensions.CopyExtension(name, target)

		if err == nil {
			fmt.Fprintf(os.Stdout, "Installed Extension: %s\n", name)
		} else {
			return err
		}

	}

	return nil

}
