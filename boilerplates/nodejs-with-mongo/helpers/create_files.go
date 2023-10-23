package helpers

import (
	"fmt"
	"html/template"
	"os"

	"github.com/harshalranjhani/go-furnace/boilerplates/nodejs-with-mongo/templates"
)

func createFile(fileName, templateContent string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating %s: %v\n", fileName, err)
		os.Exit(1)
	}
	defer file.Close()

	tmpl := template.Must(template.New("template").Parse(templateContent))
	err = tmpl.Execute(file, nil)
	if err != nil {
		fmt.Printf("Error writing to %s: %v\n", fileName, err)
		os.Exit(1)
	}
}

func CreateFiles() {
	// Create route file
	createFile(routesFolder+"/start.js", templates.RouteTemplate)

	// Create model file
	createFile(modelsFolder+"/start.js", templates.ModelTemplate)

	// Create controller file
	createFile(controllersFolder+"/start.js", templates.ControllerTemplate)

	// create server.js
	createFile(ParentFolderName+"/server.js", templates.ServerTemplate)
	fmt.Println("Project structure created successfully.")
}
