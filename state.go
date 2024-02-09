package main

import "fyne.io/fyne/v2/data/binding"

type NameValue struct {
	Name  binding.String
	Value binding.String
}

var str = binding.NewString()

var params = binding.NewUntypedList()
var headers = binding.NewUntypedList()

var selectedMethod string
