package createreactapp

import (
	"fmt"
	"os/exec"
)

const (
	craTemplate string = "typescript"
)

func CreateReactApp(projectName string, template string) {
	fmt.Println("Creating React App project...")
	if template == "" {
		fmt.Println("No template specified, using javascript setup!")
		template = "javascript"
	}
	if template != craTemplate {
		fmt.Println("Invalid template specified, using javascript setup!")
		template = "javascript"
	}
	app := "npx"
	arg0 := "create-react-app"
	arg1 := projectName
	arg2 := "--template"
	arg3 := craTemplate
	cmd, err := exec.Command(app, arg0, arg1, arg2, arg3).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmd))
}
