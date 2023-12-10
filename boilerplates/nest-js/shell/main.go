package shell

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	ParentFolderName = "nestjs-go-furnace"
)

func CreateParentFolder() {
	// Get the current working directory
	_, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Create the parent folder
	if err := os.Mkdir(ParentFolderName, 0755); err != nil {
		fmt.Println("Error creating parent folder:", err)
		return
	}
}

func ChangeDirectories() {
	// Get the desired directory
	dir := ParentFolderName

	// Change the working directory
	if err := os.Chdir(dir); err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}
}

func InitializeNPM() {
	app := "npm"
	arg0 := "init"
	arg1 := "-y"

	cmd := exec.Command(app, arg0, arg1)
	_, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func InstallDependencies() {
	app := "npm"
	arg0 := "install"
	arg1 := "-g"
	arg2 := "@nestjs/cli"

	cmd := exec.Command(app, arg0, arg1, arg2)
	_, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
