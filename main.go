package main

import (
	"fmt"
	"os"
)

// type Package struct {
// 	name    string   // The name of the package
// 	version string   // The version of the package
// 	files   []string // List of files that they are in the package
// }

func help() {
	// cmds := [...]string{"init", "install", "uninstall", "publish"}
	cmds := make(map[string]string)

	cmds["init"] = "Creates a project"
	cmds["install"] = "Installs a package"
	cmds["uninstall"] = "Uninstalls a package"
	cmds["publish"] = "Publishes a package"

	for k := range cmds {
		fmt.Printf("%s: %s\n", k, cmds[k])
	}
}

func install(pkg string, version string) {
	fmt.Printf("Installing %s@%s\n", pkg, version)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		help()
	} else {
		switch args[1] {
		case "install":
			if len(args) < 3 {
				fmt.Println("Please insert a package name and a version")
			} else {
				if len(args) < 4 {
					fmt.Println("Please insert a version")
				} else {
					install(args[2], args[3])
				}
			}
		}
	}
}
