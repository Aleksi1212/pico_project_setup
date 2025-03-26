package main

import (
	"fmt"
	"os"
)

func CreateProjectNameFile(project_name string) {
	var project_file_name string = "project_name.txt"

	_, exists, err := FileExists(project_file_name)
	Check(err)
	if exists {
		panic("Project already created")
	}

	file, err := os.Create("project_name.txt")
	Check(err)

	project_name_bytes := []byte(project_name)
	_, err = file.Write(project_name_bytes)
	Check(err)

	fmt.Println("Created " + project_file_name)
}

func CreateProjectFile(template_file, dst_file string) {
	info, exists, err := FileExists(template_file)
	Check(err)
	if !exists {
		panic("No such file as " + template_file)
	}

	err = CopyFileContents(template_file, dst_file)
	Check(err)

	if template_file == PROJECT_SH_TEMPLATE_FILE {
		err = os.Chmod(dst_file, info.Mode()|0111)
		Check(err)
	}

	fmt.Println("Created " + dst_file)
}

func CreateSRC() {
	var src_dir string = "src"
	var main_c_file string = src_dir + "/main.c"

	err := os.Mkdir(src_dir, os.ModePerm)
	Check(err)

	CreateProjectFile(MAIN_C_TEMPLATE_FILE, main_c_file)
}
