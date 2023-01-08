package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"github.com/labstack/echo"
	"github.com/zserge/lorca"
)

const (
	isElecton = true
)

// Set logger
var l = log.New(log.Writer(), log.Prefix(), log.Flags())

func main() {
	go Serve()

	if !isElecton {
		Lorca() // based Chrome
	} else {
		Electon() // based Electon
	}

}

func Electon() {

	// Create astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:           "Test",
		BaseDirectoryPath: "example",
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	// New window
	var w *astilectron.Window
	if w, err = a.NewWindow(`http://localhost:8000/`, &astilectron.WindowOptions{
		Center:         astikit.BoolPtr(true),
		Height:         astikit.IntPtr(1920),
		Width:          astikit.IntPtr(1080),
		MinHeight:      astikit.IntPtr(1080),
		MinWidth:       astikit.IntPtr(720),
		Fullscreenable: astikit.BoolPtr(true),
		Resizable:      astikit.BoolPtr(true),
	}); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	// Blocking pattern
	a.Wait()
}

func Lorca() {
	// ブラウザを起動
	ui, _ := lorca.New("http://localhost:8000", "", 1920, 1080)
	defer ui.Close()
	<-ui.Done()
}

func Serve() {
	e := echo.New()

	// Cross compile
	abs, err := filepath.Abs("./frontend/dist")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("----------", abs)
	l.Println("----------", abs)
	e.Static("/", abs)
	e.Logger.Fatal(e.Start(":8000"))
}
