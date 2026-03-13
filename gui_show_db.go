package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Окно показа базы
var showDBWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - анализатор - архив тиражей")

var showDBTable *widget.Table = widget.NewTable(
	func() (rows int, cols int) {
		return len(SlStGames), 3
	},

	func() fyne.CanvasObject {
		return widget.NewLabel("   ")
	},

	func(tci widget.TableCellID, co fyne.CanvasObject) {
		co.(*widget.Label).SetText(func() string {
			var sTemp string
			switch tci.Col {
			case 0:
				sTemp = fmt.Sprintf("%d", SlStGames[tci.Row].Game)
			case 1:
				sTemp = func(inArr []int) string {
					var sT string
					for _, v := range inArr {
						sT = fmt.Sprintf("%s %d", sT, v)
					}
					return sT
				}(SlStGames[tci.Row].Wins)
			case 2:
				sTemp = func(inArr []int) string {
					var sT string
					for _, v := range inArr {
						sT = fmt.Sprintf("%s %d", sT, v)
					}
					return sT
				}(SlStGames[tci.Row].Wrong)
			}
			return sTemp
		}())
	},
)

func showGameDB() {
	readBase(FILE_NAME)

	mainWindow.Hide()
	showDBWindow.Show()
}
