package actions

import "github.com/cookiengineer/chromium-profiles/config"
import "fmt"
import "os"
import "os/exec"
import "path/filepath"
import "time"

func Create(name string) error {

	sandbox := config.ProfilePath(name)
	folders := []string{
		filepath.Join(sandbox, "chromium"),
		filepath.Join(sandbox, "chromium-data"),
		filepath.Join(sandbox, "chromium-extensions"),
	}

	for _, folder := range folders {

		if err := os.MkdirAll(folder, 0755); err != nil {
			return err
		}

	}

	fmt.Fprintf(os.Stdout, "Created Profile: %s\n", name)

	err1 := InstallExtensions(name)

	if err1 == nil {

		fmt.Fprint(os.Stdout, "\n\n\n")
		fmt.Fprint(os.Stdout, "-> Go to chrome://extensions\n")
		fmt.Fprint(os.Stdout, "-> Activate \"Developer Mode\"\n")

		for _, extension := range bundled_extensions {
			extension_dir := filepath.Join(sandbox, "chromium-extensions", extension)
			fmt.Fprintf(os.Stdout, "-> Click on \"Load Unpacked\" and select \"%s\"\n", extension_dir)
		}

		fmt.Fprint(os.Stdout, "\n\n\n")
		time.Sleep(2 * time.Second)

		profile_dir   := filepath.Join(sandbox, "chromium")
		user_data_dir := filepath.Join(sandbox, "chromium-data")

		cmd := exec.Command(
			"chromium",
			// "--headless=new",
			"--no-first-run",
			"--no-default-browser-check",
			"--disable-gpu",
			fmt.Sprintf("--profile-directory=%s", profile_dir),
			fmt.Sprintf("--user-data-dir=%s", user_data_dir),
			"chrome://extensions",
		)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err2 := cmd.Run()

		if err2 == nil {
			return nil
		} else {
			return fmt.Errorf("chromium failed: %s", err2.Error())
		}

	} else {
		return err1
	}

}
