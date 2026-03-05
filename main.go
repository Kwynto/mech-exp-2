package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	fyneApp := app.New()
	mainWindow := fyneApp.NewWindow("Мечталлион 2 - анализатор")
	mainWindow.Resize(fyne.NewSize(400, 100))
	mainWindow.SetFixedSize(true)

	btnShowDB := widget.NewButton("Показать базу тиражей", func() {
		fmt.Println("Типа показали базу.")
	})

	btnAddGame := widget.NewButton("Добавить данные тиража", func() {
		fmt.Println("Типа добавили данные.")
	})

	btnAnalize01 := widget.NewButton("Дефектный анализ", func() {
		fmt.Println("Типа сделали дефектный анализ.")
	})

	mainWindow.SetContent(container.NewVBox(
		btnShowDB,
		btnAddGame,
		btnAnalize01,
	))

	mainWindow.ShowAndRun()
}
