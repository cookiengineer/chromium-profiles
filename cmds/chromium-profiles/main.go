package main

import "github.com/cookiengineer/chromium-profiles/actions"
import "github.com/cookiengineer/chromium-profiles/extensions"
import "fmt"
import "os"
import "strings"

func showUsage() {

	fmt.Fprint(os.Stdout, "Usage:\n")
	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprint(os.Stdout, "  chromium-profiles list\n")
	fmt.Fprint(os.Stdout, "  chromium-profiles create <name> <variant>\n")
	fmt.Fprint(os.Stdout, "  chromium-profiles launch <name> [--debug]\n")
	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprint(os.Stdout, "Flags:\n")
	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprint(os.Stdout, "--debug    Enables remote debugging on port 9222\n")
	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprint(os.Stdout, "Variants:\n")
	fmt.Fprint(os.Stdout, "\n")

	variants := extensions.Variants()

	for _, variant := range variants {
		fmt.Fprintf(os.Stdout, "  %s\n", variant)
	}

	fmt.Fprint(os.Stdout, "\n")

}

func main() {

	if len(os.Args) >= 2 {

		switch os.Args[1] {
		case "list":

			actions.List()

		case "create":

			if len(os.Args) == 4 {

				name := strings.TrimSpace(os.Args[2])
				variant := strings.TrimSpace(os.Args[3])

				if extensions.IsVariant(variant) {

					err := actions.Create(name, variant)

					if err == nil {
						os.Exit(0)
					} else {
						fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
						os.Exit(1)
					}

				} else {
					fmt.Fprintf(os.Stderr, "Error: \"%s\" is not a valid variant\n", variant)
					os.Exit(1)
				}

			} else if len(os.Args) == 3 {

				name := strings.TrimSpace(os.Args[2])
				variant := "win10-edge"
				err := actions.Create(name, variant)

				if err == nil {
					os.Exit(0)
				} else {
					fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
					os.Exit(1)
				}

			} else {
				showUsage()
				os.Exit(1)
			}

		case "launch":

			debug := false
			name := ""

			for _, arg := range os.Args[2:] {

				if arg == "--debug" || arg == "-debug" {
					debug = true
				} else if name == "" {
					name = strings.TrimSpace(arg)
				}

			}

			if name != "" {

				err := actions.Launch(name, debug)

				if err == nil {
					os.Exit(0)
				} else {
					fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
					os.Exit(1)
				}

			} else {
				showUsage()
				os.Exit(1)
			}

		default:
			showUsage()
			os.Exit(1)
		}

	} else {

		showUsage()
		os.Exit(1)

	}

}
