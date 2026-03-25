package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// Приложение
var fyneApp fyne.App = app.New()
var iconApp fyne.Resource

// Главное окно
var mainWindow fyne.Window = fyneApp.NewWindow("Мечталлион 2 - анализатор")
var btnShowDB *widget.Button = widget.NewButton("Показать базу тиражей", showGameDB)
var btnAddGame *widget.Button = widget.NewButton("Добавить данные тиража", addGame)
var btnDefAnalize *widget.Button = widget.NewButton("Дефектный анализ", makeDefAnalize)
