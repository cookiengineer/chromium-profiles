package main

import "github.com/cookiengineer/chromium-profiles/actions"
import "fmt"
import "os"

func showUsage() {

	fmt.Println("Usage:")
	fmt.Println("  chromium-profiles list")
	fmt.Println("  chromium-profiles create <name>")
	fmt.Println("  chromium-profiles launch <name>")

}

func main() {

	if len(os.Args) >= 2 {

		switch os.Args[1] {
		case "list":

			actions.List()

		case "create":

			if len(os.Args) == 3 {

				err := actions.Create(os.Args[2])

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

			if len(os.Args) == 3 {

				err := actions.Launch(os.Args[2])

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

