package ui

import (
	"fmt"

	c "github.com/1os4r/Controls/src/color"
	"github.com/1os4r/Controls/src/version"
)

func ShowMenu() {
	fmt.Println("")
	fmt.Println(c.YELLOW, "                                     Christmas v"+version.GetVersion())
	fmt.Println(c.CYAN, "                               	by 1os4r")
	fmt.Println(c.CYAN, "                               	Github: https://github.com/1os4r")
	fmt.Println("")
	fmt.Println(c.GREEN, " Please use `tab` to autocomplete commands,")
	fmt.Println(c.GREEN, " or type `exit` to quit this program.")
	fmt.Println("")
}
