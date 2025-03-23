package main

import (
	"fmt"
	"os"
)

const TEMPLATES_DIRECTORY = "/etc/pico_project_setup/"
const CMAKE_TEMPLATE_FILE = TEMPLATES_DIRECTORY + "CMakeLists.template"
const PROJECT_SH_TEMPLATE_FILE = TEMPLATES_DIRECTORY + "project_sh.template"
const MAIN_C_TEMPLATE_FILE = TEMPLATES_DIRECTORY + "main_c.template"

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("No arguments provided. Arguments: <project_name>")
	}

	CreateProjectNameFile(args[0])
	CreateProjectFile(CMAKE_TEMPLATE_FILE, "CMakeLists.txt")
	CreateProjectFile(PROJECT_SH_TEMPLATE_FILE, "project.sh")
	CreateSRC()

	fmt.Println("Created new project: " + args[0])
}
