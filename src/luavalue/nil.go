package luavalue

// LuaNil s
type LuaNil int64

// Type s
func (me *LuaNil) Type() ELuaType {
	return LuaTypeNil
}

// Value s
func (me *LuaNil) Value() interface{} {
	return nil
}

// GetBoolean s
func (me *LuaNil) GetBoolean() bool {
	return false
}

// GetInteger s
func (me *LuaNil) GetInteger() int64 {
	panic("error")
}

// GetNumber s
func (me *LuaNil) GetNumber() float64 {
	panic("error")
}

// GetString s
func (me *LuaNil) GetString() string {
	return "nil"
}

// ToLuaBoolean s
func (me *LuaNil) ToLuaBoolean() LuaBoolean {
	return NewLuaBoolean(me.GetBoolean())
}

// ToLuaInteger s
func (me *LuaNil) ToLuaInteger() LuaInteger {
	return NewLuaInteger(me.GetInteger())
}

// ToLuaNumber s
func (me *LuaNil) ToLuaNumber() LuaNumber {
	return NewLuaNumber(me.GetNumber())
}

// ToLuaString s
func (me *LuaNil) ToLuaString() LuaString {
	return NewLuaString(me.GetString())
}
