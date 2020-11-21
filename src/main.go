package main

import (
	"fmt"

	"./luavalue"
	"./state"
)

func main() {
	num := 3.99999999999
	fmt.Println(int(num))
	s := state.NewLuaStack(10)
	boolValue := luavalue.NewLuaBoolean(true)
	s.Push(boolValue)
	intValue := luavalue.NewLuaInteger(13)
	s.Push(intValue)
	numValue := luavalue.NewLuaNumber(3.14)
	s.Push(numValue)
	strValue := luavalue.NewLuaString("你好，世界")
	s.Push(strValue)
	s.Push(strValue)
	s.Push(strValue)
	s.Push(strValue)
	s.Push(strValue)
	s.Push(strValue)
	s.Push(strValue)
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
}
