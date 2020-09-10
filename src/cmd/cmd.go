package cmd

import (
	"github.com/1os4r/Controls/src/ui"
	"github.com/1os4r/Controls/src/util"
)

func Start() {
	util.DetectOS()
	util.ClearScreen()
	ui.StartMenu()
}
