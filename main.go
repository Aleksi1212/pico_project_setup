package main

import (
	"fmt"
	"os"
)

const TEMPLATES_DIRECTORY = "/etc/pico_project_setup/"
const CMAKE_TEMPLATE_FILE = TEMPLATES_DIRECTORY + "CMakeLists.template"
const MAKEFILE_TEMPLATE = TEMPLATES_DIRECTORY + "makefile.template"
const MAIN_C_TEMPLATE_FILE = TEMPLATES_DIRECTORY + "main.template"

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("No arguments provided. Arguments: <project_name>")
	}

	CreateProjectNameFile(args[0])
	CreateProjectFile(CMAKE_TEMPLATE_FILE, "CMakeLists.txt")
	CreateProjectFile(MAKEFILE_TEMPLATE, "Makefile")
	CreateSRC(args[1])

	fmt.Println("Created new project: " + args[0])
}
