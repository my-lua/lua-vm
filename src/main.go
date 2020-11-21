package main

import (
	"fmt"

	"./luavalue"
)

func main() {
	table := luavalue.NewLuaTable(0, 0)
	key := luavalue.NewLuaInteger(1)
	key2 := luavalue.NewLuaNumber(0.9999)
	value := luavalue.NewLuaNumber(3.14)
	table.Put(key, value)
	fmt.Println(table.Get(key2))
}
