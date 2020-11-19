package luavalue

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
