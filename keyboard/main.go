package main

import (
	"time"

	"github.com/sincebyte/gkeybd"
)

func main() {
	gkeybd.TypeStr("hello world!\n")
	time.Sleep(3 * time.Second)
	gkeybd.TypeStr("It looks like there is an issue \n" +
		"with your code file, `main.go`. \n" +
		"The error message indicates that \n" +
		"the first line of the file should \n" +
		"contain a 'package' keyword, but \n" +
		"instead, it has a 'EOF' (End Of \n" +
		"File) marker.\n")
	time.Sleep(3 * time.Second)
	gkeybd.TypeStr("123 world!\n")
}
