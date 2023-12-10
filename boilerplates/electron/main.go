package electron

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/harshalranjhani/go-furnace/boilerplates/electron/shell"
)

func CreateElectronApp(wg *sync.WaitGroup) {
	defer wg.Done()
	// initialize npm
	shell.InitializeNPM()

	fmt.Println("Creating Electron project...")

	// install electron
	app := "npm"
	arg0 := "install"
	arg1 := "electron"
	cmd, err := exec.Command(app, arg0, arg1).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmd))
}
