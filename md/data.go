package main

const (
	SOURCE = `# Header
## header2

hello,*world!*
-bye-,**world!**

- list 1
- list 2

> ref here

******

------

======

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
		"```\n" +
		"do not `think`, do it!\n" +
		`
then you will --free--
==here== we are
`
)
