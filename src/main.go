package main

import (
	"fmt"
	"strconv"

	"./luavalue"
)

func main() {
	var num float64 = 3.14167
	fmt.Println(strconv.FormatFloat(num, 'f', -1, 64))
	i := luavalue.NewLuaInteger(0)
	var itf luavalue.ILuaValue
	itf = &i
	fmt.Println(itf.GetBoolean())
}
