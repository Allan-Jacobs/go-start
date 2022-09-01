package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/lithammer/dedent"
	"github.com/stoewer/go-strcase"
)

type params struct {
	name        string
	moduleUrl   string
	open_vscode bool
}

func getParams() params {
	var name string
	var moduleUrl string
	var open_vscode bool

	isInteractive := flag.Bool("i", false, "Run interactively")

	flag.StringVar(&name, "n", "", "the name of the project")
	flag.StringVar(&moduleUrl, "u", "", "the module url")
	flag.BoolVar(&open_vscode, "vscode", false, "open the created project vscode")

	flag.Parse()

	if *isInteractive {
		var useGithub string
		fmt.Print("Use Github (y/n): ")
		fmt.Scanln(&useGithub)

		if useGithub == "y" {
			var repoName string
			var githubUsername string
			fmt.Print("Github username: ")
			fmt.Scanln(&githubUsername)

			fmt.Print("Repo name: ")

			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				line := scanner.Text()
				repoName = strcase.KebabCase(line)
			}

			moduleUrl = "github.com/" + githubUsername + "/" + repoName
			name = repoName
		} else {
			fmt.Print("Url: ")
			fmt.Scanln(&moduleUrl)

			fmt.Print("Project name: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				line := scanner.Text()
				name = strcase.KebabCase(line)
			}
		}
		fmt.Print("open VSCode?: (y/n) ")
		var res string
		fmt.Scanln(&res)
		open_vscode = res == "y"
	}

	return params{name, moduleUrl, open_vscode}
}

func main() {
	vals := getParams()

	if vals.name == "" {
		fmt.Println("A project name must be supplied")
		return
	}

	if vals.moduleUrl == "" {
		fmt.Println("A module url must be supplied")
		return
	}

	err := os.Mkdir(vals.name, 0775)

	if err != nil {
		fmt.Println("An Error Occurred: ", err)
		return
	}

	err = os.Chdir(vals.name)

	if err != nil {
		fmt.Println("An Error Occurred: ", err)
		return
	}

	cmd := exec.Command("go", "mod", "init", vals.moduleUrl)

	cmd.Run()

	contents := strings.Trim(dedent.Dedent(`
	package main
	
	func main() {

	}
	`), "\n")

	f, err := os.Create("main.go")

	if err != nil {
		fmt.Println("An Error Occurred: ", err)
		return
	}

	defer f.Close()

	f.WriteString(contents)

	fmt.Printf("Created Project \"%s\"\n", vals.name)

	if vals.open_vscode {
		cmd = exec.Command("code", ".", "-r", "-g", "main.go:4")

		cmd.Run()
	}

}
