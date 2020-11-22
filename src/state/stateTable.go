package state

import (
	"../luavalue"
)

// CreateTable s
func (me *LuaState) CreateTable(nArr, nRec int) {
	me.stack.Push(luavalue.NewLuaTable(nArr, nRec))
}

// NewTable s
func (me *LuaState) NewTable() {
	me.CreateTable(0, 0)
}

// tableGet s
func (me *LuaState) tableGet(table, key luavalue.ILuaValue) luavalue.ILuaValue {
	if table.Type() == luavalue.LuaTypeTable {
		return table.(*luavalue.LuaTable).Get(key)
	}
	panic("LuaState TableGet: 指定索引处的值并非LuaTable")
}

// TableGet 表Get
// 弹出栈顶的key，并从索引为index的表中获取值后推入栈顶 返回值的类型
func (me *LuaState) TableGet(index int) luavalue.ELuaType {
	value := me.tableGet(me.stack.Get(index), me.stack.Pop())
	me.stack.Push(value)
	return value.Type()
}

// TableGetField 表Get
func (me *LuaState) TableGetField(index int, key string) luavalue.ELuaType {
	value := me.tableGet(me.stack.Get(index), luavalue.NewLuaString(key))
	me.stack.Push(value)
	return value.Type()
}

// TableGetIndex 表Get
func (me *LuaState) TableGetIndex(index int, elIndex int64) luavalue.ELuaType {
	value := me.tableGet(me.stack.Get(index), luavalue.NewLuaInteger(elIndex))
	me.stack.Push(value)
	return value.Type()
}
