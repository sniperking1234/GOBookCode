package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i, arg := range os.Args[1:] {
		// 无论切片从哪开始，索引的都是从0开始计数
		fmt.Println(strconv.Itoa(i+1) + " " + arg)
	}
}
