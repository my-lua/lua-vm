package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num float64 = 3.14167
	fmt.Println(strconv.FormatFloat(num, 'f', -1, 64))
}
