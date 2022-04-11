package main

import (
	"fmt"

	"github.com/russross/blackfriday/v2"
)

var (
	SOURCE = `# Header
## header2

- list 1
- list 2

1. ordered 1
2. ordered 2

[link](#anchor)
![img](some_link)

- [ ] todo
- [x] done
`
)

func main() {
	buf := blackfriday.Run([]byte(SOURCE))
	fmt.Println(string(buf))
}
