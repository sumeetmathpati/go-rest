// Package main provides various examples of Fyne API capabilities.
package main

import (
	"fmt"
	"gorest/data"
	"gorest/tutorials"
	"log"
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var topWindow fyne.Window

func main() {

	a := app.NewWithID("io.fyne.demo")
	a.SetIcon(data.FyneLogo)
	// makeTray(a)
	logLifecycle(a)
	w := a.NewWindow("Fyne Demo")
	topWindow = w

	w.SetMainMenu(makeMenu(a, w))
	w.SetMaster()

	urlAndContent := container.NewBorder(
		container.NewVBox(makeUrlBox(), widget.NewSeparator()),
		nil,
		nil,
		nil,
		makeMainContent(w),
	)

	navAndContentSplit := container.NewHSplit(makeNav(func(t tutorials.Tutorial) {}, true), urlAndContent)
	navAndContentSplit.Offset = 0.2
	w.SetContent(navAndContentSplit)
	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}

func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)
	checkedItem := fyne.NewMenuItem("Checked", nil)
	checkedItem.Checked = true
	disabledItem := fyne.NewMenuItem("Disabled", nil)
	disabledItem.Disabled = true
	otherItem := fyne.NewMenuItem("Other", nil)
	mailItem := fyne.NewMenuItem("Mail", func() { fmt.Println("Menu New->Other->Mail") })
	mailItem.Icon = theme.MailComposeIcon()
	otherItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Project", func() { fmt.Println("Menu New->Other->Project") }),
		mailItem,
	)
	fileItem := fyne.NewMenuItem("File", func() { fmt.Println("Menu New->File") })
	fileItem.Icon = theme.FileIcon()
	dirItem := fyne.NewMenuItem("Directory", func() { fmt.Println("Menu New->Directory") })
	dirItem.Icon = theme.FolderIcon()
	newItem.ChildMenu = fyne.NewMenu("",
		fileItem,
		dirItem,
		otherItem,
	)

	openSettings := func() {
		w := a.NewWindow("Fyne Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(440, 520))
		w.Show()
	}
	settingsItem := fyne.NewMenuItem("Settings", openSettings)
	settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	settingsItem.Shortcut = settingsShortcut
	w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings()
	})

	cutShortcut := &fyne.ShortcutCut{Clipboard: w.Clipboard()}
	cutItem := fyne.NewMenuItem("Cut", func() {
		shortcutFocused(cutShortcut, w)
	})
	cutItem.Shortcut = cutShortcut
	copyShortcut := &fyne.ShortcutCopy{Clipboard: w.Clipboard()}
	copyItem := fyne.NewMenuItem("Copy", func() {
		shortcutFocused(copyShortcut, w)
	})
	copyItem.Shortcut = copyShortcut
	pasteShortcut := &fyne.ShortcutPaste{Clipboard: w.Clipboard()}
	pasteItem := fyne.NewMenuItem("Paste", func() {
		shortcutFocused(pasteShortcut, w)
	})
	pasteItem.Shortcut = pasteShortcut
	performFind := func() { fmt.Println("Menu Find") }
	findItem := fyne.NewMenuItem("Find", performFind)
	findItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyF, Modifier: fyne.KeyModifierShortcutDefault | fyne.KeyModifierAlt | fyne.KeyModifierShift | fyne.KeyModifierControl | fyne.KeyModifierSuper}
	w.Canvas().AddShortcut(findItem.Shortcut, func(shortcut fyne.Shortcut) {
		performFind()
	})

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://developer.fyne.io")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItem("Support", func() {
			u, _ := url.Parse("https://fyne.io/support/")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Sponsor", func() {
			u, _ := url.Parse("https://fyne.io/sponsor/")
			_ = a.OpenURL(u)
		}))

	// a quit item will be appended to our first (File) menu
	file := fyne.NewMenu("File", newItem, checkedItem, disabledItem)
	device := fyne.CurrentDevice()
	if !device.IsMobile() && !device.IsBrowser() {
		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}
	main := fyne.NewMainMenu(
		file,
		fyne.NewMenu("Edit", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator(), findItem),
		helpMenu,
	)
	checkedItem.Action = func() {
		checkedItem.Checked = !checkedItem.Checked
		main.Refresh()
	}
	return main
}

func makeUrlBox() *fyne.Container {
	name := widget.NewEntryWithData(str)
	name.SetPlaceHolder("Request URL")

	methodSelect := widget.NewSelect(methods, func(s string) {
		selectedMethod = strings.TrimSpace(s)
		fmt.Println(selectedMethod)
	})
	methodSelect.PlaceHolder = fmt.Sprintf("%-12s", methods[0])
	return container.NewBorder(nil, nil, methodSelect, widget.NewButton("Send", func() {}), name)
}

func makeTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		h := fyne.NewMenuItem("Hello", func() {})
		h.Icon = theme.HomeIcon()
		menu := fyne.NewMenu("Hello World", h)
		h.Action = func() {
			log.Println("System tray menu tapped")
			h.Label = "Welcome"
			menu.Refresh()
		}
		desk.SetSystemTrayMenu(menu)
	}
}

