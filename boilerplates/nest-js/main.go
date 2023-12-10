package nestjs

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/harshalranjhani/go-furnace/boilerplates/nest-js/shell"
)

func CreateNestApp(projectName string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Creating Nestjs project...")

	// create parent folder
	shell.CreateParentFolder()

	// change directories
	shell.ChangeDirectories()

	// initialize npm
	shell.InitializeNPM()

	// install dependencies
	shell.InstallDependencies()

	app := "nest"
	arg0 := "new"
	arg1 := projectName
	cmd, err := exec.Command(app, arg0, arg1).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmd))
}
