package main

import (
	"fmt"

	"github.com/gomarkdown/markdown"
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
	n := markdown.ToHTML([]byte(SOURCE), nil, nil)
	fmt.Println(string(n))
}
