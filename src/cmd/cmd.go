package cmd

import (
	"github.com/HexChristmas/Christmas/src/ui"
	"github.com/HexChristmas/Christmas/src/util"
)

func Start() {
	util.DetectOS()
	util.ClearScreen()
	ui.StartMenu()
}
