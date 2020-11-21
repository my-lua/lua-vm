package luavalue

import "strconv"

// LuaNumber Lua浮点数
type LuaNumber float64

// Type s
func (me *LuaNumber) Type() ELuaType {
	return LuaTypeBoolean
}

// Value s
func (me *LuaNumber) Value() interface{} {
	return *me
}

// ValueSrc s
func (me *LuaNumber) ValueSrc() float64 {
	return float64(*me)
}

// GetBoolean s
func (me *LuaNumber) GetBoolean() bool {
	if me.ValueSrc() != 0.0 {
		return true
	}
	return false
}

// GetInteger s
func (me *LuaNumber) GetInteger() int64 {
	return int64(me.ValueSrc())
}

// GetNumber s
func (me *LuaNumber) GetNumber() float64 {
	return me.ValueSrc()
}

// GetString s
func (me *LuaNumber) GetString() string {
	return strconv.FormatFloat(me.ValueSrc(), 'f', -1, 64)
}

// ToLuaBoolean s
func (me *LuaNumber) ToLuaBoolean() *LuaBoolean {
	return NewLuaBoolean(me.GetBoolean())
}

// ToLuaInteger s
func (me *LuaNumber) ToLuaInteger() *LuaInteger {
	return NewLuaInteger(me.GetInteger())
}

// ToLuaNumber s
func (me *LuaNumber) ToLuaNumber() *LuaNumber {
	return NewLuaNumber(me.GetNumber())
}

// ToLuaString s
func (me *LuaNumber) ToLuaString() *LuaString {
	return NewLuaString(me.GetString())
}

// NewLuaNumber s
func NewLuaNumber(value float64) *LuaNumber {
	result := LuaNumber(value)
	return &result
}
