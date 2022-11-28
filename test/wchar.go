package main

import "fmt"

func main() {
	n1, _ := fmt.Println("ascii asdnioqu172841958109(*^&%%$$%^)") // 37
	n2, _ := fmt.Println("这是一段中文！！")                              // 8
	n3, _ := fmt.Println("やばい!")                                  // 4
	n4, _ := fmt.Println("▲№§→←↑◎◇↓↓℃℃￣＼＆☆")                      // 16
	n5, _ := fmt.Println("end")                                   // 4
	fmt.Printf("%d %d %d %d %d\n", n1, n2, n3, n4, n5)

	var ch rune
	ch = '这'
	var s string = "这"
	raw := ([]byte)(s)
	fmt.Printf("%x %v\n", ch, raw)
}
