package helpers

import (
	"fmt"
	"os"
)

const (
	controllersFolder = ParentFolderName + "/controllers"
	routesFolder      = ParentFolderName + "/routes"
	modelsFolder      = ParentFolderName + "/models"
)

func createFolder(folderName string) {
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating %s folder: %v\n", folderName, err)
		os.Exit(1)
	}
}

const (
	ParentFolderName = "nodejswithmongo-go-furnace"
)

func CreateFolders() {
	// Create folders
	createFolder(ParentFolderName)
	createFolder(routesFolder)
	createFolder(modelsFolder)
	createFolder(controllersFolder)
}
