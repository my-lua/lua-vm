package value

// LuaNone s
type LuaNone int64

// Type s
func (me *LuaNone) Type() ELuaType {
	return LuaTypeNone
}

// Value s
func (me *LuaNone) Value() interface{} {
	return nil
}

// GetBoolean s
func (me *LuaNone) GetBoolean() bool {
	return false
}

// GetInteger s
func (me *LuaNone) GetInteger() int64 {
	panic("error")
}

// GetNumber s
func (me *LuaNone) GetNumber() float64 {
	panic("error")
}

// GetString s
func (me *LuaNone) GetString() string {
	return "none"
}

// ToLuaBoolean s
func (me *LuaNone) ToLuaBoolean() LuaBoolean {
	return NewLuaBoolean(me.GetBoolean())
}

// ToLuaInteger s
func (me *LuaNone) ToLuaInteger() LuaInteger {
	return NewLuaInteger(me.GetInteger())
}

// ToLuaNumber s
func (me *LuaNone) ToLuaNumber() LuaNumber {
	return NewLuaNumber(me.GetNumber())
}

// ToLuaString s
func (me *LuaNone) ToLuaString() LuaString {
	return NewLuaString(me.GetString())
}
