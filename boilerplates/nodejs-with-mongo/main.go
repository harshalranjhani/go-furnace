package nodejswithmongo

import (
	"sync"

	"github.com/fatih/color"

	"github.com/harshalranjhani/go-furnace/boilerplates/nodejs-with-mongo/helpers"
	"github.com/harshalranjhani/go-furnace/boilerplates/nodejs-with-mongo/shell"
	"gopkg.in/cheggaaa/pb.v1"
)

func CreateNodeJsWithMongo(wg *sync.WaitGroup) {
	defer wg.Done()
	progress := pb.StartNew(100)
	for i := 0; i < 100; i++ {
		progress.Increment()
	}
	// create the required folders
	helpers.CreateFolders()

	// create the required files
	helpers.CreateFiles()

	color.Green("Project initialized.")

	// change directories
	shell.ChangeDirectories()

	// initialize npm
	shell.InitializeNPM()

	// install dependencies
	shell.InstallDependencies()

	color.Green("Dependencies installed.")
	// start the server and get going!
	color.Red("Server will now be ready to accept connections on port 8080!")
	color.Cyan("Make sure you have MongoDB running on port 27017 :)")
	shell.StartServer()
}
