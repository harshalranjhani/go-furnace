package main

import (
	"sync"

	createreactapp "github.com/harshalranjhani/go-furnace/boilerplates/create-react-app"
	nestjs "github.com/harshalranjhani/go-furnace/boilerplates/nest-js"
	nodejswithmongo "github.com/harshalranjhani/go-furnace/boilerplates/nodejs-with-mongo"
	"github.com/rivo/tview"
)

var selectedItems map[string]bool
var selectedItemsText *tview.TextView

var wg *sync.WaitGroup

var optionsMap = map[string]func(){
	"CRA - Javascript":    func() { createreactapp.CreateReactApp("go-furnace", false, wg) },
	"CRA - Typescript":    func() { createreactapp.CreateReactApp("go-furnace-ts", true, wg) },
	"Nodejs With MongoDB": func() { nodejswithmongo.CreateNodeJsWithMongo(wg) },
	"Nest.js":             func() { nestjs.CreateNestApp("go-furance-nest-js", wg) },
}

func executeSelected(wg *sync.WaitGroup) {
	for item, selected := range selectedItems {
		wg.Add(1)
		if selected {
			go optionsMap[item]()
		}
	}
}

func main() {
	app := tview.NewApplication()

	wg = &sync.WaitGroup{}

	selectedItemsText = tview.NewTextView().
		SetText("Selected Items:\n").
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)

	list := tview.NewList().
		AddItem("CRA - Javascript", "Create React App Boilerplate with Javascript template.", 'a', func() { toggleItem("CRA - Javascript") }).
		AddItem("CRA - Typescript", "Create React App Boilerplate with Typescript template.", 'b', func() { toggleItem("CRA - Typescript") }).
		AddItem("Nest.js", "Create Nest.js project template.", 'c', func() { toggleItem("Nest.js") }).
		// AddItem("Nodejs with MongoDB", "Nodejs boilerplate with MongoDB setup running on localhost:8080", 'c', func() { toggleItem("Nodejs with MongoDB") }).
		AddItem("Execute", "Execute the selected commands", 'd', func() {
			executeSelected(wg)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(list, 0, 1, true).
		AddItem(selectedItemsText, 0, 1, false) // The last item should not grow with the size of the screen.

	selectedItems = make(map[string]bool)

	if err := app.SetRoot(flex, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}

	wg.Wait()
}

func toggleItem(item string) {
	if selected, exists := selectedItems[item]; exists {
		selectedItems[item] = !selected // Toggle the selection
	} else {
		selectedItems[item] = true // If not selected, mark it as selected
	}
	showSelectedOptions(selectedItems)
}

func showSelectedOptions(selectedItems map[string]bool) {
	selectedItemsText.SetText("Selected Items:\n")
	for item, selected := range selectedItems {
		if selected {
			selectedItemsText.SetText(selectedItemsText.GetText(true) + item + "\n")
		}
	}
}
