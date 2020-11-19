package luavalue

// LuaInteger Lua整数
type LuaInteger int64

// Type s
func (me *LuaInteger) Type() ELuaType {
	return LuaTypeBoolean
}

// Value s
func (me *LuaInteger) Value() interface{} {
	return *me
}

// ValueSrc s
func (me *LuaInteger) ValueSrc() int64 {
	return int64(*me)
}
