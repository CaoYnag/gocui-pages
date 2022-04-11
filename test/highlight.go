package main

import (
	"fmt"

	"github.com/d4l3k/go-highlight"
)

func test(lang, code string) {
	fmt.Printf("---------%s----------\n", lang)
	b, e := highlight.Term(lang, []byte(code))
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(string(b))
	}
}

func main() {
	test("go", `
package main

import "fmt"

func main() {
	fmt.Println("Duck!")
}
`)
	test("go", `
package main

import "fmt"

func main() {
fmt.Println("Duck!")
`)

	test("java", `
package main

import "fmt"

void main(string[] args) {
	System.out.println("Duck!")
}
#include <iostream> // wrong
`)
	test("c++", `
#include <iostream>
using namespace std;

int main(int argc, char** argv)
{
	cout << "hello,world!" << endl;
}`)
	test("md", `
#Header
## header
- list1
- list2

1. order list
2. order list

[link](https://blogs.lifesucks.cn)
![image](/home/spes/Pictures/avatar.png)
`)
	test("lua", `
local i = 1
function foo()
{}`)
	test("python", `
import numpy
def foo:
	aaa
m=1`)
}
