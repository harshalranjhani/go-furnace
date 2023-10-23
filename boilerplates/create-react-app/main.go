package createreactapp

import (
	"fmt"
	"os/exec"
	"sync"
)

const (
	craTemplate string = "typescript"
)

func CreateReactApp(projectName string, template bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Creating React App project...")
	if template {
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
	} else {
		app := "npx"
		arg0 := "create-react-app"
		arg1 := projectName
		cmd, err := exec.Command(app, arg0, arg1).Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(cmd))
	}

}
