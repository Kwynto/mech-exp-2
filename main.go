package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Kwynto/mech-exp-2/internal/intypes"
)

const FILE_NAME = "data.txt"
const TEXT_MSG_1 = "Запись добалена"

var SlStGames []intypes.TStGame

func fileRead(sName string) (string, error) {
	bRead, err := os.ReadFile(sName)
	if err != nil {
		return "", err
	}
	return string(bRead), nil
}

func convInsideFormat(sData string) []intypes.TStGame {
	var SlStGamesTemp []intypes.TStGame

	slList := strings.Split(sData, "\n")

	for _, sLine := range slList {
		slGame := strings.Split(sLine, "|")

		if len(slGame) == 1 {
			break
		}

		slSWins := strings.Split(slGame[1], "-")
		slSWrongs := strings.Split(slGame[2], "-")

		var slWinNums []int
		var slWrongNums []int

		for _, sWin := range slSWins {
			iNum, err := strconv.Atoi(sWin)
			if err != nil {
				fmt.Println("Conversion failed:", err)
			}
			slWinNums = append(slWinNums, iNum)
		}

		for _, sWrong := range slSWrongs {
			iNum, err := strconv.Atoi(sWrong)
			if err != nil {
				fmt.Println("Conversion failed:", err)
			}
			slWrongNums = append(slWrongNums, iNum)
		}

		iGame, err := strconv.Atoi(slGame[0])
		if err != nil {
			fmt.Println("Conversion failed:", err)
		}

		stGameTemp := intypes.TStGame{
			Game:  iGame,
			Wins:  slWinNums,
			Wrong: slWrongNums,
		}

		SlStGamesTemp = append(SlStGamesTemp, stGameTemp)
	}

	return SlStGamesTemp
}

func readBase(sName string) {
	sData, err := fileRead(sName)
	if err != nil {
		return
	}

	clear(SlStGames)
	SlStGames = convInsideFormat(sData)
}

// Приложение
var fyneApp fyne.App = app.New()
var iconApp fyne.Resource

// Главное окно
var mainWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - анализатор")
var btnShowDB *widget.Button = widget.NewButton("Показать базу тиражей", showGameDB)
var btnAddGame *widget.Button = widget.NewButton("Добавить данные тиража", addGame)
var btnDefAnalize *widget.Button = widget.NewButton("Дефектный анализ", makeDefAnalize)

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

// Окно добавления тиража
var addGameWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - добавить тираж")

var entAddGame *widget.Entry = widget.NewEntry()
var entAddWins *widget.Entry = widget.NewEntry()
var entAddWrongs *widget.Entry = widget.NewEntry()
var sepAddGame *widget.Separator = widget.NewSeparator()
var btnAddGameOK *widget.Button = widget.NewButton("Добаввить запись тиража", fAddGameOK)

func fAddGameOK() {
	defer readBase(FILE_NAME)

	var sNum, sWins, sWrongs string

	sNum = entAddGame.Text
	sWins = entAddWins.Text
	sWrongs = entAddWrongs.Text

	entAddGame.SetText("")
	entAddWins.SetText("")
	entAddWrongs.SetText("")

	sRecord := fmt.Sprintf("%s|%s|%s\n", sNum, sWins, sWrongs)
	fFileName, err := os.OpenFile(FILE_NAME, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer fFileName.Close()

	if _, err := fFileName.WriteString(sRecord); err != nil {
		return
	}

	dialog.ShowInformation(
		" ",
		TEXT_MSG_1,
		addGameWindow,
	)
}

func addGame() {
	mainWindow.Hide()
	addGameWindow.Show()
}

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

	defAnBtnOK *widget.Button = widget.NewButton("  OK  ", func() {
		defAnGoodWindow.Show()
		defAnBedWindow.Show()
	})
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
var defAnGoodWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - анализ - Положительный")

var showDefAnGoodTable *widget.Table = widget.NewTable(
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

// Окно отрицательного дефектного анализа
var defAnBedWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - анализ - Отрицательный")

var showDefAnBedTable *widget.Table = widget.NewTable(
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

func setStartInterface() {
	iconApp, _ = fyne.LoadResourceFromPath("Icon.png")

	mainWindow.Resize(fyne.NewSize(400, 100))
	mainWindow.SetFixedSize(true)
	mainWindow.CenterOnScreen()
	mainWindow.SetIcon(iconApp)
	mainWindow.SetContent(container.NewVBox(
		btnShowDB,
		btnAddGame,
		btnDefAnalize,
	))
	mainWindow.SetOnClosed(func() {
		fyneApp.Quit()
	})

	showDBWindow.Resize(fyne.NewSize(890, 500))
	showDBWindow.CenterOnScreen()
	showDBWindow.SetIcon(iconApp)
	showDBTable.SetColumnWidth(0, 50)
	showDBTable.SetColumnWidth(1, 735)
	showDBTable.SetColumnWidth(2, 85)

	showDBWindow.SetContent(
		showDBTable,
	)

	showDBWindow.SetCloseIntercept(func() {
		showDBWindow.Hide()
		mainWindow.Show()
	})

	addGameWindow.Resize(fyne.NewSize(400, 150))
	addGameWindow.SetFixedSize(true)
	addGameWindow.CenterOnScreen()
	addGameWindow.SetIcon(iconApp)
	entAddGame.SetPlaceHolder("Тираж")
	entAddWins.SetPlaceHolder("Выпали")
	entAddWrongs.SetPlaceHolder("НЕ выпали")

	addGameWindow.SetContent(container.NewVBox(
		entAddGame,
		entAddWins,
		entAddWrongs,
		sepAddGame,
		btnAddGameOK,
	))

	addGameWindow.SetCloseIntercept(func() {
		addGameWindow.Hide()
		mainWindow.Show()
	})

	defAnalizeWindow.Resize(fyne.NewSize(600, 200))
	defAnalizeWindow.SetFixedSize(true)
	defAnalizeWindow.CenterOnScreen()
	defAnalizeWindow.SetIcon(iconApp)

	defAnLb1.Show()
	defAnEnt1.Show()
	defAnLb3.Hide()
	defAnEnt3.Hide()
	defAnEnt4.Hide()

	defAnEnt1.SetText("0")
	defAnEnt2.SetText("2")
	defAnEnt3.SetText("0")
	defAnEnt4.SetText("0")

	defAnalizeWindow.SetContent(
		container.NewCenter(
			container.NewVBox(
				defAnTopMenu,
			),
		),
	)

	defAnalizeWindow.SetCloseIntercept(fDefAnClose)

	showDefAnBedTable.SetColumnWidth(0, 50)
	showDefAnBedTable.SetColumnWidth(1, 30)
	showDefAnBedTable.SetColumnWidth(2, 150)

	defAnGoodWindow.Resize(fyne.NewSize(300, 200))
	defAnBedWindow.Resize(fyne.NewSize(300, 200))

	defAnGoodWindow.SetContent(
		showDefAnGoodTable,
	)

	defAnBedWindow.SetContent(
		showDefAnBedTable,
	)

	defAnBedWindow.SetCloseIntercept(fDefAnClose)
	defAnGoodWindow.SetCloseIntercept(fDefAnClose)
}

func fDefAnClose() {
	defAnGoodWindow.Hide()
	defAnBedWindow.Hide()
	defAnalizeWindow.Hide()
	mainWindow.Show()
}

func main() {
	readBase(FILE_NAME)

	setStartInterface()

	mainWindow.Show()
	fyneApp.Run()
}
