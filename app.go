package main

import (
	"context"
	"dcalc/lib/calc"
	"dcalc/lib/dev"
	"fmt"
	"log"
	"os"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx     context.Context
	Menu    *menu.Menu
	state   *dev.AppState
	route   string
	version string
}

type Function func(*dev.AppState) string

// NewApp creates a new App application struct
func NewApp() *App {
	app := &App{
		Menu:    menu.NewMenu(),
		state:   dev.NewAppState(),
		version: Version,
	}

	AppMenu := app.Menu.AddSubmenu("About")
	AppMenu.AddText("About DCalc", keys.CmdOrCtrl("i"), func(_ *menu.CallbackData) {
		app.Navigate("about")
	})

	AppMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	FileMenu := app.Menu.AddSubmenu("File")
	FileMenu.AddText("Export CSV", keys.CmdOrCtrl("e"), func(_ *menu.CallbackData) {
		app.Export(ExportTypeCSV)
	})

	app.Menu.Append(menu.EditMenu())
	app.Menu.Append(menu.WindowMenu())

	HelpMenu := app.Menu.AddSubmenu("Help")
	HelpMenu.AddText("Guide", keys.Control("h"), func(_ *menu.CallbackData) {
		app.Navigate("guide")
	})

	return app
}

func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
}

func (app *App) EmitStateChanged() {
	runtime.EventsEmit(app.ctx, "stateChanged", app.state)
}

func (app *App) GetState() *dev.AppState {
	return app.state
}

func (app *App) Calculate() {
	state := app.state
	node := state.GetCurrentNode()

	result, ok := node.Calculate()
	if !ok {
		return
	}

	mainCal := fmt.Sprintf("%f %s %f", state.Value, node.Operator, result)
	mainValue, err := calc.Calculate(mainCal)
	if err != nil {
		return
	}

	state.Value = mainValue

	state.AddNode()
}

func (app *App) OnTyping(key string) {
	node := app.state.GetCurrentNode()

	if key == "`" {
		app.state.ToggleHideName()
		app.EmitStateChanged()
		return
	}

	if key == "Escape" {
		if app.state.Clearing {
			app.state.CancelClear()
		} else {
			app.state.ToClear()
		}
		app.EmitStateChanged()
		return
	}

	if key == "Backspace" {
		node.Pop()
		app.EmitStateChanged()
		return
	}

	if key == "Enter" {
		if app.state.Clearing {
			app.state.Reset()
		} else {
			app.Calculate()
		}
		app.EmitStateChanged()
		return
	}

	if key == "+" || key == "-" || key == "*" || key == "/" {
		if node.IsEmpty() {
			node.Operator = dev.NodeOperator(key)
			app.EmitStateChanged()
			return
		}
	}

	app.state.AddBuffer(key)
	app.EmitStateChanged()
}

func (app *App) GetRecords() [][]string {
	var records [][]string
	for _, node := range app.state.Nodes {
		records = append(records, app.state.NodeToRecord(&node))
	}
	records = append(records, []string{"", "", "Total = ", app.state.Helpers.FormatFloat(app.state.Value)})
	return records
}

func (app *App) Export(exportType ExportType) {

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fileName, err := runtime.SaveFileDialog(app.ctx, runtime.SaveDialogOptions{
		Title: "Export CSV",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "CSV Files (*.csv)",
				Pattern:     "*.csv",
			},
		},
		DefaultDirectory: dirname + "/Downloads",
		DefaultFilename:  "export.csv",
	})

	if err != nil {
		fmt.Println(err)
	}

	NewExporter(app).Data().ToCSV(fileName)
}

func (app *App) Navigate(route string) {
	app.route = route
	runtime.EventsEmit(app.ctx, "routeChanged", app.route)
}

func (app *App) GetVersion() string {
	return app.version
}