func makeNav(setTutorial func(tutorial tutorials.Tutorial), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()

	// tree := &widget.Tree{
	// 	ChildUIDs: func(uid string) []string {
	// 		return tutorials.TutorialIndex[uid]
	// 	},
	// 	IsBranch: func(uid string) bool {
	// 		children, ok := tutorials.TutorialIndex[uid]

	// 		return ok && len(children) > 0
	// 	},
	// 	CreateNode: func(branch bool) fyne.CanvasObject {
	// 		return widget.NewLabel("Collection Widgets")
	// 	},
	// 	UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
	// 		t, ok := tutorials.Tutorials[uid]
	// 		if !ok {
	// 			fyne.LogError("Missing tutorial panel: "+uid, nil)
	// 			return
	// 		}
	// 		obj.(*widget.Label).SetText(t.Title)
	// 		if unsupportedTutorial(t) {
	// 			obj.(*widget.Label).TextStyle = fyne.TextStyle{Italic: true}
	// 		} else {
	// 			obj.(*widget.Label).TextStyle = fyne.TextStyle{}
	// 		}
	// 	},
	// 	OnSelected: func(uid string) {
	// 		if t, ok := tutorials.Tutorials[uid]; ok {
	// 			if unsupportedTutorial(t) {
	// 				return
	// 			}
	// 			a.Preferences().SetString(preferenceCurrentTutorial, uid)
	// 			setTutorial(t)
	// 		}
	// 	},
	// }

	// if loadPrevious {
	// 	currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
	// 	tree.Select(currentPref)
	// }

	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)

	// return container.NewBorder(nil, themes, nil, nil, tree)
	return container.NewBorder(nil, themes, nil, nil, nil)
}

func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	switch sh := s.(type) {
	case *fyne.ShortcutCopy:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutCut:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutPaste:
		sh.Clipboard = w.Clipboard()
	}
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}

func makeMainContent(_ fyne.Window) fyne.CanvasObject {
	left := container.NewAppTabs(
		container.NewTabItem("Params", makeParamsWidget()),
		// container.NewTabItem("Headers", makeHeadersWidget()),
		container.NewTabItem("Auth", makeAuthWidget()),
		container.NewTabItem("Body", widget.NewLabel("Content of tab 3")),
	)
	// left.Wrapping = fyne.TextWrapWord
	// left.SetText("Long text is looooooooooooooong")
	right := widget.NewMultiLineEntry()
	return container.NewHSplit(container.NewVScroll(left), right)
}

func makeTabLocationSelect(callback func(container.TabLocation)) *widget.Select {
	locations := widget.NewSelect([]string{"Top", "Bottom", "Leading", "Trailing"}, func(s string) {
		callback(map[string]container.TabLocation{
			"Top":      container.TabLocationTop,
			"Bottom":   container.TabLocationBottom,
			"Leading":  container.TabLocationLeading,
			"Trailing": container.TabLocationTrailing,
		}[s])
	})
	locations.SetSelected("Top")
	return locations
}

func makeParamsWidget() fyne.CanvasObject {

	return container.NewBorder(
		widget.NewButton("Add Param", func() {
			params.Append(NameValue{
				Name: binding.NewString(), Value: binding.NewString(),
			})

		}),
		nil,
		nil,
		nil,

		widget.NewListWithData(
			params,
			func() fyne.CanvasObject {
				return container.NewGridWithColumns(
					2,
					widget.NewEntry(),
					widget.NewEntry(),
				)
			},
			func(i binding.DataItem, o fyne.CanvasObject) {
				v, _ := i.(binding.Untyped).Get()
				container := o.(*fyne.Container)
				container.Objects[0].(*widget.Entry).Bind(v.(NameValue).Name)
				container.Objects[1].(*widget.Entry).Bind(v.(NameValue).Value)
			},
		),
	)
}

// func makeHeadersWidget() fyne.CanvasObject {
// 	// data := []NameValue{
// 	// 	NameValue{Name: binding.NewString(), Value: binding.NewString()},
// 	// 	NameValue{Name: binding.NewString(), Value: binding.NewString()},
// 	// }

// 	return container.NewBorder(
// 		// container.NewVBox(
// 		// 	container.NewGridWithColumns(
// 		// 		3,
// 		// 		widget.NewEntry(),
// 		// 		widget.NewEntry(),
// 		// 		widget.NewButton("+", func() {}),
// 		// 	), widget.NewSeparator(),
// 		// ),
// 		widget.NewButton("Add Header", func() {}),
// 		nil,
// 		nil,
// 		nil,
// 		widget.NewList(
// 			func() int {
// 				return len([data])
// 			},
// 			func() fyne.CanvasObject {
// 				return container.NewGridWithColumns(
// 					2,
// 					widget.NewEntry(),
// 					widget.NewEntry(),
// 				)
// 			},
// 			func(id widget.ListItemID, item fyne.CanvasObject) {
// 				container := item.(*fyne.Container)
// 				container.Objects[0].(*widget.Entry).Bind(data[id].Name)
// 				container.Objects[1].(*widget.Entry).Bind(data[id].Value)
// 			},
// 		),
// 	)
// }

func makeAuthWidget() fyne.CanvasObject {

	authNames := []string{
		authDisplay(none),
		authDisplay(bearer),
		authDisplay(basic),
		authDisplay(key),
	}
	return container.NewBorder(
		// container.NewVBox(
		// 	container.NewGridWithColumns(
		// 		3,
		// 		widget.NewEntry(),
		// 		widget.NewEntry(),
		// 		widget.NewButton("+", func() {}),
		// 	), widget.NewSeparator(),
		// ),
		widget.NewSelect(authNames, func(s string) { fmt.Println("selected", s) }),
		nil,
		nil,
		nil,
		// widget.NewList(
		// 	func() int {
		// 		return len(data)
		// 	},
		// 	func() fyne.CanvasObject {
		// 		return container.NewGridWithColumns(
		// 			2,
		// 			headerName,
		// 			headerValue,
		// 		)
		// 	},
		// 	func(id widget.ListItemID, item fyne.CanvasObject) {},
		// ),
	)
}
