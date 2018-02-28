package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func projectTree() []string {
	var t []string
	t = append(t, "%s/cmd/")
	t = append(t, "%s/internal")
	t = append(t, "%s/internal/platform")
	return t
}

func webTree() []string {
	var t []string
	t = append(t, "%s/cmd/")
	t = append(t, "%s/internal")
	t = append(t, "%s/internal/platform")
	t = append(t, "%s/static")
	return t
}

var usage = `USAGE gotree [new_pkg|new_project|new_web] <name>
DESCRIPTION
        new_pkg
            Create a package
        new_project
            Create a project
        new_web
            Create a web project
EXEMPLE
    gotree new_pkg doctor_who`

func main() {
	var root string
	var name string

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Print(usage)
		os.Exit(0)
	}

	opt := args[0]
	name = args[1]

	switch opt {
	case "new_pkg":
		os.Mkdir(name, 0740)
		// Create README.md
		data := fmt.Sprintf("package %s", name)
		filename := fmt.Sprintf("%s/%s.go", name, name)
		ioutil.WriteFile(filename, []byte(data), 0740)

		filename = fmt.Sprintf("%s/%s_test.go", name, name)
		ioutil.WriteFile(filename, []byte(data), 0740)

	case "new_project":
		tree := projectTree()
		// Create Tree
		for _, t := range tree {
			root = fmt.Sprintf(t, name)
			os.MkdirAll(root, 0774)
		}
	case "new_web":
		tree := webTree()
		// Create Tree
		for _, t := range tree {
			root = fmt.Sprintf(t, name)
			os.MkdirAll(root, 0774)
		}
	default:
		fmt.Print(usage)
		os.Exit(0)
	}

	// Create README.md
	data := fmt.Sprintf("# %s", strings.Title(name))
	filename := fmt.Sprintf("%s/README.md", name)
	ioutil.WriteFile(filename, []byte(data), 0740)
}
