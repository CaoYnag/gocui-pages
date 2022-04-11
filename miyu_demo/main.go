package main

import (
	"fmt"
	"gocui-demo/miyu_demo/ui"
)

func main() {
	// e := ui.Run(ui.GetLogin("spes", ""))
	// e := ui.Run(ui.GetReg("spes"))
	// e := ui.Run(ui.GetDesktop())
	e := ui.Run(ui.GetChat("Ehh?"))
	fmt.Println(e)
}
