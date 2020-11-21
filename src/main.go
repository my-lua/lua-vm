package main

import (
	"fmt"

	"./luavalue"
)

func main() {
	table := luavalue.NewLuaTable(10, 0)
	// key := luavalue.NewLuaString("nihao")
	key2 := luavalue.NewLuaNumber(3.14)
	table.Put(key2, luavalue.NewLuaInteger(13))
	value := table.Get(key2)
	fmt.Println(value.GetString())

	// num := 3.99999999999
	// fmt.Println(int(num))

	// tmap := make(map[string]luavalue.ILuaValue, 1)

	// fmt.Println(tmap["2"])

	// s := state.NewLuaStack(10)
	// boolValue := luavalue.NewLuaBoolean(true)
	// s.Push(boolValue)
	// intValue := luavalue.NewLuaInteger(13)
	// s.Push(intValue)
	// numValue := luavalue.NewLuaNumber(3.14)
	// s.Push(numValue)
	// strValue := luavalue.NewLuaString("你好，世界")
	// s.Push(strValue)
	// s.Push(strValue)
	// s.Push(strValue)
	// s.Push(strValue)
	// s.Push(strValue)
	// s.Push(strValue)
	// s.Push(strValue)
	// s.Print()
	// s.Pop()
	// s.Print()
	// s.Pop()
	// s.Print()
	// s.Pop()
	// s.Print()
	// s.Pop()
	// s.Print()
	// s.Pop()
	// s.Print()
}
