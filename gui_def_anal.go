package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Kwynto/mech-exp-2/internal/defective"
)

// Окно дефектного анализа
var defAnalizeWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - анализатор - дефектный анализ")

var defAnIsInterval *widget.Check = widget.NewCheck("Интервал", func(b bool) {
	if b {
		defAnLb1.Hide()
		defAnEnt1.Hide()
		defAnLb3.Show()
		defAnEnt3.Show()
		defAnEnt4.Show()
	} else {
		defAnLb1.Show()
		defAnEnt1.Show()
		defAnLb3.Hide()
		defAnEnt3.Hide()
		defAnEnt4.Hide()
	}
})

var (
	defAnLb1 *widget.Label = widget.NewLabel("Кол-во тиражей:")
	defAnLb2 *widget.Label = widget.NewLabel("  Граница:")
	defAnLb3 *widget.Label = widget.NewLabel("  Тиражи:")

	defAnEnt1 *widget.Entry = widget.NewEntry() // Кол-во тиражей
	defAnEnt2 *widget.Entry = widget.NewEntry() // Граница

	// Тиражи
	defAnEnt3 *widget.Entry = widget.NewEntry()
	defAnEnt4 *widget.Entry = widget.NewEntry()

	defAnBtnOK   *widget.Button = widget.NewButton("  OK  ", fDefAnOK)
	defAnBtnInfo *widget.Button = widget.NewButton("i", func() {
		dialog.ShowInformation(
			" ",
			"0 - значение по-умолчанию, означает ДЛЯ ВСЕХ ДАННЫХ.\nДля Границы минимум должен быть 2.\nОпция Интервал предназначена для постанализа.\nБазовый анализ производится без включения опции Интервал.\n",
			defAnalizeWindow,
		)
	})
)

var defAnTopMenu *fyne.Container = container.NewHBox(
	defAnIsInterval,
	defAnLb1,
	defAnEnt1,
	defAnLb3,
	defAnEnt3,
	defAnEnt4,
	defAnLb2,
	defAnEnt2,
	defAnBtnOK,
	widget.NewLabel("   "),
	defAnBtnInfo,
)

func makeDefAnalize() {
	mainWindow.Hide()
	defAnalizeWindow.Show()
}

// Окно положительного дефектного анализа
var defAnGoodWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - анализ - Премиальные номера")

var showDefAnGoodTable *widget.Table = widget.NewTable(
	func() (rows int, cols int) {
		return len(slMainPremiumNumber), 3
	},

	func() fyne.CanvasObject {
		return widget.NewLabel(" ")
	},

	func(tci widget.TableCellID, co fyne.CanvasObject) {
		clear(slMainPremiumNumber)
		clear(slMainRiskNumbers)

		iGame, err := strconv.Atoi(defAnEnt1.Text)
		if err != nil {
			fmt.Println("Ошибка конвертации")
		}

		iBorder, err := strconv.Atoi(defAnEnt2.Text)
		if err != nil {
			fmt.Println("Ошибка конвертации")
		}

		slMainPremiumNumber, _ = defective.StartAnalize(SlStGames, iGame, iBorder)

		co.(*widget.Label).SetText(func() string {
			var sTemp string
			switch tci.Col {
			case 0:
				sTemp = fmt.Sprintf("%d", slMainPremiumNumber[tci.Row].Name)
			case 1:
				sTemp = "номер выиграл раз = "
			case 2:
				sTemp = fmt.Sprintf("%d", slMainPremiumNumber[tci.Row].Sum)
			}
			return sTemp
		}())
	},
)

// Окно отрицательного дефектного анализа
var defAnBedWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - анализ - Рискованные номера")

var showDefAnBedTable *widget.Table = widget.NewTable(
	func() (rows int, cols int) {
		return len(slMainRiskNumbers), 3
	},

	func() fyne.CanvasObject {
		return widget.NewLabel(" ")
	},

	func(tci widget.TableCellID, co fyne.CanvasObject) {
		clear(slMainPremiumNumber)
		clear(slMainRiskNumbers)

		iGame, err := strconv.Atoi(defAnEnt1.Text)
		if err != nil {
			fmt.Println("Ошибка конвертации")
		}

		iBorder, err := strconv.Atoi(defAnEnt2.Text)
		if err != nil {
			fmt.Println("Ошибка конвертации")
		}

		_, slMainRiskNumbers = defective.StartAnalize(SlStGames, iGame, iBorder)

		co.(*widget.Label).SetText(func() string {
			var sTemp string
			switch tci.Col {
			case 0:
				sTemp = fmt.Sprintf("%d", slMainRiskNumbers[tci.Row].Name)
			case 1:
				sTemp = "номер проиграл раз = "
			case 2:
				sTemp = fmt.Sprintf("%d", slMainRiskNumbers[tci.Row].Sum)
			}
			return sTemp
		}())
	},
)

func fDefAnClose() {
	defAnGoodWindow.Hide()
	defAnBedWindow.Hide()
	defAnalizeWindow.Hide()
	mainWindow.Show()
}

func fDefAnOK() {
	clear(slMainPremiumNumber)
	clear(slMainRiskNumbers)

	iGame, err := strconv.Atoi(defAnEnt1.Text)
	if err != nil {
		fmt.Println("Ошибка конвертации")
	}

	iBorder, err := strconv.Atoi(defAnEnt2.Text)
	if err != nil {
		fmt.Println("Ошибка конвертации")
	}

	slMainPremiumNumber, slMainRiskNumbers = defective.StartAnalize(SlStGames, iGame, iBorder)

	defAnGoodWindow.Show()
	defAnBedWindow.Show()
}
