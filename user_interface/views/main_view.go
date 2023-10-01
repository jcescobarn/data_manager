package views

import (
	"data_manager/user_interface"

	"github.com/rivo/tview"
)

type MainView struct {
	app       *tview.Application
	flex      *tview.Flex
	navigator *user_interface.Navigator
}

func NewMainView(app *tview.Application, navigator *user_interface.Navigator) *MainView {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	list := tview.NewList()
	list.AddItem("Configuración de conexiones a bases de datos", "Configurar conexiones hacia las distintas bases de datos", 'a', func() {
		navigator.Navigate(NewDatabaseConfigView(app, navigator))
	})
	list.AddItem("Revisar data cargada actualmente", "View actual data load into app", 'b', nil)
	list.AddItem("Convertir Datos", "Realizar conversiones básicas a los datos como campos en blanco y datos atípicos", 'c', nil)
	list.AddItem("Exportar Datos", "Exportar información cargada actualmente en el aplicativo", 'c', nil)
	list.AddItem("Salir", "Presione para salir de la aplicación", 'q', func() {
		app.Stop()
	})

	flex.AddItem(list, 0, 1, true)

	flex.SetBorder(true)
	flex.SetTitle("Menu de gestión")

	return &MainView{
		app:       app,
		flex:      flex,
		navigator: navigator,
	}
}

func (mv *MainView) Render() {
	mv.app.SetRoot(mv.flex, true).Run()
}
