package main

import (
	"fmt"

	"./luavalue"
)

func main() {
	table := luavalue.NewLuaTable(0, 0)
	key := luavalue.NewLuaString("nihao")
	value := luavalue.NewLuaNumber(3.14)
	table.Put(key, value)
	fmt.Println(table.Get(key))
}
