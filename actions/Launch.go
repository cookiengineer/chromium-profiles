package actions

import "github.com/cookiengineer/chromium-profiles/config"
import "fmt"
import "os"
import "os/exec"
import "path/filepath"
import "strings"

func Launch(name string) error {

	sandbox := config.Sandbox(name)

	stat, err0 := os.Stat(sandbox)

	if err0 == nil && stat.IsDir() {

		profile_dir    := filepath.Join(sandbox, "chromium")
		user_data_dir  := filepath.Join(sandbox, "chromium-data")
		extension_dirs := make([]string, 0)

		for _, extension := range bundled_extensions {
			extension_dirs = append(extension_dirs, filepath.Join(sandbox, "chromium-extensions", extension))
		}

		cmd := exec.Command(
			"chromium",
			fmt.Sprintf("--profile-directory=%s", profile_dir),
			fmt.Sprintf("--user-data-dir=%s", user_data_dir),
			fmt.Sprintf("--load-extension=%s", strings.Join(extension_dirs, ",")),
		)

		return cmd.Start()

	} else {
		return fmt.Errorf("profile does not exist: %s", name)
	}

}
