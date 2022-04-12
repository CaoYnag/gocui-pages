package main

import "fmt"

func main() {
	for i := 0; i <= 7; i++ {
		for _, j := range []int{1, 4, 7} {
			fmt.Printf("Hello \033[3%d;%dmcolors!\033[0m\n", i, j)
		}
	}
	for i := 0; i < 256; i++ {
		str := fmt.Sprintf("\x1b[48;5;%dm\x1b[30m%4d\x1b[0m", i, i)
		str += fmt.Sprintf("\x1b[38;5;%dm%4d\x1b[0m", i, i)

		if (i+1)%10 == 0 {
			str += "\n"
		}

		fmt.Print(str)
	}
	fmt.Println()
	fmt.Println("\x1b[1mbold\x1b[0m")
	fmt.Println("\x1b[3mitalic\x1b[0m")
	fmt.Println("\x1b[4munderline\x1b[0m")
	fmt.Println("\x1b[9mstrikethrough\x1b[0m")
	fmt.Println("\x1b[1m\x1b[3m\x1b[4m\x1b[9m \033[31;1mcolors!\033[0m\x1b[0m\x1b[0m\x1b[0m\x1b[0m")
	fmt.Println("\x1b[48;5;228m\x1b[30m \x1b[38;5;199m hello,world! \x1b[0m \x1b[0m")
	fmt.Println("\x1b[31;3;4;9;7m\x1b[34;1mhello,world\x1b[0m\x1b[0m")
	fmt.Println("\x1b[48;5;228;38;5;199;mhello\x1b[0m")
}
