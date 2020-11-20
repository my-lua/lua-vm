package state

import (
	"../luavalue"
)

// LuaStack Lua栈
type LuaStack struct {
	top   int
	slots []luavalue.ILuaValue
}

// Size 获取栈容量
func (me *LuaStack) Size() int {
	return len(me.slots)
}

// Push 向栈顶压入一个Lua值
func (me *LuaStack) Push(value luavalue.ILuaValue) {
	if me.top >= me.Size() {
		panic("LuaStack Push: 栈空间不足")
	}
	me.slots[me.top] = value
	me.top++
}

// Pop 从栈顶弹出一个Lua值
func (me *LuaStack) Pop() luavalue.ILuaValue {
	if me.top < 1 {
		panic("LuaStack Pop: 栈已为空")
	}
	me.top--
	result := me.slots[me.top]
	luaNil := luavalue.NewLuaNil()
	me.slots[me.top] = &luaNil
	return result
}

// AbsIndex 相对索引转换成为绝对索引（内部索引）
func (me *LuaStack) AbsIndex(index int) int {
	if index >= 0 {
		return index
	}
	return me.top + index + 1
}

// IndexIsValid 判断索引是否在有效范围之内（外部索引）
func (me *LuaStack) IndexIsValid(index int) bool {
	absIndex := me.AbsIndex(index)
	return absIndex > 0 && absIndex <= me.top
}

// Get 根据索引从栈内获取Lua值（外部索引）
func (me *LuaStack) Get(index int) luavalue.ILuaValue {
	if me.IndexIsValid(index) {
		return me.slots[me.AbsIndex(index)-1]
	}
	result := luavalue.NewLuaNil()
	return &result
}

// Set 设置Lua值到栈中指定索引（外部索引）
func (me *LuaStack) Set(index int, value luavalue.ILuaValue) {
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
