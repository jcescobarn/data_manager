package main

import (
	"data_manager/user_interface"
	"data_manager/user_interface/views"

	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()

	navigator := user_interface.NewNavigator(app)

	mainView := views.NewMainView(app, navigator)

	navigator.Navigate(mainView)

	if err := app.Run(); err != nil {
		panic(err)
	}

}
