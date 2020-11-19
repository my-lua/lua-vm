package luavalue

import "fmt"

// ILuaValue Lua值接口
type ILuaValue interface {
	Type() ELuaType
	Value() interface{}
}

func ssss() {
	var a interface{}
	a = 13
	b := a.(ILuaValue)
	fmt.Println(b.Value())
}
