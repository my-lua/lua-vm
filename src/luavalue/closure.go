package luavalue

import (
	"../chunk"
)

// LuaClosure Lua闭包（函数）
type LuaClosure struct {
	proto *chunk.Prototype
}

// Type s
func (me *LuaClosure) Type() ELuaType {
	return LuaTypeFunction
}

// Value s
func (me *LuaClosure) Value() interface{} {
	return *me
}

// ValueSrc s
func (me *LuaClosure) ValueSrc() LuaClosure {
	return me.Value().(LuaClosure)
}

// GetBoolean s
func (me *LuaClosure) GetBoolean() bool {
	return true
}

// GetInteger s
func (me *LuaClosure) GetInteger() int64 {
	return 1
}

// GetNumber s
func (me *LuaClosure) GetNumber() float64 {
	return 1.0
}

// GetString s
func (me *LuaClosure) GetString() string {
	return "[function]"
}

// ToLuaBoolean s
func (me *LuaClosure) ToLuaBoolean() LuaBoolean {
	return NewLuaBoolean(me.GetBoolean())
}

// ToLuaInteger s
func (me *LuaClosure) ToLuaInteger() LuaInteger {
	return NewLuaInteger(me.GetInteger())
}

// ToLuaNumber s
func (me *LuaClosure) ToLuaNumber() LuaNumber {
	return NewLuaNumber(me.GetNumber())
}

// ToLuaString s
func (me *LuaClosure) ToLuaString() LuaString {
	return NewLuaString(me.GetString())
}

// Prototype 获取函数原型
func (me *LuaClosure) Prototype() *chunk.Prototype {
	return me.proto
}

// NewLuaClosure 构造函数
func NewLuaClosure(proto *chunk.Prototype) *LuaClosure {
	return &LuaClosure{proto}
}
