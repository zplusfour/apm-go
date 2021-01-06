package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

type Package struct {
	file []string
}

func safe(e error) {
	log.Fatal(e)
}

func install(pkg string) {
	splitted := strings.Split(pkg, "@")
	res, err := http.Get("https://registry010.theboys619.repl.co/packages/" + splitted[0] + "@" + splitted[1])

	if err != nil {
		safe(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		safe(err)
	}

	data := string(body)
	pkg := new(Package)

	json.Unmarshal([]byte(data), &pkg)
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
				pkg := args[2]
				install(pkg)
			}
		}
	}
}
