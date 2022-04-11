package main

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
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
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(SOURCE), &buf); err != nil {
		panic(err)
	}
	fmt.Println(string(buf.Bytes()))
}
