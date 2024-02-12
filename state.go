package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

type NameValue struct {
	Id    string
	Name  binding.String
	Value binding.String
}

var inputUrl = binding.NewString()
var topWindow fyne.Window

var params = binding.NewUntypedList()
var headers = binding.NewUntypedList()
var selectedMethod string
var currentDir string = os.Getenv("HOME")
var tree *xwidget.FileTree
var output *widget.Entry
