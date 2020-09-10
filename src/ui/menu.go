package ui

import (
	"fmt"

	c "github.com/HexChristmas/Christmas/src/color"
	"github.com/HexChristmas/Christmas/src/version"
)

func ShowMenu() {
	fmt.Println("")
	fmt.Println(c.YELLOW, "                                     Christmas v"+version.GetVersion())
	fmt.Println(c.CYAN, "                               	by HexChristmas")
	fmt.Println(c.CYAN, "                               	Blog: https://HexChristmas.github.io")
	fmt.Println("")
	fmt.Println(c.GREEN, " Please use `tab` to autocomplete commands,")
	fmt.Println(c.GREEN, " or type `exit` to quit this program.")
	fmt.Println("")
}
