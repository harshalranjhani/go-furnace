package vite

import (
	"fmt"
	"os/exec"

	"golang.org/x/exp/slices"
)

var viteTemplates []string = []string{
	"react",
	"vue",
	"preact",
	"lit-element",
	"solid",
	"angular",
	"svelte",
}

func CreateVite(projectName string, template string) {
	app := "npm"
	arg0 := "create"
	arg1 := "vite@latest"
	arg2 := projectName
	arg3 := "--template"
	arg4 := template

	if template == "" {
		fmt.Println("No template specified, using default template: react")
		arg4 = "react"
	}

	if !slices.Contains(viteTemplates, template) {
		fmt.Println("Invalid template specified, using default template: react")
		arg4 = "react"
	}

	fmt.Println("Creating Vite project...")
	cmd, err := exec.Command(app, arg0, arg1, arg2, arg3, arg4).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmd))
}
