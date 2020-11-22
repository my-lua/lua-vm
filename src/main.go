package main

import (
	"fmt"

	"./luavalue"
)

func main() {
	table := luavalue.NewLuaTable(0, 0)
	key1 := luavalue.NewLuaInteger(1)
	key2 := luavalue.NewLuaNumber(1.1)
	key3 := luavalue.NewLuaString("1.1")
	key4 := luavalue.NewLuaNil()
	value := luavalue.NewLuaNumber(3.14)
	table.Put(key1, value)
	table.Put(key2, value)
	table.Put(key3, value)
	table.Put(key4, value)
	fmt.Println(table.Get(key1))
	fmt.Println(table.Get(key2))
	fmt.Println(table.Get(key3))
	fmt.Println(table.Get(key4))
	table.Print()
}
