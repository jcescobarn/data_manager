package user_interface

import "github.com/rivo/tview"

type View interface {
	Render()
}

type Navigator struct {
	app         *tview.Application
	currentView View
}

func NewNavigator(app *tview.Application) *Navigator {
	return &Navigator{
		app: app,
	}
}

func (n *Navigator) Navigate(view View) {
	n.currentView = view
	view.Render()
}
