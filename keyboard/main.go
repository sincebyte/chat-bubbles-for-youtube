package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/sincebyte/gkeybd"
)

func main() {
	file := os.Args[1:]

	time.Sleep(5 * time.Second)
	data, err := ioutil.ReadFile(file[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(data)
	lines := strings.Split(content, "\n") // 按照换行符分割字符串

	for _, line := range lines { // 遍历切片，并打印每一行
		gkeybd.TypeStr(line)
		gkeybd.TypeStr("\n")
		line = strings.TrimSpace(line)
		if len(line) == 0 || line[0] == '#' {
			time.Sleep(3 * time.Second)
		}
	}
}
