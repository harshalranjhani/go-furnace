package shell

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/harshalranjhani/go-furnace/boilerplates/nodejs-with-mongo/helpers"
)

func ChangeDirectories() {
	// Get the desired directory
	dir := helpers.ParentFolderName

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
	arg1 := "express"
	arg2 := "mongoose"

	cmd := exec.Command(app, arg0, arg1, arg2)
	_, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func StartServer() {
	app := "node"
	arg0 := "server.js"

	cmd := exec.Command(app, arg0)
	_, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
