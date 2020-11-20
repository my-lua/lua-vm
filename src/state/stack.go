package state

import (
	"../luavalue"
)

// LuaStack Lua栈
type LuaStack struct {
	top   int64
	slots []luavalue.ILuaValue
}

func (me *LuaStack) Push(value luavalue.ILuaValue) {
	me.slots[me.top] = value
}

func (me *LuaStack) Pop() luavalue.ILuaValue {
	me.top--
	result := me.slots[me.top]
	return result
}

// AbsIndex 相对索引转换成为绝对索引（内部索引）
func (me *LuaStack) AbsIndex(index int64) int64 {
	if index >= 0 {
		return index
	}
	return me.top + index + 1
}

// IndexIsValid 判断索引是否在有效范围之内（外部索引）
func (me *LuaStack) IndexIsValid(index int64) bool {
	absIndex := me.AbsIndex(index)
	return absIndex > 0 && absIndex <= me.top
}

// Get 根据索引从栈内获取Lua值（外部索引）
func (me *LuaStack) Get(index int64) luavalue.ILuaValue {
	if me.IndexIsValid(index) {
		return me.slots[me.AbsIndex(index)-1]
	}
	return me.slots[me.AbsIndex(index)-1]
}

// Set 设置Lua值到栈中指定索引（外部索引）
func (me *LuaStack) Set(index int64, value luavalue.ILuaValue) {
	if me.IndexIsValid(index) {
		me.slots[me.AbsIndex(index)-1] = value
		return
	}
	panic("LuaStack Set: 错误的索引")
}

// NewLuaStack 构造函数
func NewLuaStack(size int) *LuaStack {
	return &LuaStack{
		top:   0,
		slots: make([]luavalue.ILuaValue, size),
	}
}
