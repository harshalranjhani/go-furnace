package helpers

import (
	"fmt"
	"os"
)

const (
	controllersFolder = "controllers"
	routesFolder      = "routes"
	modelsFolder      = "models"
)

func createFolder(folderName string) {
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating %s folder: %v\n", folderName, err)
		os.Exit(1)
	}
}

func CreateFolders() {
	// Create folders
	createFolder(routesFolder)
	createFolder(modelsFolder)
	createFolder(controllersFolder)
}
