package main

import (
	"os"

	"fyne.io/fyne/v2/data/binding"
	xwidget "fyne.io/x/fyne/widget"
)

type NameValue struct {
	Name  binding.String
	Value binding.String
}

var inputUrl = binding.NewString()

var params = binding.NewUntypedList()
var headers = binding.NewUntypedList()
var selectedMethod string
var currentDir string = os.Getenv("HOME")
var tree *xwidget.FileTree
