package main

import (
	"fmt"
	"gocui-demo/miyu_demo/ui"
)

func main() {
	var err error
	err = ui.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ui.Run()
	fmt.Printf("Run rslt: %v\n", err)
}
