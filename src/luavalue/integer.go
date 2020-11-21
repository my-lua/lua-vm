package luavalue

import "strconv"

// LuaInteger Lua整数
type LuaInteger int64

// Type s
func (me *LuaInteger) Type() ELuaType {
	return LuaTypeInteger
}

// Value s
func (me *LuaInteger) Value() interface{} {
	return *me
}

// ValueSrc s
func (me *LuaInteger) ValueSrc() int64 {
	return int64(*me)
}

// GetBoolean s
func (me *LuaInteger) GetBoolean() bool {
	if me.ValueSrc() != 0 {
		return true
	}
	return false
}

// GetInteger s
func (me *LuaInteger) GetInteger() int64 {
	return me.ValueSrc()
}

// GetNumber s
func (me *LuaInteger) GetNumber() float64 {
	return float64(me.ValueSrc())
}

// GetString s
func (me *LuaInteger) GetString() string {
	return strconv.FormatInt(me.ValueSrc(), 10)
}

// ToLuaBoolean s
func (me *LuaInteger) ToLuaBoolean() *LuaBoolean {
	return NewLuaBoolean(me.GetBoolean())
}

// ToLuaInteger s
func (me *LuaInteger) ToLuaInteger() *LuaInteger {
	return NewLuaInteger(me.GetInteger())
}

// ToLuaNumber s
func (me *LuaInteger) ToLuaNumber() *LuaNumber {
	return NewLuaNumber(me.GetNumber())
}

// ToLuaString s
func (me *LuaInteger) ToLuaString() *LuaString {
	return NewLuaString(me.GetString())
}

// NewLuaInteger s
func NewLuaInteger(value int64) *LuaInteger {
	result := LuaInteger(value)
	return &result
}
