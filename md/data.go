package main

const (
	SOURCE = `# Header
## header2

hello,*world!*
-bye-,**world!**

- list 1
- list 2

---

1. ordered 1
2. ordered 2

[link](#anchor)
![img](some_link)

- [ ] todo
- [x] done
` +
		"```cpp\n" +
		"#include <iostream>\n" +
		"using namespace std;\n" +
		"int main(){}\n" +
		"```"

	COLOR_GREY uint = iota
	COLOR_RED
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_PURPLE
	COLOR_PINK
	COLOR_CYAN
	COLOR_WHITE
)
